package win

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

const rad90 float32 = 90 * (math32.Pi / 180)

func Ground(w *Window) {

	p := geometry.NewPlane(1, 1)

	// color not working currently
	gMat := material.NewStandard(&math32.Color{
		R: 7,
		G: 14,
		B: 48,
	})
	gMat.SetWireframe(false)
	gMat.SetSide(material.SideDouble)

	g := graphic.NewMesh(p, gMat)

	// this func takes in the angle in radians
	// 90* in radians is not normal num, so we'll convert
	g.RotateX(rad90)

	w.Add2Scene(g)

}
