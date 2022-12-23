package win

import (
	"fmt"
	"github.com/Jensen-holm/RealTimeMLB/bsbl"
	"github.com/Jensen-holm/RealTimeMLB/obj"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/util/helper"
	"log"
)

// Init -> Necessary boilerplate to create
// a g3n application before running it
func (w *Window) Init() {

	gui.Manager().Set(w.scene)
	w.cam.SetPosition(0, 0, 2)
	w.scene.Add(w.cam)
	camera.NewOrbitControl(w.cam)

	w.app.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	if w.help {
		w.scene.Add(helper.NewAxes(0.5))
	}

	// add stuff
	Ground(w)
	AddLight(
		w,
		100,
		100,
		100,
	)

	err := AddObjs(w)
	if err != nil {
		log.Fatalf("error loading obj files: %v", err)
	}

	bsbl.Init(w.scene)

}

func AddObjs(w *Window) error {

	gs, err := obj.LoadAll(
		obj.Ball(),
		obj.Soccer(),
	)

	if err != nil {
		return fmt.Errorf("error loading test object")
	}

	for _, g := range gs {
		w.Add2Scene(g)
	}
	return nil
}
