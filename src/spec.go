package specify

import(
	"fmt";
	"container/list";
	"os";
)

var behaviors *list.List;
var exitCode int;

func init() {
	behaviors = list.New();
}

type matcher interface {
	Be(interface{});
}

type should struct {
	value interface{};
}

func (self should) Be(value interface{}) {
	if self.value != value {
		self.Failf("Expected %v to be %v\n", self.value, value);
	}
}

func (self should) Failf(format string, args ...) {
	fmt.Printf(format, args);
	exitCode = 1;
}

type That struct {
	Should matcher;
}

type Expect interface {
	That(interface {}) *That;
}

type It interface {
	Should(string, func(*Expect));
}

type expect struct {
	spec func(*Expect);
}

func (self expect) run() {
	var expect Expect = self;
	self.spec(&expect);
}

func (self expect) That(value interface {}) (that *That) {
	that = &That{};
	that.Should = &should{value};
	return;
}

type it struct {
	name string;
	expectations *list.List;
}

func (self it) Should(desc string, spec func(*Expect)) {
	expectation := expect{spec};
	self.expectations.PushBack(expectation);
}

func (self it) run() {
	runList(self.expectations);
}

func Behavior(item string, spec func(*It)) {
	var it It = it{item, list.New()};
	spec(&it);
	behaviors.PushBack(it);
}

func Run() {
	runList(behaviors);
	os.Exit(exitCode);
}

type runner interface {
	run();
}

func runList(l *list.List) {
	iter := l.Iter();
	for !closed(iter) {
		item := <-iter;
		if item == nil { break; }
		runner,_ := item.(runner);
		runner.run();
	}
}