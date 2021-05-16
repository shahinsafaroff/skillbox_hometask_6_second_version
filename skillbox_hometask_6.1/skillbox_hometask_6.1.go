package main

import "fmt"

func main() {
	var number, i int
	fmt.Println("Please enter a number:  ")
	fmt.Scan(&number)
	for i=0;i<number;i++{
	fmt.Println(i)
	}
	fmt.Println(number)
}