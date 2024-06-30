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

func Hellos(names []string, ages []int) (map[string]string, error) {
	var messages = make(map[string]string)
	for i := 0; i < len(names); i++ {
		message, err := Hello(names[i], ages[i])
		if err != nil {
			return nil, err
		}
		messages[names[i]] = message
	}
	return messages, nil

	//a simple for loop for range is as
	// for _, name := range names {

	// }

}
