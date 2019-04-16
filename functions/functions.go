package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func beFriendly(name string) {
	sayHello(name)
	fmt.Println("How's the weather?")
}

func addOne(x int) int {
	return x + 1
}

func sayHelloABunch() {
	fmt.Println("Hello")
	sayHelloABunch()
}

func main() {
	beFriendly("Jack")
	beFriendly("Jill")

	x := 5
	fmt.Println(addOne(x))
	x = addOne(addOne(addOne(x)))
	fmt.Println(x)

	// sayHelloABunch()
}