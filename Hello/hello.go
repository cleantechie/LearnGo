package main

import (
	"fmt"
	"learnFunctions/function"
	"log"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello Goofers")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
	message, err := function.Hello("", 0)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(message)
	}

}
