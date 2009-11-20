package main

import "./specify"

func init() {
	specify.Behavior("Math", func(it *specify.It) {
		it.Should("add", func(expect *specify.Expect) {
			expect.That(1 + 2).Should.Be(3);
		});

		it.Should("multiply", func(expect *specify.Expect) {
			expect.That(2 * 4).Should.Be(6);
		});
	});

	specify.Behavior("Strings", func(it *specify.It) {
		it.Should("concatenate", func(expect *specify.Expect) {
			expect.That("foo" + "bar").Should.Be("bar")
		});
	});
}