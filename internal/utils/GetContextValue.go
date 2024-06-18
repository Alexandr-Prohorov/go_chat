package utils

import "net/http"

func GetContextValue(w http.ResponseWriter, r *http.Request, key string) string {
	value := r.Context().Value(key).(string)
	if value == "" {
		http.Error(w, "Username not found", http.StatusUnauthorized)
		return ""
	}

	return value
}
