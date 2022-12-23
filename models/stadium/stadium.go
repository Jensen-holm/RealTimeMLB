package stadium

import (
	"github.com/Jensen-holm/RealTimeMLB/models"
	"github.com/g3n/engine/core"
)

func Soccer() *core.Node {
	g, _ := models.ReadObj(
		"stadium/Soccer.obj",
		"stadium/Soccer.mtl",
	)
	return g
}
