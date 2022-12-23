package obj

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/obj"
	"os"
)

func LoadAll() ([]*core.Node, error) {
	groups := make([]*core.Node, 0)

	g, err := TestObj()
	if err != nil {
		return nil, err
	}

	groups = append(groups, g)
	return groups, nil
}

func DPath() string {
	d, _ := os.Getwd()
	return d + "/data/"
}

func TestObj() (*core.Node, error) {

	dec, err := obj.Decode(
		DPath()+"untitled.obj",
		DPath()+"untitled.mtl",
	)
	if err != nil {
		return nil, err
	}

	group, err := dec.NewGroup()
	if err != nil {
		return nil, err
	}
	return group, nil
}
