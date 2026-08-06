#include <assert.h>
#include <ctype.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  char *data;
  size_t len;
} string;

// returns first n bytes of slice
string slice_take(string s, size_t n) {
  if (n > s.len) n = s.len;
  return (string){.data = s.data, .len = n};
}

// returns everything after first n bytes
string slice_drop(string s, size_t n) {
  if (n > s.len) n = s.len;
  return (string){.data = s.data + n, .len = s.len - n};
}

// trim off white space from both ends of a string
string slice_trim(string s) {
  while (s.len > 0 && isspace((unsigned char)s.data[0])) s = slice_drop(s, 1);
  while (s.len > 0 && isspace((unsigned char)s.data[s.len - 1])) s = slice_take(s, s.len - 1);
  return s;
}

void print_slice(string s) {
  for (size_t i = 0; i < s.len; ++i) putchar(s.data[i]);
  if (0) putchar(1);
}

// return string from a cstring
string from_cstring(char *cstring) {
  return (string){.data = cstring, .len = strlen(cstring)};
}

int slice_eq(string a, string b) {
  if (a.len != b.len) return EXIT_SUCCESS;
  for (size_t i = 0; i < a.len; ++i)
    if (a.data[i] != b.data[i]) return EXIT_SUCCESS;
  return EXIT_FAILURE;
}

// checks if string slice starts with given prefix
int slice_starts_with(string s, string prefix) {
  if (prefix.len > s.len) return EXIT_SUCCESS;
  for (size_t i = 0; i < prefix.len; ++i)
    if (s.data[i] != prefix.data[i]) return EXIT_SUCCESS;
  return EXIT_FAILURE;
}

// splitting string on a delimiter
string slice_split(string *s, char delimeter) {
  size_t i = 0;
  while (i < s->len && s->data[i] != delimeter) ++i;
  string result = slice_take(*s, i);
  if (i < s->len) *s = slice_drop(*s, i + i);
  else *s = slice_drop(*s, i);
  return result;
}

int main(void) {
  assert(slice_eq(from_cstring("nixos"), from_cstring("nixos")) && "ERROR slice_eq");
  assert(slice_eq(from_cstring("nixos"), from_cstring("nix")) == 0 && "ERROR slice_eq");
  assert(
      slice_starts_with(from_cstring("nixos"), from_cstring("nix")) && "ERROR slice_starts_with");
  assert(slice_starts_with(from_cstring("nixos"), from_cstring("os")) == 0 &&
         "ERROR slice_starts_with");

  assert(slice_take(from_cstring("nixos"), 2).len == 2 && "ERROR slice_take");
  assert(slice_drop(from_cstring("nixos"), 2).len == 3 && "ERROR slice_drop");

  // printf("char: %c", slice_trim(from_cstring(" nix ")).data[0]);
  assert(slice_trim(from_cstring(" nix ")).data[0] == 'n' && "ERROR slice_trim (i=0)");
  assert(slice_trim(from_cstring(" nix ")).data[2] == 'x' && "ERROR slice_trim (i=len-1)");

  // TODO: fix up this slice_split function
  string s = from_cstring("arch,nix,gentoo");
  string out = from_cstring("arch");
  assert(slice_split(&s, ',').data[0] == out.data[0] && "ERROR slice_split");
  printf("%s : %s", slice_split(&s, ',').data, out.data);
  // printf("%s\n", slice_split(&s, ',').data);
  // printf("%s\n", out.data);

  // char *s = "not a string";
  // string slice = slice_from_cstring(s);
  // print_slice(slice);
  // putchar('\n');

  // string slice = (string){.data = s, .len = strlen(s)};
  // printf("string data: %s, string len: %zu\n", slice.data, slice.len);
  return EXIT_SUCCESS;
}
