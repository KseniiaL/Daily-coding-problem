//package main demonstrates simple solution to the given problem
//I'll try to update algorithm, which won't use a division
package main

import (
	"fmt"
)

//func CreateNewArr returns a sice every element in which
//is the product of all the numbers in the original array except given one
func CreateNewArr(arr []int) []int {
	//Create empty slice
	newArr := make([]int, len(arr))
	mult := 1

	//Multiply all of array items
	for i := range arr {
		mult *= arr[i]
	}

	//Update values in the slice with the result of
	//dividing the product by the element
	for i := range arr {
		newArr[i] = mult / arr[i]
	}

	return newArr
}

func main() {
	fmt.Println(CreateNewArr([]int{1, 2, 3, 4, 5}))
}
