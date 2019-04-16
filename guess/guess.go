package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100
	fmt.Printf("Think of a number between %v and %v, when ready press ENTER", low, high)
	scanner.Scan()

	tries := 1
	for {
		guess := (low + high) / 2
		fmt.Printf("Try #%v: I guess your number is %v\n", tries, guess)
		fmt.Println("Is this:")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
			tries++
		} else if response == "b" {
			low = guess + 1
			tries++
		} else if response == "c" {
			fmt.Printf("I guessed correct in %v guesses!\n", tries)
			break
		} else {
			fmt.Println("Invalid response, please try again.")
		}
	}
}