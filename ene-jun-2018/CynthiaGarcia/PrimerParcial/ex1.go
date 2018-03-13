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
		fmt.Print("error.")
	}
	fmt.Print("\noptions: \n")
	rs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("error %s", rs)
	fmt.Print("tipo")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	if strings.TrimRight(text, "\n") == "mega" {

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
	}

	if strings.TrimRight(text, "\n") == "rightparrot" {

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
	}

	defer res.Body.Close()

}
