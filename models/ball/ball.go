package ball

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

var ballRad = Cm2Ft(3.65)

// Cm2Ft -> used to compute baseball dimensions
func Cm2Ft(cm float32) float64 {
	return float64(cm) / 30.48
}

// Create a new baseball to get flown thorough the air
// maybe instead of creating a new one each time we change
// the position of the ball after each play

type Ball struct {
	R    float64
	Geom *geometry.Geometry
	Mat  *material.Standard
	Mesh *graphic.Mesh
}

// NewBall -> after this is run, the sphere returned
// must be added to the windows scene object
func NewBall() *Ball {
	s := geometry.NewSphere(
		ballRad,
		100,
		100,
	)

	mat := material.NewStandard(
		math32.NewColor("slategray"),
	)
	sphereMesh := graphic.NewMesh(s, mat)
	return &Ball{
		R:    ballRad,
		Geom: s,
		Mat:  mat,
		Mesh: sphereMesh,
	}
}
