package main 
import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os")

	const direccion= "127.0.3000"
	const buferSize =256
	const endLIne =10

	var nick string
	var in *bufio.Reader

	fun man(){
		in =bufio.NewReader(os.Stdin)
	
	for nick== " "{
		fmt.Printf("Dame tu nick:")
		buf, _, _:=  in.ReadLine()
		nick=string (buf)
	}

	var conn net.Coon
	var err error

	for { 
		fmt.Printf("Conectando a %s...ln", direccion)
	conn err = net.Dial("tcp", direccion)
	if err ==nil {break
		}
	}

	defer conn.close()

	go reciveMessages(conn)
	handleConnection(conn)
	}
	func handleConnection(conn net.Conn){
		for{
			buf, _, _:=in.ReadLine()
			if len(buf)>0{
				conn.Write(append([]byte(nick + "->"),append(buf, endLine)..))
			}
		}
	}
	func reciveMessages(conn net.Conn){
		var data []byte
		buffer := make([]byte, bufferSize)
		for{
			for{
				n, err :=conn.Read(buffer)
				if err !=nil{
					if err ==io.EOF {break}
				}
			}
			fmt.Printf("/.s\n",data[:len(data)-1])
			data=make([]byte, 0)
		}
	}

