package main

import sp "specify"

func init() {
	Describe("Math", func() {
		It("adds", func(the sp.The) {
			the.Value(1 + 1).Should(Be(2));
		});

		It("multiplies", func(the sp.The) {
			the.Value(3 * 3).Should(Be(9));
			the.Value(2 * 4).ShouldNot(Be(6));
		});
	});

	Describe("String", func() {
		It("concatenates", func(the sp.The) {
			the.Value("Doctor" + "Donna").Should(Be("DoctorDonna"));
			the.Value("foo" + "bar").ShouldNot(Be("bar"));
		});
	});
}
