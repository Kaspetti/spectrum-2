package entities
// Script containing functionality for the player. Player satisfies the Entity interface 
// as it implements the Update and Draw function. These are used in the game to render and 
// run game logic for the player.

import (
	"errors"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/jakecoffman/cp"
)


const (
    movementSpeed = 5
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

    if direction.Length() == 0 {
        return
    }

    // Normalize the direction vector to make sure the player moves at a 
    // constant speed, even when moving diagonally
    p.Body.SetPosition(p.Body.Position().Add(direction.Normalize().Mult(movementSpeed)))
}


// Draws the player sprite to the screen according to the properties of
// the body of the player.
func (p *Player) Draw(screen *ebiten.Image) {
    options := ebiten.DrawImageOptions{}
    options.GeoM.Translate(p.Body.Position().X, p.Body.Position().Y)
    // TODO: Set sprite scale according to physics simulation

    screen.DrawImage(p.Sprite, &options)
}


// Creates a new player with the given body and image(sprite). This function returns an 
// error if either the body or the image is 'nil'. Use this function to make sure the 
// body and image of the player is properly instantiated.
func NewPlayer(body *cp.Body, img *image.Image) (*Player, error) {
    if body == nil || img == nil {
        return nil, errors.New("body or img is nil")
    }

    sprite, err := ebiten.NewImageFromImage(*img, ebiten.FilterDefault)    
    if err != nil {
        return nil, err
    }

    return &Player{
        Body: body,
        Sprite: sprite,
    }, nil
}
