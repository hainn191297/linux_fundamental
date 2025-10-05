package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "net"
)

func main() {
    addr := flag.String("addr", ":8080", "listen address")
    flag.Parse()

    ln, err := net.Listen("tcp", *addr)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("listening on", *addr)
    for {
        c, err := ln.Accept()
        if err != nil {
            log.Println("accept:", err)
            continue
        }
        go handle(c)
    }
}

func handle(c net.Conn) {
    defer c.Close()
    log.Println("conn from", c.RemoteAddr())
    r := bufio.NewReader(c)
    w := bufio.NewWriter(c)
    for {
        line, err := r.ReadString('\n')
        if err != nil {
            return
        }
        if _, err := w.WriteString(line); err != nil {
            return
        }
        if err := w.Flush(); err != nil {
            return
        }
    }
}

