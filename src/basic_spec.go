package main

import "./specify"

var runner TestRunner;

func init() {
	initSpec();
	
	spec.Describe("Specification", func() {
		runner = makeTestRunner();

		spec.Before(func () {
			s := specify.New();
			s.Describe("Foo", func() {
				s.It("pass", func(it specify.It) {
					it.That(7 * 6).Should(spec.Be(42));
					it.That(1).ShouldNot(spec.Be(2));
				});

				s.It("fail", func(it specify.It) {
					it.That(7 * 6).ShouldNot(spec.Be(42));
					it.That(1).Should(spec.Be(2));
				});
			});
			s.Run(runner);
		});

		spec.It("indicates a passing test", func(it specify.It) {
			it.That(runner.PassCount()).Should(spec.Be(1));
		});

		spec.It("indicates a failing test", func(it specify.It) {
			it.That(runner.FailCount()).Should(spec.Be(1));
		})
	});
}