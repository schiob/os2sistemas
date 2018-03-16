package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "holi")

	fmt.Fprint(w, len([]rune("holi")))
}

func main() {
	http.HandleFunc("/", index)

	fmt.Println("El servidor esta a la escucha en el puerto 8085")
	http.ListenAndServe(":8085", nil)
}
