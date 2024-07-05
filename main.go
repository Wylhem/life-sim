package main

import (
	"life/games"
	"log"

	"github.com/hajimehoshi/ebiten"
)

// main is the entry point for the application.
func main() {
	game := games.NewGame()

	game.Fps = 10

	ebiten.SetWindowSize(games.SCREEN_WIDTH, games.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Game of Life")
	ebiten.SetMaxTPS(game.Fps)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
