package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 21; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("hola")
	say("adios")
}
