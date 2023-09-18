package static

import (
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/jakecoffman/cp"
	"github.com/kaspetti/spectrum-2/internal/objects"
)


type TestObject struct {
    Body *cp.Body
    Sprite *ebiten.Image
}


func (to *TestObject) Update() {
}


func (to *TestObject) Draw(screen *ebiten.Image) {
    options := ebiten.DrawImageOptions{}

    width, height := to.Sprite.Size()
    x := to.Body.Position().X-float64(width)/2
    y := to.Body.Position().Y-float64(height)/2
    options.GeoM.Translate(x, y)

    screen.DrawImage(to.Sprite, &options)

}


func (to *TestObject) GetBody() *cp.Body {
    return to.Body
}


func NewObject(space *cp.Space, spritePath string) (*TestObject, error) {
    // Load sprite to image
    f, err := os.Open(spritePath)
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

    body := cp.NewStaticBody()
    body.SetPosition(cp.Vector{
        X: 250,
        Y: 250,
    })
    space.AddBody(body)

    shape := space.AddShape(
        cp.NewBox(body, float64(img.Bounds().Dx()), float64(img.Bounds().Dy()), 0),
    )
    shape.SetCollisionType(objects.Collider)


    return &TestObject{
        Body: body,
        Sprite: sprite,
    }, nil
}  
