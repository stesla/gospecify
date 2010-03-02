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
package specify

import (
	"fmt"
	"os"
)

type equalto interface {
	EqualTo(interface{}) bool
}

type equalityMatcher struct {
	expected interface{}
}

func newEqualityMatcher(value interface{}) equalityMatcher {
	return equalityMatcher{value}
}

func (self equalityMatcher) test(val interface{}) bool {
	switch t := self.expected.(type) {
	case equalto:
		return t.EqualTo(val)
	}
	return val == self.expected
}

func (self equalityMatcher) Should(actual interface{}) os.Error {
	if !self.test(actual) {
		return os.NewError(fmt.Sprintf("expected `%v` to be equal to `%v`", actual, self.expected))
	}
	return nil
}

func (self equalityMatcher) ShouldNot(actual interface{}) os.Error {
	if self.test(actual) {
		return os.NewError(fmt.Sprintf("expected `%v` not to be equal to `%v`", actual, self.expected))
	}
	return nil
}
