/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Codec {
public:

    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        string serialized = "";
        recursiveSerialize(root, serialized);
        return serialized;
    }
    
    void recursiveSerialize(TreeNode* root, string &serialized) {
        if (root == NULL) {
            serialized += "null,";
        } else {
            serialized += to_string(root->val) + ",";
            recursiveSerialize(root->left, serialized);
            recursiveSerialize(root->right, serialized);
        }
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        queue<string> vals;
        string s;
        for (int i = 0; i < data.size(); i++) {
            if (data[i] == ',') {
                vals.push(s);
                s = "";
            } else {
                s += data[i];
            }
        }
        return recursiveDeserialize(vals);
    }
    
    TreeNode* recursiveDeserialize(queue<string>& vals) {
        TreeNode* root = new TreeNode(0);
        if (vals.empty()) {
            return NULL;
        }
        if (vals.front() == "null") {
            vals.pop();
            return NULL;
        }
        root->val = stoi(vals.front());
        vals.pop();
        root->left = recursiveDeserialize(vals);
        root->right = recursiveDeserialize(vals);
        return root;
    }
};

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.deserialize(codec.serialize(root));
