package main

import(
	"os";
	"./specify";
)

var spec specify.Specification

func initSpec() {
	if spec == nil {
		spec = specify.New();
	}
}

type TestRunner interface {
	specify.Runner;
	PassCount() int;
}

type testRunner struct {
	failCount, passCount int;
}

func (self *testRunner) Fail(err os.Error) { self.failCount++; }
func (self *testRunner) Failed() bool { return self.failCount > 0; }
func (self *testRunner) Finish() {}
func (self *testRunner) Pass() { self.passCount++; }
func (self *testRunner) PassCount() int { return self.passCount; }
func (self *testRunner) Run(test specify.Test) {
	if pass,err := test(); pass {
		self.Pass();
	} else {
		self.Fail(err);
	}
}
