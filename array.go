package main
import "fmt"
func main(){
	// var arr [3]int
	arr := [3]int{}
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)
	arr[1] = -10
	fmt.Println(arr)
	fmt.Println("size:", len(arr))
}