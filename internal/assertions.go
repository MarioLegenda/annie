package internal

import (
	"fmt"
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
