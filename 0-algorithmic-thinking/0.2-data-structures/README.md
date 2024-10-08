# Data Structures

## Random Access Memory (RAM)

Variables are stored in random access memory (RAM). We sometimes call RAM "working memory" or just "memory".

In addition to "memory," your computer has storage (sometimes called "persistent storage" or "disk"). Storage is where we keep files like mp3s, videos, Word documents, and even executable programs or apps.

Memory (or RAM) is faster but has less space, while storage (or "disk") is slower but has more space. A modern laptop might have ~500GB of storage but only ~16GB of RAM.

![Processor - Memory controller - RAM](assets/processor-memory-controller-ram.svg)

Think of RAM like a really tall bookcase with a lot of shelves. Like, billions of shelves. 

The shelves are numbered. We call a shelf's number its address. Each shelf has one byte (8 bits) of storage.

We also have a processor that does all the real work inside our computer. It's connected to a memory controller. 

The memory controller does the actual reading and writing to and from RAM. 

Even though the memory controller can jump between far-apart memory addresses quickly, programs tend to access memory that's nearby. So computers are tuned to get an extra speed boost when reading memory addresses that're close to each other by using cache which even faster than RAM. 

When the processor asks for the contents of a given memory address, the memory controller also sends the contents of a handful of nearby memory addresses. And the processor puts all of it in the cache.

## Binary numbers

The number system we usually use (the one you probably learned in elementary school) is called base 10 (decimal), because each digit has ten possible values (1, 2, 3, 4, 5, 6, 7, 8, 9, and 0).

![Decimal numbers](assets/base-10-numbers.svg)

But computers don't have digits with ten possible values. They have bits with two possible values. So they use base 2 numbers (binary).

![Binary numbers](assets/base-2-numbers.svg)

So far we've been talking about unsigned integers. Here's how some other numbers could be stored:
- Fractions: Store two numbers: the numerator and the denominator.
- Decimals: Also two numbers: 1) the number with the decimal point taken out, and 2) the position where the decimal point goes (how many digits over from the leftmost digit).
- Negative Numbers: Reserve the leftmost bit for expressing the sign of the number. 0 for positive and 1 for negative.

## Fixed-width integers

The 256 possibilities we get with 1 byte are pretty limiting. So we usually use 4 or 8 bytes (32 or 64 bits) for storing integers.

- 32-bit integers have 2^32 possible values — more than 4 billion
- 64-bit integers have 2^64 possible values — more than 10 billion billion (10^19)

If we have a 64-bit fixed-length integer, it doesn't matter if that integer is 0 or 193,457 — it still takes up the same amount of space in RAM: 64 bits.

## Arrays

An array is a collection of items stored at **contiguous** memory locations. The idea is to store multiple items of the **same size** together. Looking up the contents of a given array index is O(1) time.

For example, we can store 5 8-bit integers in a single array variable.

```
nth item's memory address = address of array start + n
```

![Array](assets/array.svg)

So how do we build an array of 64-bit (8 byte) integers on top of our 8-bit (1 byte) memory slots?

We simply give each array index 8 address slots instead of 1.

```
nth item's memory address = address of array start + (n * size of each item in bytes)
```

![Array of 64-bit integers](assets/array-of-64-bit-integers.svg)

Advantages:
- Cache-friendly. All items are stored next to each other in memory.
- Lookups are O(1) time. We can jump right to the memory address of the item we're looking for.

Tradeoffs:
- We have to know the size of the array ahead of time.
- We need enough uninterrupted free memory to store all items next to each other.

## Strings

A string is an array of characters. Each character is usually 1 byte (8 bits) long.

![String](assets/strings.svg)

## Pointers

A pointer is a variable that stores the memory address of another variable.

Array of pointers is an array of memory addresses.

Advantages: 
- Items don't have to be the same length. Each item can be any size.
- We don't need enough uninterrupted free memory to store all items next to each other.

Tradeoffs:
- An array of pointers is not cache-friendly (but lookups in this array are still O(1))

![Pointers array](assets/pointers-array.svg)

## Dynamic arrays

A dynamic array is an array that can resize itself when it runs out of space.

When a dynamic array is full, it doesn't have enough contiguous free memory to store new items. So it has to move to a new, larger location in memory.

Append an item to a dynamic array that's full:

![Dynamic array](assets/dynamic-array-1.svg)

![Dynamic array](assets/dynamic-array-2.svg)

1. Allocate a new, larger chunk of memory. Usually twice as big.
2. Copy all items from the old memory location to the new one.
3. Add the new item to the new memory location.
4. Free the old memory location.



Advantages:
- You don't need to specify the size ahead of time.

Tradeoffs:
- Som appends can be expensive.
- Expensive prepend operation.

## Linked lists

A linked list is a collection of items where each item points to the next one in the sequence.

- Each item in a linked list is called a "node" with two components: The value and a pointer to the next item
- The first item in the list is called the "head"
- The last item in the list is called the "tail"
- The last item in the list has a pointer to nothing (null)

![Linked list](assets/linked-list.svg)


It's important to have a pointer variable referencing the head of the list—otherwise we'd be unable to find our way back to the start of the list!

We'll also sometimes keep a pointer to the tail. That comes in handy when we want to add something new to the end of the linked list.

![Linked list operations](assets/linked-list-operation.svg)

Advantages:
- Prepend is O(1) time if we keep a pointer to the head.
- Append is O(1) time if we keep a pointer to the tail.
- We don't need to know the size of the list ahead of time.

Tradeoffs:
- Not cache-friendly. Because the next item in the list could be anywhere in memory.
- Lookups are O(n) time. We have to start at the head and follow the pointers until we find the item we're looking for.

## Hash tables

A hash table is a collection of key-value pairs. It uses a hash function to map keys to values.

![Hash table](assets/hash-table.svg)

Advantages:
- Lookups, inserts, and deletes are O(1) time on average.

Tradeoffs:
- Hash collisions can be a problem, two keys might hash to the same key in the hash table which means some lookups could be a bit slow O(n). 

![Hash table collisions](assets/hash-table-collisions.svg)

We can use a linked list at each index to handle collisions. There are fancy algorithms that keep the number of collisions low and keep the lengths of our linked lists nice and short. Collisions are rare enough that on average lookups in a hash table are O(1) time

