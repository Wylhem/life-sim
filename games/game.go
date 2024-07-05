package games

import (
	"fmt"
	"image/color"
	world "life/World"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/ncruces/zenity"
	"golang.org/x/image/font/basicfont"
)

type GameState int

const (
	Menu GameState = iota
	Simulation
)

type Game struct {
	World         *world.World
	State         GameState
	Paused        bool
	LiveCellCount int
	Fps           int
}

const (
	SCREEN_WIDTH  = 1000
	SCREEN_HEIGHT = 800
)

func NewGame() *Game {
	return &Game{
		State: Menu,
	}
}
func (g *Game) handleMenuClick() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= SCREEN_WIDTH/2-80 && x <= SCREEN_WIDTH/2+80 && y >= SCREEN_HEIGHT/2-10 && y <= SCREEN_HEIGHT/2+10 {
			g.State = Simulation
			g.World = world.NewWorld(SCREEN_WIDTH/10, SCREEN_HEIGHT/10, 100)
			g.Paused = true
			g.LiveCellCount = g.World.LiveCellCount()
		}
		if x >= SCREEN_WIDTH/2-80 && x <= SCREEN_WIDTH/2+80 && y >= SCREEN_HEIGHT/2+20 && y <= SCREEN_HEIGHT/2+60 {
			filename, err := zenity.SelectFile()
			if err != nil {
				if err == zenity.ErrCanceled {
					log.Println("File selection was canceled.")
					return
				}
				log.Fatalf("Failed to load world state: %v", err)
			}
			g.State = Simulation
			err = g.World.LoadFile(filename)
			if err != nil {
				log.Fatalf("Failed to load world state: %v", err)
			}
			g.Paused = true
			g.LiveCellCount = g.World.LiveCellCount()
		}
	}
}

func (g *Game) handleClick() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x < 0 || y < 0 || x >= SCREEN_WIDTH || y >= SCREEN_HEIGHT {
			return
		}
		g.World.Area[x/10][y/10].Alive = !g.World.Area[x/10][y/10].Alive
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Paused = !g.Paused
		g.LiveCellCount = g.World.LiveCellCount()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.World = world.Reset(g.World)
		g.Paused = true
		g.LiveCellCount = g.World.LiveCellCount()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		filename, err := zenity.SelectFile()
		if err != nil {
			if err == zenity.ErrCanceled {
				log.Println("File selection was canceled.")
				return
			}
			log.Fatalf("Failed to load world state: %v", err)
		}
		g.State = Simulation
		err = g.World.LoadFile(filename)
		if err != nil {
			log.Fatalf("Failed to load world state: %v", err)
		}
		g.Paused = true
		g.LiveCellCount = g.World.LiveCellCount()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		filename, err := zenity.SelectFileSave(
			zenity.Title("Save Simulation"),
			zenity.ConfirmOverwrite(),
		)
		if err != nil {
			log.Printf("Failed to open file save dialog: %v", err)
			return
		}
		err = g.World.SaveFile(filename)
		if err != nil {
			log.Printf("Failed to save world state: %v", err)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if g.Fps < 24 {
			g.Fps += 1
			ebiten.SetMaxTPS(g.Fps)
			g.Paused = false
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if g.Fps > 10 {
			g.Fps -= 1
			ebiten.SetMaxTPS(g.Fps)
		} else {
			g.Fps = 10
			g.Paused = true
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) Update(screen *ebiten.Image) error {
	switch g.State {
	case Menu:
		g.handleMenuClick()
	case Simulation:
		g.handleClick()
		if !g.Paused {
			g.World.Update()
			g.LiveCellCount = g.World.LiveCellCount()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.State {
	case Menu:
		g.drawMenu(screen)
	case Simulation:
		g.drawSimulation(screen)
	}
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	text.Draw(screen, "Nouvelle Simulation", basicfont.Face7x13, SCREEN_WIDTH/2-80, SCREEN_HEIGHT/2-40, color.White)
	text.Draw(screen, "Jouer", basicfont.Face7x13, SCREEN_WIDTH/2-80, SCREEN_HEIGHT/2, color.White)
	text.Draw(screen, "Charger", basicfont.Face7x13, SCREEN_WIDTH/2-80, SCREEN_HEIGHT/2+40, color.White)
}

func (g *Game) drawSimulation(screen *ebiten.Image) {
	for i := 0; i < len(g.World.Area); i++ {
		for j := 0; j < len(g.World.Area[i]); j++ {
			if g.World.Area[i][j].Alive {
				ebitenutil.DrawRect(screen, float64(i*10), float64(j*10), 10, 10, color.White)
			}
		}
	}

	if g.Paused {
		text.Draw(screen, "Pause", basicfont.Face7x13, 10, 20, color.White)
	}
	liveCellCountStr := fmt.Sprintf("Live Cells: %d", g.LiveCellCount)
	text.Draw(screen, liveCellCountStr, basicfont.Face7x13, 10, 40, color.White)
	fpsStr := fmt.Sprintf("FPS: %d", g.Fps)
	text.Draw(screen, fpsStr, basicfont.Face7x13, 10, 60, color.White)
	text.Draw(screen, "Space: Play/Pause", basicfont.Face7x13, 10, 80, color.White)
	text.Draw(screen, "R: Reset", basicfont.Face7x13, 10, 100, color.White)
	text.Draw(screen, "S: Save", basicfont.Face7x13, 10, 120, color.White)
	text.Draw(screen, "G: Load", basicfont.Face7x13, 10, 140, color.White)

}
