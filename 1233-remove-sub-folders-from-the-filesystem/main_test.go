package main

import (
	"fmt"
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test", t, func() {
		// root := NewNode("")
		// vals1 := []string{"a", "b", "c"}
		// vals2 := []string{"a", "b", "d"}

		// root.Add(vals1)
		// root.Add(vals2)

		// spew.Dump(root)

		folder := []string{"/a/b/c", "/a/b/ca", "/a/b/d"}

		fmt.Println(removeSubfolders(folder))

		folder = []string{"/a", "/a/b/c", "/a/b/d"}
		fmt.Println(removeSubfolders(folder))

		folder = []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
		fmt.Println(removeSubfolders(folder))

	})
}
