package main

func init() {
	initSpec();

	spec.Describe("Math", func() {
		spec.It("adds", func() {
			spec.That(1 + 1).Should().Be(3);
		});

		spec.It("multiplies", func() {
			spec.That(3 * 3).Should().Be(9);
			spec.That(2 * 4).ShouldNot().Be(6);
		});
	});

	spec.Describe("String", func() {
		spec.It("concatenates", func() {
			spec.That("Doctor" + "Donna").Should().Be("Donna");
			spec.That("foo" + "bar").ShouldNot().Be("bar");
		});
	});
}
