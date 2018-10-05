package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
	color         uint8
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{R: img.color * uint8(x) / 2, G: img.color * uint8(y) / 2, B: 255, A: 255}
}

func main() {
	m := Image{100, 100, 100}
	pic.ShowImage(m)
}
