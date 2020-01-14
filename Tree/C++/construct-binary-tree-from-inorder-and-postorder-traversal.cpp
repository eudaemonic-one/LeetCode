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
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        int ps = 0, pe = postorder.size()-1;
        map<int, int> indexTable;
        for (int i = 0; i < inorder.size(); i++) {
            indexTable[inorder[i]] = i;
        }
        return helper(inorder, postorder, indexTable, 0, inorder.size()-1, ps, pe);
    }
    
    TreeNode* helper(vector<int>& inorder, vector<int>& postorder, map<int, int>& indexTable, int is, int ie, int& ps, int& pe) {
        if (is > ie || pe < 0) {
            return NULL;
        }
        TreeNode* root = new TreeNode(postorder[pe--]);
        int idx = indexTable[root->val];
        root->right = helper(inorder, postorder, indexTable, idx+1, ie, ps, pe);
        root->left = helper(inorder, postorder, indexTable, is, idx-1, ps, pe);
        return root;
    }
};
