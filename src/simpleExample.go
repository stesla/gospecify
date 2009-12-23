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
	"fmt";
	"os";
)

type simpleExample struct {
	name	string;
	block	func(Example);
	fields	map[string]interface{};
	fail	chan os.Error;
}

func makeSimpleExample(name string, block func(Example)) *simpleExample {
	return &simpleExample{name, block, make(map[string]interface{}), make(chan os.Error)}
}

func (self *simpleExample) Run(reporter Reporter, before func(Example)) {
	if self.block == nil {
		reporter.Pending(self.name);
		return;
	}

	pass := make(chan bool);
	go func() {
		if before != nil {
			before(self)
		}

		self.block(self);
		pass <- true;
	}();

	select {
	case err := <-self.fail:
		reporter.Fail(os.NewError(fmt.Sprintf("%v - %v", self.name, err)))
	case <-pass:
		reporter.Pass()
	}
}

func (self *simpleExample) GetField(field string) (result interface{}) {
	result, _ = self.fields[field];
	return;
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
