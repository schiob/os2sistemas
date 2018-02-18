package main

import (
    "fmt"
    "net/http"
)

var texto string
func index(w http.ResponseWriter,r *http.Request){
  fmt.Scanf("%v", &texto)
  cant :=len([]rune(texto))
  fmt.Fprintln(w,"texto")
  fmt.Fprint(w,"caracteres: ",cant)
}
func main(){
  http.HandleFunc("/",index)
  fmt.Println("servidor activo")
  err:=http.ListenAndServe(":8081",nil)
  if err != nil{
    fmt.Print(err)
  }
}
