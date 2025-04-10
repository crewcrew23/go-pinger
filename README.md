# Simple Go Pinger

A simple ICMP ping utility written in Go. It allows you to send ICMP Echo requests to a target (IP or domain) and get the response time.

## Features

- Send ICMP Echo requests to a target IP address or domain.
- Customizable number of retries (`-r` flag).
- Prints round-trip time (RTT) in milliseconds.
- Displays packet loss percentage.
- Support for reading target addresses from a file (-f flag)

## Usage

```bash
pinger -t <target> -r <retry_count>
```
## Flags:
- ```-t <target> ```
Required: The target IP address or domain name you want to ping.

- ```-r <retry_count>```
Optional: The number of ping requests to send (default is 4 if not specified).

- ```-f <file>```
  Optional: path to file with ip`s or domains. <br>
  if entered flag -f, flag -t will be ignored <br>
  file should be next format.
```
8.8.8.8
amazon.com
```

## Example
``` ./pinger -t google.com -r <optional>```
```
Received reply from 74.125.205.100 in 43.4899ms
Received reply from 74.125.205.100 in 20.2975ms
Received reply from 74.125.205.100 in 20.4572ms
Received reply from 74.125.205.100 in 21.826ms

══════════════════════════════════════════════════════════════════════
🏓  --- Ping statistics for google.com ---  🏓
══════════════════════════════════════════════════════════════════════
  ▪ 4 packets transmitted, 4 received, 0.00% packet loss
  ▪ Round-Trip Time (RTT):
    • Min RTT:  46.48 ms
    • Avg RTT:  51.74 ms
    • Max RTT:  66.61 ms
══════════════════════════════════════════════════════════════════════
```

``` ./pinger -f ./ips.txt -r <optional>```
```
Received reply from 8.8.8.8 in 23.6496ms
Received reply from 54.239.28.85 in 138.7886ms
Received reply from 8.8.8.8 in 23.3856ms
Received reply from 54.239.28.85 in 126.9047ms
Received reply from 8.8.8.8 in 23.0589ms
Received reply from 54.239.28.85 in 126.3435ms
Received reply from 8.8.8.8 in 23.2269ms
Received reply from 54.239.28.85 in 126.503ms

══════════════════════════════════════════════════════════════════════
🏓  --- Ping statistics for 8.8.8.8 ---  🏓
══════════════════════════════════════════════════════════════════════
  ▪ 4 packets transmitted, 4 received, 0.00% packet loss
  ▪ Round-Trip Time (RTT):
    • Min RTT:  23.06 ms
    • Avg RTT:  23.33 ms
    • Max RTT:  23.65 ms
══════════════════════════════════════════════════════════════════════

══════════════════════════════════════════════════════════════════════
🏓  --- Ping statistics for amazon.com ---  🏓
══════════════════════════════════════════════════════════════════════
  ▪ 4 packets transmitted, 4 received, 0.00% packet loss
  ▪ Round-Trip Time (RTT):
    • Min RTT:  126.34 ms
    • Avg RTT:  129.63 ms
    • Max RTT:  138.79 ms
══════════════════════════════════════════════════════════════════════
```

# Installation
```
git clone https://github.com/crewcrew23/go-pinger.git
cd go-pinger
mkdir bin
make
cd bin
sudo ./pinger -t <target> 
```

After just  run pinger))