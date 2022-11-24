//https://securego.io/docs/rules/g104.html
//G104: Audit errors not checked
package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func a() error {
    return fmt.Errorf("This is an error")
}

func b() {
    fmt.Println("b")
    ioutil.WriteFile("foo.txt", []byte("bar"), os.ModeExclusive)
}

func c() string {
    return fmt.Sprintf("This isn't anything")
}

func main() {
    _ = a()
    a()
    b()
    c()
}
