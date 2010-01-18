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

type runner struct {
	examples       *exampleCollection
	currentExample *complexExample
}

type afterBlock struct {
	f   AfterFunc
	loc Location
}

var emptyAfter = afterBlock{func(Context) {}, nil}
var emptyBefore = func(Example) {}

func makeRunner() *runner { return &runner{examples: makeExampleCollection()} }

func (self *runner) After(f AfterFunc) {
	if self.currentExample != nil {
		block := afterBlock{f, newBlockLocation()}
		self.currentExample.AddAfter(block)
	}
}

func (self *runner) Before(block BeforeBlock) {
	if self.currentExample != nil {
		self.currentExample.AddBefore(block)
	}
}

func (self *runner) Describe(name string, block ExampleGroupBlock) {
	self.examples.Add(makeComplexExample(name, block))
}

func (self *runner) It(name string, block ExampleBlock) {
	if self.currentExample != nil {
		loc := newBlockLocation()
		self.currentExample.Add(makeSimpleExample(self.currentExample, name, block, loc))
	}
}

func (self *runner) Run(reporter Reporter) {
	self.examples.Init(self)
	self.examples.Run(reporter, emptyBefore, emptyAfter)
	reporter.Finish()
}
