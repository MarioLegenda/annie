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
		a.AddError(fmt.Sprintf("Cannot step into a node '%s'. Node is empty. Subsequent node assertions will be made on the current node", name))
	}

	d, ok := a.data[name].(map[string]interface{})

	if !ok {
		a.AddError(fmt.Sprintf("Cannot step into a node '%s'. Node is not an indexable type (map[string]interface{}). Subsequent node assertions will be made on the current node", name))

		return newNode(a, nil, nil)
	}

	if len(d) == 0 {
		a.AddError(fmt.Sprintf("Cannot step into a node '%s'. Node is not an indexable type (map[string]interface{}). Subsequent node assertions will be made on the current node", name))

		return newNode(a, nil, nil)
	}

	return newNode(a, nil, d)
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

func (a *annie) Validate(name string, callback func(value interface{}) string) anniePkg.Node {
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

func (a *annie) Close() {
	a.data = nil
	a.errors = nil
}

func (a *annie) GetData() interface{} {
	return a.data
}
