package pkg

type Stepper interface {
	StepInto(name string) Node
	StepOut() Node
}

type Validations interface {
	CannotBeEmpty(node string) Node
	IsString(node string) Node
	IsNumeric(node string) Node
	IsArray(node string) Node
	IsMap(node string) Node
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
