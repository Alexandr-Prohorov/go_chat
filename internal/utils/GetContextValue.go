package utils

import "net/http"

func GetContextValue(w http.ResponseWriter, r *http.Request) string {
	username := r.Context().Value("username").(string)
	if username == "" {
		http.Error(w, "Username not found", http.StatusUnauthorized)
		return ""
	}

	return username
}
