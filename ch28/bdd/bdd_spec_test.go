package bdd_test

import "testing"

func TestSpec(t *testing.T) {
	// only pass t into top-level Convey calls
	Convey("Given 2 even numbers", func() {
		a := 2
		b := 4
		Convey("When add the two numbers", func() {
			c := a + b
			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
