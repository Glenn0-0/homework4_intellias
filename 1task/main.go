package main

import (
	"fmt"
	/*
	"strings"
	"strconv"
	"bufio"
  	"os"
	*/
)

/*
func convertStringToArrayInt(s string) []int {
	str := strings.Fields(s)
	var intArray []int
    for i := 0; i < len(str); i++ {
		intVar, _ := strconv.Atoi(string(str[i])) 
		intArray = append(intArray, intVar) //converting string to int type and adding to the list
    }
	return intArray
}
*/

func elemCounter(s []int, n int) int {
	counter := 0 //setting the counter of a particular elem to 0
	for i := 0; i < len(s); i++ {
		if s[i] == n {counter++} //a loop through the array to count the number of elements
	}
	return counter
}

func main(){
    arr := []int{4, 1, 4, -4, 6, 3, 8, 8} //!!!
	//var arr []int

    var result []int

	// !!! uncomment code below, convertStringToArrayInt function, the rest of the import and change the arr declaration to use your own input values (!!!)
	
	/*
	fmt.Println("Enter numbers (separated by spaces) and press ENTER to delete duplicates")
	input := bufio.NewReader(os.Stdin)
	yourInput, _ := input.ReadString('\n')

	arr = convertStringToArrayInt(yourInput)
	*/

	for i := 0; i < len(arr); i++ {
		if elemCounter(result, arr[i]) == 0 { result = append(result, arr[i]) }
	}

	fmt.Println(result)

}