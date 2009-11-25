/*
Copyright (c) 2009 Samuel Tesla <samuel.tesla@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import(
	"os";
	"specify";
	t "../src/testspecify";
)

type TestRunner interface {
	t.Runner;
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
func (self *testRunner) Run(test t.Test) { t.RunTest(test, self); }

var BePassing passingMatcher;

type passingMatcher int;

type runnerTest func(TestRunner) (bool, os.Error);

func withTestRunner(val specify.Value, test runnerTest) (bool, os.Error) {
	runner, ok := val.(TestRunner);
	if !ok { return false, os.NewError("not a TestRunner") }
	return test(runner);
}

func (passingMatcher) Should(val specify.Value) (bool, os.Error) {
	return withTestRunner(val, func(runner TestRunner) (pass bool, err os.Error) {
		if pass = runner.FailCount() == 0; !pass {
			err = os.NewError("expected runner to be passing, but it was not");
		}
		return;
	});
}

func (passingMatcher) ShouldNot(val specify.Value) (bool, os.Error) {
	return withTestRunner(val, func(runner TestRunner) (pass bool, err os.Error) {
		if pass = runner.FailCount() > 0; !pass {
			err = os.NewError("expected runner to be failing, but it was not");
		}
		return;
	});
}