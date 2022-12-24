package win

import (
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
	"time"
)

func (w *Window) Rend(renderer *renderer.Renderer) error {
	w.App.Gls().Clear(
		gls.DEPTH_BUFFER_BIT |
			gls.STENCIL_BUFFER_BIT |
			gls.COLOR_BUFFER_BIT,
	)
	err := renderer.Render(w.Scene, w.Cam)
	return err
}

func (w *Window) RunApp() error {
	var err error = nil

	w.App.Run(
		func(renderer *renderer.Renderer, deltaTime time.Duration) {
			err = w.Rend(renderer)
		},
	)

	return err
}
