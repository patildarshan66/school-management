package middlewares

import (
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	vistors   map[string]int
	limit     int
	mu        sync.Mutex
	resetTime time.Duration
}

func NewRateLimiter(limit int, resetTime time.Duration) *rateLimiter {
	rl := &rateLimiter{
		limit:     limit,
		resetTime: resetTime,
		vistors:   make(map[string]int),
	}
	go rl.resetVisitorCount()
	return rl
}

func (rl *rateLimiter) resetVisitorCount() {
	for {
		time.Sleep(rl.resetTime)
		rl.mu.Lock()
		rl.vistors = make(map[string]int)
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl.mu.Lock()
		defer rl.mu.Unlock()
		visitorIp := r.RemoteAddr
		rl.vistors[visitorIp]++

		if rl.vistors[visitorIp] > rl.limit {
			http.Error(w, "Too Many Request", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
