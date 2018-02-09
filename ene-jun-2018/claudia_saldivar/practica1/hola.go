package main

import "fmt"
import "time"

var palabra string

func say(s string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	fmt.Println("palabra:")
	fmt.Scanf("%v", &palabra)
	fmt.Println("")
	go say("claudia")
	say(palabra)
}
