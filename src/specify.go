package specify

import(
	"fmt";
	"container/list";
	"os";
)

type Specification interface {
	Run();
	Describe(name string, description func());
	It(name string, description func());
	That(value Value) That;
}

type specification struct {
	currentDescribe *describe;
	currentIt *it;
	describes *list.List;
}

func (self *specification) Run() {
	report := makeReport();
	runList(self.describes, report);
	fmt.Println(report.summary());
	if report.failed > 0 {
		os.Exit(1);
	}
}

func (self *specification) Describe(name string, description func()) {
	self.currentDescribe = makeDescribe(name);
	description();
	self.describes.PushBack(self.currentDescribe);
}

func (self *specification) It(name string, description func()) {
	self.currentIt = makeIt(name);
	description();
	self.currentDescribe.addIt(self.currentIt);
}

type Value interface{}

func (self *specification) That(value Value) (result That) {
	result = makeThat(self.currentDescribe, self.currentIt, value);
	self.currentIt.addThat(result);
	return;
}

func New() Specification {
	return &specification{describes:list.New()};
}
 
type describe struct {
	name string;
	its *list.List;
}

func makeDescribe(name string) (result *describe) {
	result = &describe{name:name};
	result.its = list.New();
	return;
}

func (self *describe) addIt(it *it) {
	self.its.PushBack(it);
}

func (self *describe) run(report *report) {
	runList(self.its, report);
}

func (self *describe) String() string { return self.name; }

type it struct {
	name string;
	thats *list.List;
}

func makeIt(name string) (result *it) {
	result = &it{name:name};
	result.thats = list.New();
	return;
}

func (self *it) addThat(that That) {
	self.thats.PushBack(that);
}

func (self *it) run(report *report) {
	runList(self.thats, report);
}

func (self *it) String() string { return self.name; }

type That interface {
	SetBlock(func(Value) (bool, string));
	Should() Matcher;
	ShouldNot() Matcher;
}

type Matcher interface {
	Be(Value);
}

type that struct {
	*describe;
	*it;
	value Value;
	block func(Value) (bool, string);
}

func makeThat(describe *describe, it *it, value Value) That {
	return &that{describe:describe, it:it, value:value};
}

func (self *that) run(report *report) {
	if pass,msg := self.block(self.value); pass {
		report.pass();
	} else {
		msg = fmt.Sprintf("%v %v - %v", self.describe, self.it, msg);
		report.fail(msg);
	}
}

func (self *that) SetBlock(block func(Value) (bool, string)) { self.block = block; }
func (self *that) Should() Matcher  { return &should{self}; }
func (self *that) ShouldNot() Matcher { return &shouldNot{self}; }

type should struct { That }
func (self *should) Be(expected Value) {
	self.That.SetBlock(func(actual Value) (bool, string) {
		if actual != expected {
			error := fmt.Sprintf("expected `%v` to be `%v`", actual, expected);
			return false, error;
		}
		return true, "";
	});
}

type shouldNot struct { That }
func (self *shouldNot) Be(expected Value) {
	self.That.SetBlock(func(actual Value) (bool, string) {
		if actual == expected {
			error := fmt.Sprintf("expected `%v` not to be `%v`", actual, expected);
			return false, error;
		}
		return true, "";
	})
}

type runner interface {
	run(*report);
}

func runList(list *list.List, report *report) {
	doList(list, func(item Value) {
		runner,_ := item.(runner);
		runner.run(report);
	});
}

func doList(list *list.List, do func(Value)) {
	iter := list.Iter();
	for !closed(iter) {
		item := <-iter;
		if item == nil { break; }
		do(item);
	}
}

type report struct {
	passed, failed int;
	failures *list.List;
}

func makeReport() *report {
	return &report{failures:list.New()};
}

func (self *report) pass() { self.passed++; }
func (self *report) total() int { return self.passed + self.failed; }

func (self *report) fail(msg string) {
	self.failures.PushBack(msg);
	self.failed++;
}

func (self *report) summary() string {
	if self.failed > 0 {
		fmt.Println("\nFAILED TESTS:");
		doList(self.failures, func(item Value) { fmt.Println("-", item) });
		fmt.Println("");
	}
	return fmt.Sprintf("Passed: %v Failed: %v Total: %v", self.passed, self.failed, self.total());
}