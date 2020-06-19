package serve

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func InitRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/*", staticHandler)
	return r
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	assetsDir := "dist"
	filePath := "./" + assetsDir + r.URL.Path
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		if r.URL.Path == "" || r.URL.Path == "/" {
			http.ServeFile(w, r, "./"+assetsDir+"/index.html")
			return
		}

		//w.Header().Set(HeaderContentType, MimeApplicationJSON)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))
		return
	}
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./"+assetsDir+"/index.html")
		return
	}
	http.ServeFile(w, r, filePath)
}
