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
				s.It("pass", func() {
					s.That(7 * 6).Should().Be(42);
					s.That(1).ShouldNot().Be(2);
				});

				s.It("fail", func() {
					s.That(7 * 6).ShouldNot().Be(42);
					s.That(1).Should().Be(2);
				});
			});
			s.Run(runner);
		});

		spec.It("indicates a passing test", func() {
			spec.That(runner.PassCount()).Should().Be(1);
		});

		spec.It("indicates a failing test", func() {
			spec.That(runner.FailCount()).Should().Be(1);
		})
	});
}