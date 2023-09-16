package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jakecoffman/cp"
	"github.com/kaspetti/spectrum2/internal/entities"
)


type Game struct{
    space *cp.Space
    ents []entities.Entity
}


func (g *Game) Update(screen *ebiten.Image) error {
    for _, ent := range g.ents {
        ent.Update()
    }

    return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
    for _, ent := range g.ents {
        ent.Draw(screen)
    }
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return outsideWidth, outsideHeight
}


func main() {
    ebiten.SetWindowSize(1280, 720)
    ebiten.SetWindowTitle("Spectrum 2")
    game := &Game{}

    body := cp.NewKinematicBody()

    player, err := entities.NewPlayer(body, "assets/epicimage.png")
    if err != nil {
        panic(err)
    }
    game.ents = []entities.Entity{
        player,
    }

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}
