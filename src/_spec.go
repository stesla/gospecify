package main

import(
	"os";
	"./specify";
)

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

type TestRunner interface {
	specify.Runner;
	FailCount() int;
	PassCount() int;
}

type testRunner struct {
	specify.Runner;
	failCount, passCount int;
}

func makeTestRunner() (result *testRunner) {
	result = &testRunner{};
	result.Runner = specify.BasicRunner();
	return;
}

func (self *testRunner) Fail(err os.Error) { self.failCount++; }
func (self *testRunner) FailCount() int { return self.failCount; }
func (self *testRunner) Pass() { self.passCount++; }
func (self *testRunner) PassCount() int { return self.passCount; }
