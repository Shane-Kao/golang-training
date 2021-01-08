package main
import (
	"fmt"
	"reflect"
)
func main(){
	fmt.Println(3)
	fmt.Println(3.1415)
	fmt.Println("測試中文")
	fmt.Println(true)
	fmt.Println('a')
	var x int
	x=4
	fmt.Println(x)
	x=10
	fmt.Println(x)
	//x="Hello"
	var f float64=3.1415
	fmt.Println(f)
	var s string="Hello Golang"
	fmt.Println(s)
	var test bool=true
	fmt.Println(test)
	var c rune='b'
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(test))
	quantity:=4
	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(float64(quantity)))
	fmt.Println(float64(quantity))
}