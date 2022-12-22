package main

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
	"time"
)

func main() {

	// Create application and scene
	a := app.App()
	s := core.NewNode()

	// Create perspective camera
	cam := camera.New(1)
	cam.SetPosition(0, 0, 3)
	s.Add(cam)

	// Set up perspective camera
	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		err := renderer.Render(s, cam)
		if err != nil {
			panic(err)
		}
	})

}
