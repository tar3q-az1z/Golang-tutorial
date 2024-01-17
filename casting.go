package main
import "fmt"
func main(){
	f := float64(35.343)
	l := int64(f)
	// notOk := f + l // go does not support implicit conversions
	fmt.Printf("f: %f\n", f)
	fmt.Printf("l: %d\n", l)
	// fmt.Printf("notOk: %f\n", notOk)
	fmt.Println(f + float64(l))
}