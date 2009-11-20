PACKAGE=specify
TESTPROG=test_$(PACKAGE)

include $(GOROOT)/src/Make.$(GOARCH)

SRC=spec.go

TEST_SRC=$(wildcard *_spec.go)

all: test

clean:
	rm -rf *.$O $(TESTPROG) $(PACKAGE).$O

test: $(TESTPROG)
	./$(TESTPROG)

$(PACKAGE).$O: $(SRC)
	$(GC) -o $@ $(SRC)

$(TESTPROG).$O: $(PACKAGE).$O $(TEST_SRC)
	$(GC) -o $@ $(TEST_SRC)

$(TESTPROG): $(TESTPROG).$O
	$(LD) -o $@ $<