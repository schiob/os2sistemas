package main

import (
	"bufio"
	"fmt"
  "net/http"
	"os"
  "strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println("Len:",len(ucl))
}
if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
func main() {
	http.HandleFunc("/", index)
	fmt.Println("servidor activo")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
	fmt.Print(err)
	}
}
