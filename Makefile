GOBIN=$(HOME)/bin

all: test

clean:
	cd src; make clean

test:
	cd src; make test

package:
	cd src; make package

install: package
	cp src/specify.a $(GOROOT)/pkg/$(GOOS)_$(GOARCH)
	cp bin/specify $(GOBIN)