package slices_tut

import (
	"fmt"
)


func BasicTut() {
	numbers:= []int{10, 20, 30}

	fmt.Println("Original", numbers)
	fmt.Println("Capacity", cap(numbers))
	fmt.Println("Length", len(numbers))

	fmt.Println("-------")


	numbers = append(numbers,40)

	fmt.Println("Original", numbers)
	fmt.Println("Capacity", cap(numbers))
	fmt.Println("Length", len(numbers))

	fmt.Println("-------")


	subset :=  numbers[1:3]

	fmt.Println("Subset [1:3]:", subset)


	subset[0] = 99

	fmt.Println("subset:", subset)
	fmt.Println("numbers: ", numbers)

}