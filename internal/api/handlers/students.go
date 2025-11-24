package handlers

import "net/http"

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET request received at /students"))
	case http.MethodPost:
		w.Write([]byte("POST request received at /students"))
	case http.MethodDelete:
		w.Write([]byte("DELETE request received at /students"))
	case http.MethodPut:
		w.Write([]byte("PUT request received at /students"))
	case http.MethodPatch:
		w.Write([]byte("PATCH request received at /students"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}
