package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){
    input := "1 9 3 4 -5" //!!!
    var result string

	//converting string to an array of int
	str := strings.Fields(input)
	var res []int
    for i := 0; i < len(str); i++ {
		intVar, _ := strconv.Atoi(string(str[i])) 
		res = append(res, intVar)
    }

	//finding the max and min values
	max := res[0]
	min := res[0]
	for i := 0; i < len(res); i++ {
		if res[i] > max { max = res[i] }
		if res[i] < min { min = res[i] }
	}

	//adjusting the result string
	maxStr := strconv.Itoa(max)
	minStr := strconv.Itoa(min)
	if min == max {
		fmt.Println(maxStr)
	} else {
		result = maxStr + " " + minStr
		fmt.Println(result)
	}

}

// uncomment the code below to try your own inputs

/*
package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
  	"os"
)

func main(){
    var result string

	fmt.Println("Enter numbers (separated by spaces) and press ENTER to get the max and min values")
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')

	//converting string to an array of int
	str := strings.Fields(input)
	var res []int
    for i := 0; i < len(str); i++ {
		intVar, _ := strconv.Atoi(string(str[i])) 
		res = append(res, intVar)
    }

	//finding the max and min values
	max := res[0]
	min := res[0]
	for i := 0; i < len(res); i++ {
		if res[i] > max { max = res[i] }
		if res[i] < min { min = res[i] }
	}

	//adjusting the result string
	maxStr := strconv.Itoa(max)
	minStr := strconv.Itoa(min)
	if min == max {
		fmt.Println(maxStr)
	} else {
		result = maxStr + " " + minStr
		fmt.Println(result)
	}

}
*/