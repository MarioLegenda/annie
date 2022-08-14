package internal

import (
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

func assignIfEmpty(node internalAction, name string) bool {
	d, ok := node.GetData().(map[string]interface{})

	if !ok {
		if valueEmpty(node.GetData()) {
			node.AddError(fmt.Sprintf("Node '%s' cannot be empty", name))

			return true
		}
	}

	if valueEmpty(d[name]) {
		node.AddError(fmt.Sprintf("Node '%s' cannot be empty", name))

		return true
	}

	return false
}

func isString(node internalAction, name string) {
	d, ok := node.GetData().(map[string]interface{})

	if !ok {
		node.AddError(fmt.Sprintf("Node '%s' is not a string", name))

		return
	}

	_, ok = d[name].(string)

	if !ok {
		node.AddError(fmt.Sprintf("Node '%s' is not a string", name))
	}
}

func isNumeric(node internalAction, name string) {
	d, ok := node.GetData().(map[string]interface{})

	if !ok {
		node.AddError(fmt.Sprintf("Node '%s' is not a numeric value", name))

		return
	}

	v, ok := d[name].(string)

	if ok {
		_, err := strconv.Atoi(v)

		if err != nil {
			node.AddError(fmt.Sprintf("Node '%s' is not a numeric value", name))

			return
		}
	}

	_, ok = d[name].(int)

	if !ok {
		node.AddError(fmt.Sprintf("Node '%s' is not a numeric value", name))
	}
}

func isArray(node internalAction, name string) {
	d, ok := node.GetData().(map[string]interface{})

	if !ok {
		node.AddError(fmt.Sprintf("Node '%s' is not an array", name))
	}

	_, ok = d[name].([]interface{})

	if !ok {
		node.AddError(fmt.Sprintf("Node '%s' is not an array", name))
	}
}

func isMap(node internalAction, name string) {
	d, ok := node.GetData().(map[string]interface{})

	if !ok {
		node.AddError(fmt.Sprintf("Node '%s' is not a map", name))

		return
	}

	v, ok := d[name].(map[string]interface{})

	if !ok {
		node.AddError(fmt.Sprintf("Node '%s' is not a map", name))

		return
	}

	if len(v) == 0 {
		node.AddError(fmt.Sprintf("Node '%s' is not a map", name))

		return
	}
}
