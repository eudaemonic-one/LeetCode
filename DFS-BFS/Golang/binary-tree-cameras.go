/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
    ans := 0
    visited := make(map[*TreeNode]bool)
    visited[nil] = true
    dfs(&ans, root, nil, visited)
    return ans
}

func dfs(ans *int, root *TreeNode, parent *TreeNode, visited map[*TreeNode]bool) {
    if root != nil {
        dfs(ans, root.Left, root, visited)
        dfs(ans, root.Right, root, visited)

        _, containsRoot := visited[root]
        _, containsRootLeft := visited[root.Left]
        _, containsRootRight := visited[root.Right]

        if parent == nil && !containsRoot || !containsRootLeft || !containsRootRight {
            (*ans)++
            visited[parent] = true
            visited[root] = true
            visited[root.Left] = true
            visited[root.Right] = true
        }
    }
}
