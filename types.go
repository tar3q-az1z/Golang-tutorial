package main
import "fmt"
func main(){
	// integers
	var id, class int64
	id = 100000000000000
	class = -45
	fmt.Println("====================")
	fmt.Println(id, class)
	val := 30 // default type is int32
	p := int64(1400000000000)
	fmt.Println(val, p)
	fmt.Println("====================")

	// float
	var x float32 = 15.5
	var y float64 = 3434333434343.34243434243
	z := 2344342424324.2343423434 // default type is float64
	fmt.Println(x, y, z)
	fmt.Println("=====================")

	// string
	var s string = "looook"
	var t string = `at you` // OK
	fmt.Println(s, t)
	fmt.Println("=====================")

	// boolean
	var isOk bool = true
	fmt.Println(isOk)
	isOk = false
	fmt.Println(isOk)
}