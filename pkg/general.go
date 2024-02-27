package pkg

import (
	"github.com/veandco/go-sdl2/sdl"
)

var isInit = false

type Position struct {
	X int
	Y int
}

func NewPos(x, y int) Position {
	return Position{X: x, Y: y}
}

type Size struct {
	W uint
	H uint
}
type RGBAColor struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewColor(r, g, b uint8) RGBAColor {
	return RGBAColor{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}
func NewColorApha(r, g, b, a uint8) RGBAColor {
	return RGBAColor{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
func NewSize(w, h uint) Size {
	return Size{W: w, H: h}
}
func IsInitBackend() bool {
	return isInit
}
func EndBackend() {
	isInit = false
	sdl.Quit()
}
func InitBackend() error {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return err
	}
	isInit = true
	return nil
}

var (
	ColorWhite         = NewColor(255, 255, 255) // 255 255 255
	ColorBlack         = NewColor(0, 0, 0)       //0 0 0
	ColorRed           = NewColor(255, 0, 0)     // 255 0 0
	ColorGreen         = NewColor(0, 255, 0)     //0 255 0
	ColorBlue          = NewColor(0, 0, 255)     // 0 0 255
	ColorYellow        = NewColor(255, 255, 0)   // 255 255 0
	ColorMagenta       = NewColor(255, 0, 255)   //255 0 255
	ColorCyan          = NewColor(0, 255, 255)   //0 255 255
	ColorGray          = NewColor(128, 128, 128) //128 128 128
	ColorDarkGray      = NewColor(64, 64, 64)
	ColorLightGray     = NewColor(192, 192, 192)
	ColorBrown         = NewColor(165, 42, 42)
	ColorLightGreen    = NewColor(144, 238, 144)
	ColorOlive         = NewColor(128, 128, 0)
	ColorSkyBlue       = NewColor(135, 206, 235)
	ColorSteelBlue     = NewColor(70, 130, 180)
	ColorNavy          = NewColor(0, 0, 128)
	ColorSlateBlue     = NewColor(106, 90, 205)
	ColorLightBlue     = NewColor(173, 216, 230)
	ColorMidnightBlue  = NewColor(25, 25, 112)
	ColorPurple        = NewColor(128, 0, 128)
	ColorViolet        = NewColor(238, 130, 238)
	ColorOrange        = NewColor(255, 165, 0)
	ColorCoral         = NewColor(255, 127, 80)
	ColorGold          = NewColor(255, 215, 0)
	ColorSilver        = NewColor(192, 192, 192)
	ColorBronze        = NewColor(205, 127, 50)
	ColorPink          = NewColor(255, 192, 203)
	ColorSalmon        = NewColor(250, 128, 114)
	ColorTurquoise     = NewColor(64, 224, 208)
	ColorLimeGreen     = NewColor(0, 255, 0)
	ColorMintGreen     = NewColor(189, 252, 201)
	ColorDarkOlive     = NewColor(85, 107, 47)
	ColorRoseGold      = NewColor(255, 192, 203)
	ColorBeige         = NewColor(245, 245, 220)
	ColorIndigo        = NewColor(75, 0, 130)
	ColorSeaGreen      = NewColor(46, 139, 87)
	ColorChocolate     = NewColor(210, 105, 30)
	ColorVioletBlue    = NewColor(138, 43, 226)
	ColorMarineGreen   = NewColor(0, 128, 128)
	ColorDarkLime      = NewColor(50, 205, 50)
	ColorRoyalBlue     = NewColor(65, 105, 225)
	ColorFuchsia       = NewColor(255, 0, 255)
	ColorLemon         = NewColor(253, 233, 16)
	ColorPearlGray     = NewColor(230, 230, 250)
	ColorIvory         = NewColor(255, 255, 240)
	ColorAlmond        = NewColor(239, 222, 205)
	ColorDarkSlateGray = NewColor(47, 79, 79)
	ColorOldGold       = NewColor(207, 181, 59)
	ColorMint          = NewColor(152, 255, 152)
)
