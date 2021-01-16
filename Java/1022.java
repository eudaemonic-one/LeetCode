/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public int sumRootToLeaf(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int[] res = new int[1];
        List<Integer> path = new ArrayList<>();
        dfs(root, path, res);
        return res[0];
    }
    
    private void dfs(TreeNode root, List<Integer> path, int[] res) {
        if (root == null) {
            return;
        }
        path.add(root.val);
        if (root.left == null && root.right == null) {
            StringBuilder sb = new StringBuilder();
            for (int x: path) {
                if (x == 0) {
                    sb.append('0');
                } else {
                    sb.append('1');
                }
            }
            res[0] += Integer.parseInt(sb.toString(), 2);
        }
        dfs(root.left, path, res);
        dfs(root.right, path, res);
        path.remove(path.size()-1);
    }
}