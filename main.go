package main

import (
	"chi"
	"http"
)

func main() {
	r := chi.NewRouter()

	r.Post("/api/compile/go", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	http.ListenAndServe(":3000", r)
}
