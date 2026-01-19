#include <assert.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Arena {
  struct Arena *next;
  size_t cap;
  size_t len;
  uint8_t *buf;
} Arena;

Arena arena_init(size_t cap) {
  void *buf = malloc(sizeof(uint8_t) * cap);
  if (buf == NULL) {
    return (Arena){0};
  }
  Arena arena = {.cap = cap, .len = 0, .buf = buf, .next = NULL};
  return arena;
}

void *arena_alloc(Arena *arena, size_t len) {
  Arena *this = arena;

  while (!(this->len + len <= this->cap)) {
    if (this->next == NULL) {
      Arena *next = malloc(sizeof(Arena));
      if (next == NULL)
        return NULL;
      *next = arena_init(this->cap * 2);
      this->next = next;
    }
    this = this->next;
  }

  uint8_t *buf = &this->buf[this->len];
  this->len += len;
  return buf;
}

void arena_reset(Arena *arena) { arena->len = 0; }

void arena_free(Arena *arena) {
  arena->len = 0;
  arena->cap = 0;
  free(arena->buf);
}

int main(void) {
  Arena arena = arena_init(1024);
  void *p = arena_alloc(&arena, 100);
  void *p1 = arena_alloc(&arena, 1000);
  printf("capacity of arena: %zu, allocated bytes: %zu, pointer: %p\n",
         arena.cap, arena.len, arena.buf);
  arena_free(&arena);
  return 0;
}
