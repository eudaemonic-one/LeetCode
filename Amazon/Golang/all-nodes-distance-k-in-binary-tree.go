/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
    if target == nil {
        return []int{}
    }
    res := make([]int, 0)
    dfs(&res, root, target, K)
    return res
}

func dfs(res *[]int, root *TreeNode, target *TreeNode, K int) int {
    if root == nil {
        return -1
    }
    if root == target {
        subTreeAdd(res, root, 0, K)
        return 1
    }
    
    l := dfs(res, root.Left, target, K)
    r := dfs(res, root.Right, target, K)
    
    if l != -1 {
        if l == K {
            *res = append(*res, root.Val)
        } else if l < K {
            subTreeAdd(res, root.Right, l+1, K)
        }
        return l + 1
    } else if r != -1 {
        if r == K {
            *res = append(*res, root.Val)
        } else if r < K {
            subTreeAdd(res, root.Left, r+1, K)
        }
        return r + 1
    }
    
    return -1
}

func subTreeAdd(res *[]int, root *TreeNode, dist int, K int) {
    if root == nil {
        return
    }
    if dist == K {
        *res = append(*res, root.Val)
    } else if dist < K {
        subTreeAdd(res, root.Left, dist+1, K)
        subTreeAdd(res, root.Right, dist+1, K)
    }
}
