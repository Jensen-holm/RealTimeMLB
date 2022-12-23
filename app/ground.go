package win

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

func Ground(w *Window) {

	p := geometry.NewPlane(.5, 1)
	gMat := material.NewStandard(&math32.Color{
		R: 1,
		G: 1,
		B: 1,
	})
	gMat.SetWireframe(false)
	gMat.SetSide(material.SideBack)

	g := graphic.NewMesh(p, gMat)

	w.Add2Scene(g)

}
