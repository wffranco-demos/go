package main

import (
	"fmt"
	"go/test/mypackage"
)

func main() {
	var car = mypackage.Car{
		Brand: "Toyota",
		Model: "Corolla",
	}
	fmt.Println(car)
}
