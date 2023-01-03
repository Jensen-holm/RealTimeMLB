package stadium

import (
	"github.com/Jensen-holm/g3n-App-Skeleton/apper/model"
	"github.com/g3n/engine/core"
	"os"
)

func ImportStadium() (*core.Node, error) {
	wd, _ := os.Getwd()
	return model.Read(
		wd+"/models/SoccerStadium/Soccer.obj",
		wd+"/models/SoccerStadium/Soccer.mtl",
	)
}
