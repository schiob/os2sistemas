package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var eleccion string

func index_handler(w http.ResponseWriter, r *http.Request) {
    // MAIN SECTION HTML CODE
    fmt.Fprintf(w, "<h1>Elegiste este bonito gif!</h1>")
    fmt.Fprintf(w, "<title>Party Parrot Or DIE</title>")
    fmt.Fprintf(w, "<img src='imagenes/parrotdeseado.gif' alt='Está bonito'>")
}

func main() {
	opciones, err := http.Get("https://ppaas.herokuapp.com/baseparrots")
	if err != nil {
		fmt.Printf("Request http fallido! %s\n", err)
	} else {
		datosleidos, _ := ioutil.ReadAll(opciones.Body)
		fmt.Println("Opciones de parrot: ", string(datosleidos))
	}

	fmt.Println("Elige una de las opciones anteriores: ")
	fmt.Scanf("%s", &eleccion)

	if eleccion == "boredparrot" {
		leerbored, err := http.Get("https://ppaas.herokuapp.com/partyparrot/boredparrot")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(leerbored.Body)
			crearbored, err := os.Create("imagenes/parrotdeseado.gif")
			if err != nil {
				log.Fatal(err)
			}
			defer crearbored.Close()

			_, err = crearbored.Write(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El archivo con los datos del gif ha sido creado")
	} //endifboredparrot

	if eleccion == "congaparrot" {
		response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/congaparrot")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			crearconga, err := os.Create("congaparrot.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer crearconga.Close()

			_, err = crearconga.Write(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El archivo con los datos del gif ha sido creado")
	} //endifcongaparrot

	if eleccion == "mega" {
		response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/mega")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			crearmega, err := os.Create("imagenes/parrotdeseado.gif")
			if err != nil {
				log.Fatal(err)
			}
			defer crearmega.Close()

			_, err = crearmega.Write(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El archivo con los datos del gif ha sido creado")
	} //endifmegaparrot

	if eleccion == "middleparrot" {
		response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/middleparrot")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			crearmiddle, err := os.Create("middleparrot.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer crearmiddle.Close()

			_, err = crearmiddle.Write(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El archivo con los datos del gif ha sido creado")
	} //endifmiddleparrot

	if eleccion == "parrot" {
		response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/parrot")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			crearparrot, err := os.Create("parrot.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer crearparrot.Close()

			_, err = crearparrot.Write(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El archivo con los datos del gif ha sido creado")
	} //endifparrot

	if eleccion == "rightparrot" {
		response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/rightparrot")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			crearright, err := os.Create("rightparrot.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer crearright.Close()

			_, err = crearright.Write(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El archivo con los datos del gif ha sido creado")
	} //endifrightparrot

    http.HandleFunc("/", index_handler)
    http.Handle("/imagenes/", http.StripPrefix("/imagenes/", http.FileServer(http.Dir("./imagenes"))))
    http.ListenAndServe(":8081", nil)
}
