package main
import "fmt"
func main(){
	var x int=0
	for x<10{
		if x==7{
			break
		}
		fmt.Println(x)
		x++
	}
	var y int
	for y=0;y<5;y++{
		if y%2==0{
			continue
		}
		fmt.Println(y)
	}
	var result int=0
	for true{
		var n int
		fmt.Scanln(&n)
		if n==0{
			break
		}
		result += n
	}
	fmt.Println("總和: ", result)
}