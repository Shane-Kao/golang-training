package main
import "fmt"
func main(){
	var x int
	x=3*3+10
	fmt.Println(x)
	x=5
	x=x+1
	x+=2
	fmt.Println(x)
	x++
	fmt.Println(x)
	var b bool=4>3
	fmt.Println(b)
	b=!false
	fmt.Println(b)
	b=true&&false
	fmt.Println(b)
	b=true||false
	fmt.Println(b)
}