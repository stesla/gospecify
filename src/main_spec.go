package main

import "./specify"

var spec specify.Specification;

func initSpec() {
	if spec == nil {
		spec = specify.New();
	}
}

func main() {
	runner := specify.DotRunner();
	spec.Run(runner);
}