package main

import (
        "fmt"
        "log"
        "net"
)

func main() {
        conn, err := net.Dial("tcp", "localhost:8081")
        if err != nil {
                log.Fatalln(err)
        }

        _, err = conn.Write([]byte("Hello Server!"))
        if err != nil {
                log.Fatalln(err)
        }

        _, err = conn.Write([]byte("Como estas?"))
        if err != nil {
                log.Fatalln(err)
        }

        for {
                buffer := make([]byte, 1400)
                dataSize, err := conn.Read(buffer)
                if err != nil {
                        fmt.Println("connection closed")
                        return
                }

        data := buffer[:dataSize]
                fmt.Println("received message: ", string(data))
        }
}
