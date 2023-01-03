package baseball

import (
	phys2 "github.com/Jensen-holm/RealTimeMLB/apper/phys"
	"github.com/Jensen-holm/g3n-App-Skeleton/apper/model"
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
		phys2.Cm2Ft(3.65),
		148,
		"white",
	)
}
