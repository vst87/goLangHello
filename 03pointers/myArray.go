package main

import "fmt"

func main() {
	var numbers [3]string
	numbers[0] = "Uno"
	numbers[1] = "dos"
	numbers[2] = "tres"

	fmt.Println(numbers)

	var colors = [4]string{"Red", "Blue", "Green", "Pink"}
	fmt.Println(colors)
	fmt.Println(colors[2])
	fmt.Println(len(colors))
}
