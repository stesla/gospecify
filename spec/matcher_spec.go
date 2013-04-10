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

import (
	"bytes"
    . "github.com/doun/gospecify"
)

func init() {
	Describe("Be", func() {
		It("should match reference equality", func(e Example) {
			var a, b int
			e.Value(&a).Should(Be(&a))
			e.Value(&a).ShouldNot(Be(&b))
		})

		It("should not care about the value", func(e Example) {
			a := 42
			b := 42
			e.Value(&a).ShouldNot(Be(&b))
		})
	})

	Describe("BeNil", func() {
		It("should match nil", func(e Example) { e.Value(nil).Should(BeNil()) })
		It("should not match non-nil values", func(e Example) { e.Value(42).ShouldNot(BeNil()) })
		It("should not match zero", func(e Example) { e.Value(0).ShouldNot(BeNil()) })
		It("should not match false", func(e Example) { e.Value(false).ShouldNot(BeNil()) })
	})

	Describe("BeFalse", func() {
		It("should match false", func(e Example) { e.Value(false).Should(BeFalse()) })
		It("should not match true", func(e Example) { e.Value(true).ShouldNot(BeFalse()) })
		It("should not match nil", func(e Example) { e.Value(nil).ShouldNot(BeFalse()) })
		It("should not match zero", func(e Example) { e.Value(0).ShouldNot(BeFalse()) })
		It("should not match other values", func(e Example) { e.Value(42).ShouldNot(BeFalse()) })
	})

	Describe("BeTrue", func() {
		It("should match true", func(e Example) { e.Value(true).Should(BeTrue()) })
		It("should not match false", func(e Example) { e.Value(false).ShouldNot(BeTrue()) })
		It("should not match other values", func(e Example) { e.Value(42).ShouldNot(BeTrue()) })
	})

	Describe("BeEqualTo", func() {
		It("should match numbers", func(e Example) {
			e.Value(1).Should(BeEqualTo(1))
			e.Value(1.2).ShouldNot(BeEqualTo(2.1))
		})

		It("should match strings", func(e Example) {
			e.Value("foo").Should(BeEqualTo("foo"))
			e.Value("Doctor").ShouldNot(BeEqualTo("Donna"))
		})

		It("should match things with EqualTo()", func(e Example) { e.Value([]byte{1, 2}).Should(BeEqualTo(bslice([]byte{1, 2}))) })
	})
}

type bslice []byte

func (self bslice) EqualTo(value interface{}) bool {
	if other, ok := value.([]byte); ok {
		return bytes.Equal(([]byte)(self), other)
	}
	return false
}
