/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    res := make([][]int, 0)
	if root == nil {
		return res
	}
	path := make([]int, 0)
	dfs(root, sum, path, &res)
	return res
}

func dfs(root *TreeNode, sum int, path []int, res *[][]int) {
    if root == nil {
        return
    }
    path = append(path, root.Val)
    sum -= root.Val
    if root.Left == nil && root.Right == nil && sum == 0 {
        tmp := make([]int, len(path))
        copy(tmp, path)
		*res = append(*res, tmp)
    } else {
        if root.Left != nil {
		    dfs(root.Left, sum, path, res)
        }
	    if root.Right != nil {
		    dfs(root.Right, sum, path, res)
        }
	}
    path = path[:len(path)-1]
    sum += root.Val
}
