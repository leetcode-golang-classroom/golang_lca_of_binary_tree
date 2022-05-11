# golang_lca_of_binary_tree

Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes `p` and `q` as the lowest node in `T` that has both `p` and `q` as descendants (where we allow **a node to be a descendant of itself**).”

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.

```

**Example 2:**

![https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

```

**Example 3:**

```
Input: root = [2,1], p = 2, q = 1
Output: 2

```

**Constraints:**

- The number of nodes in the tree is in the range `[2,$10^5$]`.
- $`-10^9$ <= Node.val <= $10^9$`
- All `Node.val` are **unique**.
- `p != q`
- `p` and `q` will exist in the BST.

## 解析

題目給定一個二元搜尋樹的根結點，還有兩個結點p, q

要求 p, q 兩個結點的最小共同祖先結點

要解出這題需要先了解二元搜尋樹的特性

二元搜尋樹的特性如下

1. 每個結點的值比左子樹下每個結點的值大
2. 每個結點的值比右子樹下每個結點的值小

所以這題要找最小共同祖先

這個共同祖先 root 必須符合幾個特性

1 Min(p.Val, q.Val) ≤ root.Val

2 Max(p.Val, q.Val) ≥ root.Val

因為是二元搜尋樹所以再搜詢時

可以照以下條件搜尋

當 Min(p.Val, q.Val) > root.Val ，向右子樹搜訊

當 Max(p.Val, q.Val) < root.Val ，向左子樹搜訊

如果都不再兩個條件代表當下結點就是共同組結點，因為由根結點往下搜尋，因此會是最小子結點

## 程式碼

## 困難點

1. 需要理解 BST(二元搜尋樹)特性
2. 需要理解 [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor) 最小共同祖結點定義

## Solve Point

- [x]  Understand what problem need to be solved
- [x]  Analysis Complexity