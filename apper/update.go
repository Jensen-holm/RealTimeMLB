package apper

import (
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
	"log"
	"time"
)

var totTime float32 = 0

func (a *App) Update(r *renderer.Renderer, deltaTime time.Duration) {
	a.FrameRater.Start()

	a.A.Gls().Clear(
		gls.COLOR_BUFFER_BIT | gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT,
	)

	a.Cam.UpdatePos()

	var buf = int(totTime) >= 5
	if a.Sim != nil && buf {
		a.Sim.Update(deltaTime)
	}

	err := r.Render(a.Scene, a.Cam.Self)
	if err != nil {
		log.Fatalf("error rendering the scene: %v", err)
	}

	totTime += float32(deltaTime.Seconds())
	a.FrameRater.Wait()
}

func (a *App) Run() {
	a.A.Run(a.Update)
}
