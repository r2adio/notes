// gcc 01open-window/main.c -o test $(pkg-config --cflags --libs sdl3)
#include <SDL3/SDL.h>
#include <SDL3/SDL_main.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

// #define SDL_FLAGS (SDL_INIT_VIDEO | SDL_INIT_AUDIO)
#define SDL_FLAGS SDL_INIT_VIDEO

struct Game {
  SDL_Window *window;
  SDL_Renderer *render;
};

bool game_init_sdl() {
  if (!SDL_Init(SDL_FLAGS)) {
    fprintf(stderr, "Error initializing SDL3: %s\n", SDL_GetError());
    return false;
  }
  return true;
}

void game_free() { SDL_Quit(); }

int main() {
  bool exit_status = EXIT_FAILURE;
  if (game_init_sdl()) {
    exit_status = EXIT_SUCCESS;
  }
  game_free();
  return exit_status;
}
