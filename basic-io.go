package main
import "fmt"
func main(){
	fmt.Println(3)
	fmt.Println(3,"Hello",true)
	var x int
	fmt.Print("輸入一個數字")
	fmt.Scanln(&x)
	fmt.Println(x)
	var a int
	var b int
	fmt.Print("輸入第一個數字:")
	fmt.Scanln(&a)
	fmt.Print("輸入第二個數字:")
	fmt.Scanln(&b)
	var result1 int=a+b
	fmt.Println(result1)
	var c int
	var d int
	fmt.Print("輸入兩個數字. 用空格隔開:")
	fmt.Scanln(&c, &d)
	var result2 int=c+d
	fmt.Println(result2)	
}