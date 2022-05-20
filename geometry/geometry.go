package geometry

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Area[T Number](length, width T) T {
	return length * width
}

func Diagonal(length, width float64) float64 {
	return math.Sqrt((length * length) * (width * width))
}