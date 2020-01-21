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
    vector<string> binaryTreePaths(TreeNode* root) {
        vector<string> ans;
        vector<int> path;
        helper(root, path, ans);
        return ans;
    }
    
    void helper(TreeNode *root, vector<int> &path, vector<string> &ans) {
        if (root == NULL) {
            return;
        }
        path.push_back(root->val);
        if ((root->left == NULL) && (root->right == NULL) && (path.size() > 0)) {
            string s = to_string(path[0]);
            for (int i = 1; i < path.size(); i++)
                s += "->" + to_string(path[i]);
            ans.push_back(s);
        }
        helper(root->left, path, ans);
        helper(root->right, path, ans);
        path.pop_back();
    }
};
