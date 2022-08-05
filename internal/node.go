package internal

import (
	anniePkg "annie/pkg"
	"fmt"
)

type node struct {
	data   map[string]interface{}
	parent *node
	annie  *annie
}

func (a *node) GetData() interface{} {
	return a.data
}

func (a *node) AddError(err string) {
	a.annie.errors = append(a.annie.errors, buildError(err))
}

func (a *node) StepInto(name string) anniePkg.Node {
	if ok := valueEmpty(a.data[name]); ok {
		a.annie.errors = append(a.annie.errors, buildError(fmt.Sprintf("Cannot step into a node %s. Node is empty", name)))
	}

	d, ok := a.data[name].(map[string]interface{})

	if !ok {
		a.annie.errors = append(a.annie.errors, buildError(fmt.Sprintf("Cannot step into a node %s. Node is not an indexable type (map[string]interface{})", name)))
	}

	return &node{annie: a.annie, data: d, parent: a}
}

func (a *node) StepOut() anniePkg.Node {
	return a.parent
}

func (a *node) CannotBeEmpty(node string) anniePkg.Node {
	assignIfEmpty(a.annie, node)

	return a
}

func (a *node) IsString(node string) anniePkg.Node {
	_, ok := a.data[node].(string)

	if !ok {
		a.AddError(fmt.Sprintf("Node %s is not a string", node))
	}

	return a
}

func (a *node) IsArray(node string) anniePkg.Node {
	_, ok := a.data[node].([]interface{})

	if !ok {
		a.AddError(fmt.Sprintf("Node %s is not an array", node))
	}

	return a
}

func (a *node) IsNumeric(node string) anniePkg.Node {
	_, ok := a.data[node].(string)

	if !ok {
		a.AddError(fmt.Sprintf("Node %s is not a string", node))
	}

	return a
}

func (a *node) Equal(value interface{}) anniePkg.Node {
	return a
}
