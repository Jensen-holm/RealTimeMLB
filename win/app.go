package win

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
)

// Window -> This struct and files associated
// with it serve as a wrapper for g3n win stuff
type Window struct {
	App   *app.Application
	Scene *core.Node
	Cam   *camera.Camera
	Help  bool
}

func NewWindow() *Window {
	return &Window{
		App:   app.App(),
		Scene: core.NewNode(),
		Cam:   camera.New(1),
		Help:  false,
	}
}

func (w *Window) Add2Scene(node core.INode) {
	w.Scene.Add(node)
}

func (w *Window) ToggleHelp(h bool) {
	if h != w.Help {
		w.Help = h
	}
}
