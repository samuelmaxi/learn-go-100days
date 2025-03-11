package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.Mutex
	limit    int
	window   time.Duration
}

type visitor struct {
	lastSeen time.Time
	requests int
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		visitors: make(map[string]*visitor),
		limit:    limit,
		window:   window,
	}
}

func (r *RateLimiter) cleanup() {
	for ip, v := range r.visitors {
		if time.Since(v.lastSeen) > r.window {
			delete(r.visitors, ip)
		}
	}
}

func (r *RateLimiter) Allow(ip string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.cleanup()

	v, exists := r.visitors[ip]
	if !exists {
		r.visitors[ip] = &visitor{lastSeen: time.Now(), requests: 1}
		return true
	}

	if time.Since(v.lastSeen) > r.window {
		v.requests = 1
		v.lastSeen = time.Now()
		return true
	}

	if v.requests >= r.limit {
		return false
	}

	v.requests++
	return true
}

func rateLimitMiddleware(rl *RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			if !rl.Allow(ip) {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

}

func main() {
	limiter := NewRateLimiter(5, time.Minute)

	mux := http.DefaultServeMux
	mux.Handle("/", rateLimitMiddleware(limiter)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Word!"))
	})))

	fmt.Println("Rodando na porta :8080")

	log.Fatal(http.ListenAndServe(":8080", mux))

}
