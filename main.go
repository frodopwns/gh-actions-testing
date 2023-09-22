package main

import "fmt"

func main() {
	say("Hello, World!")
}

func say(s string) string {
	fmt.Println(s)
	return s
}
