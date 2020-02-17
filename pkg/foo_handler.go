package foo

import (
	"net/http"
	"fmt"
)

func Foo_handler (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "\nThis is imported Kiran's foo package\n")
}
