package main


import (
  "net/http"
  "log"
  "fmt"
  "io/ioutil"
)

func main() {
  resp, err := http.Get("http://localhost:8080/hello/asdfghj")
  if err != nil {
    log.Fatalf("fatal error when getting http response, desc: %v", err)
    return
  }

  defer resp.Body.Close()
  result, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(result))

}
