package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// For these functions, we can omit the value (eb/sb) because it is not being used
// EX: func (eb englishBot) --> eb not needed in this case
func (englishBot) getGreeting() string {
	// Very custom logic for generatig an English greeting :)
	return "Hello there!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating a Spanish greeting :)
	return "Hola!"
}


func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}