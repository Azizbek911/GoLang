package main

import (
    "fmt"
	"myproject/greetings"
	"myproject/values"
	"myproject/variables"
	"myproject/constants"
	"myproject/for"
)

func main() {
	fmt.Println("-------------------------------------------- Hello world:")
	greetings.Hello()

	fmt.Println("-------------------------------------------- Values:")
	values.Values()

	fmt.Println("-------------------------------------------- Variables:")
	variables.Variables()

	fmt.Println("-------------------------------------------- Constants:")
	constants.Constant()


	fmt.Println("-------------------------------------------- For:")
	for_test.For()
}
