all: test

test:
	cd src; make test

package:
	cd src; make package

install: package
	cp src/specify.a $(GOROOT)/pkg/$(GOOS)_$(GOARCH)