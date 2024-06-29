package function

import (
	"fmt"
)

func Hello(name string, age int) string {
	message := fmt.Sprintf("Hello my name is %v and my age is %v", name, age)
	return message
}
