package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
    )


var texto string
func PostHandler(w http.ResponseWriter, r *http.Request) {
    	if r.Method == "POST" {
    		leerbody, err := ioutil.ReadAll(r.Body)
    		if err != nil {
    			http.Error(w, "Error al leer el Body del Request",
    				http.StatusInternalServerError)
    		}
    		texto = string(leerbody)
    		fmt.Fprintln(w, "Se hizo un POST con: " +texto)
        fmt.Fprintln(w, texto + " tiene una longitud de", len(leerbody), "caracteres")
    	} else {
    		http.Error(w, "Error en el Request", http.StatusMethodNotAllowed)
    	}
    }

    func main(){
      http.HandleFunc("/post", PostHandler)
      fmt.Println("Servidor en ejecucion, haz un POST")
      log.Fatal(http.ListenAndServe(":8080", nil))
}
