package specify

import(
	"fmt";
	"container/list";
	"os";
)

type Test func() (bool, os.Error);

type Runner interface {
	Fail(os.Error);
	Failed() bool;
	Pass();
	Run(Test);
	Summary() string;
}

type Specification interface {
	Run(Runner);
	Describe(name string, description func());
	It(name string, description func());
	That(value Value) That;
}

type specification struct {
	currentDescribe *describe;
	currentIt *it;
	describes *list.List;
}

func (self *specification) Run(runner Runner) {
	runList(self.describes, runner);
	fmt.Println(runner.Summary());
	if runner.Failed() {
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

func (self *describe) run(runner Runner) {
	runList(self.its, runner);
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

func (self *it) run(runner Runner) {
	runList(self.thats, runner);
}

func (self *it) String() string { return self.name; }

type ValueTest func(Value) (bool, os.Error);

type That interface {
	SetBlock(ValueTest);
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
	block ValueTest;
}

func makeThat(describe *describe, it *it, value Value) That {
	return &that{describe:describe, it:it, value:value};
}

func (self *that) run(runner Runner) {
	runner.Run(func()(bool, os.Error) { return self.block(self.value) });
}

func (self *that) SetBlock(block ValueTest) { self.block = block; }
func (self *that) Should() Matcher  { return &should{self}; }
func (self *that) ShouldNot() Matcher { return &shouldNot{self}; }

type should struct { that That }
func (self *should) Be(expected Value) {
	self.that.SetBlock(func(actual Value) (bool, os.Error) {
		if actual != expected {
			error := fmt.Sprintf("expected `%v` to be `%v`", actual, expected);
			return false, os.NewError(error);
		}
		return true, nil;
	});
}

type shouldNot struct { that That }
func (self *shouldNot) Be(expected Value) {
	self.that.SetBlock(func(actual Value) (bool, os.Error) {
		if actual == expected {
			error := fmt.Sprintf("expected `%v` not to be `%v`", actual, expected);
			return false, os.NewError(error);
		}
		return true, nil;
	})
}

type test interface {
	run(Runner);
}

func runList(list *list.List, runner Runner) {
	doList(list, func(item Value) {
		test,_ := item.(test);
		test.run(runner);
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

func DotRunner() Runner { return &dotRunner{failures:list.New()}; }

type dotRunner struct {
	passed, failed int;
	failures *list.List;
}

func (self *dotRunner) Failed() bool { return self.failed > 0; }
func (self *dotRunner) Pass() { self.passed++; }
func (self *dotRunner) total() int { return self.passed + self.failed; }

func (self *dotRunner) Fail(err os.Error) {
	self.failures.PushBack(err);
	self.failed++;
}

func (self *dotRunner) Summary() string {
	if self.failed > 0 {
		fmt.Println("\nFAILED TESTS:");
		doList(self.failures, func(item Value) { fmt.Println("-", item) });
		fmt.Println("");
	}
	return fmt.Sprintf("Passed: %v Failed: %v Total: %v", self.passed, self.failed, self.total());
}

func (self *dotRunner) Run(test Test) {
	if pass,err := test(); pass {
		self.Pass();
	} else {
		self.Fail(err);
	}
}