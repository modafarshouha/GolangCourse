package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var userInput string
	var num int
	cond, valid := true, true
	data := make([]int, 0, 3)
	for cond {
		fmt.Println("Please enter your integer :")
		fmt.Scan(&userInput)
		num, cond, valid = validateInput(userInput, cond)
		if cond && valid {
			data = append(data, num)
			sort.Ints(data)
		}
		fmt.Println(data)
		fmt.Println()
	}
}

func validateInput(userInput string, cond bool) (int, bool, bool) {
	input, err := strconv.Atoi(userInput)
	var num int
	valid := true
	if err != nil {
		if strings.ToLower(userInput) == "x" {
			fmt.Println("exiting the loop")
			cond = false
		} else {
			valid = false
		}
	} else {
		num = input
	}
	return num, cond, valid
}
