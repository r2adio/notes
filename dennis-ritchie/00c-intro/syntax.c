#include <stdio.h>
int main() {
  // NUMBER INPUT
  int usf, euf;
  printf("Enter US Floor:\n");
  // calling by reference
  scanf("%d", &usf);
  euf = usf - 1;
  printf("EU Floor: %d\n", euf);

  // STRING INPUT
  char name[100]; // string is a character array
  printf("Enter name:\n");
  // no pass by reference, as name is array
  // name w/ no [] or index operator, then u're passing the address of the
  // beginning of array
  scanf("%99s", name);
  printf("hello %s\n", name);

  // LINE INPUT
  char line[100];
  printf("enter a line:\n");
  // after scanf("%99s", name); the newline from pressing Enter remains in
  // stdin, and fgets() will immediately consume it, giving you an empty string.
  getchar(); // consume leftover newline
  // scanf("%[^\n] 100s", line); // match any char that's not a new line
  fgets(line, 100, stdin); // safer approach
  printf("Line: %s\n", line);

  // READ FILE
  char statement[1000];
  FILE *hand;
  hand = fopen("test.md", "r");
  while (fgets(statement, 1000, hand) != NULL) {
    printf("%s", statement);
  }
}
