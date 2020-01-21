/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int closestValue(TreeNode* root, double target) {
        int closest = root->val;
        vector<int> inorder;
        if (root == NULL)
            return 0;
        helper(root, inorder);
        for (auto val : inorder) {
            if (abs(val-target) < abs(closest-target))
                closest = val;
        }
        return closest;
    }
    
    void helper(TreeNode *node, vector<int> &inorder) {
        if (node == NULL)
            return;
        inorder.push_back(node->val);
        helper(node->left, inorder);
        helper(node->right, inorder);
    }
};
