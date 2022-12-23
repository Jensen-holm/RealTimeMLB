package win

import (
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
)

func AddLight(w *Window) {
	// Adds directional front light
	l1 := light.NewDirectional(
		math32.NewColor("white"),
		0.6,
	)
	l1.SetPosition(100, 100, 100)
	w.Add2Scene(l1)
}
