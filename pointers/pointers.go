package main

import "fmt"

type position struct {
	x float32
	y float32
}

// can put structs inside of structs
type badGuy struct {
	name 	string
	health 	int
	pos 	position
}

// passing whole badGuy means it will make a copy, which is a lot of space
func describeBadGuy(b *badGuy) {
	x := b.pos.x // Go automatically dereferences to access members of struct, thanks Go
	y := b.pos.y // in C, would have to change . to -> e.g. b->pos.y
	name := b.name
	health := b.health
	fmt.Printf("Bad guy %v is at x: %v, y: %v with %v health remaining\n", name, x, y, health)
}

func addOne(num int) {
	num = num + 1
}

func addOneToPtr(num *int) {
	*num = *num + 1 // * before variable name dereferences the pointer, so get me the thing that the pointer points to
}

// Go lets you see them and interact with them, but doesn't let you do anything you want with them (like C)
func main() {
	x := 5
	fmt.Println(x)

	// & means the address that x lives in, in the memory of computer
	// xPtr := &x // inferred type
	
	var xPtr *int = &x // explicit type *int
	// * in front of the variable type means it's a pointer type
	
	fmt.Println(xPtr) // e.g. prints 0xc00008a000


	// you can't modify these values, so what can you do?
	addOne(x) // sent a copy of x to function
	fmt.Println(x) // printing original x, which is unchanged

	addOneToPtr(xPtr) // send the address of where x lives
	fmt.Println(x) // now the actual value has been modified

	p := position{4, 2}
	badguy := badGuy{"Nathan", 1, p}
	describeBadGuy(&badguy)
}