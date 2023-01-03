package apper

import (
	"fmt"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
	"log"
	"time"
)

var totTime float32 = 0

func (a *App) Update(r *renderer.Renderer, dt time.Duration) {
	a.FrameRater.Start()

	a.A.Gls().Clear(
		gls.COLOR_BUFFER_BIT | gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT,
	)

	if int(totTime) > 3 {
		a.Cam.UpdatePos()
		if a.Sim != nil && int(dt) > 3 {
			a.Sim.Update(dt)
		}
	}

	err := r.Render(a.Scene, a.Cam.Self)
	if err != nil {
		log.Fatalf("error rendering the scene: %v", err)
	}

	totTime += float32(dt.Seconds())
	fmt.Println(totTime)
	a.FrameRater.Wait()
}

func (a *App) Run() {
	a.A.Run(a.Update)
}
