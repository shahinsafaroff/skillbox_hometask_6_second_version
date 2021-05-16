package main
import "fmt"
func main() {
	var price, discount float64
	fmt.Println("The Original Price of item:  ")
	fmt.Scan(&price)
	fmt.Println("Discount for this item: ")
	fmt.Scan(&discount)

	if (discount<=price*0.3) && (discount<=2000.0) {
		fmt.Println("Here is final Discount Amount: ", discount)
	} else {
		fmt.Println("Discount is too much!!!")
	}
}
