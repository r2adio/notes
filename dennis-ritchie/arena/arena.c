#include <assert.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct {
  size_t cap;
  size_t len;
  uint8_t *buf;
} Arena;

Arena arena_init(size_t cap) {
  void *buf = malloc(sizeof(uint8_t) * cap);
  if (buf == NULL) {
    return (Arena){0};
  }
  Arena arena = {.cap = cap, .len = 0, .buf = buf};
  return arena;
}

void *arena_alloc(Arena *arena, size_t len) {
  assert(arena->len + len <= arena->cap);
  uint8_t *buf = &arena->buf[arena->len];
  arena->len += len;
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
  printf("allocated bytes: %zu\n", arena.len);
  arena_free(&arena);
  return 0;
}
