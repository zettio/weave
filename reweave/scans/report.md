# Vulnerability Report

```
Report date: 2024-07-30
Unique vulnerability count: 18
Images version: 2.8.8
```

## Scanner Details

```
Application:         grype
Version:             0.79.3
BuildDate:           2024-07-15T12:57:46Z
GitCommit:           45b7236e948ef973a8a6ffbac52dff28be0fd70e
GitDescription:      v0.79.3
Platform:            linux/amd64
GoVersion:           go1.22.5
Compiler:            gc
Syft Version:        v1.9.0
Supported DB Schema: 5
```

## Vulnerabilities

### weave-kube: (18) 

```
NAME        INSTALLED   FIXED-IN  TYPE       VULNERABILITY   SEVERITY 
bind-libs   9.18.27-r0            apk        CVE-2024-4076   High      
bind-libs   9.18.27-r0            apk        CVE-2024-1975   High      
bind-libs   9.18.27-r0            apk        CVE-2024-1737   High      
bind-libs   9.18.27-r0            apk        CVE-2024-0760   High      
bind-tools  9.18.27-r0            apk        CVE-2024-4076   High      
bind-tools  9.18.27-r0            apk        CVE-2024-1975   High      
bind-tools  9.18.27-r0            apk        CVE-2024-1737   High      
bind-tools  9.18.27-r0            apk        CVE-2024-0760   High      
curl        8.7.1-r0    8.9.0-r0  apk        CVE-2024-6874   Unknown   
curl        8.7.1-r0    8.9.0-r0  apk        CVE-2024-6197   Unknown   
libcrypto3  3.3.1-r0    3.3.1-r1  apk        CVE-2024-5535   Critical  
libcurl     8.7.1-r0    8.9.0-r0  apk        CVE-2024-6874   Unknown   
libcurl     8.7.1-r0    8.9.0-r0  apk        CVE-2024-6197   Unknown   
libssl3     3.3.1-r0    3.3.1-r1  apk        CVE-2024-5535   Critical  
stdlib      go1.21.9              go-module  CVE-2024-24790  Critical  
stdlib      go1.21.9              go-module  CVE-2024-24791  High      
stdlib      go1.21.9              go-module  CVE-2024-24789  Medium    
stdlib      go1.21.9              go-module  CVE-2024-24787  Medium
```

### weave-npc: (6) 

```
NAME        INSTALLED  FIXED-IN  TYPE       VULNERABILITY   SEVERITY 
libcrypto3  3.3.1-r0   3.3.1-r1  apk        CVE-2024-5535   Critical  
libssl3     3.3.1-r0   3.3.1-r1  apk        CVE-2024-5535   Critical  
stdlib      go1.21.9             go-module  CVE-2024-24790  Critical  
stdlib      go1.21.9             go-module  CVE-2024-24791  High      
stdlib      go1.21.9             go-module  CVE-2024-24789  Medium    
stdlib      go1.21.9             go-module  CVE-2024-24787  Medium
```

### weave: (18) 

```
NAME        INSTALLED   FIXED-IN  TYPE       VULNERABILITY   SEVERITY 
bind-libs   9.18.27-r0            apk        CVE-2024-4076   High      
bind-libs   9.18.27-r0            apk        CVE-2024-1975   High      
bind-libs   9.18.27-r0            apk        CVE-2024-1737   High      
bind-libs   9.18.27-r0            apk        CVE-2024-0760   High      
bind-tools  9.18.27-r0            apk        CVE-2024-4076   High      
bind-tools  9.18.27-r0            apk        CVE-2024-1975   High      
bind-tools  9.18.27-r0            apk        CVE-2024-1737   High      
bind-tools  9.18.27-r0            apk        CVE-2024-0760   High      
curl        8.7.1-r0    8.9.0-r0  apk        CVE-2024-6874   Unknown   
curl        8.7.1-r0    8.9.0-r0  apk        CVE-2024-6197   Unknown   
libcrypto3  3.3.1-r0    3.3.1-r1  apk        CVE-2024-5535   Critical  
libcurl     8.7.1-r0    8.9.0-r0  apk        CVE-2024-6874   Unknown   
libcurl     8.7.1-r0    8.9.0-r0  apk        CVE-2024-6197   Unknown   
libssl3     3.3.1-r0    3.3.1-r1  apk        CVE-2024-5535   Critical  
stdlib      go1.21.9              go-module  CVE-2024-24790  Critical  
stdlib      go1.21.9              go-module  CVE-2024-24791  High      
stdlib      go1.21.9              go-module  CVE-2024-24789  Medium    
stdlib      go1.21.9              go-module  CVE-2024-24787  Medium
```

### weaveexec: (18) 

```
NAME        INSTALLED   FIXED-IN  TYPE       VULNERABILITY   SEVERITY 
bind-libs   9.18.27-r0            apk        CVE-2024-4076   High      
bind-libs   9.18.27-r0            apk        CVE-2024-1975   High      
bind-libs   9.18.27-r0            apk        CVE-2024-1737   High      
bind-libs   9.18.27-r0            apk        CVE-2024-0760   High      
bind-tools  9.18.27-r0            apk        CVE-2024-4076   High      
bind-tools  9.18.27-r0            apk        CVE-2024-1975   High      
bind-tools  9.18.27-r0            apk        CVE-2024-1737   High      
bind-tools  9.18.27-r0            apk        CVE-2024-0760   High      
curl        8.7.1-r0    8.9.0-r0  apk        CVE-2024-6874   Unknown   
curl        8.7.1-r0    8.9.0-r0  apk        CVE-2024-6197   Unknown   
libcrypto3  3.3.1-r0    3.3.1-r1  apk        CVE-2024-5535   Critical  
libcurl     8.7.1-r0    8.9.0-r0  apk        CVE-2024-6874   Unknown   
libcurl     8.7.1-r0    8.9.0-r0  apk        CVE-2024-6197   Unknown   
libssl3     3.3.1-r0    3.3.1-r1  apk        CVE-2024-5535   Critical  
stdlib      go1.21.9              go-module  CVE-2024-24790  Critical  
stdlib      go1.21.9              go-module  CVE-2024-24791  High      
stdlib      go1.21.9              go-module  CVE-2024-24789  Medium    
stdlib      go1.21.9              go-module  CVE-2024-24787  Medium
```

