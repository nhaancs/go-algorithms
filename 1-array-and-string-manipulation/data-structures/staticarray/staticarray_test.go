package staticarray_test

import (
	"fmt"
	"github.com/nhaancs/go-algorithms/1-array-and-string-manipulation/data-structures/staticarray"
	"testing"
)

func TestStringStaticArr_New(t *testing.T) {
	t.Parallel()

	capacity := uint(3)
	arr := staticarray.New(capacity)

	if arr == nil {
		t.Error("Failed to create new static array")
	}

	if arr.Capacity() != 3 {
		t.Errorf("Expected %d but got %d", capacity, arr.Capacity())
	}
}

func TestStringStaticArr_Lookup(t *testing.T) {
	t.Parallel()

	arr := staticarray.New(2)
	arr.Append("a")

	if arr.Lookup(0) != "a" {
		t.Errorf("Expected a but got %s", arr.Lookup(0))
	}

	fmt.Println(arr)

	defer func() {
		_ = recover()
	}()

	arr.Lookup(1)
	t.Error("The code did not panic")
}

func TestStringStaticArr_Append(t *testing.T) {
	t.Parallel()

	arr := staticarray.New(3)
	arr.Append("a")
	arr.Append("b")
	arr.Append("c")

	if arr.Lookup(0) != "a" {
		t.Errorf("Expected a but got %s", arr.Lookup(0))
	}

	if arr.Lookup(1) != "b" {
		t.Errorf("Expected b but got %s", arr.Lookup(1))
	}

	if arr.Lookup(2) != "c" {
		t.Errorf("Expected c but got %s", arr.Lookup(2))
	}

	defer func() {
		_ = recover()
	}()

	arr.Append("d")
	t.Error("The code did not panic")
}

func TestStringStaticArr_Insert(t *testing.T) {
	t.Parallel()

	arr := staticarray.New(4)
	arr.Append("d")
	arr.Insert(0, "a")
	arr.Insert(1, "b")
	arr.Insert(2, "c")

	if arr.Lookup(0) != "a" {
		t.Errorf("Expected a but got %s", arr.Lookup(0))
	}

	if arr.Lookup(1) != "b" {
		t.Errorf("Expected b but got %s", arr.Lookup(1))
	}

	if arr.Lookup(2) != "c" {
		t.Errorf("Expected c but got %s", arr.Lookup(2))
	}

	if arr.Lookup(3) != "d" {
		t.Errorf("Expected d but got %s", arr.Lookup(2))
	}

	defer func() {
		if v := recover(); v == nil {
			t.Error("The code did not panic")
		}

		defer func() {
			if v := recover(); v == nil {
				t.Error("The code did not panic")
			}
		}()
		arr.Insert(4, "y")
	}()
	arr.Insert(2, "x")
}

func TestStringStaticArr_Delete(t *testing.T) {
	t.Parallel()

	arr := staticarray.New(3)
	arr.Append("a")
	arr.Append("b")

	arr.Delete(0)

	if arr.Lookup(0) != "b" {
		t.Errorf("Expected b but got %s", arr.Lookup(0))
	}

	defer func() {
		_ = recover()
	}()

	arr.Delete(1)
	t.Error("The code did not panic")
}
