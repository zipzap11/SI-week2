package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	// your code here
	// top down approach
	if len(jumps) == 0 {
		return 1
	}
	if len(jumps) == 1 {
		return jumps[0]
	}
	return helper(jumps, 0, 0, map[int]int{})

}

func helper(arr []int, idx, total int, Temp map[int]int) int {
	if idx == len(arr)-1 {
		return total 
	}

	result1 := helper(arr, idx+1, total + int(math.Abs(float64(arr[idx]-arr[idx+1]))), Temp)
	result2 := math.MaxInt32
	if idx + 2 < len(arr) {
		result2 = helper(arr, idx+2, total + int(math.Abs(float64(arr[idx]-arr[idx+2]))), Temp)
	}

	result := int(math.Min(float64(result1), float64(result2)))
	fmt.Println("function call")
	return result
}
 

func main() {

   fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

   fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}