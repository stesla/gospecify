PACKAGE=specify
TESTPROG=test_$(PACKAGE)

include $(GOROOT)/src/Make.$(GOARCH)
export GC
export LD
export O

SRC=specify.go

TEST_SRC=$(wildcard *_spec.go)

all: test

clean:
	rm -rf *.[a68] $(TESTPROG)

test: package
	../bin/specify -I . *_spec.go

package: $(PACKAGE).a

$(PACKAGE).a: $(PACKAGE).$O
	gopack grc $@ $(PACKAGE).$O

$(PACKAGE).$O: $(SRC)
	$(GC) -o $@ $(SRC)
