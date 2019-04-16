package main

import (
	"github.com/markwai/rpg/game"
	"github.com/markwai/rpg/ui2d"
)

func main() {
	ui := &ui2d.UI2d{}
	game.Run(ui)
}