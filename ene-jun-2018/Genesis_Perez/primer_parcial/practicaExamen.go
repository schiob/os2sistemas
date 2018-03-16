package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var results []string
var pp string

func main() {
	response, err := http.Get("https://ppaas.herokuapp.com/baseparrots")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Opciones de parrot: ", string(data))
	}

	fmt.Println("Elige una de las opciones anteriores: ")
	fmt.Scanf("%s", &pp)

	if pp == "boredparrot" {
		response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/boredparrot")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			crearbored, err := os.Create("boredparrot.txt")
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

	if pp == "congaparrot" {
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

	if pp == "mega" {
		response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/mega")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			crearmega, err := os.Create("megaparrot.txt")
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

	if pp == "middleparrot" {
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

	if pp == "parrot" {
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

	if pp == "rightparrot" {
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
}
