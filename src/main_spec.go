package main

import "specify";

var spec specify.Specification

func Be(value specify.Value) specify.Matcher {
	return specify.Be(value);
}

func Describe(name string, block func()) {
	initSpec();
	spec.Describe(name, block);
}

func It(name string, block func(specify.The)) {
	spec.It(name, block);
}

func initSpec() {
	if spec == nil {
		spec = specify.New();
	}
}

func main() {
	runner := specify.DotRunner();
	spec.Run(runner);
}