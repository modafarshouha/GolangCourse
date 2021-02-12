package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	var animal Animal
	var animalInput, infoInput, info string
	var stopLoop, wrongEntry bool
	var r = bufio.NewReader(os.Stdin)
	for true {
		printHeader()
		fmt.Fscanf(r, "%s %s", &animalInput, &infoInput)
		animalInput = strings.ToLower(animalInput)
		infoInput = strings.ToLower(infoInput)
		animal, stopLoop, wrongEntry = readAnimal(animalInput)
		if stopLoop {
			break
		} else if !wrongEntry {
			info, stopLoop, wrongEntry = readInfo(animal, infoInput)
			if stopLoop {
				break
			} else if !wrongEntry {
				fmt.Println("Animal: ", animalInput, "\t Information: ", info)
			}
		}
	}
}

func printHeader() {
	fmt.Println("\n********************************************")
	fmt.Println("We have these animals:")
	fmt.Println("1- Cow\t 2-Bird\t 3-Snake")
	fmt.Println("You can ask the following about any of them:")
	fmt.Println("1- Food eaten > Food")
	fmt.Println("2- Locomotion method > Move")
	fmt.Println("3- Spoken sound > Speak")
	fmt.Println("Please enter the animal name (cow, bird, snake)")
	fmt.Println("and the required information about it (food, move, speak)")
	fmt.Println("Input format is: >[animal name] [info]")
	fmt.Println("Enter 'X' to exit")
	fmt.Println("\n>")
}

func (animal *Animal) eat() string {
	return animal.food
}

func (animal *Animal) move() string {
	return animal.locomotion
}

func (animal *Animal) speak() string {
	return animal.noise
}

func readAnimal(animalInput string) (Animal, bool, bool) {
	var animal Animal
	exitLoop, wrongEntry := false, false
	switch animalInput {
	case "cow":
		animal = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	case "bird":
		animal = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	case "snake":
		animal = Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	case "x":
		exitLoop = true
	default:
		fmt.Println("Wrong entry !!!")
		wrongEntry = true
	}
	return animal, exitLoop, wrongEntry
}

func readInfo(animal Animal, infoInput string) (string, bool, bool) {
	var info string
	infoInput = strings.ToLower(infoInput)
	exitLoop, wrongEntry := false, false
	switch infoInput {
	case "food":
		info = animal.eat()
	case "move":
		info = animal.move()
	case "speak":
		info = animal.speak()
	case "x":
		exitLoop = true
	default:
		fmt.Println("Wrong entry !!!")
		wrongEntry = true
	}
	return info, exitLoop, wrongEntry
}
