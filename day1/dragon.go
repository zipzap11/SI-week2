package main

import (
	"fmt"
	"sort"
)

 

func DragonOfLoowater(dragonHead, knightHeight []int) {
  // your code here
  sort.Ints(knightHeight)
  sort.Ints(dragonHead)
  i := 0
  j := 0
  total := 0
  for i < len(dragonHead) {
	  if j >= len(knightHeight) {
		  break
	  }
	  if dragonHead[i] <= knightHeight[j] {
		  total += knightHeight[j]
		  i++
		  j++
	  } else {
		  j++
	  }
  }
  if i == len(dragonHead) {
	  fmt.Println(total)
  } else {
	  fmt.Println("knight fall")
  }
}

 

func main() {

   DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11

   DragonOfLoowater([]int{5, 10}, []int{5})         // knight fall

   DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall

   DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10

}