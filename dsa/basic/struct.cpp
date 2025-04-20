// STRUCT
// a structure that groups related variables under one name,
// variables in struct are known as "members", can be data members or member fxn
// members can be accessed with . {class member access operator
// inside main memory, struct object s1 is created inside??
// definition doesnt consume any memory bu itself, unless an obj is declared

#include "iostream"
struct student {
  // fields of struct, student
  int roll;
  char name[10];
  std::string address;
};
struct cards {
  int face;
  int shape;
  int color;
};
int main(int argc, char *argv[]) {
  // declares a struct obj, s1
  struct student s1;
  // initializes values for all fields of struct object
  s1.roll = 1;
  // s1.name = "x-furate"; // error
  // correct way to assign a string to a char array
  std::strcpy(s1.name, "x-furate");
  s1.address = "sa";

  // declaration+initializing
  struct student s2 = {12, "furate", "as"};

  struct cards deck[52] = {{1, 0, 0}, {2, 0, 0}, {3, 0, 1}};
  deck[10].face = 13;
  deck[10].shape = 3;
  deck[10].color = 1;
  std::cout << deck[0].face << '\n';
  std::cout << deck[10].face << '\n';

  return 0;
}
