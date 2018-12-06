#include <iostream>

struct Node {
  int data;
  Node* left;
  Node* right;
};

int min(Node* root) {
  int ans = 100000;
  while(root != NULL) {
    ans = root->data;
    root = root->left;
  }
  return ans;
}

int max(Node* root) {
  int ans = -1;
  while(root != NULL) {
    ans = root -> data;
    root = root->right;
  }
  return ans;
}

bool checkBST(Node* root) {
  if (root == NULL) return false;
  if (root->left != NULL) {
    bool leftChildBst = checkBST(root->left);
    if (!leftChildBst) return false;
    int maxv = max(root->left);
    if (root->data <= maxv) return false;
  }
  if (root->right != NULL) {
    bool rightChildBst = checkBST(root->right);
    if (!rightChildBst) return false;
    int minv = min(root->right);
    if (root->data >= minv) return false;
  }
  return true;
}

int main() {
  return 0;
}