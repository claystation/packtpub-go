package main

import "fmt"

func tax(cost float32, taxRate float32) float32 {
	return cost * (taxRate / 100)
}

func main() {
	items := map[float32]float32{
		0.99: 7.5,
		2.75: 1.5,
		0.87: 2,
	}
	var cost float32
	for key, value := range items {
		cost += tax(key, value)
	}

	fmt.Println("Sales Tax Total: ", cost)
}