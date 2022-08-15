package pkg

type Stepper interface {
	StepInto(name string) Node
	StepOut() Node
}

type Validations interface {
	CannotBeEmpty(args ...interface{}) Node
	IsString(args ...interface{}) Node
	IsNumeric(args ...interface{}) Node
	IsArray(args ...interface{}) Node
	IsMap(args ...interface{}) Node
}

type Evaluator interface {
	Stepper
	Validations
}

type Annie interface {
	Close()
	Errors() []error
	Evaluator
}

type Node interface {
	Evaluator
}
