package main 
import "fmt"
func main(){
	var val int
	fmt.Scan(&val)
	switch x := val; x{
		case 1:
			println("c")
			fallthrough // executes next case block
		case 2:
			println("c++")
		case 3, 4:
			println("java, python")
		default:
			println("nothing")
	}
}