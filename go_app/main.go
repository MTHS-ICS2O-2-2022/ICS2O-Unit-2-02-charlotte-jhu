// Created by: Charlotte Jhu
// Created on: March 2023
//
// This program displays area and perimeter of a rectangle

package main

import (
	"fmt"
)

func main() {
	// This function calculates the area and perimeter of a rectangle
	// with dimensions 5cm x 3cm

	// input
	fmt.Println("If a rectangle has the dimensions:")
	fmt.Println("5cm x 3cm")
	fmt.Println()

	// process
	area := 5 * 3
	perimeter := (5 + 3) * 2

	// output
	fmt.Println("Area is", area, "cm²")
	fmt.Println("Perimeter is", perimeter, "cm")

	fmt.Println("\nDone.")
}
