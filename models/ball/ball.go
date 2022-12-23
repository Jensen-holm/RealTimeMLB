package ball

import (
	"github.com/Jensen-holm/RealTimeMLB/models"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"log"
)

// Create a new baseball to get flown thorough the air
// maybe instead of creating a new one each time we change
// the position of the ball after each play

// NewBall -> after this is run, the sphere returned
// must be added to the windows scene object
func NewBall() *graphic.Mesh {
	s := geometry.NewSphere(10, 16, 16)

	m, err := texture.NewTexture2DFromImage(models.DPath() + "ball/red.jpg")
	if err != nil {
		log.Fatalf("error loading materiel for ball from image: %s", err)
	}

	var mat = material.NewStandard(math32.NewColor("red"))
	mat.AddTexture(m)

	// create a mesh for the sphere
	sphereMesh := graphic.NewMesh(s, mat)
	return sphereMesh
}
