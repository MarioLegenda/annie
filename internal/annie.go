package internal

import (
	anniePkg "annie/pkg"
	"fmt"
	"strconv"
)

type internalAction interface {
	GetData() interface{}
	AddError(err string)
}

type annie struct {
	data   map[string]interface{}
	errors []error
}

func (a *annie) StepInto(name string) anniePkg.Node {
	if ok := valueEmpty(a.data[name]); ok {
		a.errors = append(a.errors, buildError(fmt.Sprintf("Cannot step into a node %s. Node is empty", name)))
	}

	d, ok := a.data[name].(map[string]interface{})

	if !ok {
		a.errors = append(a.errors, buildError(fmt.Sprintf("Cannot step into a node %s. Node is not an indexable type (map[string]interface{})", name)))
	}

	return &node{annie: a, data: d}
}

func (a *annie) StepOut() anniePkg.Node {
	return &node{}
}

func (a *annie) CannotBeEmpty(node string) anniePkg.Node {
	assignIfEmpty(a, node)

	return a
}

func (a *annie) IsString(node string) anniePkg.Node {
	_, ok := a.data[node].(string)

	if !ok {
		a.AddError(fmt.Sprintf("Node %s is not a string", node))
	}

	return a
}

func (a *annie) IsArray(node string) anniePkg.Node {
	_, ok := a.data[node].([]interface{})

	if !ok {
		a.AddError(fmt.Sprintf("Node %s is not an array", node))
	}

	return a
}

func (a *annie) IsNumeric(node string) anniePkg.Node {
	v, _ := a.data[node].(string)

	_, err := strconv.Atoi(v)

	if err != nil {
		a.AddError(fmt.Sprintf("Node %s is not a string", node))
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

func (a *annie) AddError(err string) {
	a.errors = append(a.errors, buildError(err))
}

func (a *annie) Errors() []error {
	return a.errors
}
