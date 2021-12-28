package main

import "fmt"

func main() {
	numbers := []int{}

	for index := 0; index <= 10; index++ {
		numbers = append(numbers, index)
	}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("%v is even ",number)
	}		else {
			fmt.Printf("%v is odd ",number)
	}
	fmt.Println("")

	}
}