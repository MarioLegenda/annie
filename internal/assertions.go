package internal

import (
	"errors"
	"fmt"
	"strconv"
)

func valueEmpty(v interface{}) bool {
	if v == nil {
		return true
	}

	if v == "" {
		return true
	}

	return false
}

func assignIfEmpty(node internalAction, name string, msg string) bool {
	d, ok := node.GetData().(map[string]interface{})
	if msg == "" {
		msg = fmt.Sprintf("Node '%s' cannot be empty", name)
	}

	if !ok {
		if valueEmpty(node.GetData()) {
			node.AddError(msg)

			return true
		}
	}

	if valueEmpty(d[name]) {
		node.AddError(msg)

		return true
	}

	return false
}

func isString(node internalAction, name string, msg string) {
	d, ok := node.GetData().(map[string]interface{})
	if msg == "" {
		msg = fmt.Sprintf("Node '%s' is not a string", name)
	}

	if !ok {
		node.AddError(msg)

		return
	}

	_, ok = d[name].(string)

	if !ok {
		node.AddError(msg)
	}
}

func isNumeric(node internalAction, name string, msg string) {
	d, ok := node.GetData().(map[string]interface{})
	if msg == "" {
		msg = fmt.Sprintf("Node '%s' is not a numeric value", name)
	}

	if !ok {
		node.AddError(msg)

		return
	}

	v, ok := d[name].(string)

	if ok {
		_, err := strconv.Atoi(v)

		if err != nil {
			node.AddError(msg)

			return
		}
	}

	_, ok = d[name].(int)

	if !ok {
		node.AddError(msg)
	}
}

func isArray(node internalAction, name string, msg string) {
	d, ok := node.GetData().(map[string]interface{})

	if !ok {
		node.AddError(msg)
	}

	if msg == "" {
		msg = fmt.Sprintf("Node '%s' is not an array", name)
	}

	_, ok = d[name].([]interface{})

	if !ok {
		node.AddError(msg)
	}
}

func isMap(node internalAction, name string, msg string) {
	d, ok := node.GetData().(map[string]interface{})
	if msg == "" {
		msg = fmt.Sprintf("Node '%s' is not a map", name)
	}

	if !ok {
		node.AddError(msg)

		return
	}

	v, ok := d[name].(map[string]interface{})

	if !ok {
		node.AddError(msg)

		return
	}

	if len(v) == 0 {
		node.AddError(msg)

		return
	}
}

func extractArgs(args []interface{}) (string, string, error) {
	var n, m string

	if len(args) == 0 {
		return "", "", errors.New("Invalid number of arguments. Function signature: func (string, string)")
	}

	if len(args) > 2 {
		return "", "", errors.New("Invalid number of arguments. Function signature: func (string, string)")
	}

	if len(args) == 1 {
		name, ok := args[0].(string)

		if !ok {
			return "", "", errors.New("Invalid number of arguments. Function signature func (string, string)")
		}

		n = name
	}

	if len(args) == 2 {
		name, ok := args[1].(string)

		if !ok {
			return "", "", errors.New("Invalid number of arguments. Function signature func (string, string)")
		}

		m = name
	}

	return n, m, nil
}
