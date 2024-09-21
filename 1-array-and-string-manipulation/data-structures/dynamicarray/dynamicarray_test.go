package dynamicarray_test

import (
	"fmt"
	"github.com/nhaancs/go-algorithms/1-array-and-string-manipulation/data-structures/dynamicarray"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringDynamicArr_New(t *testing.T) {
	t.Parallel()

	arr := dynamicarray.New()
	assert.NotNil(t, arr)
	assert.Equal(t, uint(0), arr.Size())
	assert.Equal(t, uint(1), arr.Capacity())

	_, err := fmt.Println(arr)
	assert.NoError(t, err)
}

func TestStringDynamicArr_Lookup(t *testing.T) {
	t.Parallel()

	arr := dynamicarray.New()
	arr.Append("a")
	assert.Equal(t, "a", arr.Lookup(0))
	assert.Panics(t, func() { arr.Lookup(2) })

}

func TestStringDynamicArr_Append(t *testing.T) {
	t.Parallel()

	arr := dynamicarray.New()
	arr.Append("a")

	assert.Equal(t, uint(1), arr.Size())
	assert.Equal(t, uint(1), arr.Capacity())
	assert.Equal(t, "a", arr.Lookup(0))

	arr.Append("b")

	assert.Equal(t, uint(2), arr.Size())
	assert.Equal(t, uint(2), arr.Capacity())
	assert.Equal(t, "b", arr.Lookup(1))

	arr.Append("c")

	assert.Equal(t, uint(3), arr.Size())
	assert.Equal(t, uint(4), arr.Capacity())
	assert.Equal(t, "c", arr.Lookup(2))
}

func TestStringDynamicArr_Insert(t *testing.T) {
	t.Parallel()

	arr := dynamicarray.New()
	arr.Append("d")
	arr.Insert(0, "a")
	arr.Insert(1, "b")
	arr.Insert(2, "c")

	assert.Equal(t, uint(4), arr.Size())
	assert.Equal(t, uint(4), arr.Capacity())
	assert.Equal(t, "a", arr.Lookup(0))
	assert.Equal(t, "b", arr.Lookup(1))
	assert.Equal(t, "c", arr.Lookup(2))
	assert.Equal(t, "d", arr.Lookup(3))
	assert.Panics(t, func() { arr.Insert(4, "d") })
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

	assert.Equal(t, uint(2), arr.Size())
	assert.Equal(t, uint(4), arr.Capacity())
	assert.Equal(t, "b", arr.Lookup(0))
	assert.Equal(t, "c", arr.Lookup(1))
	assert.Panics(t, func() { arr.Delete(2) })
}
