package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct {
    w, h int
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
    v := uint8(x + y)
    return color.RGBA{v, v, uint8(x), 255}
}

func main() {
    m := Image{w: 100, h: 100}
    pic.ShowImage(m)
}
