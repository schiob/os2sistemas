package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "POST!= %v\n",r.PostForm)
		text := r.FormValue("texto")
		can := (text)
		fmt.Fprintf(w, "Texto = %s\n", text)
		fmt.Fprintf(w, "NUmero de caracteres = ", len(can))

}

func main() {
	http.HandleFunc("/", index)
	fmt.Printf("SERVIDOR ACTIVO")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
