package application

import "errors"

var NodeNotFound = errors.New("node not found")
var CyclicDependency = errors.New("cyclic dependency")
var NodeExists = errors.New("node with same id")
