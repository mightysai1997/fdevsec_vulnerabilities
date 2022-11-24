//https://securego.io/docs/rules/g304.html
//G304: File path provided as taint input
package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    repoFile := "/safe/path/../../private/path"
    if !strings.HasPrefix(repoFile, "/safe/path/") {
        panic(fmt.Errorf("Unsafe input"))
    }
    byContext, err := ioutil.ReadFile(repoFile)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s", string(byContext))
}
