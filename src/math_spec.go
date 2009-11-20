package main

import "./specify"

func init() {
	specify.Behavior("Math", func(it *specify.It) {

		it.Should("add integers", func(expect *specify.Expect) {

			expect.That(1 + 2).Should.Be(3);

		})

	})
}