# annie

Go yaml validator

Validates content of Yaml files.

Example Yaml file

```yaml
configuration:
  array: ['value1', 'value2', 'value3']
  simpleString: "value"
  simpleNumber: 5,
  arrayList:
    - value1
    - value2
    - value3
  complex:
    entryString: "string"
    entryNumber: 5
    entryArray: ['value1', 'value1', 5]
    entry:
      arrayList:
        - value1
        - value2
        - value3
options:
  - option1
  - option2
  - option3
```

Example usage based on yaml file:

```go

import (
	"annie"
	"log"
)

ann, err := internal.NewAnnie("test_complex_config.yml")

if err != nil {
	log.Fatal(err)
}

ann.StepInto("configuration").
	IsString("simpleString").
	IsNumeric("simpleNumber").
	StepInto("complex").
	StepInto("entry").
	IsArray("arrayList").
	StepOut().
	StepOut().
	IsArray("arrayList")

errs := ann.Errors()

if len(errs) != 0 {
	log.Fatal(errs)
}

ann.Close()
```
