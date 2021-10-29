package main

import "fmt"

 

func fibo(n int) int {
	// your code here
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	fibonacci := []int{1, 1}
	for i := 2; i < n; i++ {
		newFibo := fibonacci[i-1] + fibonacci[i-2]
		fibonacci = append(fibonacci, newFibo)
	}
	return fibonacci[len(fibonacci)-1]
}

 

func main() {

   fmt.Println(fibo(0))  // 0

   fmt.Println(fibo(1))  // 1

   fmt.Println(fibo(2))  // 1

   fmt.Println(fibo(3))  // 2

   fmt.Println(fibo(5))  // 5

   fmt.Println(fibo(6))  // 8

   fmt.Println(fibo(7))  // 13

   fmt.Println(fibo(9))  // 34

   fmt.Println(fibo(10)) // 55

}