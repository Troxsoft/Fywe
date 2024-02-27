package main

import (
	"fmt"
	"runtime"

	"github.com/Troxsoft/Fywe/pkg"
)

func init() {
	runtime.LockOSThread()
}
func main() {
	win, _ := pkg.NewWindow("hola mundo", pkg.NewPos(100, 100), pkg.NewSize(300, 300))
	rectangle := win.NewRectangle("1", pkg.NewPos(30, 30), pkg.NewSize(100, 100))
	rectangle.SetColor(pkg.ColorBrown)
	for win.IsOpen() {
		win.UpdateEvents()
		if win.GetKeyPress() != "" {

			fmt.Println("Press: ", win.GetKeyPress())
		}
		if win.GetKeyDown() != "" {

			fmt.Println("Down: ", win.GetKeyPress())
		}
		if win.GetKeyRepeat(0, 20) != "" {

			fmt.Println("Repeat 0,20: ", win.GetKeyPress())
		}
		rectangle.Draw()
		rectangle.SetPos(pkg.NewPos(rectangle.Pos().X+1, rectangle.Pos().Y))
		win.EndUpdate()
	}

}
