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

import (
	"os";
	"specify";
	t "../src/testspecify";
)

type TestingReporter interface {
	t.Reporter;
	FailingExamples() int;
	PassingExamples() int;
	PendingExamples() int;
}

func testRun(block func(t.Runner)) (reporter TestingReporter) {
	runner := t.NewRunner();
	runner.Describe("", func() { block(runner) });
	reporter = &testingReporter{};
	runner.Run(reporter);
	return;
}

type testingReporter struct {
	failing, passing, pending int;
}

func (self *testingReporter) Fail(err os.Error) {
	self.failing++
}
func (self *testingReporter) Finish()	{}
func (self *testingReporter) Pass()	{ self.passing++ }
func (self *testingReporter) Pending()	{ self.pending++ }

func (self *testingReporter) FailingExamples() int {
	return self.failing
}
func (self *testingReporter) PassingExamples() int {
	return self.passing
}
func (self *testingReporter) PendingExamples() int {
	return self.pending
}

func HavePassing(expected interface{}) specify.Matcher {
	return reporterMatcher{expected, func(r TestingReporter) interface{} { return r.PassingExamples() }}
}

func HavePending(expected interface{}) specify.Matcher {
	return reporterMatcher{expected, func(r TestingReporter) interface{} { return r.PendingExamples() }}
}

func HaveFailing(expected interface{}) specify.Matcher {
	return reporterMatcher{expected, func(r TestingReporter) interface{} { return r.FailingExamples() }}
}

type reporterMatcher struct {
	expected	interface{};
	actualFunc	func(TestingReporter) interface{};
}

func toTestingReporter(value interface{}) (reporter TestingReporter, err os.Error) {
	var ok bool;
	if reporter, ok = value.(TestingReporter); !ok {
		err = os.NewError("Not a TestingReporter")
	}
	return;
}

func (self reporterMatcher) Should(actual interface{}) (result os.Error) {
	if reporter, error := toTestingReporter(actual); error != nil {
		result = specify.Be(self.expected).Should(self.actualFunc(reporter))
	}
	return;
}
func (self reporterMatcher) ShouldNot(actual interface{}) (result os.Error) {
	if reporter, error := toTestingReporter(actual); error != nil {
		result = specify.Be(self.expected).ShouldNot(self.actualFunc(reporter))
	}
	return;
}
