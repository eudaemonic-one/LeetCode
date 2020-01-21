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
        int pVal = p->val, qVal = q->val;
        while (root) {
            if (pVal < root->val && qVal < root->val) {
                root = root->left;
            } else if (pVal > root->val && qVal > root->val) {
                root = root->right;
            } else {
                return root;
            }
        }
        return NULL;
    }
};
