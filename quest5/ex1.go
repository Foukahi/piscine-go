package main

import "fmt"

func FirstRune(s string) rune {

	r := []rune(s)
	return r[0]

}

func main() {

	fmt.Printf(string(FirstRune("Hello!")))
	fmt.Printf(string(FirstRune("Salut!")))
	fmt.Printf(string(FirstRune("Ola!")))
	fmt.Printf(string(FirstRune("\n")))

}
