// gcc 01open-window/main.c -o test $(pkg-config --cflags --libs sdl3)
#include <SDL3/SDL.h>
#include <SDL3/SDL_main.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

// #define SDL_FLAGS (SDL_INIT_VIDEO | SDL_INIT_AUDIO)
#define SDL_FLAGS SDL_INIT_VIDEO

#define WINDOW_TITLE "Open Window"
#define WINDOW_WIDTH 800
#define WINDOW_HEIGHT 600

struct Game {
  SDL_Window *window;
  SDL_Renderer *render;
};

bool game_init_sdl(struct Game *g) {
  if (!SDL_Init(SDL_FLAGS)) {
    fprintf(stderr, "Error initializing SDL3: %s\n", SDL_GetError());
    return false;
  }

  g->window = SDL_CreateWindow(WINDOW_TITLE, WINDOW_WIDTH, WINDOW_HEIGHT, 0);
  if (!g->window) {
    fprintf(stderr, "Error creatting Window: %s\n", SDL_GetError());
    return false;
  }

  g->render = SDL_CreateRenderer(g->window, NULL);
  if (!g->render) {
    fprintf(stderr, "Error creatting Render: %s\n", SDL_GetError());
    return false;
  }

  return true;
}

void game_free(struct Game *g) {
  if (g->window) {
    SDL_DestroyWindow(g->window);
    g->window = NULL;
  }
  if (g->render) {
    SDL_DestroyRenderer(g->render);
    g->render = NULL;
  }
  SDL_Quit();
}

int main() {
  bool exit_status = EXIT_FAILURE;

  struct Game game = {0};

  if (game_init_sdl(&game)) {
    exit_status = EXIT_SUCCESS;
  }
  SDL_Delay(5000);

  game_free(&game);
  return exit_status;
}
