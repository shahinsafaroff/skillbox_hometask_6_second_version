package main

import "fmt"

func main() {
	var count, i int
	fmt.Println("Please enter a number:  ")
	fmt.Scan(&count)
	for i = 0; i <= count; i++{
		fmt.Println(i)
	}
	fmt.Println(count)
}