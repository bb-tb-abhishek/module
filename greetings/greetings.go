package greetings

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello %v, Welcome to go!", name)
}
