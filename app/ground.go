package win

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

func Deg2Rad(d float32) float32 {
	return d * (math32.Pi / 180)
}

func Ground(w *Window) {

	p := geometry.NewPlane(500, 500)

	gMat := material.NewStandard(
		math32.NewColor("mediumspringgreen"),
	)

	gMat.SetWireframe(false)
	gMat.SetSide(material.SideDouble)

	g := graphic.NewMesh(p, gMat)

	g.SetPosition(0, -.01, 0)

	// makes it flat
	g.RotateX(Deg2Rad(90))
	w.Add2Scene(g)

}
