# Big-O notation

## Definition

Big-O notation is used to describe the efficiency of an algorithm in terms of time and space complexity.

In timme complexity, Big-O notation describes the amount of time taken by an algorithm to run relative to the size of the input (or "How many steps the computer has to perform for each function").

In space complexity, Big-O notation describes the amount of memory taken by an algorithm to run relative to the size of the input.

## What can cause time in a function?

- Operations (+, -, *, /)
- Comparisions (<, >, ==)
- Looping (for, while)
- Outside function calls

## What causes space complexity?

- Variables declaration
- Data structures size
- Function calls
- Dynamic memory allocations (e.g., using malloc in C or new in Go)

### Notes

Don't include space taken up by the inputs to the function in the space complexity calculation.

## Big-Os

![Big-O Complexity Chart](assets/big_o_complexity_chart.png) 

| Big-O      | Name        | Description                                                                                                                                                                         |
|------------|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| O(1)       | Constant    | The complexity is constant, regardless of the size of the input.                                                                                                                    |
| O(log n)   | Logarithmic | The complexity grows logarithmically with the size of the input. Example: searching on sorted collection (binary search).                                                           |
| O(n)       | Linear      | The complexity grows linearly with the size of the input. Example: functions having for, while loops.                                                                               |
| O(n log n) | Log Linear  | The complexity grows linearithmically with the size of the input. Example: Sorting algorithms.                                                                                      |
| O(n^2)     | Quadratic   | The complexity grows quadratically with the size of the input. Example: Functions have two nested loops, every element in a collection needs to be compared to every other element. |
| O(2^n)     | Exponential | The complexity grows exponentially with the size of the input. Example: Recursive algorithms that solve a problem of size n                                                         |
| O(n!)      | Factorial   | The complexity grows factorially with the size of the input. You are adding a loop for every element. (DONT write code have this complexity)                                        |

### Notes

- N could be the actual input, or the size of the input
- Iterating through half a collection is still O(n)
- Iterating through two separate collections is O(a+b)
- You should develop the skill to see time and space optimizations, as well as the wisdom to judge if those optimizations are worthwhile.


