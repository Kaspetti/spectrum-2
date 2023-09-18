package static

import (
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/jakecoffman/cp"
)


type TestObject struct {
    Body *cp.Body
    Shape *cp.Shape
    Sprite *ebiten.Image
}


func (to *TestObject) Update() {
}


func (to *TestObject) Draw(screen *ebiten.Image) {
    options := ebiten.DrawImageOptions{}

    x := to.Shape.BB().L
    y := to.Shape.BB().B
    options.GeoM.Translate(x, y)


    //fmt.Printf("x: %v | y: %v\n", x, y)
    // TODO: Set sprite scale according to physics simulation

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


    return &TestObject{
        Body: body,
        Sprite: sprite,
        Shape: shape,
    }, nil
}  
