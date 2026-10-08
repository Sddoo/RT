package main

import (
	"encoding/json/v2"
	"log"
	"os"

	u "test/utils"

	"github.com/veandco/go-sdl2/sdl"
)

const MapPath = "map1.json"

type Dimensions struct {
	Height   float32 `json:"height,omitzero"`
	Width    float32 `json:"width,omitzero"`
	Length   float32 `json:"length,omitzero"`
	Radius   float32 `json:"radius,omitzero"`
	SideSize float32 `json:"sideSize,omitzero"`
}

type Camera struct {
	Position  u.Point `json:"position"`
	Direction u.Point `json:"direction"`
}

type Light struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Position  u.Point `json:"position"`
	Intensity float32 `json:"intensity"`
}

type Figure struct {
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Position   u.Point    `json:"position"`
	Color      sdl.RGB888 `json:"color"`
	Dimensions Dimensions `json:"dimensions"`
	Specular   int        `json:"specular"`
	Reflection float32    `json:"reflection"`
}

type MapConfig struct {
	Camera  Camera   `json:"camera"`
	Lights  []Light  `json:"lights"`
	Figures []Figure `json:"figures"`
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
