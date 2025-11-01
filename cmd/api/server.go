package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Hello from the root path!")
	w.Write([]byte("Hello from the root route!"))
	fmt.Println("Hello from the root route!")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userId := strings.TrimSuffix(path, "/")
		fmt.Println("User ID", userId)

		queryParams := r.URL.Query()

		sortBy := queryParams.Get("sortby")
		sortOrder := queryParams.Get("sortorder")

		fmt.Printf("Sortby: %v, SortOrder: %v\n", sortBy, sortOrder)

		w.Write([]byte("GET request received at /teachers"))
		fmt.Println("GET request received at /teachers")
	case http.MethodPost:

		//Parse form data x-www-form-urlencoded
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		fmt.Println("Form Request Body:", r.Form)
		var responseMap = make(map[string]interface{})
		for k, v := range r.Form {
			responseMap[k] = v[0]
		}
		fmt.Println("Processed Form Map:", responseMap)

		//Raw body parsing
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		fmt.Println("Raw Request Body:", string(body))
		var user user
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		fmt.Println("Unmarshalled User Struct:", user)
		fmt.Println("Unmarshalled User Name:", user.Name)

		/// Access the request details
		fmt.Println("Request Body", r.Body)
		fmt.Println("Request Form", r.Form)
		fmt.Println("Request Headers:", r.Header)
		fmt.Println("Context:", r.Context())
		fmt.Println("Content Length:", r.ContentLength)
		fmt.Println("Host:", r.Host)
		fmt.Println("Request Method:", r.Method)
		fmt.Println("Protocol:", r.Proto)
		fmt.Println("Remote Address:", r.RemoteAddr)
		fmt.Println("Request URI:", r.RequestURI)
		fmt.Println("TLS Info:", r.TLS)
		fmt.Println("Trailer Headers:", r.Trailer)
		fmt.Println("Transfer Encoding:", r.TransferEncoding)
		fmt.Println("URL:", r.URL)
		fmt.Println("User Agent:", r.UserAgent())
		fmt.Println("Port:", r.URL.Port())
		fmt.Println("URL Scheme:", r.URL.Scheme)

		//Send response
		w.Write([]byte("POST request received at /teachers"))
		fmt.Println("POST request received at /teachers")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		fmt.Println("Method not allowed at /teachers")
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET request received at /execs"))
		fmt.Println("GET request received at /execs")
	case http.MethodPost:
		w.Write([]byte("POST request received at /execs"))
		fmt.Println("POST request received at /execs")
	case http.MethodDelete:
		w.Write([]byte("DELETE request received at /execs"))
		fmt.Println("DELETE request received at /execs")
	case http.MethodPut:
		w.Write([]byte("PUT request received at /execs"))
		fmt.Println("PUT request received at /execs")
	case http.MethodPatch:
		w.Write([]byte("PATCH request received at /execs"))
		fmt.Println("PATCH request received at /execs")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		fmt.Println("Method not allowed at /execs")
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET request received at /students"))
		fmt.Println("GET request received at /students")
	case http.MethodPost:
		w.Write([]byte("POST request received at /students"))
		fmt.Println("POST request received at /students")
	case http.MethodDelete:
		w.Write([]byte("DELETE request received at /students"))
		fmt.Println("DELETE request received at /students")
	case http.MethodPut:
		w.Write([]byte("PUT request received at /students"))
		fmt.Println("PUT request received at /students")
	case http.MethodPatch:
		w.Write([]byte("PATCH request received at /students"))
		fmt.Println("PATCH request received at /students")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		fmt.Println("Method not allowed at /students")
	}
}
func main() {

	port := ":3000"

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/teachers", teachersHandler)
	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/students/", studentsHandler)

	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Starting server on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
