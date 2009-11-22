package main

import "./specify"

func init() {
	initSpec();

	spec.Describe("Math", func() {
		spec.It("adds", func(it specify.It) {
			it.That(1 + 1).Should(spec.Be(2));
		});

		spec.It("multiplies", func(it specify.It) {
			it.That(3 * 3).Should(spec.Be(9));
			it.That(2 * 4).ShouldNot(spec.Be(6));
		});
	});

	spec.Describe("String", func() {
		spec.It("concatenates", func(it specify.It) {
			it.That("Doctor" + "Donna").Should(spec.Be("DoctorDonna"));
			it.That("foo" + "bar").ShouldNot(spec.Be("bar"));
		});
	});
}
