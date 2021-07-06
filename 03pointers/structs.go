package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	abc := User{"abc", "abc@xy.com", 12}
	fmt.Printf("%+v\n", abc)
	fmt.Printf("%v\n", abc)
	fmt.Printf("%v\n", abc.Name)

	var xyz = new(User)
	xyz.Name = "xyz"
	xyz.Email = "xyz@ab.com"
	xyz.Age = 13
	fmt.Printf("%+v\n", xyz)
	fmt.Printf("%v\n", xyz.Name)

	var test = &User{"test", "test@test.com", 44}
	fmt.Printf("%+v\n", test)
}
