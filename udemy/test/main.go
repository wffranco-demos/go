package main

import (
	"fmt"

	"github.com/wffranco-demos/go/udemy/test/mypackage"
)

func main() {
	var car = mypackage.Car{
		Brand: "Toyota",
		Model: "Corolla",
	}
	fmt.Println(car)
}
