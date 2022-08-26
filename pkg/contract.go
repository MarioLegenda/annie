package pkg

type Stepper interface {
	StepInto(name string) Node
	StepOut() Node
}

type Closer interface {
	Close()
}

type Errors interface {
	Errors() []error
}

type Conditional interface {
	If(name string, cond ...func(node Node) string)
}

type Validations interface {
	CannotBeEmpty(args ...interface{}) Node
	IsString(args ...interface{}) Node
	IsNumeric(args ...interface{}) Node
	IsArray(args ...interface{}) Node
	IsMap(args ...interface{}) Node
	Validate(name string, callback func(value interface{}) string) Node
}

type Evaluator interface {
	Stepper
	Validations
}

type Parent interface {
	Closer
	Errors
	Evaluator
	Conditional
}

type Node interface {
	Evaluator
	Conditional
}
