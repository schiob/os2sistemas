package main 
import "fmt"
import "time"
 
   func say(s string) {
	for i := 1; i <=10 ; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("claudia")
	say("hello")
}

