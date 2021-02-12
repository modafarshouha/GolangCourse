package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	printHeader()
	data := storeData()
	bubbleSort(data)
	printFooter(data)
}

func printHeader() {
	fmt.Println()
	fmt.Println("********************************************")
	fmt.Println("You can enter up to ten integers to be sorted")
	fmt.Println("Enter 'X' once you are done with your inputs")
	fmt.Println("********************************************")
	fmt.Println()
}

func printFooter(sorted []int) {
	fmt.Println()
	fmt.Println("********************************************")
	fmt.Println("Your sorted data is :")
	fmt.Println(sorted)
	fmt.Println("********************************************")
	fmt.Println()
}

func storeData() []int {
	data := make([]int, 0, 10)
	condition := true
	intCounter := 0
	for condition && (intCounter < 10) {
		newItem, valid, stop := readInput()
		if stop {
			condition = false
		} else if valid {
			data = append(data, newItem)
			intCounter++
		}
	}
	fmt.Println("Your entered data is :")
	fmt.Println(data)
	return data
}

func readInput() (int, bool, bool) {
	var userInput string
	fmt.Println("Please enter your integer :")
	fmt.Scan(&userInput)
	item, stop, valid := validateInput(userInput)
	return item, valid, stop
}

func validateInput(userInput string) (int, bool, bool) {
	input, err := strconv.Atoi(userInput)
	var num int
	stop, valid := false, true
	if err != nil {
		if strings.ToLower(userInput) == "x" {
			fmt.Println("exiting the loop")
			stop = true
		} else {
			valid = false
		}
	} else {
		num = input
	}
	return num, stop, valid
}

func bubbleSort(data []int) {
	sorted := false
	for !sorted {
		sorted = true
		index := 0
		for index < len(data)-1 {
			if data[index] > data[index+1] {
				swap(data, index)
				sorted = false
			}
			index++
		}
	}
}

func swap(data []int, position int) {
	temp := data[position]
	data[position] = data[position+1]
	data[position+1] = temp
}
