package main

import (
	"fmt"
	"net/http"
    "io/ioutil"
    "encoding/json"
)



var respcorrectas = [5]string {"Policia", "Bombero", "Matematicas", "Amarillo", "CDMX"}
var respuesta string
var aciertosjugador1 int
var aciertosjugador2 int

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Respuesta struct {
        Idcliente    string `json:"idcliente"`
        Respuesta    string `json:"respuesta"`
}

var conexiones = make(map[*websocket.Conn]bool) //Clientes conectados
var evaluador = make(chan Respuesta) //Para evaluar respuestas


func validarcliente() {
    resp := <-evaluador
    if resp.Idcliente == "cliente1" {
            client := 0//Declarado bien?
    } else if resp.Idcliente == "cliente2" {
            client := 1//Declarado bien?
            }

    for index, item := range respcorrectas {
            if item == resp.Respuesta {
            var msg string = "Respuesta Correcta!"
            //json.NewEncoder(w).Encode(resp)
            err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(conexiones, client)
            }
            if client == 0 {
            aciertosjugador1 = aciertosjugador1 + 1
            } else if client == 1 {
            aciertosjugador2 = aciertosjugador2 + 1
            }
            return
            }
  } 
   //var msg string = "Respuesta Incorrecta!"
    msg := "Respuesta Incorrecta!"
   err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(conexiones, client)
}

func main() {
    fmt.Println(respcorrectas)
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.Handle("/imagenes/", http.StripPrefix("/imagenes/", http.FileServer(http.Dir("./imagenes"))))
    http.HandleFunc("/validar", validar)
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
        conexiones[conn] = true

		for {
            var resp Respuesta
			// Read message from browser
			err := conn.ReadJSON(&resp)
		if err != nil {
			log.Printf("error: %v", err)
			delete(conexiones, conn)
			break
		              }
            evaluador <- resp
            //SI VA ESTA LLAVE?
}	})
    go validarcliente()
    http.ListenAndServe(":8081", nil)
}
