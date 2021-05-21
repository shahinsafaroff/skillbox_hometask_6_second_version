package main
import "fmt"
func main() {
	var number1, number2, sum int
	fmt.Println("Please enter first number:  ")
	fmt.Scan(&number1)
	fmt.Println("Please enter second number:  ")
	fmt.Scan(&number2)
	sum = number1 + number2

	for number1 < sum {
		number1++
	}
	fmt.Println("summary of 2 integers is: ", number1)
}