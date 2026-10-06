package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	CANVAS_WIDTH       = 1920
	CANVAS_HEIGHT      = 1080
	CANVAS_HALF_WIDTH  = CANVAS_WIDTH / 2
	CANVAS_HALF_HEIGHT = CANVAS_HEIGHT / 2

	VIEWPORT_WIDTH                 = 3840
	VIEWPORT_HEIGHT                = 2160
	VIEWPORT_DISTANCE              = 1
	CANVAS_VIEWPORT_WIDTH_SCALING  = float32(VIEWPORT_WIDTH/CANVAS_WIDTH) / 1000
	CANVAS_VIEWPORT_HEIGHT_SCALING = float32(VIEWPORT_HEIGHT/CANVAS_HEIGHT) / 1000
)

type Vector Point

func NewVector(start Point, end Point) Vector {
	return Vector{end.X - start.X, end.Y - start.Y, end.Z - start.Z}
}

func Dot(v1 Vector, v2 Vector) float32 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func getViewportCoordinates(canvasX float32, canvasY float32, camera Camera) Point {
	return Point{
		camera.Position.X + CANVAS_VIEWPORT_WIDTH_SCALING*canvasX,
		camera.Position.Y + CANVAS_VIEWPORT_HEIGHT_SCALING*canvasY,
		camera.Position.Z + VIEWPORT_DISTANCE,
	}
}

func rayTrace(mapConfig MapConfig, d Vector) sdl.RGB888 {
	figures := mapConfig.Figures
	for _, figure := range figures {
		oc := NewVector(mapConfig.Camera.Position, figure.Position)
		a := Dot(d, d)
		b := 2 * Dot(d, oc)
		c := Dot(oc, oc) - figure.Dimensions.Radius*figure.Dimensions.Radius

		discriminant := b*b - 4*a*c
		if discriminant >= 0 {
			return figure.Color
		}
	}
	return sdl.RGB888{}
}

func RT(mapConfig MapConfig) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("RT", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, CANVAS_WIDTH, CANVAS_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	for x := -CANVAS_HALF_WIDTH; x < CANVAS_HALF_WIDTH; x++ {
		for y := -CANVAS_HALF_HEIGHT; y < CANVAS_HALF_HEIGHT; y++ {
			viewportPoint := getViewportCoordinates(float32(x), float32(y), mapConfig.Camera)
			d := NewVector(mapConfig.Camera.Position, viewportPoint)
			color := rayTrace(mapConfig, d)
			surface.Set(x+CANVAS_HALF_WIDTH, y+CANVAS_HALF_HEIGHT, color)
		}
	}
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
				println("Quit")
				running = false
			}
		}

		sdl.Delay(33)
	}

}
