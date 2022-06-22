package main

import "fmt"

func main(){
    arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
    var result []int

	for i := 0; i < len(arr); i++ {
		counter := 0 
		for k := 0; k < len(result); k++ {
			if result[k] == arr[i] {counter++} 
		}
		if counter == 0 { result = append(result, arr[i]) }
	}

	fmt.Println(result)

}