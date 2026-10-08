package utils

import "github.com/veandco/go-sdl2/sdl"

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func RGBProduct(rgb sdl.RGB888, factor float32) sdl.RGB888 {
	return sdl.RGB888{
		R: byte(float32(rgb.R) * factor),
		G: byte(float32(rgb.G) * factor),
		B: byte(float32(rgb.B) * factor),
	}
}

func RGBSum(rgb1 sdl.RGB888, rgb2 sdl.RGB888) sdl.RGB888 {
	return sdl.RGB888{
		R: byte(rgb1.R + rgb2.R),
		G: byte(rgb1.G + rgb2.G),
		B: byte(rgb1.B + rgb2.B),
	}
}
