package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

const (
	maxLength     = 20
	fileExtension = "allNames.txt"
)

type name struct {
	fname string
	lname string
}

func main() {
	allNames := make([]name, 0)
	file, err := os.Open(fileExtension)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fileReader := bufio.NewReader(file)

	for {
		line, err := fileReader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(line)
			break
		}
		fullName := strings.Fields(line)
		first, last := fullName[0], fullName[1]
		firstName, lastName := TrimNames(first, last)

		newPerson := name{fname: firstName, lname: lastName}
		allNames = append(allNames, newPerson)
	}
	PrintAllNames(allNames)
}

func TrimNames(firstName string, lastName string) (string, string) {
	if len(firstName) > maxLength {
		firstName = firstName[:maxLength]
	}

	if len(lastName) > maxLength {
		lastName = lastName[:maxLength]
	}
	return firstName, lastName
}

func PrintAllNames(allNames interface{}) {
	i := 0
	values := reflect.ValueOf(allNames)
	for i < values.Len() {
		fmt.Println(values.Index(i).Interface())
		i++
	}
}
