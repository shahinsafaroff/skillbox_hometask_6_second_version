package main
import "fmt"
func main() {
	var bag1, bag2, bag3, volumeOfBag1, volumeOfBag2, volumeOfBag3 int
	fmt.Println("Please enter First Bags volume: ")
	fmt.Scan(&volumeOfBag1)
	fmt.Println("Please enter Second Bags volume: ")
	fmt.Scan(&volumeOfBag2)
	fmt.Println("Please enter Third Bags volume: ")
	fmt.Scan(&volumeOfBag3)
	for  {
		if bag1 < volumeOfBag1 {
			bag1++
			continue
		}
		if bag2 < volumeOfBag2 {
			bag2++
			continue
		}
		if bag3 < volumeOfBag3 {
			bag3++
			continue
		}
		break
			}
	fmt.Println("Baskets are filled in: ")
	fmt.Println(bag1, bag2, bag3)
}
