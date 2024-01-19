package main
import "fmt"
func main(){
	// var arr = []string{} // same as below
	arr := []string{} // a slice
	fmt.Println(arr)
	arr = append(arr, "az")
	fmt.Println(arr)
	num := [5]string{"1", "2","3","4","5"}
	arr = num[0:3]
	fmt.Println(arr)
	arr = append(arr, "6", "7")
	fmt.Println(arr)
	arr2 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"};
	copy(arr2, arr)
	fmt.Println(arr2)
	fmt.Println(len(arr2))
	arr3 := make([]int, 3, 100);
	fmt.Println(arr3)
}