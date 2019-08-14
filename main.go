package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	// sb := spanishBot{}
	printGreeting(eb)
	// printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting)
// }

func (englishBot) getGreeting() string {
	// VERY custom logic for eng greeting
	return "Hey man!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for spanish greeting
	return "Hola tio!"
}
