package main

// Importamos las librerías necesarias, aunque con solo guardar se importan automáticamente :D
import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// La primera variable es un mapa donde la clave es en realidad un puntero a un WebSocket, el valor es un booleano.
// La segunda variable es un canal que actuará como una cola de mensajes enviados por los clientes.
var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan Message)           // Broadcast channel

// Este es solo un objeto con métodos para tomar una conexión HTTP normal y actualizarla a un WebSocket
var upgrader = websocket.Upgrader{}

// Message Definiremos un objeto para guardar nuestros mensajes, para interactuar con el servicio ***Gravatar*** que nos proporcionará un avatar único.
type Message struct {
        Email    string `json:"email"`
        Username string `json:"username"`
        Message  string `json:"message"`
}

// Esta función manejará nuestras conexiones WebSocket entrantes
func handleConnections(w http.ResponseWriter, r *http.Request) {

	// El método Upgrade() permite cambiar nuesra solicitud GET inicial a una completa en WebSocket, si hay un error lo mostramos en consola pero no salimos.
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Para cerrar la conexión una vez termina la función
	defer ws.Close()

	// Registramos nuestro nuevo cliente al agregarlo al mapa global de "clients" que fue creado anteriormente.
	clients[ws] = true

	// Bucle infinito que espera continuamente que se escriba  un nuevo mensaje en el WebSocket, lo desserializa de JSON a un objeto Message y luego lo arroja al canal de difusión.
	for {
		var msg Message

		// Read in a new message as JSON and map it to a Message object
		// Si hay un error, registramos ese error y eliminamos ese cliente de nuestro mapa global de clients
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

// Ciclo que lee continuamente desde el canal broadcast y luego transmite el mesaje a todos nuestros clientes a través de su respectiva conexión WebSocket. Si hay un error cerramos la conexión y elminamos del mapa "clients".
func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	// Para que al entrar en la aplicación el cliente tome el index.html
	fs := http.FileServer(http.Dir("../practica3"))
	http.Handle("/", fs)

	// "/ws" es la ruta que nos ayudará a manejar cualquier solicitud para iniciar un WebSocket.
	http.HandleFunc("/ws", handleConnections)

	// Creamos una goroutina que nos ayudará a manejar los mensajes
	go handleMessages()

	// Iniciamos el servidor en localhost en el puerto 8080 y un mensaje que muetre por consola
	log.Println("http server started on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
