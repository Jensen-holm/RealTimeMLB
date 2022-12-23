package win

import (
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
	"time"
)

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
