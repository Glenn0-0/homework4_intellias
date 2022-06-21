package main

import "fmt"

var (
	inputNum, i int
	count int = 0
)

func main() {

	fmt.Println("Enter a number to see its first 10 multiples:")
	fmt.Scan(&inputNum)
	i = inputNum
	for {
		if i%inputNum == 0 && count != 10 {
			fmt.Print(i, " ")
			count++
			i++
		} else if count == 10 {
			break
		} else {
			i++
			continue
		}
	}

}