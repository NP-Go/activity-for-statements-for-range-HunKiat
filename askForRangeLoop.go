package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Insert code for only even here

	var input1, input2 string

	fmt.Println("Enter 1st No:")
	fmt.Scanln(&input1)

	fmt.Println("Enter 2nd No:")
	fmt.Scanln(&input2)

	// stage 1: check if nos are integer

	_, err := strconv.Atoi(input1)
	for err != nil {
		_, err = strconv.Atoi(input1)
		if err != nil {
			fmt.Println("Enter 1st Integer No:")
			fmt.Scanln(&input1)
		}
	}

	_, err2 := strconv.Atoi(input2)
	for err2 != nil {
		_, err2 = strconv.Atoi(input2)
		if err2 != nil {
			fmt.Println("Enter 2nd Integer No:")
			fmt.Scanln(&input2)
		}
	}

	// stage 2: Order nos smaller and bigger
	inputInt1, _ := strconv.Atoi(input1)
	inputInt2, _ := strconv.Atoi(input2)
	if inputInt2 < inputInt1 {
		// swap the values, so tt input 1 is lesser than 2
		inputInt1, inputInt2 = inputInt2, inputInt1
		fmt.Println(".... I'm swapping your nos from smaller to bigger")
	}

	// stage 3: print requests
	fmt.Println("Printing all the odd nos you entered...")
	for i := inputInt1; i <= inputInt2; i++ {
		// for odd requests
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
