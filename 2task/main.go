package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){
    input := "1 9 3 4 -5" 
    var result string

	str := strings.Fields(input)
	var res []int
    for i := 0; i < len(str); i++ {
		intVar, _ := strconv.Atoi(string(str[i])) 
		res = append(res, intVar)
    }

	max := res[0]
	min := res[0]
	for i := 0; i < len(res); i++ {
		if res[i] > max { max = res[i] }
		if res[i] < min { min = res[i] }
	}

	maxStr := strconv.Itoa(max)
	minStr := strconv.Itoa(min)
	if min == max {
		fmt.Println(maxStr)
	} else {
		result = maxStr + " " + minStr
		fmt.Println(result)
	}

}