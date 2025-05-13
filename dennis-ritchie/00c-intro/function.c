#include <stdio.h>
// pass by value
int value(int a, int b) {
  int c = a * b;
  return c;
}
// pass by reference
int reference() { return 0; }

int main() {
  int a, b;
  int retval = value(6, 7);
  printf("Answer: %d\n", retval);
  return 0;
}
