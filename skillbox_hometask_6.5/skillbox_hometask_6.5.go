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
	for bag1= 0; bag1<= volumeOfBag1; bag1 ++ {
		fmt.Println(bag1)
		if bag1 == volumeOfBag1 {
			continue
		}
	}
	for bag2= 0; bag2<= volumeOfBag2; bag2 ++ {
		fmt.Println(bag2)
		if bag2 == volumeOfBag2 {
			continue
		}
	}
	for bag3= 0; bag3<= volumeOfBag3; bag3 ++ {
		fmt.Println(bag3)
		if bag3 == volumeOfBag3 {
			break
		}
	}
}