package jtrans

import "strings"

// ValueType  the type of transform value
type ValueType string

const (
	Mapping  ValueType = "mapping"
	Constant ValueType = "constant"
)

// Map contains a transformation of one field value
type Map struct {
	Type ValueType
	From []string // depending on type a constant or mapping by key
	To   []string
}

// M provides a convenience variable and attached functions
var M = Map{}

// Build is a convenience method for building a map using namespace notation
func (Map) Build(vt ValueType, from, to string) *Map {
	return &Map{
		Type: vt,
		From: strings.Split(from, "."),
		To:   strings.Split(to, "."),
	}
}
