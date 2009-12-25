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
	"fmt";
	"os";
	"specify";
	"strings";
)

func HavePassing(expected interface{}) reporterMatcher {
	return reporterMatcher{expected, func(r TestingReporter) interface{} { return r.PassingCount() }}
}

func HavePending(expected interface{}) reporterMatcher {
	return reporterMatcher{expected, func(r TestingReporter) interface{} { return r.PendingCount() }}
}

func HaveFailing(expected interface{}) reporterMatcher {
	return reporterMatcher{expected, func(r TestingReporter) interface{} { return r.FailingCount() }}
}

func HaveErrors(expected interface{}) reporterMatcher {
	return reporterMatcher{expected, func(r TestingReporter) interface{} { return r.ErrorCount() }}
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
		result = error
	} else {
		result = specify.Be(self.expected).Should(self.actualFunc(reporter))
	}
	return;
}
func (self reporterMatcher) ShouldNot(actual interface{}) (result os.Error) {
	if reporter, error := toTestingReporter(actual); error != nil {
		result = error
	} else {
		result = specify.Be(self.expected).ShouldNot(self.actualFunc(reporter))
	}
	return;
}

func HaveFailureIncluding(s string) haveFailureIncluding {
	return (haveFailureIncluding)(s)
}

type haveFailureIncluding string

func (s haveFailureIncluding) Match(r TestingReporter) bool {
	for report := range r.EachFailure() {
		if strings.Count(report.Title(), (string)(s)) > 0 {
			return true
		}
	}
	return false;
}
func (s haveFailureIncluding) Should(val interface{}) os.Error {
	if reporter, error := toTestingReporter(val); error != nil {
		return error
	} else {
		if !s.Match(reporter) {
			return os.NewError(fmt.Sprintf("expected failing example including `%v`", s))
		}
	}
	return nil;
}
func (haveFailureIncluding) ShouldNot(val interface{}) os.Error {
	return os.NewError("matcher not implemented")
}

func HavePendingIncluding(s string) havePendingIncluding {
	return (havePendingIncluding)(s)
}

type havePendingIncluding string

func (s havePendingIncluding) Match(r TestingReporter) bool {
	for report := range r.EachPending() {
		if strings.Count(report.Title(), (string)(s)) > 0 {
			return true
		}
	}
	return false;
}
func (s havePendingIncluding) Should(val interface{}) os.Error {
	if reporter, error := toTestingReporter(val); error != nil {
		return error
	} else {
		if !s.Match(reporter) {
			return os.NewError(fmt.Sprintf("expected pending example including `%v`", s))
		}
	}
	return nil;
}
func (havePendingIncluding) ShouldNot(val interface{}) os.Error {
	return os.NewError("matcher not implemented")
}

func HaveErrorIncluding(s string) haveErrorIncluding {
	return (haveErrorIncluding)(s)
}

type haveErrorIncluding string

func (s haveErrorIncluding) Match(r TestingReporter) bool {
	for report := range r.EachError() {
		if strings.Count(report.Title(), (string)(s)) > 0 {
			return true
		}
	}
	return false;
}
func (s haveErrorIncluding) Should(val interface{}) os.Error {
	if reporter, error := toTestingReporter(val); error != nil {
		return error
	} else {
		if !s.Match(reporter) {
			return os.NewError(fmt.Sprintf("expected error example including `%v`", s))
		}
	}
	return nil;
}
func (haveErrorIncluding) ShouldNot(val interface{}) os.Error {
	return os.NewError("matcher not implemented")
}

func HaveFailureAt(loc string) haveFailureAt	{ return (haveFailureAt)(loc) }

type haveFailureAt string

func (loc haveFailureAt) Match(r TestingReporter) bool {
	for report := range r.EachFailure() {
		if strings.HasSuffix(report.Location().String(), (string)(loc)) {
			return true
		}
	}
	return false;
}
func (loc haveFailureAt) Should(val interface{}) os.Error {
	if reporter, error := toTestingReporter(val); error != nil {
		return error
	} else {
		if !loc.Match(reporter) {
			return os.NewError(fmt.Sprintf("expected failure at `%v`", loc))
		}
	}
	return nil;
}
func (haveFailureAt) ShouldNot(val interface{}) os.Error {
	return os.NewError("matcher not implemented")
}

func HavePendingAt(loc string) havePendingAt	{ return (havePendingAt)(loc) }

type havePendingAt string

func (loc havePendingAt) Match(r TestingReporter) bool {
	for report := range r.EachPending() {
		if strings.HasSuffix(report.Location().String(), (string)(loc)) {
			return true
		}
	}
	return false;
}
func (loc havePendingAt) Should(val interface{}) os.Error {
	if reporter, error := toTestingReporter(val); error != nil {
		return error
	} else {
		if !loc.Match(reporter) {
			return os.NewError(fmt.Sprintf("expected failure at `%v`", loc))
		}
	}
	return nil;
}
func (havePendingAt) ShouldNot(val interface{}) os.Error {
	return os.NewError("matcher not implemented")
}

func HaveErrorAt(loc string) haveErrorAt	{ return (haveErrorAt)(loc) }

type haveErrorAt string

func (loc haveErrorAt) Match(r TestingReporter) bool {
	for report := range r.EachError() {
		if strings.HasSuffix(report.Location().String(), (string)(loc)) {
			return true
		}
	}
	return false;
}
func (loc haveErrorAt) Should(val interface{}) os.Error {
	if reporter, error := toTestingReporter(val); error != nil {
		return error
	} else {
		if !loc.Match(reporter) {
			return os.NewError(fmt.Sprintf("expected failure at `%v`", loc))
		}
	}
	return nil;
}
func (haveErrorAt) ShouldNot(val interface{}) os.Error {
	return os.NewError("matcher not implemented")
}
