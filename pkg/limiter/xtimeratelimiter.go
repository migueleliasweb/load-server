package limiter

import (
	"net/http"

	"golang.org/x/time/rate"
)

type Limiter struct {
	limiter rate.Limiter
	next    http.Handler
}

func (l *Limiter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !l.limiter.Allow() {
		http.Error(
			w,
			http.StatusText(http.StatusTooManyRequests),
			http.StatusTooManyRequests,
		)

		return
	}

	l.next.ServeHTTP(w, r)
}

func AddLimiter(limit int, next http.Handler) http.Handler {
	return &Limiter{
		limiter: *rate.NewLimiter(rate.Limit(limit), limit),
		next:    next,
	}
}
