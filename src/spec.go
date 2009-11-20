package specify

import "fmt"

type Matcher interface {
	Be(interface{});
}

type Expectation struct {
	Should Matcher;
}

type Expect interface {
	That(interface {}) *Expectation;
}

type It interface {
	Should(item string, spec func(*Expect));
}

func Run() {
	fmt.Println("Specs Would Run Here");
}

func Behavior(item string, spec func(*It)) {
}
