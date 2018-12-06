#include <stdio.h>

struct Node {
  int data;
  struct Node *next;
  Node(int i, struct Node* ni) {
    data = i;
    next = ni;
  }
};

bool has_cycle(Node *head) {
  if (head == NULL) return false;
  if (head->next == head) return true;
  Node *A = head;
  Node *B = head->next;
  while(A != NULL && B != NULL) {
    if (A == B) return true;
    A = A->next;
    B = B->next;
    if (B == NULL) return false;
    B = B->next;
  }
  return false;
}

int main() {
  Node *head = new Node(1, NULL);
  Node *first = new Node(2, NULL);
  Node *second = new Node(3, NULL);
  Node *third = new Node(4, NULL);
  head->next = first;
  first->next = second;
  second->next = third;
  third->next = first;
  printf("ans is %s\n", has_cycle(head) ? "true" : "false");
  return 0;
}
