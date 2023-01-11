package baseball

import (
	"math"
)

// We need a function to somehow get the launch angle and the spray angle into a vector as well
// to finish setting up the apply force function for each time a hit is scraped. It is kind of hard
// to test anything while there are no games going on in the mlb right now.

// CalculateVector -> Chat gpt gave me this function, excited to see if it will work
func CalculateVector(la, ev, spray float64) (float32, float32, float32) {

	// convert from degrees to radians
	laRad := la * (math.Pi / 180)
	spRad := spray * (math.Pi / 180)

	x := float32(ev * math.Cos(laRad))
	y := float32(ev * math.Sin(laRad) * math.Cos(spRad))
	z := float32(ev * math.Sin(laRad) * math.Sin(spRad))

	return x, y, z
}
