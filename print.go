package main
import "fmt"
func main(){

	msg := "hello"
	x := int64(333)
	y := 34.34434324345

	// use of fmt.Print() 
	fmt.Print(msg, " ")
	fmt.Print(x, y)
	fmt.Println("\n#############")

	// use of fmt.Println() -> print newline at end
	fmt.Println(msg)
	fmt.Println(msg, x) // space separated automatically when comma used
	fmt.Println("message:", msg)
	fmt.Println("###############")

	// use of fmt.Printf() -> print like c/c++
	// %d -> int, %f/%g -> real, %s -> string, %t -> bool
	fmt.Printf("message = %s\n", msg)
	fmt.Printf("float value = %0.3f\n", y)
	fmt.Printf("int value = %d\n", x)
	fmt.Printf("msg is %s with value %d\n", msg, x)
}