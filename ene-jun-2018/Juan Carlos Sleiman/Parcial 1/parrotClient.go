package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	res, err := http.Get("https://ppaas.herokuapp.com/baseparrots")
	if err != nil {
		fmt.Print("It's not possible to connect.")
	}
	fmt.Print("\nThese are the parrot options: \n")
	rs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", rs)
	fmt.Print("\nType the parrot option you want to see: \n")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n') //text is where the typed parrot is saved
	//if strings.TrimRight(input, "\n") == "a"
	if strings.TrimRight(text, "\n") == "boredparrot" {

		res, err := http.Get("https://ppaas.herokuapp.com/partyparrot/boredparrot")
		if err != nil {
			log.Fatal(err)
		}
		//rs, err := ioutil.ReadAll(res.Body)
		//res.Body.Close()
		//if err != nil {
		//	log.Fatal(err)
		//}
		out, err := os.Create("./boredparrot.gif")
		if err != nil {
			// code
		}
		defer out.Close()
		io.Copy(out, res.Body)
		//fmt.Printf("%s", rs)

	} else if strings.TrimRight(text, "\n") == "congaparrot" {

		res, err := http.Get("https://ppaas.herokuapp.com/partyparrot/congaparrot")
		if err != nil {
			log.Fatal(err)
		}
		out, err := os.Create("./congaparrot.gif")
		if err != nil {
			// code
		}
		defer out.Close()
		io.Copy(out, res.Body)
	} else if strings.TrimRight(text, "\n") == "mega" {

		res, err := http.Get("https://ppaas.herokuapp.com/partyparrot/mega")
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create("./megaparrot.gif")
		if err != nil {
			// code
		}
		defer out.Close()
		io.Copy(out, res.Body)
	} else if strings.TrimRight(text, "\n") == "middleparrot" {

		res, err := http.Get("https://ppaas.herokuapp.com/partyparrot/middleparrot")
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create("./middleparrot.gif")
		if err != nil {
			// code
		}
		defer out.Close()
		io.Copy(out, res.Body)

	} else if strings.TrimRight(text, "\n") == "rightparrot" {

		res, err := http.Get("https://ppaas.herokuapp.com/partyparrot/rightparrot")
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create("./rightparrot.gif")
		if err != nil {
			// code
		}
		defer out.Close()
		io.Copy(out, res.Body)
	} else if strings.TrimRight(text, "\n") == "parrot" {

		res, err := http.Get("https://ppaas.herokuapp.com/partyparrot/parrot")
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create("./justparrot.gif")
		if err != nil {
			// code
		}
		defer out.Close()
		io.Copy(out, res.Body)

	} else {
		print("That option is not listed\n")
	}

	defer res.Body.Close()
}
