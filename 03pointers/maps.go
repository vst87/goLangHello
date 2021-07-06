package main

import "fmt"

func main() {
	score := make(map[string]int)
	score["abc"] = 100
	score["xyz"] = 1
	score["pqr"] = 2
	score["lmn"] = 2
	score["ijk"] = 4
	fmt.Println(score)
	getScore := score["tqr"]
	fmt.Println(getScore)

	delete(score, "pqr")
	fmt.Println(score)

	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k, v)
	}
}
