package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is not")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 10; num < 0 {
		fmt.Println("Why do this?")
	} else if num < 10 {
		fmt.Println(num, "Why declare variables in conditional?")
	} else {
		fmt.Println(num, "I will use gofmt -w anyway.")
	}
}
