package main

import (
	"fmt"
	"net/http"
    "io/ioutil"
    "encoding/json"
)

type Persona struct {
    ID        string   `json:"id,omitempty"`
    Email string   `json:"email,omitempty"`
    Nombre  string   `json:"nombre,omitempty"`
}

var cont int = 0
var jugadores []Persona
var respcorrectas = [5]string {"Montevideo", "Ballena Azul", "1810", "Tierra", "CDMX"}
var respuesta string


func validar(w http.ResponseWriter, r *http.Request) {
    //Convertir el body del request a string y asignarlo a variable
    if r.Method == "POST" {
    		body, err := ioutil.ReadAll(r.Body)
    		if err != nil {
    			http.Error(w, "Error reading request body", http.StatusInternalServerError)
    		}
    		respuesta = string(body)
    	} else {
    		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    	}
    //Validar respuesta correcta
    for _, item := range respcorrectas {
        if item == respuesta {
            //fmt.Fprintf(w, "Respuesta Correcta!")
            var resp string = "Respuesta Correcta!"
            cont = cont + 1
            json.NewEncoder(w).Encode(resp)
            return
        }
    }
    //fmt.Fprintf(w, "Respuesta Incorrecta")
    resp := "Respuesta Incorrecta"
    json.NewEncoder(w).Encode(resp)
      
}

func main() {
    fmt.Println(respcorrectas)
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.Handle("/imagenes/", http.StripPrefix("/imagenes/", http.FileServer(http.Dir("./imagenes"))))
    http.HandleFunc("/validar", validar)
    http.ListenAndServe(":8081", nil)
}
