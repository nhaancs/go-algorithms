# Binary Search 

## Definition
Binary search is an algorithm; its input is a sorted list of elements. If an element you’re looking for is in that list, binary search returns the position where it’s located. Otherwise, binary search returns null.

## Example
- Phonebook lookup
- Dictionary lookup
- Number guessing

## Big O

In the worst case, binary search will take log2(n) steps to run.

### Logarithms
log10(100) mean how many 10s do we multiply together to get 100? The answer is 2: 10x10. So log10(100) = 2

- 10^2 = 100 <=> log10(100) = 2
- 10^3 = 1000 <=> log10(1000) = 3
- 2^3 = 8 <=> log2(8) = 3
- 2^4 = 16 <=> log2(16) = 4

When we talk about running time in Big O notation, log always means log2.

For a list of 8 elements, log 8 == 3, because 2^3 == 8. So for a list of 8 numbers, you would have to check 3 numbers at most. For a list of 1024 elements, log 1024 = 10, because 2^10 == 1024. So for a list of 1024 numbers, you’d have to check 10 numbers at most (instead of 1024).
 