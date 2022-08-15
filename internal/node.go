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
		a.annie.errors = append(a.annie.errors, buildError(fmt.Sprintf("Cannot step into a node '%s'. Node is empty", name)))
	}

	d, ok := a.data[name].(map[string]interface{})

	if !ok {
		a.annie.errors = append(a.annie.errors, buildError(fmt.Sprintf("Cannot step into a node '%s'. Node is not an indexable type (map[string]interface{})", name)))

		return a
	}

	if len(d) == 0 {
		a.annie.errors = append(a.annie.errors, buildError(fmt.Sprintf("Cannot step into a node '%s'. Node is not an indexable type (map[string]interface{})", name)))

		return a
	}

	return &node{annie: a.annie, data: d, parent: a}
}

func (a *node) StepOut() anniePkg.Node {
	a.data = nil

	if a.parent == nil {
		return a.annie
	}

	p := &node{
		data:   a.parent.data,
		parent: a.parent.parent,
		annie:  a.parent.annie,
	}

	a.parent = nil

	return p
}

func (a *node) CannotBeEmpty(node string) anniePkg.Node {
	if a.data != nil {
		assignIfEmpty(a, node, "")
	}

	return a
}

func (a *node) IsMap(name string) anniePkg.Node {
	assignIfEmpty(a, name, "")
	isMap(a, name, "")

	return a
}

func (a *node) IsString(node string) anniePkg.Node {
	isString(a, node, "")

	return a
}

func (a *node) IsArray(node string) anniePkg.Node {
	isArray(a, node, "")

	return a
}

func (a *node) IsNumeric(node string) anniePkg.Node {
	isNumeric(a, node, "")

	return a
}

func (a *node) Equal(value interface{}) anniePkg.Node {
	return a
}
