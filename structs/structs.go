package main

import "fmt"

// structs are a way to package together different types
// type says we are making a new type
// position is the name of the type
// struct says it's going to be a struct
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

func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Printf("Bad guy %v is at x: %v, y: %v with %v health remaining\n", b.name, x, y, b.health)
}

func main() {
	// var p position
	// p.x = 5
	// p.y = 4

	p := position{3, 2}

	fmt.Println(p)
	// access fields of the struct with .
	fmt.Println(p.x)
	fmt.Println(p.y)

	b := badGuy{"Nathan", 1, p}
	whereIsBadGuy(b)

}