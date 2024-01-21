package main
import (
	"fmt"
	"strings"
)
func main(){
	str := "Go Programming"
	for _, j := range str {
		fmt.Print(string(j), " ")
	}
	println("")
	// length of string using len() function
	fmt.Println("size:", len(str))
	// concatenation of strings using +
	fmt.Println("Doing " + str) // concat.
	cmp := "Go"
	// compare two strings using Compare function of strings package
	val := strings.Compare(str, cmp)
	fmt.Println(str, "||", cmp)
	if val == 0{
		fmt.Println("both equal") // can be checked by == operator
	}else if val == 1{
		fmt.Println("first one greater")
	}else{
		fmt.Println("second one greater")
	}
	// check if substring exists
	if strings.Contains(str, "Go"){
		fmt.Println("str contains Go")
	}
	// replacing some characters
	cmp = strings.Replace(cmp, "G", "M", 1)
	fmt.Println(cmp)
	// to upper case
	fmt.Println(strings.ToUpper(str))
	// split string to slice
	fmt.Println(strings.Split(str, " "))
	slice := []string{}
	slice = append(slice, "strawberry", "orange")
	fmt.Println(slice)
	// slice to string
	fmt.Println(strings.Join(slice, " "))
}