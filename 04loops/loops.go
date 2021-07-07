package main

import "fmt"

func main() {
	start := 1
	things := []string{"abc", "ijk", "pqr", "xyz"}
	//fmt.Println(things)

	for i := 0; i < 5; i++ {
		fmt.Println(i + start)
	}
	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}

	for i := range things {
		fmt.Println(things[i])
	}

	for start < 100 {
		start += start
		if start == 32 {
			//break
			//continue
			goto abclabel
		}
		fmt.Println("Start is now", start)
	}

abclabel:
	fmt.Println("www.google.in")

}
