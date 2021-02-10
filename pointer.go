package main
import "fmt"
func main(){
	var x int=5
	fmt.Println("原來的資料", x)
	fmt.Println("資料的記憶體位置", &x)
	var xPtr *int=&x
	fmt.Println("反解指標回原來的資料", *xPtr)
	var word string="Hello"
	var wordPtr *string=&word
	fmt.Println(word)
	fmt.Println(wordPtr)
	fmt.Println(*wordPtr)
}
