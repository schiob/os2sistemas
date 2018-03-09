
package main
	
	import( 
	"bytes"
	"io"
	"log"
	"net"
	"os"
	)

	const direccion = "127.0.3000"
	const bufferSize =256
	const endline= 10

	var cliente []net.Conn

	fun main(){
		cliente= make([].net.Conn, 0)
		listener, err :=net.Listen("top", direccion)
		if err != nil {
			log.Fatal("al escucha fallido" + direccion)
			os.Exit(1)
		}
		for {
			conn, _:= listener.Accept()
		cliente= append(clientes, conn)
		go handleconection(conn)
	}

    }     
   func handleconection(conn net.Conn){
	defer conn.close()
	var data []byte
	buffer :=make([]byte, bufferSize)
   
    for{
    	for{
    	n, err:=conn.Read(buffer)	
		if err != nil {
			if err== io.EOF { 
		break
	}
       }
       buffer=bytes.Trim(buffer[:n], "\x00")
       data=append(data, buffer...)
       if data[len(data)-1]==endline {
       break
   		}
   }

       SendToOthercliente(conn, data)
       data= make([]byte, 0)
   }
}

func SendToOtherClients(sender net.Conn, data[]byte){
	for i :=0; i<len(cliente);i++{
		if cliente [i] != sender {
			cliente [i].write(data)
		}
	}
}

