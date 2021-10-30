package main

import (
	"fmt"
)



func Lfreq(s string)  {
	result := make(map[rune]int)
	channel := make(chan rune)
	for _, r := range s {
		go func()  {
			channel <- r
		}()
		result[<- channel]++
	}
	for key, val := range result {
		fmt.Printf("%s: %d\n", string(key), val)
	}
}

func main() {
	Lfreq("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
}
