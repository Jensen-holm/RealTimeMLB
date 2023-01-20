package baseball

import (
	"math"
)

// CalculateVector ->
func CalculateVector(la, ev, spray float64) (float32, float32, float32) {

	// convert from degrees to radians
	laRad := la * (math.Pi / 180)
	spRad := spray * (math.Pi / 180)

	x := float32(ev * math.Cos(laRad))
	y := float32(ev * math.Sin(laRad) * math.Cos(spRad))
	z := float32(ev * math.Sin(laRad) * math.Sin(spRad))

	return x, y, z
}
