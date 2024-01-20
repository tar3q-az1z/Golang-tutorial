package main
import "fmt"
func main(){
	type student struct{
		name string
		id int
	}
	
	stu := student{"aziz", 100}
	fmt.Println(stu)
	fmt.Println(stu.name, stu.id)
	stu.name = "alpeh"
	fmt.Println(stu.name)
}