package main
import "fmt"
func main(){
	sum := 0
	for i:= 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)
	num := [5] int{1, 2, 3, 4, 5}
	for i := range num { // i being index of sequence
		print(num[i], " ")
	}
	println("")
	for _, j := range num{ // j being the element of sequence
		print(j, " ")
	}
}