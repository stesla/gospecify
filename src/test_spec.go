package main

import "./specify"

func init() {
	specify.Behavior("Math", func() {
		specify.It("adds", func() {
			specify.That(1 + 2).Should.Be(3);
		});

		specify.It("multiplies", func() {
			specify.That(2 * 4).Should.Be(6);
		});
	});

	specify.Behavior("Strings", func() {
		specify.It("concatenates", func() {
			specify.That("foo" + "bar").Should.Be("bar")
		});
	});
}
