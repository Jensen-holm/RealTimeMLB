package bsbl

import (
	"github.com/Jensen-holm/RealTimeMLB/models/ball"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/experimental/physics"
	"github.com/g3n/engine/experimental/physics/object"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/math32"
)

var (
	gConst  = &math32.Vector3{Y: -0.98}
	gravity = physics.NewConstantForceField(gConst)
)

type Sim struct {
	self *physics.Simulation
}

func NewSim(scene *core.Node) *Sim {
	sim := new(Sim)
	sim.self = physics.NewSimulation(scene)
	sim.self.AddForceField(gravity)
	return sim
}

func (s *Sim) InitGravity(b *ball.Ball) {
	nb := object.NewBody(b.Mesh())
	s.self.AddBody(nb, "Baseball")

	nb.SetPosition(0, float32(b.R), 0)
	nb.SetVelocity(math32.NewVector3(-.5, 0, 0))
	nb.SetAngularVelocity(math32.NewVector3(0, 0, 1))
}

func (s *Sim) Pitch(
	velo,
	plateX,
	plateY float32,
	ball *graphic.Mesh,
) {

	return
}

func (s *Sim) Hit(
	ev,
	la float32,
	ball *graphic.Mesh,
) {

	return
}
