package main
import "fmt"
func test() (int,error) {
    return 0, nil
}
func main() {
    v, _ := test()
    fmt.Println(v)
}
