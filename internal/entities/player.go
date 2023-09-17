package entities

// Script containing functionality for the player. Player satisfies the Entity interface
// as it implements the Update and Draw function. These are used in the game to render and
// run game logic for the player.

import (
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/jakecoffman/cp"
)


const (
    acceleration = 100
    maxSpeed = 500
    dragCoefficient = 0.85
)


type Player struct {
    Body *cp.Body
    Sprite *ebiten.Image
    Shape *cp.Shape
}


// Update updates the player each frame
func (p *Player) Update() {
    direction := cp.Vector{}

    if (ebiten.IsKeyPressed(ebiten.KeyW)) {
        direction.Y--
    }
    if (ebiten.IsKeyPressed(ebiten.KeyS)) {
        direction.Y++
    }
    if (ebiten.IsKeyPressed(ebiten.KeyA)) {
        direction.X--
    }
    if (ebiten.IsKeyPressed(ebiten.KeyD)) {
        direction.X++
    }

    // If the direction vector is empty (no inputs from the player)
    // we add drag to the players velocity
    if direction.Length() == 0 {
        velocityAfterDrag := p.Body.Velocity().Mult(dragCoefficient)
        p.Body.SetVelocityVector(velocityAfterDrag)
    } else {
        // Normalize the direction vector to make sure the player moves at a 
        // constant speed, even when moving diagonally
        newVelocity := p.Body.Velocity().Add(direction.Normalize().Mult(acceleration))
        p.Body.SetVelocityVector(newVelocity)
    }


    // Make sure the velocity does not exceed the max speed for the player
    if p.Body.Velocity().Length() > maxSpeed {
        p.Body.SetVelocityVector(p.Body.Velocity().Normalize().Mult(maxSpeed))
    }

}


// Draw draws the player sprite to the screen according to the properties of
// the body of the player.
func (p *Player) Draw(screen *ebiten.Image) {
    options := ebiten.DrawImageOptions{}

    x := p.Shape.BB().L
    y := p.Shape.BB().B
    options.GeoM.Translate(x, y)
    // TODO: Set sprite scale according to physics simulation

    screen.DrawImage(p.Sprite, &options)
}


func (p *Player) GetBB() cp.BB {
    return p.Shape.BB()
}


// NewPlayer creates a new player with the given image path. 
// Returns an error if the sprite cannot 
// be loaded with the given image path. Use this function to make sure the 
// body and image of the player is properly instantiated.
func NewPlayer(space *cp.Space, imgPath string) (*Player, error) {
    // Load sprite to image
    f, err := os.Open(imgPath)
    if err != nil {
        return nil, err
    }
    defer f.Close()


    img, err := png.Decode(f)
    if err != nil {
        return nil, err
    }

    sprite, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)    
    if err != nil {
        return nil, err
    }

    body := space.AddBody(cp.NewKinematicBody())
    shape := space.AddShape(
        cp.NewBox(body, float64(img.Bounds().Dx()), float64(img.Bounds().Dy()), 0),
    )

    return &Player{
        Body: body,
        Sprite: sprite,
        Shape: shape,
    }, nil
}
