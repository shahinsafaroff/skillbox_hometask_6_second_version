package main
import "fmt"
func main() {
	var number1, number2, sum, i int
	fmt.Println("Please enter first number:  ")
	fmt.Scan(&number1)
	fmt.Println("Please enter second number:  ")
	fmt.Scan(&number2)
	sum = number1 + number2

	for i = number1; i < sum; i++ {
		fmt.Println(i)
	}
	fmt.Println("summary of 2 integers is: ", sum)
}