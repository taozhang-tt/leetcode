package main

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestIsValid(t *testing.T) {
	Convey("TestIsValid", t, func() {
		s := "()"
		So(isValid(s), ShouldBeTrue)

		s = "()[]{}"
		So(isValid(s), ShouldBeTrue)

		s = "(]"
		So(isValid(s), ShouldBeFalse)

		s = "([)]"
		So(isValid(s), ShouldBeFalse)

		s = "{[]}"
		So(isValid(s), ShouldBeTrue)

		s = "]"
		So(isValid(s), ShouldBeFalse)

		s = "["
		So(isValid(s), ShouldBeFalse)
	})
}

func TestIsValid2(t *testing.T) {
	Convey("TestIsValid2", t, func() {
		s := "()"
		So(isValid2(s), ShouldBeTrue)

		s = "()[]{}"
		So(isValid2(s), ShouldBeTrue)

		s = "(]"
		So(isValid2(s), ShouldBeFalse)

		s = "([)]"
		So(isValid2(s), ShouldBeFalse)

		s = "{[]}"
		So(isValid2(s), ShouldBeTrue)

		s = "]"
		So(isValid2(s), ShouldBeFalse)

		s = "["
		So(isValid2(s), ShouldBeFalse)
	})
}
