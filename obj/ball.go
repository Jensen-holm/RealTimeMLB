package obj

import "github.com/g3n/engine/core"

func Ball() *core.Node {
	g, _ := ReadObj("baseball.obj", "baseball.mtl")
	return g
}
