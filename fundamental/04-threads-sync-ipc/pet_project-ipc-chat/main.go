package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

const socketPath = "/tmp/ipc-chat.sock"

func main() {
    if len(os.Args) < 2 {
        fmt.Println("usage: ipc-chat server | client <msg>")
        os.Exit(2)
    }
    mode := os.Args[1]
    switch mode {
    case "server":
        _ = os.Remove(socketPath)
        ln, err := net.Listen("unix", socketPath)
        if err != nil {
            fmt.Fprintln(os.Stderr, "listen error:", err)
            os.Exit(1)
        }
        defer ln.Close()
        fmt.Println("server listening on", socketPath)
        for {
            conn, err := ln.Accept()
            if err != nil {
                fmt.Fprintln(os.Stderr, "accept error:", err)
                continue
            }
            go func(c net.Conn) {
                defer c.Close()
                r := bufio.NewScanner(c)
                for r.Scan() {
                    fmt.Println("recv:", r.Text())
                }
            }(conn)
        }
    case "client":
        if len(os.Args) < 3 {
            fmt.Println("usage: ipc-chat client <msg>")
            os.Exit(2)
        }
        c, err := net.Dial("unix", socketPath)
        if err != nil {
            fmt.Fprintln(os.Stderr, "dial error:", err)
            os.Exit(1)
        }
        defer c.Close()
        fmt.Fprintln(c, os.Args[2])
    default:
        fmt.Println("unknown mode:", mode)
        os.Exit(2)
    }
}

