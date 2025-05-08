#include <stdio.h>

void minmax() {
  int first = 1;
  int val, minval, maxval;
  while (scanf("%d", &val) != EOF) {
    if (first || val > maxval)
      maxval = val;
    if (first || val < minval)
      minval = val;
    first = 0;
  }
  printf("maximum: %d\n", maxval);
  printf("minimum: %d\n", minval);
}

int main() {
  int i;
  for (i = 0; i < 5; i++) {
    printf("%d\n", i);
  }
  minmax();
}
