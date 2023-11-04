package main

import "fmt"

func main() {

	table(5)
	goArrays()

}

func table(multiplier int) {
	start := 1

	for start <= 10 {
		product := multiplier * start
		fmt.Printf("%d * %d = %d\n", multiplier, start, product)
		start++
	}
}

func goArrays() {
	languages := [3]string{"Go", "Java", "Python"}

	length := len(languages)

	fmt.Printf("Theere are %d languages present.\n", length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d. %s\n", i+1, languages[i])
	}
}
