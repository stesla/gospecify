package specify

import "fmt"

type Matcher interface {
	Be(interface{});
}

type should struct {
	Value interface{};
}

func (self should) Be(value interface{}) {
	if self.Value != value {
		fmt.Printf("Expected %v to be %v\n", self.Value, value);
	}
}

type Expectation struct {
	Should *should;
}

type Expect interface {
	That(interface {}) *Expectation;
}

type It interface {
	Should(string, func(*Expect));
}

func Run() {
}

type expect struct {
	name string;
}

func (self expect) That(value interface {}) (result *Expectation) {
	result = &Expectation{&should{value}};
	return;
}

type it struct {
	name string;
}

func (self it) Should(item string, spec func(*Expect)) {
	var e Expect = expect{item};
	spec(&e);
}

func Behavior(item string, spec func(*It)) {
	var i It = it{item};
	spec(&i);
}
