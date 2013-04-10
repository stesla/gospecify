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
package specify

import . "github.com/doun/gospecify"
func init() {
	Describe("Math", func() {
		It("adds", func(e Example) { e.Value(1 + 1).Should(Be(2)) })

		It("multiplies", func(e Example) {
			e.Value(3 * 3).Should(Be(9))
			e.Value(2 * 4).ShouldNot(Be(6))
		})
	})

	Describe("String", func() {
		It("concatenates", func(e Example) {
			e.Value("Doctor" + "Donna").Should(Be("DoctorDonna"))
			e.Value("foo" + "bar").ShouldNot(Be("bar"))
		})
	})
}
