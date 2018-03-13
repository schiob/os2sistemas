package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var eleccion string

func index_handler(w http.ResponseWriter, r *http.Request) {
    // HTML para visualizar el parrot
    fmt.Fprintf(w, "<h1>Elegiste este bonito Parrot!</h1>")
    fmt.Fprintf(w, "<title>Party Parrot Or DIE</title>")
    fmt.Fprintf(w, "<img src='imagenes/parrotdeseado.gif' alt='Está bonito'>")
}
func leer() {
      leerparrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/"+eleccion)
		if err != nil {
			fmt.Printf("Request http fallido! \n", err)
		} else {
			datosleidos, _ := ioutil.ReadAll(leerparrot.Body)
			crearparrot, err := os.Create("imagenes/parrotdeseado.gif")
			if err != nil {
				fmt.Printf("Error creando al parrot! \n", err)
			}
			defer crearparrot.Close()

			_, err = crearparrot.Write(datosleidos)
			if err != nil {
				fmt.Printf("Error escribiendo los datos! \n", err)
			}
		}
		fmt.Println("El gif del parrot que seleccionaste ha sido creado!")
        fmt.Println("Accede a localhost:8081 para visualizar :)")
        }

func main() {
	opciones, err := http.Get("https://ppaas.herokuapp.com/baseparrots")
	if err != nil {
		fmt.Printf("Request http fallido! %s\n", err)
	} else {
		datosleidos, _ := ioutil.ReadAll(opciones.Body)
		fmt.Println("Puedes elegir entre todos estos parrots: \n", string(datosleidos))
	}

	fmt.Println("Elige un parrot: ")
	fmt.Scanf("%s", &eleccion)

    leer()

    http.HandleFunc("/", index_handler)
    http.Handle("/imagenes/", http.StripPrefix("/imagenes/", http.FileServer(http.Dir("./imagenes"))))
    http.ListenAndServe(":8081", nil)
}
