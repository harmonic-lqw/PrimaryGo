package week1

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	srcSlice := []int{1, 2}
	tarSlice := Insert[int](0, 999, srcSlice)
	if !reflect.DeepEqual(tarSlice, []int{999, 1, 2}) {
		t.Fatal("failed")
	}
	tarSlice = Insert[int](1, 999, srcSlice)
	if !reflect.DeepEqual(tarSlice, []int{1, 999, 2}) {
		t.Fatal("failed")
	}
	tarSlice = Insert[int](2, 999, srcSlice)
	if !reflect.DeepEqual(tarSlice, []int{1, 2, 999}) {
		t.Fatal("failed")
	}
}

func TestDelete(t *testing.T) {
	srcSlice := []int{1, 2, 3, 4, 5, 6}
	tarSlice := Delete[int](0, srcSlice)
	if !reflect.DeepEqual(tarSlice, []int{2, 3, 4, 5, 6}) {
		t.Fatal("failed")
	}

	srcSlice = []int{1, 2, 3, 4, 5, 6}
	tarSlice = Delete[int](len(srcSlice)-1, srcSlice)
	if !reflect.DeepEqual(tarSlice, []int{1, 2, 3, 4, 5}) {
		t.Log(tarSlice)
		t.Fatal("failed")
	}

	srcSlice = []int{1, 2, 3, 4, 5, 6}
	tarSlice = Delete[int](3, srcSlice)
	if !reflect.DeepEqual(tarSlice, []int{1, 2, 3, 5, 6}) {
		t.Fatal("failed")
	}

	srcSlice = []int{1, 2, 3, 4, 5, 6}
	tarSlice = srcSlice
	for i := 0; i < 5; i++ {
		tarSlice = Delete[int](0, tarSlice)
	}
	if cap(tarSlice) != 2 {
		t.Fatal("reduce capacity failed")
	}
	println(len(tarSlice), cap(tarSlice))
}
