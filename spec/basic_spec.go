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
	Describe("Running", func() {
		Before(func(e Example) {
			reporter := testRun("basic", func(r t.Runner) {
				r.It("pass", func(e t.Example) {
					e.Value(7 * 6).Should(t.Be(42))
					e.Value(1).ShouldNot(t.Be(2))
				})

				r.It("fail", func(e t.Example) {
					e.Value(7 * 6).ShouldNot(t.Be(42))
					e.Value(1).Should(t.Be(2))
				})

				r.It("pending", nil)

				r.It("error", func(e t.Example) { e.Error(makeError("catastrophe")) })
			})
			e.SetField("reporter", reporter)
		})

		/* Passing Examples */
		It("counts passing examples", func(e Example) { e.Field("reporter").Should(HavePassing(1)) })

		/* Failing Examples */
		It("counts failing examples", func(e Example) { e.Field("reporter").Should(HaveFailing(1)) })
		It("reports the example name when failing", func(e Example) { e.Field("reporter").Should(HaveFailureIncluding("basic fail")) })
		It("reports the file and line of a failing expectation", func(e Example) { e.Field("reporter").Should(HaveFailureAt("basic_spec.go:39")) })

		/* Pending Examples */
		It("counts pending examples", func(e Example) { e.Field("reporter").Should(HavePending(1)) })
		It("reports the example name when pending", func(e Example) { e.Field("reporter").Should(HavePendingIncluding("basic pending")) })
		It("reports the file and line of a pending example", func(e Example) { e.Field("reporter").Should(HavePendingAt("basic_spec.go:43")) })

		/* Errored Examples */
		It("counts errored examples", func(e Example) { e.Field("reporter").Should(HaveErrors(1)) })
		It("reports the example name when errored", func(e Example) { e.Field("reporter").Should(HaveErrorIncluding("basic error")) })
		It("reports the file and line of an error", func(e Example) { e.Field("reporter").Should(HaveErrorAt("basic_spec.go:45")) })
	})
}
