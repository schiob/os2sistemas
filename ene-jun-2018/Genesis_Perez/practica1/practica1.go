package main 

import ( 
	"fmt"
	"time"
)

var palabra string 

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func main () {
	fmt.Println("Needed word: ")
	fmt.Scanf("%s",&palabra)
	fmt.Println()
	go say(palabra)
	say("hola")
}
