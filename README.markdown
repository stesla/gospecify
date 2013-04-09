# gospecify

This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.

## Installation

The Makefile assumes that you have the environment variables that are typically set for using the Go language (GOROOT, GOARCH, GOOS, and optionally GOBIN). Please refer to the GO installation instructions (http://golang.org/doc/install.html) for more information on properly setting them.
Just use go get to install
## Usage
copy spec/spec_func.go to your spec folder, and then u can write your tests as spec/example_spec.go
when finished:

     $go build -o spec
     $spec -sp.run=.*
or if you want to run with go test, just check spec/spec_test.go to see what to do
sp.run accept regex as run filter.

## Contributing

Contributions are always welcome. Just clone the git repo and hack away. You can submit pull requests or email me patches.

* GitHub: http://github.com/stesla/gospecify
* Email: samuel.tesla@gmail.com

Happy Testing!
