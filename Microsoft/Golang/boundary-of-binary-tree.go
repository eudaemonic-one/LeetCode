/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func boundaryOfBinaryTree(root *TreeNode) []int {
    res := make([]int, 0)
    if root == nil {
        return res
    }
    leftBoundary := make([]int, 0)
    leaves := make([]int, 0)
    rightBoundary := make([]int, 0)
    leftBoundary, leaves, rightBoundary = dfs(root, leftBoundary, leaves, rightBoundary, 0)
    res = append(res, leftBoundary...)
    res = append(res, leaves...)
    res = append(res, rightBoundary...)
    return res
}

func dfs(root *TreeNode, left, leaves, right []int, flag int) ([]int, []int, []int) {
    if root == nil {
        return left, leaves, right
    }
    if isRightBoundary(flag) {
        right = append([]int{root.Val}, right...)
    } else if isLeftBoundary(flag) || isRoot(flag) {
        left = append(left, root.Val)
    } else if isLeaf(root) {
        leaves = append(leaves, root.Val)
    }
    left, leaves, right = dfs(root.Left, left, leaves, right, leftChildFlag(root, flag))
    left, leaves, right = dfs(root.Right, left, leaves, right, rightChildFlag(root, flag))
    return left, leaves, right
}

func isRightBoundary(flag int) bool {
    return flag == 2
}

func isLeftBoundary(flag int) bool {
    return flag == 1
}

func isRoot(flag int) bool {
    return flag == 0
}

func isLeaf(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return root.Left == nil && root.Right == nil
}

func leftChildFlag(root *TreeNode, flag int) int {
    if isLeftBoundary(flag) || isRoot(flag) {
        return 1
    } else if isRightBoundary(flag) && root.Right == nil {
        return 2
    } else {
        return 3
    }
}

func rightChildFlag(root *TreeNode, flag int) int {
    if isRightBoundary(flag) || isRoot(flag) {
        return 2
    } else if isLeftBoundary(flag) && root.Left == nil {
        return 1
    } else {
        return 3
    }
}
