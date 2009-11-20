package main

import "./specify"

func init() {
	specify.Behavior("Math", func() {
		specify.It("adds", func() {
			specify.That(1 + 1).Should.Be(2);
		});

		specify.It("multiplies", func() {
			specify.That(3 * 3).Should.Be(9);
			specify.That(2 * 4).ShouldNot.Be(6);
		});
	});

	specify.Behavior("Strings", func() {
		specify.It("concatenates", func() {
			specify.That("Doctor" + "Donna").Should.Be("DoctorDonna");
			specify.That("foo" + "bar").ShouldNot.Be("bar")
		});
	});
}
