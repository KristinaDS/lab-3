package main
import (
	"fmt"
	"math"
)
var k, p, v float64 =  1296, 6, 6
func main(){
	var k, p, v float64
	fmt.Println("Введите k - коэф упругости: ")
	fmt.Scan(&k)
	fmt.Println("Введите p - плотность: ")
	fmt.Scan(&p)
	fmt.Println("Введите v - объем: ")
	fmt.Scan(&v)
	m := M(p, v)
	w := W(k, m)
    fmt.Println("Период колебаний математичсекого маятника: ", T(w))
}
func M(p, v float64) float64 {
    return p * v
}
func W(k, m float64) float64 {
    return math.Sqrt(k / m)
}
func T(w float64) float64 {
    return 6 / w
}

