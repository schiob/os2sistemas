package main

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	fmt.Print(r.Body)
	fmt.Fprint(w, "¡Que onda, escribe algo%s!", html.EscapeString(r.URL.Path))
	buf := &bytes.Buffer{}
	nRead, err := io.Copy(buf, r.Body)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Cantidad de caracteres", strconv.FormatInt(nRead, 10))
	fmt.Fprintln(w, "El texto: ", r.Body)
}

func main() {
	http.HandleFunc("/1", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
