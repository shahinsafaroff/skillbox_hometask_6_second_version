package main
import "fmt"
func main() {
	var number1, number2, number3, i int
	number1= 10
	number2 = 100
	number3 = 1000
	for i = 0; i <=1000; i ++ {
		if number1 < 10 {
			number1++
		}
		if number2 < 100 {
			number2++
		}
		if number3 < 1000 {
			number3 ++
		}
	}
	fmt.Println(number1,number2, number3)
}