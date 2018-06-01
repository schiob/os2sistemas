package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const End = 100 * time.Millisecond

func hello(i int) {
	time.Sleep(10 * time.Duration(i%5) * time.Millisecond)
	fmt.Println("Hola")
}
func palabras(i int, text string) {
	time.Sleep(10 * time.Duration(i%5) * time.Millisecond)
	fmt.Print(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	for i := 1; i <= 10; i++ {
		go hello(i)
		go palabras(i, text)
	}
	time.Sleep(End)
}
