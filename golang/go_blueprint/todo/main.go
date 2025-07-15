package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

const port string = ":3000"

func main() {
	// component := greetTemplate("evening")
	// component.Render(context.Background(), os.Stdout)
	// fmt.Printf("component: %v\n", component)

	// greetTemplate("click me").Render(context.Background(), os.Stdout)
	page().Render(context.Background(), os.Stdout)

	http.Handle("/", templ.Handler(page()))

	fmt.Println("listen to port: 3000")
	http.ListenAndServe(port, nil)
}
