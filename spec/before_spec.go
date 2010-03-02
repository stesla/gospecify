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
package main

import (
	. "specify"
	t "../src/testspecify"
)

func init() {
	Describe("Before", func() {
		It("should run the block before each test", func(e Example) {
			reporter := testRun("", func(r t.Runner) {
				r.Before(func(e t.Example) { e.SetField("value", 42) })

				r.It("should set the value this time", func(e t.Example) {
					e.Field("value").Should(t.Be(42))
					e.SetField("value", 24)
				})

				r.It("should set this time too", func(e t.Example) {
					e.Field("value").Should(t.Be(42))
					e.SetField("value", 24)
				})
			})

			e.Value(reporter).Should(HavePassing(2))
		})

		It("should fail examples if assertions fail in the before block", func(e Example) {
			reporter := testRun("", func(r t.Runner) {
				r.Before(func(e t.Example) { e.Value(1).Should(Be(2)) })
				r.It("should fail in before", func(t.Example) {})
			})

			e.Value(reporter).Should(HaveFailing(1))
		})
	})
}
