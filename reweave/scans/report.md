# Vulnerability Report

```
Report date: 2024-10-16
Unique vulnerability count: 22
Images version: 2.8.9
```

## Scanner Details

```
Application:         grype
Version:             0.82.1
BuildDate:           2024-10-15T13:54:04Z
GitCommit:           50815e59c973cfd0c0247cbc2af00fa37f7cda5d
GitDescription:      v0.82.1
Platform:            linux/amd64
GoVersion:           go1.23.2
Compiler:            gc
Syft Version:        v1.14.1
Supported DB Schema: 5
```

## Vulnerabilities

### weave-kube: (22) 

```
NAME        INSTALLED   FIXED-IN   TYPE       VULNERABILITY   SEVERITY 
bind-libs   9.18.27-r0             apk        CVE-2024-4076   High      
bind-libs   9.18.27-r0             apk        CVE-2024-1975   High      
bind-libs   9.18.27-r0             apk        CVE-2024-1737   High      
bind-libs   9.18.27-r0             apk        CVE-2024-0760   High      
bind-tools  9.18.27-r0             apk        CVE-2024-4076   High      
bind-tools  9.18.27-r0             apk        CVE-2024-1975   High      
bind-tools  9.18.27-r0             apk        CVE-2024-1737   High      
bind-tools  9.18.27-r0             apk        CVE-2024-0760   High      
curl        8.9.0-r0    8.10.0-r0  apk        CVE-2024-8096   Medium    
curl        8.9.0-r0    8.9.1-r0   apk        CVE-2024-7264   Medium    
libcrypto3  3.3.1-r3    3.3.2-r0   apk        CVE-2024-6119   High      
libcurl     8.9.0-r0    8.10.0-r0  apk        CVE-2024-8096   Medium    
libcurl     8.9.0-r0    8.9.1-r0   apk        CVE-2024-7264   Medium    
libssl3     3.3.1-r3    3.3.2-r0   apk        CVE-2024-6119   High      
stdlib      go1.22.2               go-module  CVE-2024-24790  Critical  
stdlib      go1.22.2               go-module  CVE-2024-34158  High      
stdlib      go1.22.2               go-module  CVE-2024-34156  High      
stdlib      go1.22.2               go-module  CVE-2024-24791  High      
stdlib      go1.22.2               go-module  CVE-2024-24789  Medium    
stdlib      go1.22.2               go-module  CVE-2024-24787  Medium    
stdlib      go1.22.2               go-module  CVE-2024-34155  Unknown   
stdlib      go1.22.2               go-module  CVE-2024-24788  Unknown
```

### weave-npc: (10) 

```
NAME        INSTALLED  FIXED-IN  TYPE       VULNERABILITY   SEVERITY 
libcrypto3  3.3.1-r3   3.3.2-r0  apk        CVE-2024-6119   High      
libssl3     3.3.1-r3   3.3.2-r0  apk        CVE-2024-6119   High      
stdlib      go1.22.2             go-module  CVE-2024-24790  Critical  
stdlib      go1.22.2             go-module  CVE-2024-34158  High      
stdlib      go1.22.2             go-module  CVE-2024-34156  High      
stdlib      go1.22.2             go-module  CVE-2024-24791  High      
stdlib      go1.22.2             go-module  CVE-2024-24789  Medium    
stdlib      go1.22.2             go-module  CVE-2024-24787  Medium    
stdlib      go1.22.2             go-module  CVE-2024-34155  Unknown   
stdlib      go1.22.2             go-module  CVE-2024-24788  Unknown
```

### weave: (22) 

```
NAME        INSTALLED   FIXED-IN   TYPE       VULNERABILITY   SEVERITY 
bind-libs   9.18.27-r0             apk        CVE-2024-4076   High      
bind-libs   9.18.27-r0             apk        CVE-2024-1975   High      
bind-libs   9.18.27-r0             apk        CVE-2024-1737   High      
bind-libs   9.18.27-r0             apk        CVE-2024-0760   High      
bind-tools  9.18.27-r0             apk        CVE-2024-4076   High      
bind-tools  9.18.27-r0             apk        CVE-2024-1975   High      
bind-tools  9.18.27-r0             apk        CVE-2024-1737   High      
bind-tools  9.18.27-r0             apk        CVE-2024-0760   High      
curl        8.9.0-r0    8.10.0-r0  apk        CVE-2024-8096   Medium    
curl        8.9.0-r0    8.9.1-r0   apk        CVE-2024-7264   Medium    
libcrypto3  3.3.1-r3    3.3.2-r0   apk        CVE-2024-6119   High      
libcurl     8.9.0-r0    8.10.0-r0  apk        CVE-2024-8096   Medium    
libcurl     8.9.0-r0    8.9.1-r0   apk        CVE-2024-7264   Medium    
libssl3     3.3.1-r3    3.3.2-r0   apk        CVE-2024-6119   High      
stdlib      go1.22.2               go-module  CVE-2024-24790  Critical  
stdlib      go1.22.2               go-module  CVE-2024-34158  High      
stdlib      go1.22.2               go-module  CVE-2024-34156  High      
stdlib      go1.22.2               go-module  CVE-2024-24791  High      
stdlib      go1.22.2               go-module  CVE-2024-24789  Medium    
stdlib      go1.22.2               go-module  CVE-2024-24787  Medium    
stdlib      go1.22.2               go-module  CVE-2024-34155  Unknown   
stdlib      go1.22.2               go-module  CVE-2024-24788  Unknown
```

### weaveexec: (22) 

```
NAME        INSTALLED   FIXED-IN   TYPE       VULNERABILITY   SEVERITY 
bind-libs   9.18.27-r0             apk        CVE-2024-4076   High      
bind-libs   9.18.27-r0             apk        CVE-2024-1975   High      
bind-libs   9.18.27-r0             apk        CVE-2024-1737   High      
bind-libs   9.18.27-r0             apk        CVE-2024-0760   High      
bind-tools  9.18.27-r0             apk        CVE-2024-4076   High      
bind-tools  9.18.27-r0             apk        CVE-2024-1975   High      
bind-tools  9.18.27-r0             apk        CVE-2024-1737   High      
bind-tools  9.18.27-r0             apk        CVE-2024-0760   High      
curl        8.9.0-r0    8.10.0-r0  apk        CVE-2024-8096   Medium    
curl        8.9.0-r0    8.9.1-r0   apk        CVE-2024-7264   Medium    
libcrypto3  3.3.1-r3    3.3.2-r0   apk        CVE-2024-6119   High      
libcurl     8.9.0-r0    8.10.0-r0  apk        CVE-2024-8096   Medium    
libcurl     8.9.0-r0    8.9.1-r0   apk        CVE-2024-7264   Medium    
libssl3     3.3.1-r3    3.3.2-r0   apk        CVE-2024-6119   High      
stdlib      go1.22.2               go-module  CVE-2024-24790  Critical  
stdlib      go1.22.2               go-module  CVE-2024-34158  High      
stdlib      go1.22.2               go-module  CVE-2024-34156  High      
stdlib      go1.22.2               go-module  CVE-2024-24791  High      
stdlib      go1.22.2               go-module  CVE-2024-24789  Medium    
stdlib      go1.22.2               go-module  CVE-2024-24787  Medium    
stdlib      go1.22.2               go-module  CVE-2024-34155  Unknown   
stdlib      go1.22.2               go-module  CVE-2024-24788  Unknown
```

