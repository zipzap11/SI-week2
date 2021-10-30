package main

import (
	"fmt"
	"sync"
)



func countFreq(s string)  {
	result := make(map[rune]int)
	someMapMutex := sync.RWMutex{}
	// channel := make(chan rune)
	for _, r := range s {
		go func(r rune)  {
			// channel <- r
			someMapMutex.Lock()
			result[r]++
			someMapMutex.Unlock()
		}(r)
		// result[<-channel]++
	}
	// close(channel)
	for key, val := range result {
		fmt.Printf("%s: %d\n", string(key), val)
	}
}

func main() {
	countFreq("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
}
