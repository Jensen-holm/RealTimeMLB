package baseball

import (
	"github.com/Jensen-holm/g3n-Wrapper/apper/phys"
	"math"
)

// Spray2Vector -> This will convert the spray angle scraped
// from baseball savant and convert it to an x and y vector that
// we will be using in the 'apply-force' function to simulate the hit
func Spray2Vector(laDeg float32) (float32, float32) {
	rad := float64(phys.Deg2Rad(laDeg))
	x := math.Cos(rad)
	y := math.Sin(rad)
	return float32(x), float32(y)
}

// We need a function to somehow get the launch angle and the spray angle into a vector as well
// to finish setting up the apply force function for each time a hit is scraped. It is kind of hard
// to test anything while there are no games going on in the mlb right now.
