/*
Copyright (c) 2009-2010 Samuel Tesla <samuel.tesla@gmail.com>

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
	"fmt"
	"os"
	"specify"
	"strings"

	t "../src/testspecify"
)

func HavePassing(expected interface{}) reporterMatcher {
	return reporterMatcher{expected, func(r t.ReporterSummary) interface{} { return r.PassingCount() }}
}

func HavePending(expected interface{}) reporterMatcher {
	return reporterMatcher{expected, func(r t.ReporterSummary) interface{} { return r.PendingCount() }}
}

func HaveFailing(expected interface{}) reporterMatcher {
	return reporterMatcher{expected, func(r t.ReporterSummary) interface{} { return r.FailingCount() }}
}

func HaveErrors(expected interface{}) reporterMatcher {
	return reporterMatcher{expected, func(r t.ReporterSummary) interface{} { return r.ErrorCount() }}
}

func HaveFailureIncluding(s string) eachMatcher {
	return eachMatcher{s, "failing example", eachFailure, matchTitle}
}

func HavePendingIncluding(s string) eachMatcher {
	return eachMatcher{s, "pending example", eachPending, matchTitle}
}

func HaveErrorIncluding(s string) eachMatcher {
	return eachMatcher{s, "error", eachError, matchTitle}
}

func HaveFailureAt(loc string) eachMatcher {
	return eachMatcher{loc, "failure", eachFailure, matchLocation}
}

func HavePendingAt(loc string) eachMatcher {
	return eachMatcher{loc, "pending example", eachPending, matchLocation}
}

func HaveErrorAt(loc string) eachMatcher {
	return eachMatcher{loc, "error", eachError, matchLocation}
}

func toReporterSummary(value interface{}) (reporter t.ReporterSummary, err os.Error) {
	var ok bool
	if reporter, ok = value.(t.ReporterSummary); !ok {
		err = os.NewError("Not a t.ReporterSummary")
	}
	return
}

type reporterMatcher struct {
	expected   interface{}
	actualFunc func(t.ReporterSummary) interface{}
}

func (self reporterMatcher) Should(actual interface{}) (result os.Error) {
	if reporter, error := toReporterSummary(actual); error != nil {
		result = error
	} else {
		result = specify.Be(self.expected).Should(self.actualFunc(reporter))
	}
	return
}
func (self reporterMatcher) ShouldNot(actual interface{}) (result os.Error) {
	if reporter, error := toReporterSummary(actual); error != nil {
		result = error
	} else {
		result = specify.Be(self.expected).ShouldNot(self.actualFunc(reporter))
	}
	return
}

type eachMatcher struct {
	s, message string
	each       func(t.ReporterSummary) <-chan t.Report
	f          func(t.Report, string) bool
}

func (self eachMatcher) match(r t.ReporterSummary) bool {
	for report := range self.each(r) {
		if self.f(report, self.s) {
			return true
		}
	}
	return false
}
func (self eachMatcher) Should(val interface{}) os.Error {
	if reporter, error := toReporterSummary(val); error != nil {
		return error
	} else {
		if !self.match(reporter) {
			return os.NewError(fmt.Sprintf("expected %v including `%v`", self.message, self.s))
		}
	}
	return nil
}
func (eachMatcher) ShouldNot(val interface{}) os.Error {
	return os.NewError("matcher not implemented")
}

func matchTitle(r t.Report, s string) bool { return strings.Count(r.Title(), s) > 0 }

func matchLocation(r t.Report, s string) bool {
	return strings.HasSuffix(r.Location().String(), s)
}

func eachFailure(r t.ReporterSummary) <-chan t.Report {
	return r.EachFailure()
}

func eachPending(r t.ReporterSummary) <-chan t.Report {
	return r.EachPending()
}

func eachError(r t.ReporterSummary) <-chan t.Report {
	return r.EachError()
}
