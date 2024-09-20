# Array and string manipulation

## Array

An array organizes items sequentially, one after another in memory.

Each position in the array has an index, starting at 0.

### Common operations

| Operation | Worst case |
| --------- |------------|
| Space     | O(n)       |
| Lookup    | O(1)       |
| Append    | O(1)       |
| Insert    | O(n)       |
| Delete    | O(n)       |

### Strengths

- Fast lookups
- Fast appends
- Cache-friendly

### Weaknesses

- Fixed size
- Costly inserts and deletes. You have to "scoot over" the other elements to fill in or close gaps, which takes worst-case O(n) time.

### Data structures built on arrays

- Don't want to specify the size of your array upfront? One option: use a **dynamic array**
- Want to look up items by something other than an index? Use a **hash table**

### Array slicing

Array slicing involves taking a subset from an array and allocating a new array with those elements.

**Careful**: there's a hidden time and space cost here! It's tempting to think of slicing as just "getting elements," but in reality you are:
- allocating a new array
- copying the elements from the original array to the new array

This takes O(n) time and O(n) space, where n is the number of elements in the resulting array.

### In-place and out-of-place algorithms

An in-place algorithm operates directly on its input and changes it, instead of creating and returning a new object. This is sometimes called destructive, since the original input is "destroyed" (or modified) during the function call.

```go
func squareArrayInPlace(arr []int) {
    for i := 0; i < len(arr); i++ {
        arr[i] = arr[i] * arr[i]
    }
}
```

An out-of-place algorithm doesn't make changes directly to its input. It makes a copy instead. The original input stays intact. This is sometimes called non-destructive.

```go
func squareArrayOutOfPlace(arr []int) []int {
    newArr := make([]int, len(arr))
    for i := 0; i < len(arr); i++ {
        newArr[i] = arr[i] * arr[i]
    }
    return newArr
}
```

Working in-place is a good way to save time and space. But be careful: an in-place algorithm can cause side effects. 

Generally, out-of-place algorithms are considered safer because they avoid side effects. But be aware that they require extra space.

## Dynamic array

A dynamic array is an array with a big improvement: automatic resizing.

### Common operations

| Operation | Average case | Worst case |
| --------- |-------------- |------------|
| Space     | O(n)         | O(n)      |
| Lookup    | O(1)         | O(1)       |
| Append    | O(1)         | O(n)       |
| Insert    | O(n)         | O(n)       |
| Delete    | O(n)         | O(n)       |

### Strengths

- Fast lookups
- Variable size
- Cache-friendly

### Weaknesses

- Appending is usually amortized O(1), but can be O(n)
- Costly inserts and deletes

### Size vs capacity

A dynamic array's **size** is the number of elements it actually contains, while its **capacity** is the maximum number of elements it can contain without resizing.

What if we try to append an item but our array's capacity is already full?

To make room, dynamic arrays automatically make a new, bigger underlying array. Usually twice as big. Each item has to be individually copied into the new array.

## Code examples and exercises

- [Code example of static array implementation in Go](data-structures/staticarray)
- [Code example of dynamic array implementation in Go](data-structures/dynamicarray)
- [Exercises](exercises)





