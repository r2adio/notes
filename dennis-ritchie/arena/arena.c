#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ALIGNMENT 8

typedef struct Arena {
  struct Arena *next;
  size_t cap;
  size_t len;
  uint8_t *buf;
} Arena;

static size_t align_size(size_t size) {
  return (size + ALIGNMENT - 1) & ~(ALIGNMENT - 1);
}

Arena *arena_init(size_t cap) {
  Arena *arena = malloc(sizeof(Arena));
  if (arena == NULL) {
    return NULL;
  }

  arena->buf = malloc(cap);
  if (arena->buf == NULL) {
    free(arena);
    return NULL;
  }

  arena->cap = cap;
  arena->len = 0;
  arena->next = NULL;
  return arena;
}

void *arena_alloc(Arena *arena, size_t len) {
  if (arena == NULL) {
    return NULL;
  }

  size_t aligned_len = align_size(len);
  Arena *this = arena;

  while (this->len + aligned_len > this->cap) {
    if (this->next == NULL) {
      this->next = arena_init(this->cap * 2);
      if (this->next == NULL) {
        return NULL;
      }
    }
    this = this->next;
  }

  uint8_t *buf = &this->buf[this->len];
  this->len += aligned_len;
  return buf;
}

void arena_reset(Arena *arena) {
  if (arena == NULL)
    return;

  Arena *this = arena;
  while (this != NULL) {
    this->len = 0;
    this = this->next;
  }
}

size_t arena_total_allocated(Arena *arena) {
  if (arena == NULL) {
    return 0;
  }

  size_t total = 0;
  Arena *this = arena;
  while (this != NULL) {
    total += this->len;
    this = this->next;
  }
  return total;
}

size_t arena_total_capacity(Arena *arena) {
  if (arena == NULL) {
    return 0;
  }

  size_t total = 0;
  Arena *this = arena;
  while (this != NULL) {
    total += this->cap;
    this = this->next;
  }
  return total;
}

void arena_free(Arena *arena) {
  Arena *this = arena;
  while (this != NULL) {
    Arena *next = this->next;
    free(this->buf);
    free(this);
    this = next;
  }
}

int main(void) {
  Arena *arena = arena_init(1024);
  if (arena == NULL) {
    printf("Failed to initialize arena\n");
    return 1;
  }

  void *p = arena_alloc(arena, 100);
  if (p == NULL) {
    printf("Failed to allocate 100 bytes\n");
    arena_free(arena);
    return 1;
  }

  void *p1 = arena_alloc(arena, 1000);
  if (p1 == NULL) {
    printf("Failed to allocate 1000 bytes\n");
    arena_free(arena);
    return 1;
  }

  printf("First arena - capacity: %zu, allocated: %zu, pointer: %p\n",
         arena->cap, arena->len, arena->buf);

  if (arena->next != NULL) {
    printf("Second arena - capacity: %zu, allocated: %zu, pointer: %p\n",
           arena->next->cap, arena->next->len, arena->next->buf);
  }

  printf("Total capacity: %zu, Total allocated: %zu\n",
         arena_total_capacity(arena), arena_total_allocated(arena));

  arena_free(arena);
  return 0;
}
