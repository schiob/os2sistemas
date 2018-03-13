package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var parrot string

func main() {
	getParrot, err := http.Get("https://ppaas.herokuapp.com/baseparrots")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		datos, _ := ioutil.ReadAll(getParrot.Body)
		fmt.Println("Opciones de parrot: ", string(datos))
	}

	fmt.Println("Escoge una opcion: ")
	fmt.Scanf("%s", &parrot)

	switch parrot {
	case "boredparrot" :
		getParrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/boredparrot")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			datos, _ := ioutil.ReadAll(getParrot.Body)
			crearbored, err := os.Create("boredparrot.gif")
			if err != nil {
				log.Fatal(err)
			}
			defer crearbored.Close()

			_, err = crearbored.Write(datos)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El gif se creo correctamente")

	case "congaparrot" :
		getParrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/congaparrot")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			datos, _ := ioutil.ReadAll(getParrot.Body)
			crearconga, err := os.Create("congaparrot.gif")
			if err != nil {
				log.Fatal(err)
			}
			defer crearconga.Close()

			_, err = crearconga.Write(datos)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El gif se creo correctamente")

	case "mega" :
		getParrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/mega")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			datos, _ := ioutil.ReadAll(getParrot.Body)
			crearmega, err := os.Create("megaparrot.gif")
			if err != nil {
				log.Fatal(err)
			}
			defer crearmega.Close()

			_, err = crearmega.Write(datos)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El gif se creo correctamente")

	case "middleparrot" :
		getParrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/middleparrot")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			datos, _ := ioutil.ReadAll(getParrot.Body)
			crearmiddle, err := os.Create("middleparrot.gif")
			if err != nil {
				log.Fatal(err)
			}
			defer crearmiddle.Close()

			_, err = crearmiddle.Write(datos)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El gif se creo correctamente")

	case "parrot" :
		getParrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/parrot")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			datos, _ := ioutil.ReadAll(getParrot.Body)
			crearparrot, err := os.Create("parrot.gif")
			if err != nil {
				log.Fatal(err)
			}
			defer crearparrot.Close()

			_, err = crearparrot.Write(datos)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El gif se creo correctamente")

	case "rightparrot" :
		getParrot, err := http.Get("https://ppaas.herokuapp.com/partyparrot/rightparrot")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			datos, _ := ioutil.ReadAll(getParrot.Body)
			crearright, err := os.Create("rightparrot.gif")
			if err != nil {
				log.Fatal(err)
			}
			defer crearright.Close()

			_, err = crearright.Write(datos)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("El gif se creo correctamente")
	}
}
