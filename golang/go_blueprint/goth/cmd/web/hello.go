package web

import (
	"log"
	"net/http"
)

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //r.ParseForm(): parse encoded form data; and loads entire form into memory
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return //stops the execution
	}

	name := r.FormValue("name")
	component := HelloPost(name)           //component
	err = component.Render(r.Context(), w) //component rendering
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
		return
	}
}
