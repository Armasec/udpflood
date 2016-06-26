# UDPFlood
A simple UDP flooder written in Golang. **Use this for educational purposes only.**

## Usage
```bash
./udpflood -h google.com -p 443 --threads 1 --size 65507
```

## Build
```bash
git clone https://github.com/cameronpoe/udpflood
cd udpflood
export GOPATH=`pwd`
go get ./...
go build
```
