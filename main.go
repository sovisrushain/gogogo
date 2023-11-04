package main

import "fmt"

func main() {

	table(5)
	goArrays()
	goSlice()
	goMap()

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

func goSlice() {

	listOfNumbers := [5]int{1, 2, 3, 4, 5}
	evenNumbers := []int{}
	oddNumbers := []int{}

	for i := 0; i < len(listOfNumbers); i++ {
		if listOfNumbers[i]%2 == 0 {
			evenNumbers = append(evenNumbers, listOfNumbers[i])
		} else {
			oddNumbers = append(oddNumbers, listOfNumbers[i])
		}
	}

	fmt.Printf("There are %d even numbers and %d odd numbers are present\n",
		len(evenNumbers),
		len(oddNumbers))

}

func goMap() {

	marks := map[string]int{
		"Math":    98,
		"Science": 95,
		"History": 90,
	}

	marks["Geography"] = 88

	fmt.Println(marks)

	delete(marks, "History")

	fmt.Println(marks)

	for sub, score := range marks {
		fmt.Printf("%s: %d\n", sub, score)
	}

}
