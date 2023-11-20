package week1

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	srcSlice := []int{1, 2}
	tarSlice := Insert[int](0, 999, srcSlice)
	if !reflect.DeepEqual(tarSlice, []int{999, 1, 2}) {
		t.Log("failed")
	}
	tarSlice = Insert[int](1, 999, srcSlice)
	if !reflect.DeepEqual(tarSlice, []int{1, 999, 2}) {
		t.Log("failed")
	}
	tarSlice = Insert[int](2, 999, srcSlice)
	if !reflect.DeepEqual(tarSlice, []int{1, 2, 999}) {
		t.Log("failed")
	}
}
