package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func main() {
	u1 := User{}
	fmt.Println(u1)

	u2 := User{
		Id:   1,
		Name: "John",
	}
	fmt.Println(u2)

	u3 := User{1, "John"}
	fmt.Println(u3)

	u4 := User{Name: "John", Id: 1}
	fmt.Println(u4)

	u5 := new(User)
	fmt.Println(u5)
}
