package main

import (
	"bytes"
	"io"
	"log"
	"net"
	"os"
)

const addr = "127.0.0.1:3000"
const bufferSize = 256
const endLine = 10

var clientes []net.Conn

func main() {
	clientes = make([]net.Conn, 0)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Can't listen to " + addr)
		os.Exit(1)
	}

	for {
		conn, _ := listener.Accept()
		clientes = append(clientes, conn)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var data []byte
	buffer := make([]byte, bufferSize)

	for {
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				}
			}
			buffer = bytes.Trim(buffer[:n], "\x00")
			data = append(data, buffer...)
			if data[len(data)-1] == endLine {
				break
			}
		}
		sendToOtherClients(conn, data)
		data = make([]byte, 0)
	}
}

func sendToOtherClients(sender net.Conn, data []byte) {
	for i := 0; i < len(clientes); i++ {
		if clientes[i] != sender {
			clientes[i].Write(data)
		}
	}
}
