package main

import "fmt"

 
var Memoize = make(map[int]int)
func fibo(n int) int {
	// your code here
	fmt.Println("function call") // check how many print with and without Memoize and with Memoize
	val, exist := Memoize[n] // comment this line to check difference

	if exist {
		return val
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	num := fibo(n-1) + fibo(n-2)
	Memoize[n] = num // comment this line to check difference
	return num
}

 

func main() {

   fmt.Println(fibo(0))  // 0

   fmt.Println(fibo(1))  // 1

   fmt.Println(fibo(2))  // 1

   fmt.Println(fibo(3))  // 2

   fmt.Println(fibo(5))  // 5

   fmt.Println(fibo(6))  // 8

   fmt.Println(fibo(7))  // 13

   fmt.Println(fibo(9))  // 13

   fmt.Println(fibo(10)) // 55

}