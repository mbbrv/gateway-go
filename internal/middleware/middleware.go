package middleware

// RateLimiterRepository is an interface to store rate limiting data
type RateLimiterRepository interface {
	// Get returns the current rate limit for the given identifier
	Get(identifier string) (int, error)

	// Incr increments the rate limit for the given identifier
	Incr(identifier string) error
}

type Middleware struct {
	rateLimiterRepository RateLimiterRepository
}

// NewMiddleware creates a new Middleware
func NewMiddleware(repository RateLimiterRepository) *Middleware {
	return &Middleware{
		rateLimiterRepository: repository,
	}
}
