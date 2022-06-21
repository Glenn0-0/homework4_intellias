package main

import "fmt"

var (
	inputNum, i int
)

func main() {

	fmt.Println("Enter a whole number to check if it is prime:")
	fmt.Scan(&inputNum)
	if inputNum == 1 { 
		fmt.Println("Entered number is not prime!")
		return
	}
	for i = 2; i <= inputNum/2; i++ {
		if inputNum%i == 0 {
			fmt.Println("Entered number is not prime!")
			return
		}
	}
	fmt.Println("Entered number IS PRIME!")

}