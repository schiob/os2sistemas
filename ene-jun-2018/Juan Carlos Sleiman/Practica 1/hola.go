 package main

import (
"fmt"
"time"
)

func say(s string) {
for i := 0; i < 10; i++ {
time.Sleep(100 * time.Millisecond)
fmt.Println(s)
}
}

func main() {
var text string
fmt.Print()
fmt.Scan(&text)
go say("hola")
say(text)
}
