package main

import (
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/jakecoffman/cp"
	"github.com/kaspetti/spectrum2/internal/entities"
)


type Game struct{
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
    return 320, 240
}


func main() {
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Spectrum 2")
    game := &Game{}

    body := cp.NewKinematicBody()
    f, err := os.Open("assets/epicimage.png")
    if err != nil {
        panic(err)
    }
    img, err := png.Decode(f)
    if err != nil {
        panic(err)
    }

    player, err := entities.NewPlayer(body, &img)
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
