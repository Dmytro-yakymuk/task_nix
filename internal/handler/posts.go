package handler

import (
	"fmt"
	"net/http"
)

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}
