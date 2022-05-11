package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	small := p
	large := q
	if p.Val > q.Val {
		small = q
		large = p
	}
	if root.Val > large.Val {
		return lowestCommonAncestor(root.Left, small, large)
	}
	if root.Val < small.Val {
		return lowestCommonAncestor(root.Right, small, large)
	}
	return root
}
