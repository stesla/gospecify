package main

import sp "specify"

var runner TestRunner;

func init() {
	Describe("Specification", func() {
		runner = makeTestRunner();

		spec.Before(func () {
			s := sp.New();
			s.Describe("Foo", func() {
				s.It("pass", func(the sp.The) {
					the.Value(7 * 6).Should(Be(42));
					the.Value(1).ShouldNot(Be(2));
				});

				s.It("fail", func(the sp.The) {
					the.Value(7 * 6).ShouldNot(Be(42));
					the.Value(1).Should(Be(2));
				});
			});
			s.Run(runner);
		});

		It("indicates a passing test", func(the sp.The) {
			the.Value(runner.PassCount()).Should(Be(1));
		});

		It("indicates a failing test", func(the sp.The) {
			the.Value(runner.FailCount()).Should(Be(1));
		})
	});
}