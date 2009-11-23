package specify

import(
	"fmt";
	"container/list";
	"os";
)

type Test func() (bool, os.Error);

type Runner interface {
	Fail(os.Error);
	Finish();
	Pass();
	Run(Test);
}

type Specification interface {
	Before(func());
	Run(Runner);
	Describe(name string, description func());
	It(name string, description ItBlock);
}

type specification struct {
	currentDescribe *describe;
	describes *list.List;
}

func (self *specification) Run(runner Runner) {
	doList(self.describes, func(item Value) {
		describe,_ := item.(builder);
		describe.build();
	});
	runList(self.describes, runner);
	runner.Finish();
}

func (self *specification) Before(action func()) {
	self.currentDescribe.addBefore(action);
}

func (self *specification) Describe(name string, description func()) {
	describe := makeDescribe(self, name, description);
	self.describes.PushBack(describe);
}

type ItBlock func(the The);

func (self *specification) It(name string, description ItBlock) {
	it := makeIt(name);
	it.description = description;
	it.describe = self.currentDescribe;
	self.currentDescribe.addIt(it);
}

func Be(expected Value) Matcher {
	return &beMatcher{};
}

type Value interface{}

func New() Specification {
	return &specification{describes:list.New()};
}
 
type describe struct {
	name string;
	block func();
	its *list.List;
	spec *specification;
	beforeAction func();
}

func makeDescribe(spec *specification, name string, block func()) *describe {
	return &describe{name:name, block:block, its:list.New(), spec:spec};
}

func (self *describe) addBefore(action func()) {
	self.beforeAction = action;
}

func (self *describe) addIt(it *it) {
	self.its.PushBack(it);
}

func (self *describe) build() {
	self.spec.currentDescribe = self;
	self.block();
}

func (self *describe) run(runner Runner) {
	if self.beforeAction != nil {
		self.beforeAction();
	}
	runList(self.its, runner);
}

func (self *describe) String() string { return self.name; }

type The interface {
	Value(Value) That;
}

type it struct {
	name string;
	description ItBlock;
	*describe;
	*itRunner;
}

func makeIt(name string) (result *it) {
	result = &it{name:name};
	result.itRunner = &itRunner{};
	return;
}

func (self *it) run(runner Runner) {
	self.description(self);
	if self.itRunner.failed {
		runner.Fail(self.itRunner.error);
	} else {
		runner.Pass();
	}
}

func (self *it) Value(value Value) That {
	return makeThat(self.describe, self, value);
}

func (self *it) String() string { return self.name; }

type ValueTest func(Value) (bool, os.Error);

type That interface {
	Should(Matcher);
	ShouldNot(Matcher);
}

type Matcher interface {
	Should(Value) (bool, os.Error);
	ShouldNot(Value) (bool, os.Error);
}

type beMatcher struct {
	expected Value;
}

func (self *beMatcher) Should(actual Value) (pass bool, err os.Error) {
	if actual != self.expected {
		error := fmt.Sprintf("expected `%v` to be `%v`", actual, self.expected);
		return false, os.NewError(error);
	}
	return true, nil;
}

func (self *beMatcher) ShouldNot(actual Value) (pass bool, err os.Error) {
	if actual == self.expected {
		error := fmt.Sprintf("expected `%v` not to be `%v`", actual, self.expected);
		return false, os.NewError(error);
	}
	return true, nil;
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

func (self *that) Should(matcher Matcher) {
	self.it.Run(func() (bool, os.Error) {
		return matcher.Should(self.value);
	});
}

func (self *that) ShouldNot(matcher Matcher) {
	self.it.Run(func() (bool, os.Error) {
		return matcher.ShouldNot(self.value);
	});
}

type test interface {
	run(Runner);
}

type builder interface {
	build();
}

func runList(list *list.List, runner Runner) {
	done := make([] chan bool, list.Len());
	i := 0;
	doList(list, func(item Value) {
		ch := make(chan bool);
		done[i] = ch;
		i++;
		go func() {
			test,_ := item.(test);
			test.run(runner);
			ch <- true;
		}();
	});
	for _,ch := range done { <-ch; }
}

func doList(list *list.List, do func(Value)) {
	iter := list.Iter();
	for !closed(iter) {
		item := <-iter;
		if item == nil { break; }
		do(item);
	}
}

func BasicRunner() Runner { return &basicRunner{} }

type basicRunner struct {}

func (self *basicRunner) Fail(os.Error) {}
func (self *basicRunner) Finish() {}
func (self *basicRunner) Pass() {}
func (self *basicRunner) Run(test Test) {
	if pass,err := test(); pass {
		self.Pass();
	} else {
		self.Fail(err);
	}
}

func DotRunner() Runner { return &dotRunner{failures:list.New()}; }

type dotRunner struct {
	basicRunner;
	passed, failed int;
	failures *list.List;
}

func (self *dotRunner) Pass() {
	self.passed++;
	fmt.Print(".");
}

func (self *dotRunner) Fail(err os.Error) {
	self.failures.PushBack(err);
	self.failed++;
	fmt.Print("F");
}

func (self *dotRunner) Finish() {
	fmt.Println("");
	fmt.Print(self.summary());
	fmt.Println("");
}

func (self *dotRunner) total() int { return self.passed + self.failed; }
func (self *dotRunner) summary() string {
	if self.failed > 0 {
		fmt.Println("FAILED TESTS:");
		doList(self.failures, func(item Value) { fmt.Println("-", item) });
	}
	return fmt.Sprintf("Passed: %v Failed: %v Total: %v", self.passed, self.failed, self.total());
}

type itRunner struct {
	basicRunner;
	failed bool;
	error os.Error;
}

func (self *itRunner) Fail(err os.Error) {
	if self.failed { return }
	self.failed = true;
	self.error = err;
}
