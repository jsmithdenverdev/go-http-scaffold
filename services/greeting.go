package services

import "fmt"

type GreetingService struct {
}

func (s GreetingService) Greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
