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
	"fmt"
	"runtime"
)

type location struct {
	file string
	line int
}

const (
	assertionDepth int = 5
	errorDepth     int = 3
)

var blockDepth int = 3

func AdjustBlockDepth(delta int)     { blockDepth += delta }
func newAssertionLocation() location { return newLocation(assertionDepth) }
func newErrorLocation() location     { return newLocation(errorDepth) }
func newBlockLocation() location     { return newLocation(blockDepth) }

func newLocation(depth int) location {
	if _, file, line, ok := runtime.Caller(depth); ok {
		return location{file, line}
	}
	panic("newLocation")
}

func (self location) String() string { return fmt.Sprintf("%v:%d", self.file, self.line) }
