package main
import "fmt"
func main(){
	// var x int = 10 // explicit 
	// var x = 10 // implicitly
	x := 15 // declaration with initialization; shorthand
	fmt.Println(x)
	var y = 13
	var z int // default value is 0
	fmt.Println(y)
	fmt.Println(z)
	// x = "okkk" // can't change type of variable
	// fmt.Println(x)
	s, n := "cases", 100
	// var s, n = "cases", 100 also valid
	fmt.Println(s, n)
	
	// const declarations
	const PI = 3.14159;
	// const PI := 3.1416 // := doesn't work for constant
	fmt.Println(PI)
}