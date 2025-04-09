# Simple Go Pinger

A simple ICMP ping utility written in Go. It allows you to send ICMP Echo requests to a target (IP or domain) and get the response time.

## Features

- Send ICMP Echo requests to a target IP address or domain.
- Customizable number of retries (`-r` flag).
- Prints round-trip time (RTT) in milliseconds.
- Displays packet loss percentage.

## Usage

```bash
pinger -t <target> -r <retry_count>
```
## Flags:
- ```-t <target> ```
Required: The target IP address or domain name you want to ping.

- ```-r <retry_count>```
Optional: The number of ping requests to send (default is 4 if not specified).

## Example
``` pinger -t google.com ```
```
Received reply from 74.125.205.100 in 43.4899ms
Received reply from 74.125.205.100 in 20.2975ms
Received reply from 74.125.205.100 in 20.4572ms
Received reply from 74.125.205.100 in 21.826ms

--- google.com ping statistics ---
4 packets transmitted, 4 received, 0.00% packet loss
rtt min/avg/max = 20.30/26.52/43.49 ms
```

# Installation
```git clone https://github.com/crewcrew23/go-pinger.git``` <br>
```cd go-pinger``` <br>
```make``` or ``` go build ./cmd/pinger ``` <br>

After just  run pinger))