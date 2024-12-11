package csshandle

import (
	"net/http"
	"path/filepath"
)

// Register Decoration
func RegisterHandleCSS(w http.ResponseWriter, r *http.Request) {
	filepath := filepath.Join("Style", "register.css")
	http.ServeFile(w, r, filepath)
}

// Login Decoration
func LoginHandleCss(w http.ResponseWriter, r *http.Request) {
	filepath := filepath.Join("Style", "login.css")
	http.ServeFile(w, r, filepath)
}

// Index Decoration
func IndexHandleCSS(w http.ResponseWriter, r *http.Request) {
	filepath := filepath.Join("Style", "index.css")
	http.ServeFile(w, r, filepath)
}
