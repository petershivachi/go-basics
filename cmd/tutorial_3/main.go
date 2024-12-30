package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue = "Hello World"
	printMe(printValue)

	var numerator = 11
	var denominator = 2
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err)
	} else if remainder == 0 {
		fmt.Printf("Result : %d", result)
	} else {
		fmt.Printf("Result %d : Remainder %d", result, remainder)
		switch remainder {
		case 1, 2:
			fmt.Println("The division close")
		default:
			fmt.Println("The division wasn't close")

		}
	}

	switch {
	case err != nil:
		fmt.Println(err)
	case remainder == 0:
		fmt.Printf("Result : %d", result)
	default:
		fmt.Printf("Result %d : Remainder %d", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator

	return result, remainder, err
}
