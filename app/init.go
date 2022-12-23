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

	// might want to add this later
	// Set up callback to update viewport and camera aspect ratio when the window is resized
	// want to make this its own function
	//onResize := func(evName string, ev interface{}) {
	//	// Get framebuffer size and update viewport accordingly
	//	width, height := w.app.GetSize()
	//	w.app.Gls().Viewport(0, 0, int32(width), int32(height))
	//	// Update the camera's aspect ratio
	//	w.cam.SetAspect(float32(width) / float32(height))
	//}
	//w.app.Subscribe(window.OnWindowSize, onResize)
	//onResize("", nil)

	w.app.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	if w.help {
		w.scene.Add(helper.NewAxes(0.5))
	}

	// add stuff
	//Ground(w)
	AddLight(w)

	err := AddObjs(w)
	if err != nil {
		log.Fatalf("error loading obj files: %v", err)
	}

	bsbl.Init(w.scene)

}

func AddObjs(w *Window) error {

	gs, err := obj.LoadAll()
	if err != nil {
		return fmt.Errorf("error loading test object")
	}

	for _, g := range gs {
		w.Add2Scene(g)
	}
	return nil
}
