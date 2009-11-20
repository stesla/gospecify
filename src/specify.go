package specify

import(
	"fmt";
	"container/list";
)

func Run() {
	runList(behaviors);
}

func Behavior(name string, description func()) {
	currentBehavior = makeBehavior(name);
	description();
	behaviors.PushBack(currentBehavior);
}

func It(name string, description func()) {
	currentIt = makeIt(name, description);
	currentBehavior.addIt(currentIt);
}

type Value interface{}

func That(value Value) (result *that) {
	result = &that{};
	result.Should = &should{common{currentBehavior, currentIt, value}};
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

func (self *behavior) run() {
	fmt.Println("Running", self.name);
	runList(self.its);
}

type it struct {
	name string;
	description func();
}

func makeIt(name string, description func()) (result *it) {
	result = &it{name, description};
	return;
}

func (self *it) run() {
	fmt.Println("  ", self.name);
	self.description();
}

type that struct {
	Should matcher;
}

type matcher interface {
	Be(Value);
}

type common struct {
	*behavior;
	*it;
	value Value;
}

type should struct {
	common;
}

func (self *should) Be(value Value) {
	if self.common.value != value {
		fmt.Print("\t");
		fmt.Printf("expected `%v` to be `%v`", self.common.value, value);
		fmt.Println("");
	}
}

type runner interface {
	run();
}

func runList(list *list.List) {
	doList(list, func(item Value) {
		runner,_ := item.(runner);
		runner.run();
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