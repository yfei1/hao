package main

import (
  "net"
  "net/http"
  "fmt"
  "os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    name, err := os.Hostname()
    if err != nil {
       fmt.Printf("Oops: %v\n", err)
       return
    }

addrs, err := net.LookupHost(name)
if err != nil {
    fmt.Printf("Oops: %v\n", err)
    return
}
  title := r.URL.Path
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", title, addrs)
}

func main() {
  http.HandleFunc("/hello/", helloHandler)
  http.ListenAndServe(":8080", nil)
}
