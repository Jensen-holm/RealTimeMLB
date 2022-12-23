package win

import (
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
)

func AddLight(w *Window, posX, posY, posZ float32) {
	// Adds directional front light
	l1 := light.NewDirectional(
		math32.NewColor("white"),
		1,
	)
	l1.SetPosition(
		posX,
		posY,
		posZ,
	)
	w.Add2Scene(l1)
}
