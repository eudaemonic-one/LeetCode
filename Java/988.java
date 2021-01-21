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
    public String smallestFromLeaf(TreeNode root) {
        if (root == null) {
            return "";
        }
        String[] res = new String[]{"~"};
        StringBuilder sb = new StringBuilder();
        dfs(root, sb, res);
        return res[0];
    }
    
    private void dfs(TreeNode root, StringBuilder sb, String[] res) {
        if (root == null) {
            return;
        }
        sb.append((char)('a'+root.val));
        if (root.left == null && root.right == null) {
            sb.reverse();
            String str = sb.toString();
            sb.reverse();
            if (str.compareTo(res[0]) < 0) {
                res[0] = str;
            }
        }
        dfs(root.left, sb, res);
        dfs(root.right, sb, res);
        sb.deleteCharAt(sb.length()-1);
    }
}