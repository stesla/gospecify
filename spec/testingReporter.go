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
	"container/list";

	t "../src/testspecify";
)

type TestingReporter interface {
	t.Reporter;
	FailingCount() int;
	PassingCount() int;
	PendingCount() int;
	EachFailure() <-chan t.Report;
	EachPending() <-chan t.Report;
}

func testRun(name string, block func(t.Runner)) (reporter TestingReporter) {
	runner := t.NewRunner();
	runner.Describe(name, func() { block(runner) });
	reporter = newTestingReporter();
	runner.Run(reporter);
	return;
}

func newTestingReporter() *testingReporter {
	return &testingReporter{failing: list.New(), pending: list.New()}
}

type testingReporter struct {
	passing			int;
	failing, pending	*list.List;
}

func (self *testingReporter) Fail(r t.Report)	{ self.failing.PushBack(r) }
func (self *testingReporter) Finish()		{}
func (self *testingReporter) Pass()		{ self.passing++ }
func (self *testingReporter) Pending(r t.Report) {
	self.pending.PushBack(r)
}

func (self *testingReporter) FailingCount() int {
	return self.failing.Len()
}
func (self *testingReporter) PassingCount() int {
	return self.passing
}
func (self *testingReporter) PendingCount() int {
	return self.pending.Len()
}
func (self *testingReporter) EachFailure() <-chan t.Report {
	return eachReport(self.failing)
}
func (self *testingReporter) EachPending() <-chan t.Report {
	return eachReport(self.pending)
}

func eachReport(l *list.List) <-chan t.Report {
	ch := make(chan t.Report, l.Len());
	for val := range l.Iter() {
		if name, ok := val.(t.Report); !ok {
			panic("typecast error")
		} else {
			ch <- name
		}
	}
	close(ch);
	return ch;
}
