// entities contain functionality for all entities in the game. This includes 
// both player and other entities. All entities implement the Entity interface
// which require an Update and a Draw function to be implemented. This is later 
// called in the ebiten Update and Draw functions.
package entities

import "github.com/hajimehoshi/ebiten"

type Entity interface {
	Update()

	Draw(screen *ebiten.Image)
}
