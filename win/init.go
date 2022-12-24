package win

import (
	"fmt"
	"github.com/Jensen-holm/RealTimeMLB/bsbl"
	"github.com/Jensen-holm/RealTimeMLB/models"
	"github.com/Jensen-holm/RealTimeMLB/models/ball"
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
	//w.AddGround()
	b := ball.NewBall()
	w.Add2Scene(b.Mesh)

	gui.Manager().Set(w.Scene)
	w.AddLight(100, 100, 100, "white")
	w.AddLight(-100, 100, -100, "white")

	err := w.AddObjs()
	if err != nil {
		log.Fatalf("error loading obj files: %v", err)
	}

	sim := bsbl.NewSim(w.Scene)
	sim.InitGravity(b)

}

func (w *Window) AddHelp() {
	if w.Help {
		w.Add2Scene(helper.NewAxes(10))
	}
}

func (w *Window) AddBg() {
	w.App.Gls().ClearColor(.5, .75, 2, .5)
}

func (w *Window) AddCam() {
	w.Cam.SetPosition(0, 0, 2)
	w.Scene.Add(w.Cam)
	camera.NewOrbitControl(w.Cam)
}

func (w *Window) AddObjs() error {

	gs, err := models.LoadAll(
	// Going without the stadium for testing physics
	//stadium.Soccer(),
	)

	if err != nil {
		return fmt.Errorf("error loading test object")
	}

	for _, g := range gs {
		w.Add2Scene(g)
	}
	return nil
}
