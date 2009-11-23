PACKAGE=specify
TESTPROG=test_$(PACKAGE)

include $(GOROOT)/src/Make.$(GOARCH)

SRC=specify.go

TEST_SRC=$(wildcard *_spec.go)

all: test

clean:
	rm -rf *.[a68] $(TESTPROG)

test: $(TESTPROG)
	./$(TESTPROG)

package: $(PACKAGE).a

$(PACKAGE).a: $(PACKAGE).$O
	gopack grc $@ $(PACKAGE).$O

$(PACKAGE).$O: $(SRC)
	$(GC) -o $@ $(SRC)

$(TESTPROG).$O: $(PACKAGE).$O $(TEST_SRC)
	$(GC) -o $@ -I. $(TEST_SRC)

$(TESTPROG): $(TESTPROG).$O
	$(LD) -o $@ $<
