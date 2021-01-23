package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeeting(eb)
	printGreeeting(sb)

}

func printGreeeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "getGreeting_englishBot"
}

func (spanishBot) getGreeting() string {
	return "getGreeting_spanishBot"
}
