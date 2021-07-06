package main

import "fmt"

func main() {

	switch rating := 4; rating {
	case 1:
		fmt.Println("One")
	case 2, 3:
		fmt.Println("Two or Three")
	case 4:
		fmt.Println("Four")
		fallthrough
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Others")
	}
	//fmt.Println(rating)

	temp := -5
	switch {
	case temp < 0:
		fmt.Println("Its cold!!")
	case temp == 0:
		fmt.Println("Its Zero!!")
	case temp > 20:
		fmt.Println("Its 20 degrees!!")
	}

}
