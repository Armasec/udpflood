# UDPFlood
A simple UDP flooder written in Golang. **Use this for educational purposes only.**

## Usage
```bash
./udpflood -h google.com -p 443 -t 1 -s 65507
```

## Build
```bash
git clone https://github.com/Armasec/udpflood
cd udpflood
export GOPATH=`pwd`
go get ./...
go build
```
