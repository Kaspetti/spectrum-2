package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jakecoffman/cp"
	"github.com/kaspetti/spectrum-2/internal/objects"
	"github.com/kaspetti/spectrum-2/internal/objects/dynamic"
	"github.com/kaspetti/spectrum-2/internal/objects/static"
)


const (
    debug = true
)


type Game struct{
    Space *cp.Space
    Objects []objects.Object
}


func (g *Game) Update(screen *ebiten.Image) error {
    for _, obj := range g.Objects {
        obj.Update()
    }

    timeStep := 1.0 / float64(ebiten.MaxTPS())
	g.Space.Step(timeStep)

	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
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

    player, err := dynamic.NewPlayer(space, "assets/epicimage.png")
    if err != nil {
        panic(err)
    }

    object, err := static.NewObject(space, "assets/testobject.png")
    if err != nil {
        panic(err)
    }

    trigger := static.NewTrigger(
        space,
        cp.Vector{
            X: 100,
            Y: 100,
        },
        cp.Vector{
            X: 50,
            Y: 50,
        } ,
    )

    return &Game {
        Space: space,
        Objects: []objects.Object{
            player,
            object,
            trigger,
        },
        Trigger: 0,
        Collider: 1,
    }
}


// TODO: Change entities and objects to inherit from same interface
func drawShape(screen *ebiten.Image, object objects.Object) {
    object.GetBody().EachShape(func(s *cp.Shape) {
        bb := s.BB()

        x1 := bb.L
        x2 := bb.R
        y1 := bb.B
        y2 := bb.T

        ebitenutil.DrawLine(screen, x1, y1, x2, y1, color.RGBA{0, 255, 0, 255})
        ebitenutil.DrawLine(screen, x2, y1, x2, y2, color.RGBA{0, 255, 0, 255})
        ebitenutil.DrawLine(screen, x2, y2, x1, y2, color.RGBA{0, 255, 0, 255})
        ebitenutil.DrawLine(screen, x1, y2, x1, y1, color.RGBA{0, 255, 0, 255})
    })
}


func main() {
    ebiten.SetWindowSize(1280, 720)
    ebiten.SetWindowTitle("Spectrum 2")

    game := newGame()

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}
