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
	"container/list";
	"fmt";
	"os";
)

type dotReporter struct {
	passing, failing, pending	int;
	failures			*list.List;
}

func makeDotReporter() Reporter	{ return &dotReporter{failures: list.New()} }

func (self *dotReporter) Fail(err os.Error) {
	self.failing++;
	self.failures.PushBack(err);
	fmt.Print("F");
}

func (self *dotReporter) Finish() {
	fmt.Println("");
	for i := range self.failures.Iter() {
		fmt.Println("-", i)
	}
	fmt.Printf("Passing: %v Failing: %v Pending: %v\n", self.passing, self.failing, self.pending);
}

func (self *dotReporter) Pass() {
	self.passing++;
	fmt.Print(".");
}

func (self *dotReporter) Pending() {
	self.pending++;
	fmt.Print("*");
}
