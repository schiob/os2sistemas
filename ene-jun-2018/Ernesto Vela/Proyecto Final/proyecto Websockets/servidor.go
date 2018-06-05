package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	//"io/ioutil"
	//"encoding/json"
	"github.com/gorilla/websocket"
)

var cont int = 0

type Respuesta struct {
	Respuesta   string `json:"respuesta"`
	numaciertos string `json:"respuesta"`
}
type Jugadores struct {
	idjugador string
	conec     *websocket.Conn
}

var people []Jugadores
var respcorrectas = [5]string{"Montevideo", "Ballena Azul", "1810", "Tierra", "CDMX"}
var aciertosjugador1 int = 0
var aciertosjugador2 int = 0

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var conexiones = make(map[*websocket.Conn]bool) //Clientes conectados
var evaluador = make(chan Respuesta)            //Para evaluar respuestas

func validarvictoria() {
	if aciertosjugador1 > aciertosjugador2 {
		var g string = "GANASTE!"
		var p string = "PERDISTE!"
		people[0].conec.WriteJSON(g)
		people[1].conec.WriteJSON(p)
	} else if aciertosjugador2 > aciertosjugador1 {
		var g string = "GANASTE!"
		var p string = "PERDISTE!"
		people[1].conec.WriteJSON(g)
		people[0].conec.WriteJSON(p)
	} else {
		var e string = "EMPATE!"
		people[1].conec.WriteJSON(e)
		people[0].conec.WriteJSON(e)
	}
}

func validarcliente() {
	//Tomar el mensaje del canal y almacenarlo
	res := <-evaluador
	//Si la conexion tiene un idjugador 1 sumas el numero de aciertos a la variable
	//Se asigna a una variable declarada la conexion para mayor comodidad
	//Se convierte a int el String de numaciertos y se guarda en variable
	if people[cont-1].idjugador == "1" {
		conex := people[cont-1].conec(*websocket.Conn) //Declarado bien?
		i := strconv.Atoi(res.numaciertos)
		aciertosjugador1 = aciertosjugador1 + i
		//Si la conexion tiene un idjugador 2 sumas el numero de aciertos a la variable
	} else if people[cont-1].idjugador == "2" {
		conex := people[cont-1].conec() //Declarado bien?
		i := strconv.Atoi(res.numaciertos)
		aciertosjugador2 = aciertosjugador2 + i
	}
	//recorrer el arreglo de respuestas correctas hasta que se encuentre el mismo valor
	//Creamos una variable string y la enviamos según su conexion escribiendo un JSON
	for index, item := range respcorrectas {
		if item == res.Respuesta {
			var msg string = "Respuesta Correcta!"
			//json.NewEncoder(w).Encode(resp)
			err := conex.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				conex.Close()
				delete(conexiones, conex)
			}
			//if people[cont-1].idjugador == "1" {
			//aciertosjugador1 = aciertosjugador1 + 1
			//} else if people[cont-1].idjugador == "2" {
			//aciertosjugador2 = aciertosjugador2 + 1
		}
		return
	}

	//var msg string = "Respuesta Incorrecta!"
	var msg string = "Respuesta Incorrecta!"
	err := conex.WriteJSON(msg)
	if err != nil {
		log.Printf("error: %v", err)
		conex.Close()
		delete(conexiones, conex)
	}
}

//funcion main
func main() {
	fmt.Println(respcorrectas)
	//Servir el index.html
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.Handle("/imagenes/", http.StripPrefix("/imagenes/", http.FileServer(http.Dir("./imagenes"))))
	//http.HandleFunc("/validar", validar)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		cont = cont + 1
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		// agregar el jugador a conexiones
		conexiones[conn] = true
		people = append(people, Jugador{idjugador: string(cont), conec: conn})
		//Loop infinito que espera recibir un Json con la estructura Respuesta
		for {
			var resp Respuesta
			// Leer mensaje del navegador y almacenar en resp
			err := people[cont-1].conec.ReadJSON(&resp)
			//err := conn.ReadJSON(&resp)
			if err != nil {
				log.Printf("error: %v", err)
				//elimina la conexion si hay error
				delete(conexiones, conn)
				break
			}
			//Enviar el nuevo mensaje al canal de evaluacion
			evaluador <- resp
		}
	})
	go validarcliente()
	go validarvictoria()
	http.ListenAndServe(":8081", nil)
}
