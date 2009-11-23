# gospecify

This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.

## Installation

The Makefile assumes that you have the environment variables that are typically set for using the Go language (GOROOT, GOARCH, GOOS, and optionally GOBIN). Please refer to the GO installation instructions (http://golang.org/doc/install.html) for more information on properly setting them.

Once those variables are set you can:

     $ make test
     $ make install              # This will install the specify script in $HOME/bin
     $ make install GOBIN=$GOBIN # This will install the specify script in $GOBIN

## Usage

Take a look at src/example_spec.go for a simple example of how to write specifications using gospecify. Just put the code in package main and import your own code in. You can then use the specify command to compile and run your specs.

     $ specify *_spec.go

Or if you need to specify a package path you can do this:

     $ specify -I/path/to/pkg *_spec.go

You can look at src/Makefile to see how gospecify runs the command to test itself.

## Contributing

Contributions are always welcome. Just clone the git repo and hack away. You can submit pull requests or email me patches.

GitHub: http://github.com/stesla/gospecify
Email: samuel.tesla@gmail.com

Happy Testing!
