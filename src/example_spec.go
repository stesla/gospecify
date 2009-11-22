package main

func init() {
	initSpec();

	spec.Describe("Math", func() {
		spec.It("adds", func() {
			spec.That(1 + 1).Should(spec.Be(2));
		});

		spec.It("multiplies", func() {
			spec.That(3 * 3).Should(spec.Be(9));
			spec.That(2 * 4).ShouldNot(spec.Be(6));
		});
	});

	spec.Describe("String", func() {
		spec.It("concatenates", func() {
			spec.That("Doctor" + "Donna").Should(spec.Be("DoctorDonna"));
			spec.That("foo" + "bar").ShouldNot(spec.Be("bar"));
		});
	});
}
