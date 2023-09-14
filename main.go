package main

import (
	"image"
	"image/png"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
    speed = 5
)


type Game struct {
    player Player
}


type Player struct {
    sprite image.Image
    scale Vector2
    position Vector2
    facingRight bool
}

type Vector2 struct {
    x float64
    y float64
}


func (v *Vector2) Normalize() {
    l := math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))

    if l == 0 {
        v = &Vector2{}
    }

    v.x /= l
    v.y /= l
}

func (v1 Vector2) Add(v2 Vector2) Vector2 {
    return Vector2{
        x: v1.x + v2.x,
        y: v1.y + v2.y,
    }
}


func (v Vector2) Mul(s float64) Vector2 {
    return Vector2{
        x: v.x * s,
        y: v.y * s,
    }
}


func (g *Game) Update(screen *ebiten.Image) error {
    var direction Vector2
    if ebiten.IsKeyPressed(ebiten.KeyS) {
        direction.y++
    }
    if ebiten.IsKeyPressed(ebiten.KeyW) {
        direction.y--
    }
    if ebiten.IsKeyPressed(ebiten.KeyA) {
        direction.x--
    }
    if ebiten.IsKeyPressed(ebiten.KeyD) {
        direction.x++
    }
    
    direction.Normalize()
    if direction.x != 0 {
        g.player.facingRight = direction.x > 0
    }

    g.player.position = g.player.position.Add(direction.Mul(speed))

    return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
    eImage, err := ebiten.NewImageFromImage(g.player.sprite, ebiten.FilterDefault)
    if err != nil {
        panic(err)
    }

    options := ebiten.DrawImageOptions{}
    
    facingRight := 1.0
    if g.player.facingRight { 
        facingRight = -1.0
        if g.player.scale.x > 0 {

        }
    }


    options.GeoM.Scale(g.player.scale.x * facingRight, g.player.scale.y)
    options.GeoM.Translate(g.player.position.x, g.player.position.y)

    if err := screen.DrawImage(eImage, &options); err != nil {
        panic(err)
    }
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 320, 240
}


func main() {
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Hello, World!")

    game := &Game{}
    f, err := os.Open("epicimage.png")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    image, err := png.Decode(f)
    if err != nil {
        panic(err)
    }

    player := Player {
        sprite: image,
        scale: Vector2 {
            x: 0.2,
            y: 0.2,
        },
        position: Vector2{
            x: 0,
            y: 0,
        },
        facingRight: false,
    }

    game.player = player

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}
