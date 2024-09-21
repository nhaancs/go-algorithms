package staticarray_test

import (
	"fmt"
	"github.com/nhaancs/go-algorithms/1-array-and-string-manipulation/data-structures/staticarray"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringStaticArr_New(t *testing.T) {
	t.Parallel()

	capacity := uint(3)
	arr := staticarray.New(capacity)

	assert.NotNil(t, arr)
	assert.Equal(t, capacity, arr.Capacity())

	_, err := fmt.Println(arr)
	assert.NoError(t, err)
}

func TestStringStaticArr_Lookup(t *testing.T) {
	t.Parallel()

	arr := staticarray.New(2)
	arr.Append("a")

	assert.Equal(t, "a", arr.Lookup(0))
	assert.Panics(t, func() { arr.Lookup(1) })
}

func TestStringStaticArr_Append(t *testing.T) {
	t.Parallel()

	arr := staticarray.New(3)
	arr.Append("a")
	arr.Append("b")
	arr.Append("c")

	assert.Equal(t, "a", arr.Lookup(0))
	assert.Equal(t, "b", arr.Lookup(1))
	assert.Equal(t, "c", arr.Lookup(2))
	assert.Panics(t, func() { arr.Append("d") })
}

func TestStringStaticArr_Insert(t *testing.T) {
	t.Parallel()

	arr := staticarray.New(4)
	arr.Append("d")
	arr.Insert(0, "a")
	arr.Insert(1, "b")
	arr.Insert(2, "c")

	assert.Equal(t, "a", arr.Lookup(0))
	assert.Equal(t, "b", arr.Lookup(1))
	assert.Equal(t, "c", arr.Lookup(2))
	assert.Equal(t, "d", arr.Lookup(3))
	assert.Panics(t, func() { arr.Insert(4, "e") })
	assert.Panics(t, func() { arr.Insert(2, "x") })
}

func TestStringStaticArr_Delete(t *testing.T) {
	t.Parallel()

	arr := staticarray.New(3)
	arr.Append("a")
	arr.Append("b")

	arr.Delete(0)

	assert.Equal(t, "b", arr.Lookup(0))
	assert.Panics(t, func() { arr.Delete(1) })
}
