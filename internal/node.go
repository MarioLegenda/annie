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
		a.AddError(fmt.Sprintf("Cannot step into a node '%s'. Node is empty. Subsequent node assertions will be made on the current node", name))
	}

	d, ok := a.data[name].(map[string]interface{})

	if !ok {
		a.AddError(fmt.Sprintf("Cannot step into a node '%s'. Node is not an indexable type (map[string]interface{}). Subsequent node assertions will be made on the current node", name))

		return a
	}

	if len(d) == 0 {
		a.AddError(fmt.Sprintf("Cannot step into a node '%s'. Node is not an indexable type (map[string]interface{}). Subsequent node assertions will be made on the current node", name))

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

func (a *node) CannotBeEmpty(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	assignIfEmpty(a, name, msg)

	return a
}

func (a *node) IsMap(args ...interface{}) anniePkg.Node {
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
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isString(a, name, msg)

	return a
}

func (a *node) IsArray(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isArray(a, name, msg)

	return a
}

func (a *node) IsNumeric(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isNumeric(a, name, msg)

	return a
}

func (a *node) Validate(name string, callback func(value interface{}) string) anniePkg.Node {
	d, ok := a.GetData().(map[string]interface{})

	if !ok {
		a.AddError(fmt.Sprintf("Could not validate '%s'. Current node must be a map[string]interface{} so value could not be extracted.", name))

		return a
	}

	msg := callback(d[name])

	if msg != "" {
		a.AddError(msg)
	}

	return a
}
