package static

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jakecoffman/cp"
	"github.com/kaspetti/spectrum-2/internal/objects"
)


type TestTrigger struct {
    Body *cp.Body
}


func (t *TestTrigger) Draw(screen *ebiten.Image) {}

func (t *TestTrigger) Update() {}

func (t *TestTrigger) GetBody() *cp.Body {
    return t.Body
}


func TestHandler(ctA, ctB cp.CollisionType) *cp.CollisionHandler {
    return nil
}


func NewTrigger(space *cp.Space, position cp.Vector, dimensions cp.Vector) *TestTrigger {
    body := cp.NewStaticBody() 
    body.SetPosition(position)

    bb := cp.BB{
        L: -dimensions.X/2.0,
        R: dimensions.X/2.0,
        T: -dimensions.Y/2.0,
        B: dimensions.Y/2.0,
    }

    space.AddBody(body)
    shape := space.AddShape(cp.NewBox2(body, bb, 0))
    shape.SetSensor(true)

    space.NewCollisionHandler(objects.Trigger, objects.Collider)

    return &TestTrigger {
        Body: body,
    }
}

