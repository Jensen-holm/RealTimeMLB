package win

import (
	"fmt"
	"github.com/Jensen-holm/RealTimeMLB/bsbl"
	"github.com/Jensen-holm/RealTimeMLB/models"
	"github.com/Jensen-holm/RealTimeMLB/models/stadium"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/util/helper"
	"log"
)

// Init -> Necessary boilerplate to create
// a g3n application before running it, and it
// adds everything to the scene as well
func (w *Window) Init() {

	w.AddCam()
	w.AddHelp()
	w.AddBg()
	w.AddGround()
	w.AddLight(100, 100, 100, "white")
	w.AddLight(-100, 100, -100, "white")
	w.AddShaders()
	// init the physics simulation
	bsbl.Init(w.scene)

	gui.Manager().Set(w.scene)

	err := w.AddObjs()
	if err != nil {
		log.Fatalf("error loading obj files: %v", err)
	}

}

func (w *Window) AddHelp() {
	if w.help {
		w.Add2Scene(helper.NewAxes(10))
	}
}

func (w *Window) AddBg() {
	w.app.Gls().ClearColor(.5, .75, 2, .5)
}

func (w *Window) AddCam() {
	w.cam.SetPosition(0, 0, 2)
	w.scene.Add(w.cam)
	camera.NewOrbitControl(w.cam)
}

func (w *Window) AddObjs() error {

	gs, err := models.LoadAll(
		stadium.Soccer(),
	)

	if err != nil {
		return fmt.Errorf("error loading test object")
	}

	for _, g := range gs {
		w.Add2Scene(g)
	}
	return nil
}
