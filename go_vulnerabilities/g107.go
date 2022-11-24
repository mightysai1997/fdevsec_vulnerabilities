//G107: Url provided to HTTP request as taint input
//https://securego.io/docs/rules/g107.html
package main
import (
    "net/http"
    "io/ioutil"
    "fmt"
    "os"
)
func main() {
    url := os.Getenv("tainted_url")
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
            panic(err)
    }
    fmt.Printf("%s", body)
}
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

var url string = "https://www.google.com"

func main() {

    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s", body)
}
