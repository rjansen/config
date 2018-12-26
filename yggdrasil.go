package migi

import (
	"errors"

	"github.com/rjansen/yggdrasil"
)

var (
	ErrInvalidReference = errors.New("Invalid Options Reference")
	optionsPath         = yggdrasil.NewPath("/migi/options")
)

func Register(roots *yggdrasil.Roots, options Options) error {
	return roots.Register(optionsPath, options)
}

func Reference(tree yggdrasil.Tree) (Options, error) {
	reference, err := tree.Reference(optionsPath)
	if err != nil {
		return nil, err
	}
	if reference == nil {
		return nil, nil
	}
	options, is := reference.(Options)
	if !is {
		return nil, ErrInvalidReference
	}
	return options, nil
}

func MustReference(tree yggdrasil.Tree) Options {
	options, err := Reference(tree)
	if err != nil {
		panic(err)
	}
	return options
}
