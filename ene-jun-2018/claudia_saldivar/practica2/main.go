package main

import (
    "fmt"
    "net/http"
)

func index(w http.ResponseWriter,r *http.Request){
  texto:="SERVIDOR"
  cant :=len([]rune(texto))
  fmt.Fprintln(w,"texto:",texto)
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
