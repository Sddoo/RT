package utils

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type Vector struct {
	X float32
	Y float32
	Z float32
}

func NewVector(start Point, end Point) Vector {
	return Vector{end.X - start.X, end.Y - start.Y, end.Z - start.Z}
}

func Dot(v1 Vector, v2 Vector) float32 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}
