package application

import "errors"

var (
	ErrNodeNotFound     = errors.New("node not found")
	ErrCyclicDependency = errors.New("cyclic dependency")
	ErrNodeExists       = errors.New("node with same id")
)
