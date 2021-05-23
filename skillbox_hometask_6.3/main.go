package main
import "fmt"
func main() {
	var price, discount float64

	fmt.Println("The Original Price of item:  ")
	fmt.Scan(&price)

	fmt.Println("Discount for this item: ")
	fmt.Scan(&discount)

	for discount > price*0.3 || discount > 2000 {
		fmt.Println("Discount is too much!!!")
		fmt.Println("Discount for this item: ")
		fmt.Scan(&discount)
	}

	fmt.Println("Price:", price)
	fmt.Println("Discount", discount)
	fmt.Println("Total:", price-discount)
	}
