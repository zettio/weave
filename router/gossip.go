package router

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"sync"
	"time"
)

const GossipInterval = 30 * time.Second

// GossipData is some data structure keyed by a GossipKey, where each
// part of the structure can be encoded and updated independently.

type GossipKey interface{}
type GossipKeySet map[GossipKey]bool

type GossipData interface {
	AllKeys() GossipKeySet
	Encode(keys GossipKeySet) []byte
	// merge in state and return "set of names where I've just learnt something new",
	OnUpdate(buf []byte) (GossipKeySet, error)
}

type Gossip interface {
	// specific message from one peer to another
	// intermediate peers relay it using unicast topology.
	GossipUnicast(dstPeerName PeerName, buf []byte) error
	// send a message to every peer, relayed using broadcast topology.
	GossipBroadcast(buf []byte) error
}

type Gossiper interface {
	OnGossipUnicast(sender PeerName, msg []byte) error
	OnGossipBroadcast(msg []byte) error
}

// Runs a goroutine to accumulate GossipData changes that need to be sent to one destination,
// and send them when the outgoing socket is not blocked.
type gossipUpdateSender struct {
	sync.Mutex
	pending    GossipKeySet   // which data needs sent
	data       GossipData     // so we can get the latest version of the data
	conn       Connection     // we send the data down this connection
	gossipChan *GossipChannel // to tag the outgoing message
	sendChan   chan<- bool    // channel to sending goroutine
}

func (c *GossipChannel) makeSender(data GossipData, conn Connection) *gossipUpdateSender {
	sendChan := make(chan bool)
	sender := &gossipUpdateSender{pending: make(map[GossipKey]bool), data: data, conn: conn, gossipChan: c, sendChan: sendChan}
	go sender.sendingLoop(sendChan)
	return sender
}

func (sender *gossipUpdateSender) sendAllPending() {
	sender.Lock()
	buf := sender.data.Encode(sender.pending)
	sender.pending = make(GossipKeySet) // Clear out the map
	sender.Unlock()
	sender.conn.(ProtocolSender).SendProtocolMsg(sender.gossipChan.gossipMsg(buf))
}

func (sender *gossipUpdateSender) sendingLoop(sendingChan <-chan bool) {
	for {
		if val := <-sendingChan; !val { // receive zero value when chan is closed
			break
		}
		sender.sendAllPending()
	}
}

type senderMap map[Connection]*gossipUpdateSender

type GossipChannel struct {
	sync.Mutex
	ourself  *LocalPeer
	name     string
	hash     uint32
	gossiper Gossiper
	data     GossipData
	senders  senderMap
}

func (router *Router) NewGossip(channelName string, g Gossiper, d GossipData) Gossip {
	channelHash := hash(channelName)
	channel := &GossipChannel{ourself: router.Ourself, name: channelName, hash: channelHash, gossiper: g, data: d, senders: make(senderMap)}
	router.GossipChannels[channelHash] = channel
	return channel
}

func (router *Router) SendAllGossip() {
	for _, channel := range router.GossipChannels {
		channel.SendGossipUpdateFor(channel.data.AllKeys())
		channel.garbageCollectSenders()
	}
}

func (router *Router) SendAllGossipDown(conn Connection) {
	for _, channel := range router.GossipChannels {
		protocolMsg := channel.gossipMsg(channel.data.Encode(channel.data.AllKeys()))
		conn.(ProtocolSender).SendProtocolMsg(protocolMsg)
	}
}

func (c *GossipChannel) garbageCollectSenders() {
	c.Lock()
	defer c.Unlock()
	newSenders := make(senderMap)
	c.ourself.ForEachConnection(func(_ PeerName, conn Connection) {
		newSenders[conn] = c.senders[conn]
		delete(c.senders, conn)
	})
	for _, sender := range c.senders {
		close(sender.sendChan)
	}
	c.senders = newSenders
}

func (router *Router) handleGossip(payload []byte, onok func(*GossipChannel, PeerName, []byte, *gob.Decoder) error) error {
	decoder := gob.NewDecoder(bytes.NewReader(payload))
	var channelHash uint32
	if err := decoder.Decode(&channelHash); err != nil {
		return err
	}
	channel, found := router.GossipChannels[channelHash]
	if !found {
		return fmt.Errorf("[gossip] received unknown channel with hash %v", channelHash)
	}
	var srcName PeerName
	if err := decoder.Decode(&srcName); err != nil {
		return err
	}
	if err := onok(channel, srcName, payload, decoder); err != nil {
		return err
	}
	return nil
}

func deliverGossipUnicast(channel *GossipChannel, srcName PeerName, origPayload []byte, dec *gob.Decoder) error {
	var destName PeerName
	if err := dec.Decode(&destName); err != nil {
		return err
	}
	if channel.ourself.Name == destName {
		var payload []byte
		if err := dec.Decode(&payload); err != nil {
			return err
		}
		return channel.gossiper.OnGossipUnicast(srcName, payload)
	} else {
		return channel.relayGossipUnicast(destName, origPayload)
	}
}

func deliverGossipBroadcast(channel *GossipChannel, srcName PeerName, origPayload []byte, dec *gob.Decoder) error {
	var payload []byte
	if err := dec.Decode(&payload); err != nil {
		return err
	}
	if err := channel.gossiper.OnGossipBroadcast(payload); err != nil {
		return err
	}
	return channel.relayGossipBroadcast(srcName, origPayload)
}

func deliverGossip(channel *GossipChannel, srcName PeerName, _ []byte, dec *gob.Decoder) error {
	var payload []byte
	if err := dec.Decode(&payload); err != nil {
		return err
	}
	if updateSet, err := channel.data.OnUpdate(payload); err != nil {
		return err
	} else if updateSet != nil {
		channel.SendGossipUpdateFor(updateSet)
	}
	return nil
}

func (c *GossipChannel) SendGossipUpdateFor(keys GossipKeySet) {
	c.Lock()
	defer c.Unlock()
	c.ourself.ForEachConnection(func(_ PeerName, conn Connection) {
		sender, found := c.senders[conn]
		if !found {
			sender = c.makeSender(c.data, conn)
			c.senders[conn] = sender
		}
		sender.Lock()
		for updateKey, _ := range keys {
			sender.pending[updateKey] = true
		}
		sender.Unlock()
		select { // non-blocking send
		case sender.sendChan <- true:
		default:
		}
	})
}

func (c *GossipChannel) gossipMsg(buf []byte) ProtocolMsg {
	return ProtocolMsg{ProtocolGossip, GobEncode(c.hash, c.ourself.Name, buf)}
}

func (c *GossipChannel) GossipUnicast(dstPeerName PeerName, buf []byte) error {
	return c.relayGossipUnicast(dstPeerName, GobEncode(c.hash, c.ourself.Name, dstPeerName, buf))
}

func (c *GossipChannel) GossipBroadcast(buf []byte) error {
	return c.relayGossipBroadcast(c.ourself.Name, GobEncode(c.hash, c.ourself.Name, buf))
}

func (c *GossipChannel) relayGossipUnicast(dstPeerName PeerName, msg []byte) error {
	if relayPeerName, found := c.ourself.Router.Routes.Unicast(dstPeerName); !found {
		c.log("unknown relay destination:", dstPeerName)
	} else if conn, found := c.ourself.ConnectionTo(relayPeerName); !found {
		c.log("unable to find connection to relay peer", relayPeerName)
	} else {
		conn.(ProtocolSender).SendProtocolMsg(ProtocolMsg{ProtocolGossipUnicast, msg})
	}
	return nil
}

func (c *GossipChannel) relayGossipBroadcast(srcName PeerName, msg []byte) error {
	if srcPeer, found := c.ourself.Router.Peers.Fetch(srcName); !found {
		c.log("unable to relay broadcast from unknown peer", srcName)
	} else {
		protocolMsg := ProtocolMsg{ProtocolGossipBroadcast, msg}
		for _, conn := range c.ourself.NextBroadcastHops(srcPeer) {
			conn.SendProtocolMsg(protocolMsg)
		}
	}
	return nil
}

func (c *GossipChannel) log(args ...interface{}) {
	log.Println(append(append([]interface{}{}, "[gossip "+c.name+"]:"), args...)...)
}
