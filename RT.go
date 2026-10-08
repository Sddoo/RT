package main

import (
	"math"
	u "test/utils"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	CANVAS_WIDTH       = 1000
	CANVAS_HEIGHT      = 1000
	CANVAS_HALF_WIDTH  = CANVAS_WIDTH / 2
	CANVAS_HALF_HEIGHT = CANVAS_HEIGHT / 2

	VIEWPORT_WIDTH                 = 1.0
	VIEWPORT_HEIGHT                = 1.0
	VIEWPORT_DISTANCE              = 1.0
	CANVAS_VIEWPORT_WIDTH_SCALING  = VIEWPORT_WIDTH / CANVAS_WIDTH
	CANVAS_VIEWPORT_HEIGHT_SCALING = VIEWPORT_HEIGHT / CANVAS_HEIGHT
)

var BACKGROUND_COLOR = sdl.RGB888{0, 0, 0}

func getViewportCoordinates(canvasX float32, canvasY float32, camera Camera) u.Point {
	return u.Point{
		X: camera.Position.X + CANVAS_VIEWPORT_WIDTH_SCALING*canvasX,
		Y: camera.Position.Y + CANVAS_VIEWPORT_HEIGHT_SCALING*canvasY,
		Z: camera.Position.Z + VIEWPORT_DISTANCE,
	}
}

func intersectSphere(figure Figure, d u.Vector, cameraPosition u.Point) (t1 float32, t2 float32) {
	co := u.NewVector(figure.Position, cameraPosition)
	a := u.Dot(d, d)
	b := 2 * u.Dot(co, d)
	c := u.Dot(co, co) - figure.Dimensions.Radius*figure.Dimensions.Radius

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return
	}
	t1 = (-b + float32(math.Sqrt(float64(discriminant)))) / (2 * a)
	t2 = (-b - float32(math.Sqrt(float64(discriminant)))) / (2 * a)
	return t1, t2
}

func closestIntersection(figures []Figure, d u.Vector, o u.Point, tMin float32, tMax float32) (closestFigure Figure, closestT float32) {
	var t1, t2 float32
	closestT = float32(math.MaxFloat32)

	for _, figure := range figures {
		switch figure.Type {
		case "sphere":
			t1, t2 = intersectSphere(figure, d, o)
		}
		if t1 > tMin && t1 < tMax && t1 < closestT {
			closestT = t1
			closestFigure = figure
		}
		if t2 > tMin && t2 < tMax && t2 < closestT {
			closestT = t2
			closestFigure = figure
		}
	}

	return closestFigure, closestT
}

func computeLighting(mapConfig MapConfig, N u.Vector, P u.Point, V u.Vector, figure Figure) float32 {
	i := float32(0.0)

	for _, light := range mapConfig.Lights {
		switch light.Type {
		case "point":
			L := u.NewVector(P, light.Position)
			shadowSphere, shadowT := closestIntersection(mapConfig.Figures, P, L, 0.001, 1.0)

			// diffuse
			calcIntensity := light.Intensity * u.Dot(N, L) / (u.Len(N) * u.Len(L))
			if calcIntensity >= 0 {
				i += calcIntensity
			}

			// specular reflection
			if figure.Specular != -1 {
				R := u.Sum(u.Prod(u.Prod(N, 2), u.Dot(N, L)), u.Prod(L, -1))
				rDotV := u.Dot(R, V)
				if rDotV > 0 {
					i += light.Intensity * float32(math.Pow(float64(rDotV/(u.Len(R)*u.Len(V))), float64(figure.Specular)))
				}
			}
		case "ambient":
			i += light.Intensity
		}
	}

	return i
}

func rayTrace(mapConfig MapConfig, d u.Vector) sdl.RGB888 {
	closestFigure, closestT := closestIntersection(mapConfig.Figures, d, mapConfig.Camera.Position, VIEWPORT_DISTANCE, math.MaxFloat32)

	if closestT == float32(math.MaxFloat32) {
		return BACKGROUND_COLOR
	}

	P := u.Point(u.Sum(u.Vector(mapConfig.Camera.Position), u.Prod(d, closestT)))
	N := u.NewVector(closestFigure.Position, P)
	N = u.Devision(N, u.Len(N))
	i := computeLighting(mapConfig.Lights, N, P, u.Prod(d, -1), closestFigure)
	return u.RGBProduct(closestFigure.Color, i)
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
			d := u.NewVector(mapConfig.Camera.Position, viewportPoint)
			color := rayTrace(mapConfig, d)
			surface.Set(x+CANVAS_HALF_WIDTH, CANVAS_HEIGHT-1-(y+CANVAS_HALF_HEIGHT), color)
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
