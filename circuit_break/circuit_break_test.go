package circuit_break

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

// State represents the circuit breaker state
type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

// String returns the string representation of the state
func (s State) String() string {
	switch s {
	case StateClosed:
		return "CLOSED"
	case StateOpen:
		return "OPEN"
	case StateHalfOpen:
		return "HALF-OPEN"
	default:
		return "UNKNOWN"
	}
}

// CircuitBreaker implements the circuit breaker pattern
type CircuitBreaker struct {
	maxFailures  int
	resetTimeout time.Duration

	mu               sync.Mutex
	state            State
	failures         int
	lastFailTime     time.Time
	halfOpenRequests int
}

var (
	ErrCircuitOpen     = errors.New("circuit breaker is open")
	ErrTooManyRequests = errors.New("too many requests in half-open state")
)

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures:  maxFailures,
		resetTimeout: resetTimeout,
		state:        StateClosed,
	}
}

// Execute runs the given function if the circuit is closed or half-open
func (cb *CircuitBreaker) Execute(fn func() error) error {
	if err := cb.beforeRequest(); err != nil {
		return err
	}

	err := fn()
	cb.afterRequest(err)
	return err
}

func (cb *CircuitBreaker) beforeRequest() error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case StateOpen:
		if time.Since(cb.lastFailTime) > cb.resetTimeout {
			cb.state = StateHalfOpen
			cb.halfOpenRequests = 0
			return nil
		}
		return ErrCircuitOpen
	case StateHalfOpen:
		// Only allow one request at a time in half-open state
		if cb.halfOpenRequests > 0 {
			return ErrTooManyRequests
		}
		cb.halfOpenRequests++
		return nil
	case StateClosed:
		return nil
	default:
		return nil
	}
}

func (cb *CircuitBreaker) afterRequest(err error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.onFailure()
	} else {
		cb.onSuccess()
	}
}

func (cb *CircuitBreaker) onSuccess() {
	if cb.state == StateHalfOpen {
		cb.state = StateClosed
		cb.failures = 0
		cb.halfOpenRequests = 0
	}
}

func (cb *CircuitBreaker) onFailure() {
	cb.failures++
	cb.lastFailTime = time.Now()

	if cb.failures >= cb.maxFailures {
		cb.state = StateOpen
		cb.halfOpenRequests = 0
	}
}

// Example usage
func Test(t *testing.T) {
	cb := NewCircuitBreaker(3, 5*time.Second)

	// Simulate a failing service
	failingService := func() error {
		// Simulate random failures
		return errors.New("service unavailable")
	}

	for i := 1; i <= 10; i++ {
		err := cb.Execute(failingService)

		if err != nil {
			if errors.Is(err, ErrCircuitOpen) {
				fmt.Printf("Request %d: Circuit is OPEN, request blocked\n", i)
			} else if errors.Is(err, ErrTooManyRequests) {
				fmt.Printf("Request %d: Too many requests in HALF-OPEN state\n", i)
			} else {
				fmt.Printf("Request %d: Failed (%v), State: %s\n", i, err, cb.state)
			}
		} else {
			fmt.Printf("Request %d: Success, State: %s\n", i, cb.state)
		}

		time.Sleep(1 * time.Second)
	}
}
