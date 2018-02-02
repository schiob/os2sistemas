 package main

import (
"bufio"
"fmt"
"os"
"time"
)

var text string

func say(s string) {
for i := 0; i < 10; i++ {
time.Sleep(100 * time.Millisecond)
fmt.Println(s)
fmt.Println(text)
}
}

func main() {
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n')
go say("hola")
say(text)
}
