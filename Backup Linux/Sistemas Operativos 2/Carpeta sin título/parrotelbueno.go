package main

import (
	"fmt"
	"io/ioutil"
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
		fmt.Println("Puedes elegir entre todos estos parrots: \n", string(datosleidos))
	}

	fmt.Println("Elige un parrot: ")
	fmt.Scanf("%s", &eleccion)


	if eleccion == "boredparrot" {
		leerparrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/boredparrot")
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
		fmt.Println("El gif del boredparrot ha sido creado!")
        fmt.Println("Accede a localhost:8081 para visualizar :)")
	} //endifboredparrot

	if eleccion == "congaparrot" {
		leerparrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/congaparrot")
		if err != nil {
			fmt.Printf("Request http fallido! %s\n", err)
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
		fmt.Println("El gif del congaparrot ha sido creado!")
        fmt.Println("Accede a localhost:8081 para visualizar :)")
	} //endifcongaparrot

	if eleccion == "mega" {
		leerparrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/mega")
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
		fmt.Println("El gif del megaparrot ha sido creado!")
        fmt.Println("Accede a localhost:8081 para visualizar :)")
	} //endifmegaparrot

	if eleccion == "middleparrot" {
		leerparrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/middleparrot")
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
		fmt.Println("El gif del middleparrot ha sido creado!")
        fmt.Println("Accede a localhost:8081 para visualizar :)")
	} //endifmiddleparrot

	if eleccion == "parrot" {
		leerparrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/parrot")
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
		fmt.Println("El gif del parrot ha sido creado!")
        fmt.Println("Accede a localhost:8081 para visualizar :)")
	} //endifparrot

	if eleccion == "rightparrot" {
		leerparrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/rightparrot")
		if err != nil {
			fmt.Printf("Request http fallido! %s\n", err)
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
		fmt.Println("El gif del parrot ha sido creado!")
        fmt.Println("Accede a localhost:8081 para visualizar :)")
	} //endifrightparrot

    http.HandleFunc("/", index_handler)
    http.Handle("/imagenes/", http.StripPrefix("/imagenes/", http.FileServer(http.Dir("./imagenes"))))
    http.ListenAndServe(":8081", nil)
}
