package main

import (
	"runtime"

	"github.com/Troxsoft/Fywe/pkg"
)

func init() {
	runtime.LockOSThread()
}
func main() {
	win2, _ := pkg.NewWindow("hola mundo 2", pkg.NewPos(100, 100), pkg.NewSize(300, 300))
	win, _ := pkg.NewWindow("hola mundo", pkg.NewPos(100, 100), pkg.NewSize(300, 300))
	win2.SetBackColor(pkg.NewColor(40, 40, 40))
	rectangle := win.NewRectangle("1", pkg.NewPos(30, 30), pkg.NewSize(100, 100))
	rectangle2 := win2.NewRectangle("2", pkg.NewPos(30, 30), pkg.NewSize(100, 100))

	for win.IsOpen() && win2.IsOpen() {
		win2.UpdateEvents()
		rectangle2.Draw()

		win2.EndUpdate()
		win.UpdateEvents()

		rectangle.Draw()
		win.EndUpdate()

	}

}
