package main

import "fmt"

func main() {

	/*var p *int
	if p != nil {
		fmt.Println("P value :", *p)
	} else {
		fmt.Println("P is nil")
	}*/

	var xyz float64 = 100
	xyzRef := &xyz

	fmt.Println(xyz)
	xyz = xyz * 2
	fmt.Println(*xyzRef)

}
