package greetings

import "fmt"

func HelloGioh(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome from Gioh!", name)
	return message
}