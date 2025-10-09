package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	W, H int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{
		R: v,
		G: v,
		B: 255,
		A: 255,
	}
}
func main() {
	m := Image{
		W: 256,
		H: 256,
	}
	pic.ShowImage(m)
}
