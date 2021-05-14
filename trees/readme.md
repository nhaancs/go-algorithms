# Trees
Applications: DOM, abstract syntax tree in programming languages, Facebook comments,...

![tree](./images/tree.jpg)

## Binary tree

![binary tree](./images/binary_tree.jpg)

### Perfect binary tree vs full binary tree

![perfect vs full binary tree](./images/perfect_vs_full_binary_tree.jpg)

Perfect binary tree attributes:
- The number of total nodes on each level doubled as we move down the tree
- The number of nodes on the last level is equal to the sum of the number of nodes on the other levels plus one.
- The number of nodes on each level = 2^level - 1, the first level start from 1. For example: 2^1 - 1, 2^2 - 1, 2^3 - 1,...

## Binary search tree 
Properties:
- Each node has a distinguished key which is unique in the tree.
- A node can only have up to two children.
- The key of any node is greater than all keys occurring in its left subtree and less than all keys occurring in its right subtree.

Pros: 
- Ordered
- Fast searchings and lookups
- Better than O(n)
- Flexible size
- Better than hash tables because it has parent-children structure that preserves the relationships.
Cons:
- Searchings and lookups are inefficient on a very unbalanced binary search tree => O(n)
- No O(1) operations

### BigO
- Lookup O(log N)
- Insert O(log N) // have to search for where to insert item first
- Delete O(log N) // have to search for where to delete item first

```
// maxsteps is the max level of the tree
maxnodes = 2^maxsteps - 1 
<=> 2^maxsteps = maxnodes // drop -1
<=> maxsteps = log maxnodes // 10^2 = 100 <=> 2 = log 100
```