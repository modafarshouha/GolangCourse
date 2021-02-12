//This is not my solution!!!

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func CreateAnimal(animalList *[]Animal, name string, animalType string) {
	switch animalType {
	case "cow":
		*animalList = append(*animalList, Cow{name})
	case "snake":
		*animalList = append(*animalList, Snake{name})
	case "bird":
		*animalList = append(*animalList, Bird{name})
	}
	fmt.Println("Created it!")
}

func AnimalAction(animal Animal, animalAction string) {
	switch animalAction {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func printAction(animalList []Animal, animalName string, animalAction string) {
	for _, animal := range animalList {
		cowValue, cowOk := animal.(Cow)
		if cowOk {
			if animalName == cowValue.name {
				AnimalAction(cowValue, animalAction)
			}
		}
		snakeValue, snakeOk := animal.(Snake)
		if snakeOk {
			if animalName == snakeValue.name {
				AnimalAction(snakeValue, animalAction)
			}
		}
		birdValue, birdOk := animal.(Bird)
		if birdOk {
			if animalName == birdValue.name {
				AnimalAction(birdValue, animalAction)
			}
		}
	}
}

func TrimNLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func main() {
	in := bufio.NewReader(os.Stdin) // create bufio Reader
	var animalList []Animal
	for {
		fmt.Println("\nPlease enter command:")
		fmt.Print(">")

		userInput, _ := in.ReadString('\n') //prompt user for input string
		inputStrings := strings.Split(userInput, " ")
		commandType := TrimNLower(inputStrings[0])

		if commandType == "newanimal" {
			animalName := TrimNLower(inputStrings[1])
			animalType := TrimNLower(inputStrings[2])

			CreateAnimal(&animalList, animalName, animalType)
		} else if commandType == "query" {
			animalName := TrimNLower(inputStrings[1])
			animalAction := TrimNLower(inputStrings[2])

			printAction(animalList, animalName, animalAction)
		} else {
			fmt.Println("\nWrong command!")
		}

	}
}
