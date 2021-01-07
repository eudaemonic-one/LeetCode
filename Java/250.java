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
    private int count = 0;
    
    public int countUnivalSubtrees(TreeNode root) {
        count = 0;
        helper(root, 0);
        return count;
    }
    
    private boolean helper(TreeNode node, int val) {
        if (node == null) {
            return true;
        }
        if (!helper(node.left, node.val) | !helper(node.right, node.val)) {
            return false;
        }
        count++;
        return node.val == val;
    }
}