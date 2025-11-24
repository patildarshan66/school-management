package utils

import "net/http"

// Middleware is a function that takes and returns an http.Handler
type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
