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
	passing		int;
	failures	*list.List;
	pending		*list.List;
}

func makeDotReporter() (result *dotReporter) {
	result = &dotReporter{};
	result.failures = list.New();
	result.pending = list.New();
	return;
}

func (self *dotReporter) Fail(err os.Error) {
	self.failures.PushBack(err);
	fmt.Print("F");
}

func printList(label string, l *list.List) {
	if l.Len() == 0 {
		return
	}
	fmt.Printf("\n%v:\n", label);
	for i := range l.Iter() {
		fmt.Println("-", i)
	}
}

func (self *dotReporter) Finish() {
	fmt.Printf("\nPassing: %v Failing: %v Pending: %v\n", self.passing, self.failures.Len(), self.pending.Len());
	printList("Failing Examples", self.failures);
	printList("Pending Examples", self.pending);
}

func (self *dotReporter) Pass() {
	self.passing++;
	fmt.Print(".");
}

func (self *dotReporter) Pending(name string) {
	self.pending.PushBack(name);
	fmt.Print("*");
}
