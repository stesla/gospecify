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

import "container/list"

type basicReporter struct {
	passing                  int
	failing, pending, errors *list.List
}

func NewBasicReporter() *basicReporter {
	return &basicReporter{failing: list.New(), pending: list.New(), errors: list.New()}
}

func (self *basicReporter) ErrorCount() int   { return self.errors.Len() }
func (self *basicReporter) FailingCount() int { return self.failing.Len() }
func (self *basicReporter) PassingCount() int { return self.passing }
func (self *basicReporter) PendingCount() int { return self.pending.Len() }
func (self *basicReporter) Error(r Report)    { self.errors.PushBack(r) }

func (self *basicReporter) EachError() <-chan Report {
	return eachReport(self.errors)
}
func (self *basicReporter) EachFailure() <-chan Report {
	return eachReport(self.failing)
}
func (self *basicReporter) EachPending() <-chan Report {
	return eachReport(self.pending)
}

func (self *basicReporter) Fail(r Report) { self.failing.PushBack(r) }

func (self *basicReporter) Finish() {}

func (self *basicReporter) Pass(r Report) { self.passing++ }

func (self *basicReporter) Pending(r Report) { self.pending.PushBack(r) }

func eachReport(l *list.List) <-chan Report {
	ch := make(chan Report, l.Len())
	for val := range l.Iter() {
		if name, ok := val.(Report); !ok {
			panic("typecast error")
		} else {
			ch <- name
		}
	}
	close(ch)
	return ch
}
