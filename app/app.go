package app

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"time"
)

// App -> This struct and the code in this file
// serve as a wrapper for the g3n app framework
type App struct {
	self  *app.Application
	scene *core.Node
	cam   *camera.Camera
	help  bool
}

func NewApp() *App {
	return &App{
		self:  app.App(),
		scene: core.NewNode(),
		cam:   camera.New(1),
		help:  true,
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

func (a *App) ToggleHelp(h bool) {
	a.help = h
}

func (a *App) AddLight(
	posX,
	posY,
	posZ float32,
	ambColor,
	ptColor *math32.Color,
	ambIntensity,
	ptIntensity float32) {

	// not sure if this ambient thing is good in this func
	a.scene.Add(light.NewAmbient(
		ambColor, ambIntensity,
	))
	ptLight := light.NewPoint(
		ptColor, ptIntensity,
	)
	ptLight.SetPosition(posX, posY, posZ)

	a.scene.Add(ptLight)
}

func (a *App) SetBG(R, G, B, A float32) {
	a.self.Gls().ClearColor(R, G, B, A)
}

// Prep -> Must be run right before RunApp
func (a *App) Prep() {
	camera.NewOrbitControl(a.cam)
	gui.Manager().Set(a.scene)

	if a.help {
		a.scene.Add(helper.NewAxes(0.5))
	}

	// Create a blue torus and add it to the scene
	geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	mat := material.NewStandard(math32.NewColor("DarkBlue"))
	mesh := graphic.NewMesh(geom, mat)
	a.scene.Add(mesh)
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
