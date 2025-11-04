package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "schoolmanagement/internal/api/middlewares"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the root route!"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET request received at /teachers"))
	case http.MethodPost:
		//Send response
		w.Write([]byte("POST request received at /teachers"))
	case http.MethodDelete:
		w.Write([]byte("DELETE request received at /teachers"))
	case http.MethodPut:
		w.Write([]byte("PUT request received at /teachers"))
	case http.MethodPatch:
		w.Write([]byte("PATCH request received at /teachers"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
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

func studentsHandler(w http.ResponseWriter, r *http.Request) {
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
func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// hppOptions := mw.HPPOptions{
	// 	Whitelist:                []string{"name", "city"},
	// 	CheckQuery:               true,
	// 	CheckBody:                true,
	// 	CheckBodyOnlyContentType: "application/x-www-form-urlencoded",
	// }

	// rl := mw.NewRateLimiter(5, time.Minute)

	// secureMux := applyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTime, rl.MiddleWare, mw.Cors)
	// mw.Cors(rl.MiddleWare(mw.ResponseTime(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))

	secureMux := mw.SecurityHeaders(mux)

	server := &http.Server{
		Addr:      port,
		TLSConfig: tlsConfig,
		Handler:   secureMux,
	}

	fmt.Println("Starting server on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}

// Middleware is a function that takes and returns an http.Handler
type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
