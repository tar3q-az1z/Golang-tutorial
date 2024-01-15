package main
import "fmt"
// operator in golang
func main(){
	var x, y int 
	fmt.Println("enter the value of x, y:")
	fmt.Scan(&x, &y)
	fmt.Println("x + y:", x + y)
	fmt.Println("x * y:", x * y)
	fmt.Println("x / y:", x / y)
	fmt.Println("x - y:", x - y)
	fmt.Println("x % y:", x % y)
	y++ // y = y + 1
	fmt.Println("y++:", y)
	x += 5 // x = x + 5
	fmt.Println("x += 5:", x)
	fmt.Println("x == y?:", x == y)
	fmt.Println("x >= y?:", x >= y)
	fmt.Println("x >= 0 && y >= 0:", x >= 0 && y >= 0)
	fmt.Println("x >> 2:", x >> 2) // like c++
}