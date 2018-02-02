package main 
import(
"fmt"
"time")
var word string
func main(){
fmt.Println("Escribe una palabra: ")
fmt.Scanf("%v%",&word)
go repetir("Hola")
repetir(word)
}
func repetir(s string){
for i:=0;i<10;i++{
time.Sleep(100 * time.Millisecond)
fmt.Println(s)
}
}
