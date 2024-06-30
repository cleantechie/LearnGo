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
	message, err := function.Hello("Harsha", 22)

	names := []string{"Harsha_1", "Harsha_2", "Harsha_3"}
	ages := []int{22, 22, 22}
	messages, errs := function.Hellos(names, ages)
	if err != nil || errs != nil {
		log.Fatal(err)
		log.Fatal(errs)
	} else {
		fmt.Println(message)
		fmt.Println(messages)
	}

}
