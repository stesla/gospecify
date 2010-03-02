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
package specify

import (
	"fmt"
	"os"
)

type simpleExample struct {
	parent *complexExample
	name   string
	block  ExampleBlock
	fields map[string]interface{}
	error  chan Report
	fail   chan Report
	loc    Location
}

func makeSimpleExample(parent *complexExample, name string, block ExampleBlock, loc Location) *simpleExample {
	return &simpleExample{parent, name, block, make(map[string]interface{}), make(chan Report), make(chan Report), loc}
}

func (self *simpleExample) Title() string {
	return fmt.Sprintf("%v %v", self.parent.name, self.name)
}

func (self *simpleExample) Run(reporter Reporter, before BeforeBlock, after afterBlock) {
	if self.block == nil {
		reporter.Pending(newReport(self.Title(), os.NewError("not implemented"), self.loc))
		return
	}

	pass := make(chan bool)
	go func() {
		if before != nil {
			before(self)
		}
		self.block(self)
		after.f(self)
		pass <- true
	}()

	select {
	case report := <-self.error:
		reporter.Error(report)
	case report := <-self.fail:
		reporter.Fail(report)
	case <-pass:
		reporter.Pass(newReport(self.Title(), nil, nil))
	}
}

func (self *simpleExample) Error(err os.Error) {
	self.error <- newReport(self.Title(), err, newErrorLocation())
}

func (self *simpleExample) GetField(field string) (result interface{}) {
	result, _ = self.fields[field]
	return
}

func (self *simpleExample) Field(field string) Assertion {
	return self.Value(self.GetField(field))
}

func (self *simpleExample) SetField(field string, value interface{}) {
	self.fields[field] = value
}

func (self *simpleExample) Value(value interface{}) Assertion {
	return assertion{self, value}
}
