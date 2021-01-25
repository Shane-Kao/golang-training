package main
import "fmt"
func test(msg string){
	fmt.Println(msg)
}

func add(n1 int, n2 int){
	var result int=n1+n2
	fmt.Println(result)
}

func sum(max int){
	var result int=0
	var n int
	for n=1;n<=max;n++{
		result+=n
	}
	fmt.Println(result)
}

func main(){
	test("Hello")
	add(3,4)
	sum(100)
}