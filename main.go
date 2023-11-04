package main

import (
	"fmt"
	"strings"
)

func main() {

	goFun()
	goString()
	goArrays()
	goSlice()
	goMap()
	goStruct1()
	goStruct2()

}

func goFun() {

	start := 1
	multiplier := 5

	for start <= 10 {
		product := multiplier * start
		fmt.Printf("%d * %d = %d\n", multiplier, start, product)
		start++
	}

}

func goString() {

	goL := "Golang"
	golang := "Golang"

	fmt.Println(strings.Compare(golang, goL))
	fmt.Println(strings.Contains(golang, goL))
	fmt.Println(strings.Replace(goL, "G", "g", 1))
	fmt.Println(strings.ToUpper(golang))

	greet := "Hello World"
	greetArray := strings.Split(greet, " ")

	for index, item := range greetArray {
		fmt.Printf("%d. %s\n", index+1, item)
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

	marks["Math"] = 99
	marks["Geography"] = 88
	fmt.Println(marks)

	delete(marks, "History")
	fmt.Println(marks)

	for sub, score := range marks {
		fmt.Printf("%s: %d\n", sub, score)
	}

}

func goStruct1() {

	type Person struct {
		name string
		age  int
	}

	Luffy := Person{"Luffy", 16}
	Zoro := Person{"Zoro", 17}
	Brook := Person{name: "Brook"}

	fmt.Println(Luffy.name)
	fmt.Println(Zoro)
	fmt.Println(Brook)

}

func goStruct2() {

	type Area func(int, int) int

	type Volume struct {
		length int
		width  int
		height int

		area Area
	}

	cube := Volume{
		length: 2,
		width:  3,
		height: 4,
		area: func(length int, width int) int {
			return length * width
		},
	}

	fmt.Printf("Area = %d\n", cube.area(cube.length, cube.width))
}
