package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jakecoffman/cp"
	"github.com/kaspetti/spectrum-2/internal/entities"
	"github.com/kaspetti/spectrum-2/internal/objects"
)


const (
    debug = true
)


type Game struct{
    Space *cp.Space
    Ents []entities.Entity
    Objects []objects.Object
}


func (g *Game) Update(screen *ebiten.Image) error {
    for _, ent := range g.Ents {
        ent.Update()
    }
    for _, obj := range g.Objects {
        obj.Update()
    }

    timeStep := 1.0 / float64(ebiten.MaxTPS())
	g.Space.Step(timeStep)

	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
    for _, ent := range g.Ents {
        ent.Draw(screen)

        if debug {
            drawShape(screen, ent)
        }
    }

    for _, obj := range g.Objects {
        obj.Draw(screen)

        if debug {
            drawShape(screen, obj)
        }
    }
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return outsideWidth, outsideHeight
}


func newGame() *Game {
    space := cp.NewSpace()

    player, err := entities.NewPlayer(space, "assets/epicimage.png")
    if err != nil {
        panic(err)
    }

    object, err := objects.NewObject(space, "assets/testobject.png")
    if err != nil {
        panic(err)
    }

    return &Game {
        Space: space,
        Ents: []entities.Entity{
            player,
        },
        Objects: []objects.Object{
            object,
        },
    }
}


// TODO: Change entities and objects to inherit from same interface
func drawShape(screen *ebiten.Image, player entities.Entity) {
    x1 := player.GetBB().L
    x2 := player.GetBB().R
    y1 := player.GetBB().B
    y2 := player.GetBB().T

    ebitenutil.DrawLine(screen, x1, y1, x2, y1, color.RGBA{0, 255, 0, 255})
    ebitenutil.DrawLine(screen, x2, y1, x2, y2, color.RGBA{0, 255, 0, 255})
    ebitenutil.DrawLine(screen, x2, y2, x1, y2, color.RGBA{0, 255, 0, 255})
    ebitenutil.DrawLine(screen, x1, y2, x1, y1, color.RGBA{0, 255, 0, 255})
}


func main() {
    ebiten.SetWindowSize(1280, 720)
    ebiten.SetWindowTitle("Spectrum 2")

    game := newGame()

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}
