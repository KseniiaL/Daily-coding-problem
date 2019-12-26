//package main demonstrates simple solution to the given problem
package main

import (
	"fmt"
)

//func countSteps takes number of stairs as the input 
//and returns a number of unique ways to climb tha staircase
//using bottom-up approach technique
func countSteps(num int) int {
	//for small number there is only one vay to climb
	if num == 0 || num == 1 {
		return 1
	}

	//creating an array, which will hold values of previous numbers
	arr := make([]int, num+1)
	//set up two wirst values as mentioned above	
	arr[0] = 1
	arr[1] = 1

	for i:= 2; i<=num; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[num]
}

func main() {
	fmt.Println(countSteps(4))
}