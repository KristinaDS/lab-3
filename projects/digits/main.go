package main
import "fmt"
func main() {
    var input string
    fmt.Scan(&input)
    maxDigit := 0
    for _, char := range input {
        digit := int(char - '0') 
        if digit > maxDigit {
            maxDigit = digit
        }
    }
    fmt.Println(maxDigit)
}