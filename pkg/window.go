package pkg

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var numWindows = 0

type Window struct {
	window               *sdl.Window
	render               *sdl.Renderer
	event                sdl.Event
	quit                 bool
	backcolor            RGBAColor
	isDestroyAutoOrOther bool
}

func (win *Window) Sdl2Event() sdl.Event {
	return win.event
}
func (win *Window) Sdl2Render() *sdl.Renderer {
	return win.render
}

func (win *Window) Sdl2Window() *sdl.Window {
	return win.window
}
func (win *Window) Destroy() {
	win.window.Destroy()
	win.quit = true
	numWindows -= 1
	if numWindows <= 0 {
		EndBackend()
		win.isDestroyAutoOrOther = true
	}
}
func (win *Window) BackColor() RGBAColor {
	return win.backcolor
}
func (win *Window) SetBackColor(color RGBAColor) {
	win.backcolor = color
}
func (win *Window) UpdateEvents() error {
	var err error

	// win.render.Destroy()
	// win.render, err = sdl.CreateRenderer(win.window, -1, 0)
	// if err != nil {
	// 	return err
	// }
	win.event = sdl.PollEvent()

	if win.event != nil {
		switch win.event.(type) {
		case *sdl.QuitEvent:
			win.quit = true
			win.isDestroyAutoOrOther = false
			break
		case *sdl.WindowEvent:

			eveWindow := win.event.(*sdl.WindowEvent)
			if eveWindow.Event == sdl.WINDOWEVENT_CLOSE {
				idWin, err := win.window.GetID()
				if err != nil {
					return err
				}
				if idWin == eveWindow.WindowID {

					win.quit = true

					win.isDestroyAutoOrOther = false
				}
			}
			break
		}
		// if win.event.GetType() == sdl.QUIT {

		// }
	}
	if err != nil {
		return err
	}
	if err = win.render.SetDrawColor(win.backcolor.R, win.backcolor.G, win.backcolor.B, win.backcolor.A); err != nil {

		return err
	}
	err = win.render.Clear()
	if err != nil {

		return err
	}
	return nil
}
func (win *Window) Close() {
	win.quit = true
}
func (win *Window) EndUpdate() {
	win.render.Present()
	if win.quit == true && win.isDestroyAutoOrOther == false {
		win.Destroy()
	}
	//sdl.Delay(16)

}
func (win *Window) IsOpen() bool {
	return !win.quit
}
func (win *Window) SetTitle(title string) {
	win.window.SetTitle(title)
}
func (win *Window) NewRectangle(id string, pos Position, size Size) *Rectangle {
	return NewRectangle(id, win, pos, size, NewColor(255, 0, 0))
}
func (win *Window) GetID() uint32 {
	g, _ := win.window.GetID()
	return g
}
func (win *Window) GetID2() (uint32, error) {
	return win.window.GetID()
}
func (win *Window) GetKeyRepeat(start uint8, end uint8) string {
	if win.event != nil {
		switch win.event.(type) {
		case *sdl.KeyboardEvent:
			keyEvent := win.event.(*sdl.KeyboardEvent)
			if keyEvent.State == sdl.PRESSED && keyEvent.Repeat > start && keyEvent.Repeat < end && keyEvent.WindowID == win.GetID() {
				key := keyEvent.Keysym.Sym
				return string(key)
			}
			break
		}
	}
	return ""
}
func (win *Window) GetKeyDown() string {
	if win.event != nil {
		switch win.event.(type) {
		case *sdl.KeyboardEvent:
			keyEvent := win.event.(*sdl.KeyboardEvent)
			if keyEvent.State == sdl.PRESSED && keyEvent.WindowID == win.GetID() {
				key := keyEvent.Keysym.Sym
				return string(key)
			}
			break
		}
	}
	return ""
}
func (win *Window) GetKeyPress() string {
	if win.event != nil {
		switch win.event.(type) {
		case *sdl.KeyboardEvent:
			keyEvent := win.event.(*sdl.KeyboardEvent)
			if keyEvent.State == sdl.PRESSED && keyEvent.Repeat == 0 && keyEvent.WindowID == win.GetID() {
				key := keyEvent.Keysym.Sym
				return string(key)
			}
			break
		}
	}
	return ""
}
func (win *Window) Pos() Position {
	x, y := win.window.GetPosition()
	return NewPos(int(x), int(y))
}
func (win *Window) SetPos(pos Position) {
	win.window.SetPosition(int32(pos.X), int32(pos.Y))
}
func (win *Window) Title() string {
	return win.window.GetTitle()
}
func NewWindow(title string, pos Position, size Size) (*Window, error) {
	if IsInitBackend() == false {
		err := InitBackend()
		if err != nil {
			return nil, err
		}
	}
	win, err := sdl.CreateWindow(title, int32(pos.X), int32(pos.Y), int32(size.W), int32(size.H), sdl.WINDOW_RESIZABLE)

	if err != nil {
		return nil, err
	}
	if win == nil {
		return nil, fmt.Errorf("Window is nil :(")
	}

	render, err := sdl.CreateRenderer(win, -1, 0)
	if err != nil {
		return nil, err
	}
	err = render.RenderSetVSync(true)

	if err != nil {
		return nil, err
	}
	numWindows += 1
	return &Window{
		window:               win,
		render:               render,
		event:                sdl.PollEvent(),
		quit:                 false,
		backcolor:            NewColor(0, 255, 0),
		isDestroyAutoOrOther: false,
	}, nil
}
