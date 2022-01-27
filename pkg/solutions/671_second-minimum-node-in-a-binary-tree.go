package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	//节点为叶子节点
	if root.Left == nil && root.Right == nil {
		return -1
	}
	//节点的两个子节点值相等
	if root.Left.Val == root.Right.Val {
		if root.Left.Val == root.Val {
			lm := findSecondMinimumValue(root.Left)
			rm := findSecondMinimumValue(root.Right)
			if lm*rm < 0 {
				return max(lm, rm)
			} else {
				return min(lm, rm)
			}
		} else {
			return root.Left.Val
		}
	} else {
		minx := min(root.Left.Val, root.Right.Val)
		if minx > root.Val {
			return min(root.Left.Val, root.Right.Val)
		} else { // minx == root.Val
			maxn := max(root.Left.Val, root.Right.Val)
			var tmpNode *TreeNode
			if root.Left.Val < root.Right.Val {
				tmpNode = root.Left
			} else {
				tmpNode = root.Right
			}
			value := findSecondMinimumValue(tmpNode)
			if value < 0 {
				return maxn
			} else {
				return min(value, maxn)
			}
		}
	}
}

func FindSecondMinimumValue(root *TreeNode) int {
	return findSecondMinimumValue(root)
}

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }

// func max(a, b int) int {
//     if a < b {
//         return b
//     }
//     return a
// }
