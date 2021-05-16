package main
import "fmt"
func main() {
	var number1, number2, number3 int
	number1= 0
	number2 = 0
	number3 = 0
	for number1= 0; number1<=10; number1++ {
		fmt.Println(number1)
		if number1 == 10 {
			break
		}
	}
	for number2= 0; number2<=100; number2++ {
		fmt.Println(number2)
		if number2 == 100 {
			break
		}
	}
	for number3= 0; number3<=1000; number3++ {
		fmt.Println(number3)
		if number3 == 1000 {
			break
		}
	}
	}