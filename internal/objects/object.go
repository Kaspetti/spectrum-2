package objects

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jakecoffman/cp"
)


type Object interface {
    Update()

    Draw(*ebiten.Image)

    GetBB() cp.BB
}
