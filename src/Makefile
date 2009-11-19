PACKAGE=spec
TESTPROG=test_$(PACKAGE)

include $(GOROOT)/src/Make.$(GOARCH)

SRC=spec.go

TEST_SRC=test_main.go

all: test

test: $(TESTPROG)
	./$(TESTPROG)

$(PACKAGE).$O: $(SRC)
	$(GC) -o $@ $(SRC)

$(TESTPROG).$O: $(PACKAGE).$O $(TEST_SRC)
	$(GC) -o $@ $(TEST_SRC)

$(TESTPROG): $(TESTPROG).$O
	$(LD) -o $@ $<