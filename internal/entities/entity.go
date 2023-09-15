package entities

import "github.com/hajimehoshi/ebiten"

type Entity interface {
	Update()

	Draw(screen *ebiten.Image)
}
