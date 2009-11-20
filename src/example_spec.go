package main

func init() {
	initSpec();

	spec.Behavior("Math", func() {
		spec.It("adds", func() {
			spec.That(1 + 1).Should.Be(2);
		});

		spec.It("multiplies", func() {
			spec.That(3 * 3).Should.Be(9);
			spec.That(2 * 4).ShouldNot.Be(6);
		});
	});

	spec.Behavior("Strings", func() {
		spec.It("concatenates", func() {
			spec.That("Doctor" + "Donna").Should.Be("DoctorDonna");
			spec.That("foo" + "bar").ShouldNot.Be("bar")
		});
	});
}
