package main
import "fmt"
func main() {
    var input string
    fmt.Scan(&input)
    for _, char := range input {
        fmt.Print(int((char - '0') * (char - '0')))
    }
}