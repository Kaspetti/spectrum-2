package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jakecoffman/cp"
	"github.com/kaspetti/spectrum2/internal/entities"
)


type Game struct{
    Space *cp.Space
    Ents []entities.Entity
    time float64
}


func (g *Game) Update(screen *ebiten.Image) error {
    for _, ent := range g.Ents {
        ent.Update()
    }

    timeStep := 1.0 / float64(ebiten.MaxTPS())
	g.time += timeStep
	g.Space.Step(timeStep)

	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
    for _, ent := range g.Ents {
        ent.Draw(screen)
    }
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return outsideWidth, outsideHeight
}


func NewGame() *Game {
    space := cp.NewSpace()
    body := space.AddBody(cp.NewKinematicBody())

    player, err := entities.NewPlayer(body, "assets/epicimage.png")
    if err != nil {
        panic(err)
    }

    return &Game {
        Space: space,
        Ents: []entities.Entity{
            player,
        },
    }
}


func main() {
    ebiten.SetWindowSize(1280, 720)
    ebiten.SetWindowTitle("Spectrum 2")

    game := NewGame()

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}
