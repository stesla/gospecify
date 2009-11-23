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

import sp "specify"

var runner TestRunner;

func init() {
	Describe("Specification", func() {
		runner = makeTestRunner();

		Before(func () {
			s := sp.New();
			s.Describe("Foo", func() {
				s.It("pass", func(the sp.The) {
					the.Value(7 * 6).Should(Be(42));
					the.Value(1).ShouldNot(Be(2));
				});

				s.It("fail", func(the sp.The) {
					the.Value(7 * 6).ShouldNot(Be(42));
					the.Value(1).Should(Be(2));
				});
			});
			s.Run(runner);
		});

		It("indicates a passing test", func(the sp.The) {
			the.Value(runner.PassCount()).Should(Be(1));
		});

		It("indicates a failing test", func(the sp.The) {
			the.Value(runner.FailCount()).Should(Be(1));
		})
	});
}