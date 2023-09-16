package entities

// Script containing functionality for the player. Player satisfies the Entity interface
// as it implements the Update and Draw function. These are used in the game to render and
// run game logic for the player.

import (
	"errors"
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
}


// Updates the player each frame
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
    // we add drag to the players velocity and returns
    if direction.Length() == 0 {
        velocityAfterDrag := p.Body.Velocity().Mult(dragCoefficient)
        p.Body.SetVelocityVector(velocityAfterDrag)
        return
    }

    // Normalize the direction vector to make sure the player moves at a 
    // constant speed, even when moving diagonally
    newVelocity := p.Body.Velocity().Add(direction.Normalize().Mult(acceleration))

    // Make sure the new veoloctiy does not exceed the max speed for the player
    if newVelocity.Length() > maxSpeed {
        newVelocity = newVelocity.Normalize().Mult(maxSpeed)
    }

    p.Body.SetVelocityVector(newVelocity)
}


// Draws the player sprite to the screen according to the properties of
// the body of the player.
func (p *Player) Draw(screen *ebiten.Image) {
    options := ebiten.DrawImageOptions{}
    options.GeoM.Translate(p.Body.Position().X, p.Body.Position().Y)
    // TODO: Set sprite scale according to physics simulation

    screen.DrawImage(p.Sprite, &options)
}


// Creates a new player with the given body and image path. 
// Returns an error if the body is nil or if the sprite cannot 
// be loaded with the given image path. Use this function to make sure the 
// body and image of the player is properly instantiated.
func NewPlayer(body *cp.Body, imgPath string) (*Player, error) {
    if body == nil {
        return nil, errors.New("Provided body of player is nil")
    }

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

    return &Player{
        Body: body,
        Sprite: sprite,
    }, nil
}
