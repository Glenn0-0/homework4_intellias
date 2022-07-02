package main

import "fmt"

func main() {
	var inputNum, i int
	var count int = 0
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