package main
import "fmt"
func main(){
	for i:= 1; i <= 3; i++ {
		if(i == 3){
			break
		}
		if(i == 2){
			continue
		}
		fmt.Println(i)
	}
}