package internal

import (
	anniePkg "annie/pkg"
	"fmt"
)

type internalAction interface {
	GetData() interface{}
	AddError(err string)
}

type parent struct {
	data   map[string]interface{}
	errors []error
}

func (a *parent) AddError(err string) {
	a.errors = append(a.errors, buildError(err))
}

func (a *parent) Errors() []error {
	return a.errors
}

func (a *parent) StepInto(name string) anniePkg.Node {
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

func (a *parent) StepOut() anniePkg.Node {
	return a
}

func (a *parent) CannotBeEmpty(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	assignIfEmpty(a, name, msg)

	return a
}

func (a *parent) IsMap(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	assignIfEmpty(a, name, msg)
	isMap(a, name, msg)

	return a
}

func (a *parent) IsString(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isString(a, name, msg)

	return a
}

func (a *parent) IsArray(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isArray(a, name, msg)

	return a
}

func (a *parent) IsNumeric(args ...interface{}) anniePkg.Node {
	name, msg, err := extractArgs(args)

	if err != nil {
		a.AddError(err.Error())

		return a
	}

	isNumeric(a, name, msg)

	return a
}

func (a *parent) Validate(name string, callback func(value interface{}) string) anniePkg.Node {
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

func (a *parent) If(name string, cond ...func(node anniePkg.Node) string) {
	if len(cond) == 0 {
		a.AddError(fmt.Sprintf("Invalid '%s' node 'If' method usage. 'If' method must have at least one condition function", name))

		return
	}

	d, ok := a.GetData().(map[string]interface{})
	if !ok {
		a.AddError(fmt.Sprintf("Invalid 's' node. 'If' method must be used with a map[string]interface{} type value, not promitive types", name))

		return
	}

	passed := false
	errs := make([]string, 0)
	for _, c := range cond {
		msg := c(newNode(a, nil, d))

		if msg != "" {
			errs = append(errs, msg)
		} else {
			passed = true
		}
	}

	if !passed {
		for _, err := range errs {
			a.AddError(err)
		}
	}
}

func (a *parent) Close() {
	a.data = nil
	a.errors = nil
}

func (a *parent) GetData() interface{} {
	return a.data
}
