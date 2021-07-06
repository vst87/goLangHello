package main

import "fmt"

func main() {
	var abc string = "Hello"
	fmt.Println(abc)

	var xyz string
	xyz = "Hello World"
	fmt.Println(xyz)

	bot := "I'm a bot"
	fmt.Println(bot)

	botRating := 3
	fmt.Println(botRating)

	botRatings := 3.
	fmt.Printf("%v, %T", botRatings, botRatings)

	var Ironman, IronMan1 string = "\naaa", "bbb"
	fmt.Println(Ironman, IronMan1)

	var defaultValue string
	fmt.Println(defaultValue)

	var (
		spiderman = "Im Spiderman"
		age       = 18
		powers    = "web slings, spidy sense"
		antman    = "Im Antman"
	)
	fmt.Println(spiderman, age, powers, antman)

}
