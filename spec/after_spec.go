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
	. "specify";
	t "../src/testspecify";
)

func init() {
	Describe("After", func() {
		It("should run the block after each test", func(the Example) {
			ch := make(chan bool, 1);
			testRun("", func(r t.Runner) {
				r.After(func(t.Example) { ch <- true });
				r.It("should pass", func(the t.Example) { /* pass */ });
			});
			_, ok := <-ch
			the.Value(ok).Should(Be(true));
		});

		It("should fail a test if the after block fails", func(the Example) {
			reporter := testRun("", func(r t.Runner) {
				r.After(func(the t.Example) { the.Value(1).Should(Be(2)) });
				r.It("should pass", func(the t.Example) { /* pass */ });
			});
			the.Value(reporter).Should(HaveFailing(1));
		});
	})
}
