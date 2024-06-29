package main

import (
	"fmt"

	"learnFunctions/function"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello Goofers")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
	message := function.Hello("Harsha", 23)
	fmt.Println(message)
}
