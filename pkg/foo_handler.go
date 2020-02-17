package foo

import (
	"net/http"
	"fmt"
)

func Foo_handler (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "\n ====  v2 =====  This is UPGRADED version of Kiran's foo package\n")
}
