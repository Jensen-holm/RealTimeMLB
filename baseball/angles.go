package baseball

import (
	"math"
)

// We need a function to somehow get the launch angle and the spray angle into a vector as well
// to finish setting up the apply force function for each time a hit is scraped. It is kind of hard
// to test anything while there are no games going on in the mlb right now.

// CalculateVector -> Chat gpt gave me this function, excited to see if it will work
func CalculateVector(launchAngle, exitVelocity, sprayAngle float64) (float32, float32, float32) {
	launchAngleRadians := launchAngle * (math.Pi / 180)
	sprayAngleRadians := sprayAngle * (math.Pi / 180)
	x := exitVelocity * math.Cos(launchAngleRadians)
	y := exitVelocity * math.Sin(launchAngleRadians) * math.Cos(sprayAngleRadians)
	z := exitVelocity * math.Sin(launchAngleRadians) * math.Sin(sprayAngleRadians)
	return float32(x), float32(y), float32(z)
}
