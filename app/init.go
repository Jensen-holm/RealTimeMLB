package win

import (
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/util/helper"
)

// Init -> Necessary boilerplate to create
// a g3n application before running it
func (w *Window) Init() {

	gui.Manager().Set(w.scene)
	w.cam.SetPosition(0, 0, 2)
	w.scene.Add(w.cam)
	camera.NewOrbitControl(w.cam)

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

	// add the ground to the scene
	Ground(w)
}
