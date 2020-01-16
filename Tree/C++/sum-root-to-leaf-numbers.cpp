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
    int sumNumbers(TreeNode* root) {
        int sum = 0;
        vector<int> path;
        if (root == NULL) {
            return 0;
        }
        helper(root, path, sum);
        return sum;
    }
    
    void helper(TreeNode* root, vector<int> path, int &sum) {
        int base = 0, pathSum = 0;
        if (root == NULL) {
            return;
        }
        path.push_back(root->val);
        if ((root->left == NULL) && (root->right == NULL)) {
            for (int i = path.size()-1; i >= 0; i--) {
                pathSum += path[i] * pow(10, base++);
            }
            sum += pathSum;
        }
        helper(root->left, path, sum);
        helper(root->right, path, sum);
        path.pop_back();
        return;
    }
};
