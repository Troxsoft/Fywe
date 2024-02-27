package pkg

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Rectangle struct {
	pos      Position
	size     Size
	color    RGBAColor
	id       string
	window   *Window
	sdl2Rect *sdl.Rect
}

func (rec *Rectangle) SetColor(color RGBAColor) {
	rec.color = color
}
func (rec *Rectangle) Color() RGBAColor {
	return rec.color
}
func (rec *Rectangle) Pos() Position {
	return rec.pos
}
func (rec *Rectangle) SetPos(newPos Position) {
	rec.pos = newPos
	rec.sdl2Rect.X = int32(newPos.X)
	rec.sdl2Rect.Y = int32(newPos.Y)
}

func (rec *Rectangle) Size() Size {
	return rec.size
}
func (rec *Rectangle) SetSize(newSize Size) {
	rec.size = newSize
	rec.sdl2Rect.W = int32(rec.size.W)
	rec.sdl2Rect.H = int32(rec.size.H)
}
func (rec *Rectangle) Window() *Window {
	return rec.window
}
func (rec *Rectangle) ID() string {
	return rec.id
}
func NewRectangle(id string, window *Window, pos Position, size Size, color RGBAColor) *Rectangle {
	return &Rectangle{
		id:     id,
		pos:    pos,
		size:   size,
		color:  color,
		window: window,
		sdl2Rect: &sdl.Rect{
			X: int32(pos.X),
			Y: int32(pos.Y),
			W: int32(size.W),
			H: int32(size.H),
		},
	}
}
func (rec *Rectangle) Draw() error {
	rec.window.Sdl2Render().SetDrawColor(rec.color.R, rec.color.G, rec.color.B, rec.color.A)
	err := rec.window.Sdl2Render().FillRect(rec.sdl2Rect)
	if err != nil {
		return err
	}
	return nil
}
