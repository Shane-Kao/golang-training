package main
import "fmt"

func main(){
	var money int
	fmt.Println("請問想領多少錢?")
	fmt.Scanln(&money)
	if money<100{
		fmt.Println("Too Few")
	}else if money<=100000{
		fmt.Println("OK")
	}else{
		fmt.Println("Too Much")
	}
}