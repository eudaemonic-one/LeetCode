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
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        int ps = 0, pe = preorder.size()-1;
        map<int, int> indexTable;
        for (int i = 0; i < inorder.size(); i++) {
            indexTable[inorder[i]] = i;
        }
        return helper(preorder, inorder, indexTable, ps, pe, 0, inorder.size()-1);
    }
    
    TreeNode* helper(vector<int>& preorder, vector<int>& inorder, map<int, int>& indexTable, int& ps, int& pe, int is, int ie) {
        if (ps > pe || is > ie) {
            return NULL;
        }
        TreeNode *root = new TreeNode(preorder[ps++]);
        int idx = indexTable[root->val];
        root->left = helper(preorder, inorder, indexTable, ps, pe, is, idx-1);
        root->right = helper(preorder, inorder, indexTable, ps, pe, idx+1, ie);
        return root;
    }
};
