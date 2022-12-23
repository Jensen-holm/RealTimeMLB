package win

import (
	"github.com/Jensen-holm/RealTimeMLB/rend"
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
	"time"
)

// Window -> This struct and files associated
// with it serve as a wrapper for g3n app stuff
type Window struct {
	app   *app.Application
	scene *core.Node
	cam   *camera.Camera
	help  bool
}

func NewWindow() *Window {
	return &Window{
		app:   app.App(),
		scene: core.NewNode(),
		cam:   camera.New(1),
		help:  true,
	}
}

func (w *Window) App() *app.Application {
	return w.app
}

func (w *Window) Scene() *core.Node {
	return w.scene
}

func (w *Window) Cam() *camera.Camera {
	return w.cam
}

func (w *Window) Add2Scene(node core.INode) {
	w.scene.Add(node)
}

func (w *Window) ToggleHelp(h bool) {
	if h != w.help {
		w.help = h
	}
}

// Init -> Necessary boilerplate to create
// a g3n application before running it
func (w *Window) Init() {

	gui.Manager().Set(w.scene)
	w.cam.SetPosition(0, 0, 2)
	w.scene.Add(w.cam)
	camera.NewOrbitControl(w.cam)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	// want to make this its own function
	onResize := func(evName string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := w.app.GetSize()
		w.app.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		w.cam.SetAspect(float32(width) / float32(height))
	}
	w.app.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	w.app.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	if w.help {
		w.scene.Add(helper.NewAxes(0.5))
	}

	rend.Ground(w)
}

func (w *Window) RunApp() error {
	var err error = nil

	w.app.Run(
		func(renderer *renderer.Renderer, deltaTime time.Duration) {
			w.app.Gls().Clear(
				gls.DEPTH_BUFFER_BIT |
					gls.STENCIL_BUFFER_BIT |
					gls.COLOR_BUFFER_BIT,
			)
			err = renderer.Render(w.scene, w.cam)
		},
	)

	return err
}
