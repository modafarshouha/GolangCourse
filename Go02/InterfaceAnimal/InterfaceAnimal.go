//This file has a bug in appending step (it appends duplicated elements instead of one)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	eat()
	move()
	speak()
}

type Cow struct {
	name, food, locomotion, noise string
}

func (cow Cow) eat() {
	fmt.Println(cow.food)
}
func (cow Cow) move() {
	fmt.Println(cow.locomotion)
}
func (cow Cow) speak() {
	fmt.Println(cow.noise)
}

type Bird struct {
	name, food, locomotion, noise string
}

func (bird Bird) eat() {
	fmt.Println(bird.food)
}
func (bird Bird) move() {
	fmt.Println(bird.locomotion)
}
func (bird Bird) speak() {
	fmt.Println(bird.noise)
}

type Snake struct {
	name, food, locomotion, noise string
}

func (snake Snake) eat() {
	fmt.Println(snake.food)
}
func (snake Snake) move() {
	fmt.Println(snake.locomotion)
}
func (snake Snake) speak() {
	fmt.Println(snake.noise)
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

var objects []interface{}

// var objects = make([]interface{}, 0)

func main() {
	var request, name, typeOrInfo string
	var r = bufio.NewReader(os.Stdin)
	for true {
		printHeader()
		fmt.Fscanf(r, "%s %s %s", &request, &name, &typeOrInfo)
		request = strings.ToLower(request)
		name = strings.ToLower(name)
		typeOrInfo = strings.ToLower(typeOrInfo)
		if request == "x" {
			break
		} else if request == "newanimal" {
			createAnimal(name, typeOrInfo)
		} else if request == "query" {
			queryAnimal(name, typeOrInfo)
		}
	}
}

func printHeader() {
	fmt.Println("\n********************************************")
	fmt.Println("We have these animals types:")
	fmt.Println("1- 'cow'\t 2-'bird'\t 3-'snake'")
	fmt.Println("You can ask the following about any of them:")
	fmt.Println("1- Food eaten > 'food'")
	fmt.Println("2- Locomotion method > 'move'")
	fmt.Println("3- Spoken sound > 'speak'")
	fmt.Println("You can create new animal > 'newanimal'")
	fmt.Println("or query information of an existing one > 'query'")
	fmt.Println("and the required information about it (food, move, speak)")
	fmt.Println()
	fmt.Println("We have this list of animals now:")
	fmt.Println(objects)
	fmt.Println("Input format is: >[request] [animal name] [type/info]")
	fmt.Println("Enter 'X' to exit")
	fmt.Println("\n>")
}

func createAnimal(name, Type string) {
	switch Type {
	case "cow":
		animal := Cow{name: name, food: "grass", locomotion: "walk", noise: "moo"}
		objects = append(objects, animal)
	case "bird":
		animal := Bird{name: name, food: "worms", locomotion: "fly", noise: "peep"}
		objects = append(objects, animal)
	case "snake":
		animal := Snake{name: name, food: "mice", locomotion: "slither", noise: "hsss"}
		objects = append(objects, animal)
	}
}

func queryAnimal(name, Info string) {
	for _, animal := range objects {
		animalType := typeof(animal)
		switch animalType {
		case "main.Cow":
			if animal.(Cow).name == name {
				fmt.Println("IN COW")
				switch Info {
				case "food":
					animal.(Cow).eat()
					fmt.Println("IN FOOD")
				case "move":
					animal.(Cow).move()
				case "speak":
					animal.(Cow).speak()
				}
			}
		case "main.Bird":
			if animal.(Bird).name == name {
				switch Info {
				case "food":
					animal.(Bird).eat()
				case "move":
					animal.(Bird).move()
				case "speak":
					animal.(Bird).speak()
				}
			}
		case "main.Snake":
			if animal.(Snake).name == name {
				switch Info {
				case "food":
					animal.(Snake).eat()
				case "move":
					animal.(Snake).move()
				case "speak":
					animal.(Snake).speak()
				}
			}
		}
	}
}
