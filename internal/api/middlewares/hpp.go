package middlewares

import (
	"net/http"
	"strings"
)

type HPPOptions struct {
	Whitelist                []string
	CheckQuery               bool
	CheckBody                bool
	CheckBodyOnlyContentType string
}

func Hpp(options HPPOptions) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if options.CheckBody && r.Method == http.MethodPost && isCorrectContentType(r, options.CheckBodyOnlyContentType) {
				filterBodyParams(r, options.Whitelist)
			}

			if options.CheckQuery && r.URL.Query() != nil {
				filterQueryParams(r, options.Whitelist)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func isCorrectContentType(r *http.Request, contentType string) bool {
	return strings.Contains(r.Header.Get("Content-Type"), contentType)
}

func filterBodyParams(r *http.Request, whitelist []string) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	for k, v := range r.Form {
		if len(v) > 1 {
			r.Form.Set(k, v[0])
		}
		if !isInWhitelist(k, whitelist) {
			delete(r.Form, k)
		}
	}
}

func isInWhitelist(param string, whitelist []string) bool {
	for _, whitelistedParam := range whitelist {
		if param == whitelistedParam {
			return true
		}
	}
	return false
}

func filterQueryParams(r *http.Request, whitelist []string) {
	query := r.URL.Query()

	for k, v := range query {
		if len(v) > 1 {
			query.Set(k, v[0])
		}
		if !isInWhitelist(k, whitelist) {
			query.Del(k)
		}
	}
	r.URL.RawQuery = query.Encode()
}
