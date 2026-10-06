package main

import (
	"encoding/json/v2"
	"log"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const MapPath = "map1.json"

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type Dimensions struct {
	Height   float32 `json:"height,omitzero"`
	Width    float32 `json:"width,omitzero"`
	Length   float32 `json:"length,omitzero"`
	Radius   float32 `json:"radius,omitzero"`
	SideSize float32 `json:"sideSize,omitzero"`
}

type Camera struct {
	Position  Point `json:"position"`
	Direction Point `json:"direction"`
}

// type Room struct {
// 	Width  float32 `json:"width"`
// 	Height float32 `json:"height"`
// 	Length float32 `json:"length"`
// }

type Spotlight struct {
	Position Point `json:"position"`
}

type Figure struct {
	Name       string     `json:"name"`
	Position   Point      `json:"position"`
	Color      sdl.RGB888 `json:"color"`
	Dimensions Dimensions `json:"dimensions"`
}

type MapConfig struct {
	// Room      Room      `json:"room"`
	Camera    Camera    `json:"camera"`
	Spotlight Spotlight `json:"spotlight"`
	Figures   []Figure  `json:"figures"`
}

func NewMapConfig(filePath string) MapConfig {
	var mapConfig MapConfig

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error while reading map file: ", err)
	}

	if err := json.Unmarshal(data, &mapConfig); err != nil {
		log.Fatal("Error while unmarshalling map file data: ", err)
	}

	return mapConfig
}

func main() {
	mapConfig := NewMapConfig(MapPath)
	RT(mapConfig)
}
