package internal

import (
	anniePkg "annie/pkg"
	"fmt"
)

type internalAction interface {
	GetData() interface{}
	AddError(err string)
}

type annie struct {
	data   map[string]interface{}
	errors []error
}

func (a *annie) AddError(err string) {
	a.errors = append(a.errors, buildError(err))
}

func (a *annie) Errors() []error {
	return a.errors
}

func (a *annie) StepInto(name string) anniePkg.Node {
	if ok := valueEmpty(a.data[name]); ok {
		a.AddError(fmt.Sprintf("Cannot step into a node '%s'. Node is empty", name))
	}

	d, ok := a.data[name].(map[string]interface{})

	if !ok {
		a.AddError(fmt.Sprintf("Cannot step into a node '%s'. Node is not an indexable type (map[string]interface{})", name))

		return &node{annie: a, data: nil, ignore: true}
	}

	if len(d) == 0 {
		a.AddError(fmt.Sprintf("Cannot step into a node '%s'. Node is not an indexable type (map[string]interface{})", name))

		return &node{annie: a, data: nil, ignore: true}
	}

	return &node{annie: a, data: d}
}

func (a *annie) StepOut() anniePkg.Node {
	return a
}

func (a *annie) CannotBeEmpty(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	assignIfEmpty(a, name, msg)

	return a
}

func (a *annie) IsMap(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	assignIfEmpty(a, name, msg)
	isMap(a, name, msg)

	return a
}

func (a *annie) IsString(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isString(a, name, msg)

	return a
}

func (a *annie) IsArray(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isArray(a, name, msg)

	return a
}

func (a *annie) IsNumeric(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isNumeric(a, name, msg)

	return a
}

func (a *annie) Close() {
	a.data = nil
	a.errors = nil
}

func (a *annie) GetData() interface{} {
	return a.data
}
