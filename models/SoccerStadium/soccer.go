package stadium

import (
	"github.com/Jensen-holm/RealTimeMLB/apper/model"
	"github.com/g3n/engine/core"
)

func ImportStadium() (*core.Node, error) {
	return model.Read("Soccer.obj", "Soccer.mtl")
}
