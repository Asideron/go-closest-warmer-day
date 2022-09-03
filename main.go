package main

import (
	"github.com/Asideron/go-closest-warmer-day/stack"

	"fmt"
)

type dayInfo struct {
	temperature int
	position    int
}

func main() {
	daysTemperatures := []int{13, 12, 15, 11, 9, 12, 16}
	results := make([]int, len(daysTemperatures))

	s := stack.NewStack[*dayInfo]()
	for i := len(daysTemperatures) - 1; i >= 0; i-- {
		search := true
		for search {
			if day, ok := s.Head(); ok {
				if day.temperature > daysTemperatures[i] {
					results[i] = day.position - i // A warmer day was found.
					s.Push(&dayInfo{
						temperature: daysTemperatures[i],
						position:    i,
					})
					search = false
				} else {
					s.Pop() // Keep searching for a warmer day.
				}
			} else {
				results[i] = 0 // No warmer day was met.
				s.Push(&dayInfo{
					temperature: daysTemperatures[i],
					position:    i,
				})
				search = false
			}
		}
	}

	fmt.Println("Result:", results)
}
