package main

import (
	"fmt"
	"sort"
)

func main() {
	/*	var things = []string{"male", "female", "Unisex"}
			fmt.Println(things)
			things = append(things, "any")
			fmt.Println(things)
			things = append(things[1:len(things)])
				fmt.Println(things)
			things = append(things[1 : len(things)-1])
			fmt.Println(things)

		heros := make([]string, 3, 3)
		heros[0] = "thor"
		heros[1] = "ironman"
		heros[2] = "spiderman"
		//fmt.Println(heros)
		heros = append(heros, "deadpool")
		//fmt.Println(heros)
		fmt.Println(cap(heros))
	*/

	myints := []int{4, 3, 5, 8, 1}
	isSorted := sort.IntsAreSorted(myints)
	//fmt.Println("Sorted Before: ", isSorted)
	if isSorted == false {
		sort.Ints(myints)
		isSorted := sort.IntsAreSorted(myints)
		fmt.Println("Sorting Array done: ", isSorted)
	} else {
		fmt.Println("Sorted Array:  ", isSorted)
	}

}
