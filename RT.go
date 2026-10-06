package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	CANVAS_WIDTH                           = 1920
	CANVAS_HEIGHT                          = 1080
	VIEWPORT_WIDTH                         = 3840
	VIEWPORT_HEIGHT                        = 2160
	VIEWPORT_DISTANCE                      = 1
	CANVAS_VIEWPORT_WIDTH_SCALING  float32 = VIEWPORT_WIDTH / CANVAS_WIDTH
	CANVAS_VIEWPORT_HEIGHT_SCALING float32 = VIEWPORT_HEIGHT / CANVAS_HEIGHT
)

type Vector Point

func NewVector(start Point, end Point) Vector {
	return Vector{end.X - start.X, end.Y - start.Y, end.Z - start.Z}
}

func getViewportCoordinates(canvasX float32, canvasY float32, cameraZ float32) (viewportX float32, viewportY float32, viewportZ float32) {
	return CANVAS_VIEWPORT_WIDTH_SCALING * canvasX, CANVAS_VIEWPORT_HEIGHT_SCALING * canvasY, cameraZ + VIEWPORT_DISTANCE
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

	// surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	for x := range CANVAS_WIDTH {
		for y := range CANVAS_HEIGHT {
			newX, newY := getViewportCoordinates(float32(x), float32(y), mapConfig.Camera.Position.Z)
			v := NewVector(mapConfig.Camera.Position)
			fmt.Println(newX, newY)
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
