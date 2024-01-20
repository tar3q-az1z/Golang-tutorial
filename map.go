package main
import "fmt"
func main(){
	mp := map[int]string{} // empty map of int to string
	fmt.Println(mp)
	mp[0] = "aziz"
	mp[1] = "sunflower"
	mp[2] = "jui"
	mp[3] = "cordova"
	fmt.Println(mp) // print
	mp[3] = "rose"
	fmt.Println(mp[3]) // value of key 3
	mp[4] = "dragonfly"
	fmt.Println(mp)
	delete(mp, 4) // delete value of key 4
	fmt.Println(mp)
}