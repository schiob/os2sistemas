package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func congaparrot() {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/congaparrot")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		createconga, err := os.Create("congaparrot.gif")
		if err != nil {
			log.Fatal(err)
		}
		defer createconga.Close()

		_, err = createconga.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Se ha creado un archivo .gif")
}

func boredparrot() {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/boredparrot")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		createbored, err := os.Create("boredparrot.gif")
		if err != nil {
			log.Fatal(err)
		}
		defer createbored.Close()

		_, err = createbored.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Se ha creado un archivo .gif")
}

func megaparrot() {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/mega")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		createmega, err := os.Create("megaparrot.gif")
		if err != nil {
			log.Fatal(err)
		}
		defer createmega.Close()

		_, err = createmega.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Se ha creado un archivo .gif")
}

func middleparrot() {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/middleparrot")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		createmiddle, err := os.Create("middleparrot.gif")
		if err != nil {
			log.Fatal(err)
		}
		defer createmiddle.Close()

		_, err = createmiddle.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Se ha creado un archivo .gif")
}

func rightparrot() {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/rightparrot")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		createright, err := os.Create("rightparrot.gif")
		if err != nil {
			log.Fatal(err)
		}
		defer createright.Close()

		_, err = createright.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Se ha creado un archivo .gif")
}

func parrot() {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/parrot")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		createparrot, err := os.Create("parrot.gif")
		if err != nil {
			log.Fatal(err)
		}
		defer createparrot.Close()

		_, err = createparrot.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Se ha creado un archivo .gif")
}

var results []string
var parr string

func main() {
	response, err := http.Get("https://ppaas.herokuapp.com/baseparrots")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Opciones de parrot: ", string(data))
	}

	fmt.Println("Elige una de las opciones anteriores: ")
	fmt.Scanf("%s", &parr)

	if parr == "boredparrot" {
		fmt.Println("Has escogido: " + parr)
		boredparrot()
	} else if parr == "congaparrot" {
		fmt.Println("Has escogido: " + parr)
		congaparrot()
	} else if parr == "mega" {
		fmt.Println("Has escogido: " + parr)
		megaparrot()
	} else if parr == "middleparrot" {
		fmt.Println("Has escogido: " + parr)
		middleparrot()
	} else if parr == "parrot" {
		fmt.Println("Has escogido: " + parr)
		parrot()
	} else if parr == "rightparrot" {
		fmt.Println("Has escogido: " + parr)
		rightparrot()
	} else {
		fmt.Println("Tienes un error ortográfico. Vuelve a intentarlo: ")
		main()
	}
}
