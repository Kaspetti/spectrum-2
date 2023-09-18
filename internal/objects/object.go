package objects

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jakecoffman/cp"
)

const (
    // Collision Types
    Trigger = 0
    Collider = 1
)

type Object interface {
    Update()

    Draw(*ebiten.Image)

    GetBody() *cp.Body 
}
