package main

import "fmt"

func main() {
	helloworld()
	result := addme(1, 3)
	resultadd := addme1(4, 3)
	fmt.Println("Result:", result)
	fmt.Println("Result:", resultadd)
	resultvalues, mylength, myname := additionvalues(1, 4, 5, 5, 12, 5)
	fmt.Println("Result:", resultvalues, mylength, myname)
}

func helloworld() {
	fmt.Println("Hello World")
}

func addme(v1 int, v2 int) int {
	return v1 + v2
}

func addme1(v1, v2 int) int {
	return v1 + v2
}

func additionvalues(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum = sum + values[i]
	}
	length := len(values)
	name := "Just name"
	return sum, length, name
}
