package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	//Frontend
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name: ")
	name, _ = reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our center: ")

	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Backend
	//fmt.Printf("%v, %v", name, mynumRating)
	fmt.Printf("Hello %v, \n Thanks for rating our xyz store %v star rating\n\n Your rating is recorded at %v \n\n", name, mynumRating, time.Now().Format(time.Stamp))
	if mynumRating == 5 {
		fmt.Println("Cheers!!")
	} else if mynumRating == 4 || mynumRating == 3 {
		fmt.Println("Good go!")
	} else if mynumRating < 3 {
		fmt.Println("Need to work!")
	}
}
