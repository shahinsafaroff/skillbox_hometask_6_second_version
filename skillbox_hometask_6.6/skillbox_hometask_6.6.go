package main

import "fmt"

func main() {
	minFloor := 1
	maxFloor := 24
	currentFloor := 1
	direction := 1

	maxPassengers:= 2
	totalPassengers := 0

	passenger1 := 4
	passenger2 := 7
	passenger3 := 10

	for passenger1 != minFloor || passenger2 !=minFloor || passenger3 != minFloor || currentFloor != minFloor{
		currentFloor += direction
		fmt.Printf("The Elevator  is in floor:  %v,  Quantity of passengers:  %v\n",  currentFloor, totalPassengers)
		if direction == 1 && currentFloor == maxFloor{
			direction = -1
		} else if currentFloor == minFloor {
			direction = 1
			totalPassengers = 0
		} else if direction == -1 {
			if passenger1 == currentFloor && totalPassengers < maxPassengers {
				fmt.Printf("We  took first passenger from floor %v\n", currentFloor)
				passenger1 =  1
				totalPassengers++
			}
			if passenger2 == currentFloor && totalPassengers < maxPassengers {
				fmt.Printf("We  took second passenger from floor %v\n", currentFloor)
				passenger2=  1
				totalPassengers++
			}
			if passenger3 == currentFloor && totalPassengers < maxPassengers {
				fmt.Printf("We  took third passenger from floor %v\n", currentFloor)
				passenger3 =  1
				totalPassengers++
			}
		}
	}
	fmt.Println("All passengers are in First Floor")

}