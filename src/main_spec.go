package main

import "./specify"

var spec specify.Specification;

func initSpec() {
	if spec == nil {
		spec = specify.New();
	}
}

func main() {
	spec.Run();
}