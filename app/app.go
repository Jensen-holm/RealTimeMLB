package win

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
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
		help:  false,
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
