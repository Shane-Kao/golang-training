package main
import "fmt"
func main(){
	var numbers [3]int
	fmt.Println(numbers)
	numbers[0]=15
	numbers[1]=115
	fmt.Println(numbers)
	fmt.Println(numbers[1])
	var names [2]string=[2]string{"John", "Mary"}
	fmt.Println(names)
	var grades [5]int=[5]int{60, 90, 75,10,33}
	var index int
	var total int=0
	for index=0;index<len(grades); index++{
		fmt.Println(grades[index])
		total+=grades[index]
	}
	fmt.Println(total/len(grades))
}