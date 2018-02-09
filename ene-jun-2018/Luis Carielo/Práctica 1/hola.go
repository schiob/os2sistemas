package main

import (
	"fmt"
	"time"
)

func mensaje(m string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(m)
	}
}

func main() {
	var palabra string
	fmt.Print()
	fmt.Scan(&palabra)
	go mensaje("hola")
	mensaje(palabra)
}
