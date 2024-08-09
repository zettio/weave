# Vulnerability Report

```
Report date: 2024-08-09
Unique vulnerability count: 20
Images version: 2.8.8
```

## Scanner Details

```
Application:         grype
Version:             0.79.4
BuildDate:           2024-07-31T15:24:34Z
GitCommit:           0cf393938995de1bf1284a7d156a0ce97816a396
GitDescription:      v0.79.4
Platform:            linux/amd64
GoVersion:           go1.22.5
Compiler:            gc
Syft Version:        v1.10.0
Supported DB Schema: 5
```

## Vulnerabilities

### weave-kube: (20) 

```
NAME                      INSTALLED             FIXED-IN  TYPE       VULNERABILITY        SEVERITY 
bind-libs                 9.18.27-r0                      apk        CVE-2024-4076        High      
bind-libs                 9.18.27-r0                      apk        CVE-2024-1975        High      
bind-libs                 9.18.27-r0                      apk        CVE-2024-1737        High      
bind-libs                 9.18.27-r0                      apk        CVE-2024-0760        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-4076        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-1975        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-1737        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-0760        High      
curl                      8.7.1-r0              8.9.0-r0  apk        CVE-2024-6197        High      
curl                      8.7.1-r0              8.9.0-r0  apk        CVE-2024-6874        Low       
curl                      8.7.1-r0                        apk        CVE-2024-7264        Unknown   
github.com/docker/docker  v24.0.9+incompatible  25.0.6    go-module  GHSA-v23v-6jw2-98fq  Critical  
libcrypto3                3.3.1-r0              3.3.1-r1  apk        CVE-2024-5535        Critical  
libcurl                   8.7.1-r0              8.9.0-r0  apk        CVE-2024-6197        High      
libcurl                   8.7.1-r0              8.9.0-r0  apk        CVE-2024-6874        Low       
libssl3                   3.3.1-r0              3.3.1-r1  apk        CVE-2024-5535        Critical  
stdlib                    go1.21.9                        go-module  CVE-2024-24790       Critical  
stdlib                    go1.21.9                        go-module  CVE-2024-24791       High      
stdlib                    go1.21.9                        go-module  CVE-2024-24789       Medium    
stdlib                    go1.21.9                        go-module  CVE-2024-24787       Medium
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

### weave: (20) 

```
NAME                      INSTALLED             FIXED-IN  TYPE       VULNERABILITY        SEVERITY 
bind-libs                 9.18.27-r0                      apk        CVE-2024-4076        High      
bind-libs                 9.18.27-r0                      apk        CVE-2024-1975        High      
bind-libs                 9.18.27-r0                      apk        CVE-2024-1737        High      
bind-libs                 9.18.27-r0                      apk        CVE-2024-0760        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-4076        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-1975        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-1737        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-0760        High      
curl                      8.7.1-r0              8.9.0-r0  apk        CVE-2024-6197        High      
curl                      8.7.1-r0              8.9.0-r0  apk        CVE-2024-6874        Low       
curl                      8.7.1-r0                        apk        CVE-2024-7264        Unknown   
github.com/docker/docker  v24.0.9+incompatible  25.0.6    go-module  GHSA-v23v-6jw2-98fq  Critical  
libcrypto3                3.3.1-r0              3.3.1-r1  apk        CVE-2024-5535        Critical  
libcurl                   8.7.1-r0              8.9.0-r0  apk        CVE-2024-6197        High      
libcurl                   8.7.1-r0              8.9.0-r0  apk        CVE-2024-6874        Low       
libssl3                   3.3.1-r0              3.3.1-r1  apk        CVE-2024-5535        Critical  
stdlib                    go1.21.9                        go-module  CVE-2024-24790       Critical  
stdlib                    go1.21.9                        go-module  CVE-2024-24791       High      
stdlib                    go1.21.9                        go-module  CVE-2024-24789       Medium    
stdlib                    go1.21.9                        go-module  CVE-2024-24787       Medium
```

### weaveexec: (20) 

```
NAME                      INSTALLED             FIXED-IN  TYPE       VULNERABILITY        SEVERITY 
bind-libs                 9.18.27-r0                      apk        CVE-2024-4076        High      
bind-libs                 9.18.27-r0                      apk        CVE-2024-1975        High      
bind-libs                 9.18.27-r0                      apk        CVE-2024-1737        High      
bind-libs                 9.18.27-r0                      apk        CVE-2024-0760        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-4076        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-1975        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-1737        High      
bind-tools                9.18.27-r0                      apk        CVE-2024-0760        High      
curl                      8.7.1-r0              8.9.0-r0  apk        CVE-2024-6197        High      
curl                      8.7.1-r0              8.9.0-r0  apk        CVE-2024-6874        Low       
curl                      8.7.1-r0                        apk        CVE-2024-7264        Unknown   
github.com/docker/docker  v24.0.9+incompatible  25.0.6    go-module  GHSA-v23v-6jw2-98fq  Critical  
libcrypto3                3.3.1-r0              3.3.1-r1  apk        CVE-2024-5535        Critical  
libcurl                   8.7.1-r0              8.9.0-r0  apk        CVE-2024-6197        High      
libcurl                   8.7.1-r0              8.9.0-r0  apk        CVE-2024-6874        Low       
libssl3                   3.3.1-r0              3.3.1-r1  apk        CVE-2024-5535        Critical  
stdlib                    go1.21.9                        go-module  CVE-2024-24790       Critical  
stdlib                    go1.21.9                        go-module  CVE-2024-24791       High      
stdlib                    go1.21.9                        go-module  CVE-2024-24789       Medium    
stdlib                    go1.21.9                        go-module  CVE-2024-24787       Medium
```

