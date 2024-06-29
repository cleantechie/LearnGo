package function

import (
	"errors"
	"fmt"
)

func Hello(name string, age int) (string, error) {
	if name == "" {
		return "", errors.New("please provide name")
	}
	if age <= 0 {
		return "", errors.New("please provide valid age")
	}
	message := fmt.Sprintf("Hello my name is %v and my age is %v", name, age)
	return message, nil
}
