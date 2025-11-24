package handlers

import "net/http"

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET request received at /execs"))
	case http.MethodPost:
		w.Write([]byte("POST request received at /execs"))
	case http.MethodDelete:
		w.Write([]byte("DELETE request received at /execs"))
	case http.MethodPut:
		w.Write([]byte("PUT request received at /execs"))
	case http.MethodPatch:
		w.Write([]byte("PATCH request received at /execs"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}
