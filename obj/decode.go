package obj

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/obj"
	"os"
)

func LoadAll() ([]*core.Node, error) {
	groups := make([]*core.Node, 0)

	g, err := Ball()
	if err != nil {
		return nil, err
	}

	g1, err := Bat()
	if err != nil {
		return nil, err
	}

	g2, err := Soccer()
	if err != nil {
		return nil, err
	}

	groups = append(groups, g)
	groups = append(groups, g1)
	groups = append(groups, g2)
	return groups, nil
}

func DPath() string {
	d, _ := os.Getwd()
	return d + "/data/"
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

func Ball() (*core.Node, error) {
	g, err := ReadObj("baseball.obj", "baseball.mtl")
	if err != nil {
		return nil, err
	}

	return g, nil
}

func Bat() (*core.Node, error) {
	g, err := ReadObj("basebbat.obj", "basebbat.mtl")
	if err != nil {
		return nil, err
	}

	return g, nil
}

func Soccer() (*core.Node, error) {
	g, err := ReadObj("Soccer.obj", "Soccer.mtl")
	if err != nil {
		return nil, err
	}

	return g, nil
}
