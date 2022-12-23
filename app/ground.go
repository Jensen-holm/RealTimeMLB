package win

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

func Ground(w *Window) {

	p := geometry.NewPlane(100, 100)
	gMat := material.NewStandard(&math32.Color{
		R: 100,
		G: 100,
		B: 100,
	})
	gMat.SetWireframe(false)
	gMat.SetSide(material.SideDouble)

	g := graphic.NewMesh(p, gMat)

	w.Add2Scene(g)

}
