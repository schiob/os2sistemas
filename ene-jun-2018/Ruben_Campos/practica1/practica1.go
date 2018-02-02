package main

import (
	"fmt"
	"time"
)

var algo string

func repite(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("Escribe algo: ")
	fmt.Scanf("%v", &algo)
	fmt.Println("")
	go repite("hola")
	repite(algo)
}
