package main

import (
	"fmt"
	"math"
)

func main() {
	var acc, v_i, d_i, time float64
	fmt.Println("Please enter the equation parameters:")
	fmt.Println("Acceleration: ")
	fmt.Scan(&acc)
	fmt.Println("Initial velocity: ")
	fmt.Scan(&v_i)
	fmt.Println("Initial displacement: ")
	fmt.Scan(&d_i)
	fmt.Println("Time: ")
	fmt.Scan(&time)

	displaceTime := genDisplaceFn(acc, v_i, d_i)
	fmt.Println("Displacement after ", time, " seconds is: ", displaceTime(time))
}

func genDisplaceFn(acceleration, velocity_o, displace_o float64) func(float64) float64 {
	fn := func(t float64) float64 {
		displace_t := 0.5*acceleration*math.Pow(t, 2) + velocity_o*t + displace_o
		return displace_t
	}
	return fn
}
