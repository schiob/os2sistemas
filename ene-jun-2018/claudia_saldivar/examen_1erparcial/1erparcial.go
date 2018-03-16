package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func partyparrot(res http.ResponseWriter, req *http.Request) {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		res.Header().Set("Content-Type", "image/gif")

		io.WriteString(res, string(data))
		fmt.Println(string(data))
	}

}
func congaparrot(res http.ResponseWriter, req *http.Request) {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/congaparrot")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		res.Header().Set("Content-Type", "image/gif")

		io.WriteString(res, string(data))
		fmt.Println(string(data))
	}

}
func boredparrot(res http.ResponseWriter, req *http.Request) {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/boredparrot")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		res.Header().Set("Content-Type", "image/gif")

		io.WriteString(res, string(data))
		fmt.Println(string(data))
	}

}
func megaparrot(res http.ResponseWriter, req *http.Request) {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/mega")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		res.Header().Set("Content-Type", "image/gif")

		io.WriteString(res, string(data))
		fmt.Println(string(data))
	}

}
func middleparrot(res http.ResponseWriter, req *http.Request) {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/middleparrot")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		res.Header().Set("Content-Type", "image/gif")

		io.WriteString(res, string(data))
		fmt.Println(string(data))
	}

}
func rightparrot(res http.ResponseWriter, req *http.Request) {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/rightparrot")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		res.Header().Set("Content-Type", "image/gif")

		io.WriteString(res, string(data))
		fmt.Println(string(data))
	}
}
func parrot(res http.ResponseWriter, req *http.Request) {
	response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/parrot")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		res.Header().Set("Content-Type", "image/gif")

		io.WriteString(res, string(data))
		fmt.Println(string(data))
	}

}

func main() {

	http.ListenAndServe(":8081", nil)

	fmt.Print("opciones disponibles ")
	switch os := runtime.GOOS; os {
	case "boredparrot":
		http.HandleFunc("/partyparrot", partyparrot)
		if file == "" {
			t.Fail()
		}
	case "congaparrot":
		http.HandleFunc("/congaparrot", congaparrot)
		if file == "" {
			t.Fail()
		}
	case "mega":
		http.HandleFunc("/megaparrot", megaparrot)
		if file == "" {
			t.Fail()
		}
	case "middleparrot":
		http.HandleFunc("/middleparrot", middleparrot)
		if file == "" {
			t.Fail()
		}
	case "parrot":
		http.HandleFunc("/parrot", parrot)
		if file == "" {
			t.Fail()
		}
	case "rightparrot":
		http.HandleFunc("/rightparrot", rightparrot)
		if file == "" {
			t.Fail()
		}
	}
}
