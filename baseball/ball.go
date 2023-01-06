package baseball

import (
	"github.com/Jensen-holm/g3n-Wrapper/apper/model"
	"github.com/Jensen-holm/g3n-Wrapper/apper/phys"
)

// NewBaseball -> Creates a standard baseball with its
// proper mass (g), radius (cm). Location must be adjusted
// if necessary after creating the baseball and not
// inside this function
func NewBaseball() *model.Sphere {
	return model.NewSphere(
		0,
		1,
		0,

		// for visibility, the ball is 2x
		// the size of a normal baseball
		phys.Cm2Ft(3.65*2),

		// 148 g was too much, dividing it by
		// 10 gives more realistic results
		148/10,
		"white",
	)
}
