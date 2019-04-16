package ui2d

import (
	"github.com/markwai/rpg/game"
	"fmt"
)

type UI2d struct{
}

func (* UI2d) Draw(level *game.Level) {
	fmt.Println("Did something")
}