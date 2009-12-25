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
)

type dotReporter struct {
	*basicReporter;
}

func makeDotReporter() (result *dotReporter)	{ return &dotReporter{NewBasicReporter()} }

func (self *dotReporter) Error(r Report) {
	self.basicReporter.Error(r);
	fmt.Print("E");
}

func (self *dotReporter) Fail(r Report) {
	self.basicReporter.Fail(r);
	fmt.Print("F");
}

func printList(label string, reports <-chan Report) {
	fmt.Printf("\n%v:\n", label);
	for r := range reports {
		fmt.Printf("\n- %v - %v\n  %v\n", r.Title(), r.Error(), r.Location())
	}
}

func (self *dotReporter) Finish() {
	fmt.Printf("\nPassing: %v  Failing: %v  Pending: %v  Errors: %v\n", self.PassingCount(), self.FailingCount(), self.PendingCount(), self.ErrorCount());
	if self.ErrorCount() > 0 {
		printList("Errors", self.EachError())
	}
	if self.FailingCount() > 0 {
		printList("Failing Examples", self.EachFailure())
	}
	if self.PendingCount() > 0 {
		printList("Pending Examples", self.EachPending())
	}
}

func (self *dotReporter) Pass() {
	self.basicReporter.Pass();
	fmt.Print(".");
}

func (self *dotReporter) Pending(r Report) {
	self.basicReporter.Pending(r);
	fmt.Print("*");
}
