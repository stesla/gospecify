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
	expect *expect;
	value interface{};
}

func (self should) Be(value interface{}) {
	if self.value != value {
		self.fail("Expected %v to be %v", self.value, value);
	}
}

func (self should) fail(format string, args ...) {
	self.expect.fail(format, args);
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
	failure string;
}

func (self *expect) run() {
	var expect Expect = self;
	self.spec(&expect);
}

func (self *expect) That(value interface {}) (that *That) {
	that = &That{};
	that.Should = &should{self, value};
	return;
}

func (self *expect) fail(format string, args ...) {
	if self.pass() {
		self.failure = fmt.Sprintf(format, args);
	}
}

func (self *expect) pass() bool {
	return self.failure == "";
}

func (self expect) log(log *log) {
	if self.pass() {
		log.pass();
	} else {
		log.fail(self.failure);
	}
}

type it struct {
	name string;
	expectations *list.List;
}

func (self *it) Should(desc string, spec func(*Expect)) {
	expectation := expect{spec:spec};
	self.expectations.PushBack(&expectation);
}

func (self *it) run() {
	runList(self.expectations);
}

func (self *it) log(log *log) {
	logList(log, self.expectations);
}

func Behavior(item string, spec func(*It)) {
	var it It = &it{item, list.New()};
	spec(&it);
	behaviors.PushBack(it);
}

func Run() {
	var log log;
	runList(behaviors);
	logList(&log, behaviors);
	fmt.Println("");
	fmt.Println(log.summary());
	os.Exit(exitCode);
}

type log struct {
	successes, failures int;
}

func (self *log) fail(message string) {
	self.failures++;
	fmt.Println(message);
}

func (self *log) pass() { self.successes++; }

func (self *log) summary() string {
	return fmt.Sprintf("Passed: %v Failed: %v Total: %v", self.successes, self.failures, self.total());
}

func (self *log) total() int {
	return self.successes + self.failures;
}

type runner interface {
	run();
}

type logger interface {
	log(*log);
}

func logList(log *log, l *list.List) {
	doList(l, func(item interface{}) {
		logger,ok := item.(logger);
		if !ok {
			fmt.Fprintf(os.Stderr, "Not a logger\n");
			os.Exit(1);
		}
		logger.log(log);
	});
}

func runList(l *list.List) {
	doList(l, func(item interface{}) {
		runner,ok := item.(runner);
		if !ok {
			fmt.Fprintf(os.Stderr, "Not a runner\n");
			os.Exit(1);
		}
		runner.run();
	});
}

func doList(l *list.List, do func(interface {})) {
	iter := l.Iter();
	for !closed(iter) {
		item := <-iter;
		if item == nil { break; }
		do(item);
	}
}