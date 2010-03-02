# Copyright (c) 2009-2010 Samuel Tesla <samuel.tesla@gmail.com>

# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:

# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.

# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.

PACKAGE=specify
TESTPROG=test_$(PACKAGE)

include $(GOROOT)/src/Make.$(GOARCH)
export GC
export LD
export O

SRC=$(wildcard *.go)

all: test

clean:
	rm -rf *.[a68] $(TESTPROG) _test

package: $(PACKAGE).a

$(PACKAGE).a: $(PACKAGE).$O
	gopack grc $@ $(PACKAGE).$O

$(PACKAGE).$O: $(SRC)
	$(GC) -o $@ $(SRC)

testpackage: test$(PACKAGE).a

test$(PACKAGE).a: test$(PACKAGE).$O
	gopack grc $@ test$(PACKAGE).$O

test$(PACKAGE).$O: _test $(SRC)
	for i in $(SRC); do sed -e 's/package $(PACKAGE)/package test$(PACKAGE)/' < $$i > _test/$$i; done
	$(GC) -o $@ _test/*.go

_test:
	mkdir _test
