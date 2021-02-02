package main

import "fmt"

func main() {
	var float_value int
	fmt.Printf("Enter Float :")
	_, _ = fmt.Scan(&float_value)
	integer_value := int(float_value)
	fmt.Printf("The resulted integer is %d", integer_value)
}
