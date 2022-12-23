package ball

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

// Create a new baseball to get flown thorough the air
// maybe instead of creating a new one each time we change
// the position of the ball after each play

// NewBall -> after this is run, the sphere returned
// must be added to the windows scene object
func NewBall() *graphic.Mesh {
	s := geometry.NewSphere(10, 16, 16)

	mat := material.NewStandard(
		math32.NewColor("red"),
	)

	// create a mesh for the sphere
	sphereMesh := graphic.NewMesh(s, mat)
	return sphereMesh
}
