package specify

import(
	"fmt";
	"container/list";
	"os";
)

func Run() {
	report := makeReport();
	runList(behaviors, report);
	fmt.Println(report.summary());
	if report.failed > 0 {
		os.Exit(1);
	}
}

func Behavior(name string, description func()) {
	currentBehavior = makeBehavior(name);
	description();
	behaviors.PushBack(currentBehavior);
}

func It(name string, description func()) {
	currentIt = makeIt(name);
	description();
	currentBehavior.addIt(currentIt);
}

type Value interface{}

func That(value Value) (result *that) {
	result = makeThat(currentBehavior, currentIt, value);
	currentIt.addThat(result);
	return;
}

var currentBehavior *behavior;
var currentIt *it;
var behaviors *list.List;

func init() {
	behaviors = list.New();
}
 
type behavior struct {
	name string;
	its *list.List;
}

func makeBehavior(name string) (result *behavior) {
	result = &behavior{name:name};
	result.its = list.New();
	return;
}

func (self *behavior) addIt(it *it) {
	self.its.PushBack(it);
}

func (self *behavior) run(report *report) {
	runList(self.its, report);
}

type it struct {
	name string;
	thats *list.List;
}

func makeIt(name string) (result *it) {
	result = &it{name:name};
	result.thats = list.New();
	return;
}

func (self *it) addThat(that *that) {
	self.thats.PushBack(that);
}

func (self *it) run(report *report) {
	runList(self.thats, report);
}

type that struct {
	Should Matcher;
}

func makeThat(behavior *behavior, it *it, value Value) *that {
	matcher := &matcher{behavior:behavior, it:it, value:value};
	return &that{&should{matcher:matcher}}
}

func (self *that) run(report *report) {
	runner,_ := self.Should.(runner);
	runner.run(report);
}

type Matcher interface {
	Be(Value);
}

type matcher struct {
	*behavior;
	*it;
	value Value;
	block func() (bool, string);
}

func (self *matcher) run(report *report) {
	if pass,msg := self.block(); pass {
		report.pass();
	} else {
		report.fail(msg);
	}
}

type should struct { *matcher }
func (self *should) Be(value Value) {
	self.matcher.block = func() (bool, string) {
		if self.matcher.value != value {
			error := fmt.Sprintf("%v - %v - expected `%v` to be `%v`", self.matcher.behavior.name, self.matcher.it.name, self.matcher.value, value);
			return false, error;
		}
		return true, "";
	}
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