package serve

import (
"net/http"
"os"

"github.com/go-chi/chi"
)

func InitServeRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/*", staticHandler)
	return r
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	assetsDir := "dist"
	if _, err := os.Stat("./" + assetsDir + r.URL.Path); os.IsNotExist(err) {
		http.ServeFile(w, r, "./"+assetsDir+"/index.html")
	} else {
		http.ServeFile(w, r, "./"+assetsDir+r.URL.Path)
	}
}

