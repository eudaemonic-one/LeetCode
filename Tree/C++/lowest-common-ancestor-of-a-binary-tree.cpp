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
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        TreeNode *ans;
        helper(root, p->val, q->val, ans);
        return ans;
    }
    
    bool helper(TreeNode *root, int &pVal, int &qVal, TreeNode *&ans) {
        bool left, curr, right;
        if (root == NULL)
            return false;
        left = helper(root->left, pVal, qVal, ans);
        curr = ((root->val == pVal) || (root->val == qVal));
        right = helper(root->right, pVal, qVal, ans);
        if ((left && curr) || (curr && right) || (left && right))
            ans = root;
        return left || curr || right;
    }
};
