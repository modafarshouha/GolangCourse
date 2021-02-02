package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("Please enter your string :")
	userInput := bufio.NewReader(os.Stdin)
	line, _ := userInput.ReadString('\n')
	userString := strings.ToLower(strings.TrimSpace(line))
	stringLength := utf8.RuneCountInString(userString)

	firstLetter := string(userString[0])
	lastLetter := string(userString[stringLength-1])
	aAvailable := strings.Contains(userString, "a")

	cond := (firstLetter == "i") && (lastLetter == "n") && aAvailable

	if cond {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
