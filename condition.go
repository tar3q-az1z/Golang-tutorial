package main
import "fmt"
func main(){
	var x, y int
	fmt.Scan(&x, &y)
	if(x == y){
		fmt.Println("equal")
	}else if(x > y){
		fmt.Println("greater")
	}else{
		fmt.Println("smaller")
	}
}