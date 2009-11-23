package main

import(
	"os";
	"specify";
)

type TestRunner interface {
	specify.Runner;
	FailCount() int;
	PassCount() int;
}

type testRunner struct {
	failCount, passCount int;
}

func makeTestRunner() (result *testRunner) {
	return &testRunner{};
}

func (self *testRunner) Fail(err os.Error) { self.failCount++; }
func (self *testRunner) FailCount() int { return self.failCount; }
func (self *testRunner) Finish() {}
func (self *testRunner) Pass() { self.passCount++; }
func (self *testRunner) PassCount() int { return self.passCount; }
func (self *testRunner) Run(test specify.Test) { specify.RunTest(test, self); }