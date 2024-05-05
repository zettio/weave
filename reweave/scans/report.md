# Vulnerability Report

```
Report date: 2024-05-05
Unique vulnerability count: 18
Images version: 2.8.7
```

## Scanner Details

```
Application:         grype
Version:             0.77.2
BuildDate:           2024-05-01T16:20:45Z
GitCommit:           bd16101ad0ed30c38e95d0992d0ad53f709dc5df
GitDescription:      v0.77.2
Platform:            linux/amd64
GoVersion:           go1.21.9
Compiler:            gc
Syft Version:        v1.3.0
Supported DB Schema: 5
```

## Vulnerabilities

### weave-kube: (18) 

```
NAME           INSTALLED   FIXED-IN  TYPE  VULNERABILITY   SEVERITY 
busybox        1.36.1-r15            apk   CVE-2023-42366  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42365  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42364  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42363  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42366  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42365  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42364  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42363  Medium    
curl           8.5.0-r0              apk   CVE-2024-0853   Medium    
curl           8.5.0-r0              apk   CVE-2024-2466   Unknown   
curl           8.5.0-r0              apk   CVE-2024-2398   Unknown   
curl           8.5.0-r0              apk   CVE-2024-2004   Unknown   
libuv          1.47.0-r0             apk   CVE-2024-24806  High      
nghttp2-libs   1.58.0-r0             apk   CVE-2024-28182  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42366  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42365  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42364  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42363  Medium
```

### weave-npc: (12) 

```
NAME           INSTALLED   FIXED-IN  TYPE  VULNERABILITY   SEVERITY 
busybox        1.36.1-r15            apk   CVE-2023-42366  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42365  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42364  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42363  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42366  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42365  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42364  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42363  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42366  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42365  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42364  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42363  Medium
```

### weave: (18) 

```
NAME           INSTALLED   FIXED-IN  TYPE  VULNERABILITY   SEVERITY 
busybox        1.36.1-r15            apk   CVE-2023-42366  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42365  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42364  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42363  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42366  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42365  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42364  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42363  Medium    
curl           8.5.0-r0              apk   CVE-2024-0853   Medium    
curl           8.5.0-r0              apk   CVE-2024-2466   Unknown   
curl           8.5.0-r0              apk   CVE-2024-2398   Unknown   
curl           8.5.0-r0              apk   CVE-2024-2004   Unknown   
libuv          1.47.0-r0             apk   CVE-2024-24806  High      
nghttp2-libs   1.58.0-r0             apk   CVE-2024-28182  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42366  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42365  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42364  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42363  Medium
```

### weaveexec: (18) 

```
NAME           INSTALLED   FIXED-IN  TYPE  VULNERABILITY   SEVERITY 
busybox        1.36.1-r15            apk   CVE-2023-42366  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42365  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42364  Medium    
busybox        1.36.1-r15            apk   CVE-2023-42363  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42366  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42365  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42364  Medium    
busybox-binsh  1.36.1-r15            apk   CVE-2023-42363  Medium    
curl           8.5.0-r0              apk   CVE-2024-0853   Medium    
curl           8.5.0-r0              apk   CVE-2024-2466   Unknown   
curl           8.5.0-r0              apk   CVE-2024-2398   Unknown   
curl           8.5.0-r0              apk   CVE-2024-2004   Unknown   
libuv          1.47.0-r0             apk   CVE-2024-24806  High      
nghttp2-libs   1.58.0-r0             apk   CVE-2024-28182  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42366  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42365  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42364  Medium    
ssl_client     1.36.1-r15            apk   CVE-2023-42363  Medium
```

