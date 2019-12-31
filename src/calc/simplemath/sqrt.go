package simplemath

import "math"

func Sqrt(i int) int {
	r := math.Sqrt(float64(i))
	return int(r)
}
