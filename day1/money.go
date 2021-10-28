package main

import "fmt"

 

func moneyCoins(money int) []int {
  // your code here
	coins := []int{1,10,20,50,100,200,500,1000,2000,5000,10000}
	answer := []int{}
	idx := len(coins)-1
	for money > 0 && idx >= 0 {
		if coins[idx] <= money {
			tmp := money / coins[idx]
			money = money % coins[idx]
			for i := 0; i < tmp; i++ {
				answer = append(answer, coins[idx])
			}
		}
		idx--
	}
	return answer
}

 

func main() {

   fmt.Println(moneyCoins(123))   // [100 20 1 1 1]

   fmt.Println(moneyCoins(432))   // [200 200 20 10 1 1]

   fmt.Println(moneyCoins(543))   // [500, 20, 20, 1, 1, 1]

   fmt.Println(moneyCoins(7752))  // [5000, 2000, 500, 200, 50, 1, 1]

   fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]

}