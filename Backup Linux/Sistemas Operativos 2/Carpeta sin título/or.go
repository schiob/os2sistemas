package main
import (

      "io"
    "net/http"
    "fmt"
        "io/ioutil"

)
func partyparrot(res http.ResponseWriter, req *http.Request){
  response, err := http.Get("https://ppaas.herokuapp.com/partyparrot")
      if err != nil {
          fmt.Printf("The HTTP request failed with error %s\n", err)
      } else {
          data, _ := ioutil.ReadAll(response.Body)
          res.Header().Set("Content-Type","image/gif")

             io.WriteString(res,string(data))
          fmt.Println(string(data))
      }

  }
  func congaparrot(res http.ResponseWriter, req *http.Request){
    response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/congaparrot")
        if err != nil {
            fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            res.Header().Set("Content-Type","image/gif")

               io.WriteString(res,string(data))
            fmt.Println(string(data))
        }

    }
    func boredparrot(res http.ResponseWriter, req *http.Request){
      response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/boredparrot")
          if err != nil {
              fmt.Printf("The HTTP request failed with error %s\n", err)
          } else {
              data, _ := ioutil.ReadAll(response.Body)
              res.Header().Set("Content-Type","image/gif")

                 io.WriteString(res,string(data))
              fmt.Println(string(data))
          }

      }
      func megaparrot(res http.ResponseWriter, req *http.Request){
        response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/mega")
            if err != nil {
                fmt.Printf("The HTTP request failed with error %s\n", err)
            } else {
                data, _ := ioutil.ReadAll(response.Body)
                res.Header().Set("Content-Type","image/gif")

                   io.WriteString(res,string(data))
                fmt.Println(string(data))
            }

        }
        func middleparrot(res http.ResponseWriter, req *http.Request){
          response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/middleparrot")
              if err != nil {
                  fmt.Printf("The HTTP request failed with error %s\n", err)
              } else {
                  data, _ := ioutil.ReadAll(response.Body)
                  res.Header().Set("Content-Type","image/gif")

                     io.WriteString(res,string(data))
                  fmt.Println(string(data))
              }

          }
          func rightparrot(res http.ResponseWriter, req *http.Request){
            response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/rightparrot")
                if err != nil {
                    fmt.Printf("The HTTP request failed with error %s\n", err)
                } else {
                    data, _ := ioutil.ReadAll(response.Body)
                    res.Header().Set("Content-Type","image/gif")

                       io.WriteString(res,string(data))
                    fmt.Println(string(data))
                }
            }
            func parrot(res http.ResponseWriter, req *http.Request){
              response, err := http.Get("https://ppaas.herokuapp.com/partyparrot/parrot")
                  if err != nil {
                      fmt.Printf("The HTTP request failed with error %s\n", err)
                  } else {
                     data, _ := ioutil.ReadAll(response.Body)
                      res.Header().Set("Content-Type","image/gif")

                        io.WriteString(res,string(data))
                     fmt.Println(string(data))
                  }

              }


func main() {
http.HandleFunc("/partyparrot", partyparrot)
http.HandleFunc("/congaparrot", congaparrot)
http.HandleFunc("/boredparrot", boredparrot)
http.HandleFunc("/megaparrot", megaparrot)
http.HandleFunc("/middleparrot", middleparrot)
http.HandleFunc("/rightparrot", rightparrot)
http.HandleFunc("/parrot", parrot)
http.ListenAndServe(":8081", nil)
}
