package main
import "fmt"
func main(){
	s := "hello world"
	for i, j := range s{
		fmt.Println(i, "->", string(j))
	}
}