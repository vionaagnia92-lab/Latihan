package main

import "fmt"
import "net/http"

func main() {
handlerIndex := func(w http.ResponseWriter, r *http.Request) {
fmt. Fprintln(w, "Hallo nama saya viona Agnia")
}
http.HandleFunc("/", handlerIndex)

fmt.Println("server started at 43.173.13.18:9020")
http.ListenAndServe(":9020", nil)

}
