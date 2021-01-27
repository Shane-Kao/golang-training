package main
import "fmt"


func add(n1 int, n2 int) int{
	var result int=n1+n2
	return result
}

func test(n1 int, n2 int) (int,int){
	var result1 int=n1+n2
	var result2 int=n1-n2
	return result1,result2
}

func main(){
	x:=add(3, 4)
	fmt.Println(x)
	n1, n2:=test(5,2)
	fmt.Println(n1,n2)
}