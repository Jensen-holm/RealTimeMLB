package main

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
	"log"
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

	err := RunApp(a, cam, s)
	if err != nil {
		log.Fatalf("error running the app: %s", err)
	}
}

func RunApp(
	app *app.Application,
	camera *camera.Camera,
	scene *core.Node) error {

	var err error = nil
	app.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		app.Gls().Clear(
			gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT,
		)
		err = renderer.Render(scene, camera)
	})
	return err
}
