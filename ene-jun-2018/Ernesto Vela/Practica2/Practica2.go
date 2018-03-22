package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
    )

func PostHandler(w http.ResponseWriter, r *http.Request) {
    	if r.Method == "POST" {
    		body, err := ioutil.ReadAll(r.Body)
    		if err != nil {
    			http.Error(w, "Error leyendo el body request",
    				http.StatusInternalServerError)
    		}

        fmt.Fprintln(w, "Numero de caracteres:", len(body))
    	} else {
    		http.Error(w, "ERROR", http.StatusMethodNotAllowed)
    	}
    }

    func main(){
      http.HandleFunc("/post", PostHandler)
      fmt.Println("El servidor corriendo")
      log.Fatal(http.ListenAndServe(":8081", nil))
}
