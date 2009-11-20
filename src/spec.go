package specify

import(
	"fmt";
	"container/list";
)

var expectations *list.List;

func init() {
	expectations = list.New();
}

type Matcher interface {
	Be(interface{});
}

type should struct {
	value interface{};
}

func (self should) Be(value interface{}) {
	if self.value != value {
		fmt.Printf("Expected %v to be %v\n", self.value, value);
	}
}

type Expectation struct {
	Should Matcher;
}

type Expect interface {
	That(interface {}) *Expectation;
}

type It interface {
	Should(string, func(*Expect));
}

func Run() {
	iter := expectations.Iter();
	for !closed(iter) {
		item := <-iter;
		if item == nil { break; }
		test,_ := item.(func());
		test();
	}
}

type expect struct {
	name string;
}

func (self expect) That(value interface {}) (result *Expectation) {
	result = &Expectation{};
	result.Should = &should{value};
	return;
}

type it struct {
	name string;
}

func (self it) Should(item string, spec func(*Expect)) {
	var e Expect = expect{item};
	expectations.PushBack(func() { spec(&e) });
}

func Behavior(item string, spec func(*It)) {
	var i It = it{item};
	spec(&i);
}
