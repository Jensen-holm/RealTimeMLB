package obj

import "github.com/g3n/engine/core"

func Soccer() *core.Node {
	g, _ := ReadObj(
		"Soccer.obj",
		"Soccer.mtl",
	)
	return g
}
