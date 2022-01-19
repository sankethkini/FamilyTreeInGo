package application

import "errors"

var (
	NodeNotFoundErr     = errors.New("node not found")
	CyclicDependencyErr = errors.New("cyclic dependency")
	NodeExistsErr       = errors.New("node with same id")
)
