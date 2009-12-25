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

import "os"

type AfterFunc func() os.Error
type BeforeBlock func(Example)
type ExampleBlock func(Example)
type ExampleGroupBlock func()

type Runner interface {
	After(AfterFunc);
	Before(BeforeBlock);
	Describe(string, ExampleGroupBlock);
	It(string, ExampleBlock);
	Run(Reporter);
}

func NewRunner() Runner	{ return makeRunner() }

type Location interface {
	String() string;
}

type Report interface {
	Title() string;
	Error() os.Error;
	Location() Location;
}

type Reporter interface {
	Fail(Report);
	Finish();
	Pass();
	Pending(Report);
}

func DotReporter() Reporter	{ return makeDotReporter() }

type Example interface {
	GetField(string) interface{};
	Field(string) Assertion;
	SetField(string, interface{});
	Value(interface{}) Assertion;
}

type Assertion interface {
	Should(Matcher);
	ShouldNot(Matcher);
}

type Matcher interface {
	Should(interface{}) os.Error;
	ShouldNot(interface{}) os.Error;
}

func Be(value interface{}) Matcher	{ return makeBeMatcher(value) }
