package main

import "fmt"

func main() {

	table(5)

}

func table(multiplier int) {
	start := 1

	for start <= 10 {
		product := multiplier * start
		fmt.Printf("%d * %d = %d\n", multiplier, start, product)
		start++
	}
}
