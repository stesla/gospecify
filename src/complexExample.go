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

type complexExample struct {
	name        string
	afterBlock  afterBlock
	beforeBlock BeforeBlock
	block       ExampleGroupBlock
	*exampleCollection
}

func makeComplexExample(name string, block ExampleGroupBlock) *complexExample {
	return &complexExample{name, emptyAfter, emptyBefore, block, makeExampleCollection()}
}

func (self *complexExample) AddBefore(block BeforeBlock) {
	self.beforeBlock = block
}

func (self *complexExample) AddAfter(block afterBlock) {
	self.afterBlock = block
}

func (self *complexExample) Init() { self.block() }
func (self *complexExample) Run(reporter Reporter, _ BeforeBlock, _ afterBlock) {
	/* TODO: Nested describes get weird with before blocks */
	self.exampleCollection.Run(reporter, self.beforeBlock, self.afterBlock)
}
