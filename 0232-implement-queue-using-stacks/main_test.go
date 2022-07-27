package main

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestMyQueue(t *testing.T) {
	Convey("TestMyQueue", t, func() {
		q := Constructor()
		q.Push(1)
		q.Push(2)
		So(q.Peek(), ShouldEqual, 1)
		So(q.Pop(), ShouldEqual, 1)
		So(q.Empty(), ShouldBeFalse)
	})
}
