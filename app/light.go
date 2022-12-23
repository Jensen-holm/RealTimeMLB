package win

import (
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
)

func AddLight(w *Window) {
	// Adds directional front light
	dir1 := light.NewDirectional(&math32.Color{R: 1, G: 1, B: 1}, 0.6)
	dir1.SetPosition(100, 100, 100)
	w.Add2Scene(dir1)
}
