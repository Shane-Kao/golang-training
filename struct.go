package main
import "fmt"
type Person struct{
	name string
	age int
}
func main(){
	var p1 Person=Person{"Shane", 18}
	fmt.Println(p1.name, p1.age)
	var p2 Person=Person{age:33, name:"Jack"}
	p2.name="Jerry"
	fmt.Println(p2.name, p2.age)
}