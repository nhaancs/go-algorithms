package dynamicarray_test

import (
	"fmt"
	"github.com/nhaancs/go-algorithms/1-array-and-string-manipulation/data-structures/dynamicarray"
	"testing"
)

func TestStringDynamicArr_New(t *testing.T) {
	t.Parallel()

	arr := dynamicarray.New()

	if arr == nil {
		t.Error("Failed to create new static array")
	}

	if arr.Size() != 0 {
		t.Errorf("Expected 0 but got %d", arr.Size())
	}

	if arr.Capacity() != 1 {
		t.Errorf("Expected 1 but got %d", arr.Capacity())
	}
}

func TestStringDynamicArr_Lookup(t *testing.T) {
	t.Parallel()

	arr := dynamicarray.New()
	arr.Append("a")

	if arr.Lookup(0) != "a" {
		t.Errorf("Expected a but got %s", arr.Lookup(0))
	}

	fmt.Println(arr)

	defer func() {
		_ = recover()
	}()

	arr.Lookup(2)
	t.Error("The code did not panic")
}

func TestStringDynamicArr_Append(t *testing.T) {
	t.Parallel()

	arr := dynamicarray.New()
	arr.Append("a")

	if arr.Size() != 1 {
		t.Errorf("Expected 1 but got %d", arr.Size())
	}

	if arr.Capacity() != 1 {
		t.Errorf("Expected 1 but got %d", arr.Capacity())
	}

	if arr.Lookup(0) != "a" {
		t.Errorf("Expected a but got %s", arr.Lookup(0))
	}

	arr.Append("b")

	if arr.Size() != 2 {
		t.Errorf("Expected 2 but got %d", arr.Size())
	}

	if arr.Capacity() != 2 {
		t.Errorf("Expected 2 but got %d", arr.Capacity())
	}

	if arr.Lookup(1) != "b" {
		t.Errorf("Expected b but got %s", arr.Lookup(1))
	}

	arr.Append("c")

	if arr.Size() != 3 {
		t.Errorf("Expected 3 but got %d", arr.Size())
	}

	if arr.Capacity() != 4 {
		t.Errorf("Expected 4 but got %d", arr.Capacity())
	}

	if arr.Lookup(2) != "c" {
		t.Errorf("Expected c but got %s", arr.Lookup(2))
	}

	fmt.Println(arr)
}

func TestStringDynamicArr_Insert(t *testing.T) {
	t.Parallel()

	arr := dynamicarray.New()
	arr.Append("d")
	arr.Insert(0, "a")
	arr.Insert(1, "b")
	arr.Insert(2, "c")

	if arr.Size() != 4 {
		t.Errorf("Expected 4 but got %d", arr.Size())
	}

	if arr.Capacity() != 4 {
		t.Errorf("Expected 4 but got %d", arr.Capacity())
	}

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
		t.Errorf("Expected d but got %s", arr.Lookup(3))
	}

	defer func() {
		_ = recover()
	}()
	arr.Insert(4, "d")
	t.Error("The code did not panic")
}

func TestStringDynamicArr_Delete(t *testing.T) {
	t.Parallel()

	arr := dynamicarray.New()
	arr.Append("a")
	arr.Append("b")
	arr.Append("c")
	arr.Append("d")

	arr.Delete(0)
	arr.Delete(2)

	if arr.Size() != 2 {
		t.Errorf("Expected 2 but got %d", arr.Size())
	}

	if arr.Lookup(0) != "b" {
		t.Errorf("Expected a but got %s", arr.Lookup(0))
	}

	if arr.Lookup(1) != "c" {
		t.Errorf("Expected empty string but got %s", arr.Lookup(1))
	}

	defer func() {
		_ = recover()
	}()

	arr.Delete(2)
	t.Error("The code did not panic")
}
