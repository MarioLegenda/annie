package internal

import (
	anniePkg "annie/pkg"
	"bytes"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func NewAnnie(path string) (anniePkg.Annie, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, buildError(fmt.Sprintf("Provided path %s does not exist", path))
	}

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, buildError(fmt.Sprintf("An error occurred trying to read provided file %s: %s", path, err.Error()))
	}

	trimmed := bytes.Trim(data, " ")

	if len(trimmed) == 0 {
		return nil, buildError(fmt.Sprintf("Config file '%s' is empty", path))
	}

	var unwrap map[string]interface{}

	err = yaml.Unmarshal(trimmed, &unwrap)
	if err != nil {
		return nil, buildError(fmt.Sprintf("An error occurred trying to unmarshal provided yaml file %s: %s", path, err.Error()))
	}

	return newAnnie(unwrap), nil
}

func newAnnie(data map[string]interface{}) *annie {
	return &annie{data: data, errors: make([]error, 0)}
}

func buildError(err string) error {
	return errors.New(fmt.Sprintf("Annie: %s", err))
}
