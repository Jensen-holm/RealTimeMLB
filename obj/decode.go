package obj

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/obj"
	"os"
)

func LoadAll(gs ...*core.Node) ([]*core.Node, error) {
	groups := make([]*core.Node, 0)
	for _, g := range gs {
		groups = append(groups, g)
	}
	return groups, nil
}

func DPath() string {
	d, _ := os.Getwd()
	return d + "/models/"
}

func ReadObj(objPath, mtlPath string) (*core.Node, error) {

	dec, err := obj.Decode(
		DPath()+objPath,
		DPath()+mtlPath,
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
