package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

func main() {
	fmt.Println("asdkjfh")

	asdf := func(w http.ResponseWriter, r *http.Request) {}
	hello().Render(context.Background(), os.Stdout)

	http.HandleFunc("/", asdf)
	http.Handle("/hello", templ.Handler(hello()))
	http.ListenAndServe(":3000", nil)
}
