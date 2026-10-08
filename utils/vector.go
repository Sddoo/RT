package utils

import "math"

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

func Len(v Vector) float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func Sum(v1 Vector, v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func Prod(v Vector, factor float32) Vector {
	return Vector{v.X * factor, v.Y * factor, v.Z * factor}
}

func Devision(v Vector, divider float32) Vector {
	return Vector{v.X / divider, v.Y / divider, v.Z / divider}
}
