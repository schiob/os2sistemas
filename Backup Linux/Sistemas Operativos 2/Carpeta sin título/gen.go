package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "time"
  "encoding/json"
    )

var results []string

func GetHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}

// PostHandler converts post request body to string
func PostHandler(w http.ResponseWriter, r *http.Request) {
    	if r.Method == "POST" {
    		body, err := ioutil.ReadAll(r.Body)
    		if err != nil {
    			http.Error(w, "Error reading request body",
    				http.StatusInternalServerError)
    		}
    		results = append(results, string(body))
    		fmt.Fprintln(w, "POST done")
        fmt.Fprintln(w, "Longitud:", len(body))
    	} else {
    		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    	}
    }

    func main(){
      results = append(results, time.Now().Format(time.RFC1123Z))
      http.HandleFunc("/", GetHandler)
      http.HandleFunc("/post", PostHandler)
      fmt.Println("El servidor se encuentra en ejecución")
      log.Fatal(http.ListenAndServe(":8081", nil))
}
