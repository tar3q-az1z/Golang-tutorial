package main
import "fmt"
func main(){
	var s string
	var n int 

	// input using Scan() function
	// space, newline can be used to give input
	// fmt.Scan(&s, &n)
	// fmt.Println(s, n)
	// fmt.Println("------------")

	// input using Scanln() function
	// takes input until newline found
	// fmt.Scanln(&x, &y)
	// fmt.Println(s, n)

	// input using Scanf() function
	// take input until newline found using format specifier
	fmt.Scanf("%s %d\n", &s, &n)
	fmt.Println(s, n)
	
}