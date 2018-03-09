package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
        fmt.Printf("de web: %s\n", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    var s string
    go func(){
      for {
        _, err := fmt.Scanln(&s)
        if err != nil{
          log.Fatal(err)
        }
        fmt.Printf("got: %s\n", s)
      }
    }()

    log.Fatal(http.ListenAndServe(":8081", nil))
}
