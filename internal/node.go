package internal

import (
	anniePkg "annie/pkg"
	"fmt"
)

type node struct {
	data   map[string]interface{}
	parent *node
	annie  *annie
	ignore bool
}

func (a *node) GetData() interface{} {
	return a.data
}

func (a *node) AddError(err string) {
	a.annie.errors = append(a.annie.errors, buildError(err))
}

func (a *node) StepInto(name string) anniePkg.Node {
	if a.ignore {
		return a
	}

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
	if a.ignore {
		return a
	}

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

func (a *node) CannotBeEmpty(args ...interface{}) anniePkg.Node {
	if a.ignore {
		return &node{annie: a.annie, data: nil, parent: nil}
	}

	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	assignIfEmpty(a, name, msg)

	return a
}

func (a *node) IsMap(args ...interface{}) anniePkg.Node {
	if a.ignore {
		return &node{annie: a.annie, data: nil, parent: nil}
	}

	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	assignIfEmpty(a, name, msg)
	isMap(a, name, msg)

	return a
}

func (a *node) IsString(args ...interface{}) anniePkg.Node {
	if a.ignore {
		return &node{annie: a.annie, data: nil, parent: nil}
	}

	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isString(a, name, msg)

	return a
}

func (a *node) IsArray(args ...interface{}) anniePkg.Node {
	if a.ignore {
		return &node{annie: a.annie, data: nil, parent: nil}
	}

	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isArray(a, name, msg)

	return a
}

func (a *node) IsNumeric(args ...interface{}) anniePkg.Node {
	if a.ignore {
		return &node{annie: a.annie, data: nil, parent: nil}
	}

	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isNumeric(a, name, msg)

	return a
}
