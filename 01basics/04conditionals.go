package main

import "fmt"

func main() {
	var isLoggedIn bool = true
	if isLoggedIn {
		fmt.Println("Show cart page")
		var len, err = fmt.Println("Show cart page")
		if err == nil {
			fmt.Println(len)
			fmt.Printf("Length is %v, %T\n", len, len)
			fmt.Println(`I'm testing`)
		}
	} else {
		fmt.Println("Show user login page")
	}
}
