package main
import "fmt"

func add1(x int){
	x+=10
}

func add2(x *int){
	*x+=10
}


func main(){
	var a int=10
	add1(a) //Pass by Value
	fmt.Println("Main Function", a)
	add2(&a) //Pass by Pointer
	fmt.Println("Main Function", a)
}