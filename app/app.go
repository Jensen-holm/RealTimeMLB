package app

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
	"time"
)

// App -> This struct and the code in this file
// serve as a wrapper for the g3n app framework
type App struct {
	self  *app.Application
	scene *core.Node
	cam   *camera.Camera
}

func NewApp() *App {
	return &App{
		self:  app.App(),
		scene: core.NewNode(),
		cam:   camera.New(1),
	}
}

func (a *App) Cam() *camera.Camera {
	return a.cam
}

func (a *App) Scene() *core.Node {
	return a.scene
}

func (a *App) Self() *app.Application {
	return a.self
}

func (a *App) RunApp() error {
	var err error = nil
	a.self.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.self.Gls().Clear(
			gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT,
		)
		err = renderer.Render(a.scene, a.cam)
	})
	return err
}
