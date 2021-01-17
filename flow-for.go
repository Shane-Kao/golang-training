package main
import "fmt"
func main(){
	for false{
		fmt.Println("Hello")
	}
	var x int=0
	for x<3{
		fmt.Println(x)
		x++
	}
	var y int
	for y=0;y<5;y+=2{
		fmt.Println(y)
	}
	var result int=0
	var z int=1
	for z<=50{
		result=result+z
		z++
	}
	fmt.Println(result)
}