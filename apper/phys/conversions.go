package phys

const mileFt = 5280

// Kg2G -> converts kilograms to grams
func Kg2G(kg float32) float32 {
	return kg * 1000
}

// Mph2Fps -> converts miles per hour
// into feet per second
func Mph2Fps(mph float32) float32 {
	ftHr := mph * mileFt
	return (ftHr / (60)) / 60
}

// Cm2Ft -> used in the construction of the
// baseball, radius is known in centimeters
func Cm2Ft(cm float32) float32 {
	return (cm * 0.393701) / 12
}
