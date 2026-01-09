package circuit_break

import (
	"errors"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name         string
		maxFailures  int
		resetTimeout time.Duration
		setup        func(*CircuitBreaker)
		fn           func() error
		wantErr      error
		wantState    State
		wantFailures int
	}{
		{
			name:         "closed_state_success",
			maxFailures:  3,
			resetTimeout: 5 * time.Second,
			fn:           func() error { return nil },
			wantState:    StateClosed,
			wantFailures: 0,
		},
		{
			name:         "closed_state_failure",
			maxFailures:  3,
			resetTimeout: 5 * time.Second,
			fn:           func() error { return errors.New("service error") },
			wantErr:      errors.New("service error"),
			wantState:    StateClosed,
			wantFailures: 1,
		},
		{
			name:         "transition_to_open",
			maxFailures:  2,
			resetTimeout: 5 * time.Second,
			setup: func(cb *CircuitBreaker) {
				cb.Execute(func() error { return errors.New("fail") })
			},
			fn:           func() error { return errors.New("fail") },
			wantErr:      errors.New("fail"),
			wantState:    StateOpen,
			wantFailures: 2,
		},
		{
			name:         "open_state_blocks_requests",
			maxFailures:  1,
			resetTimeout: 5 * time.Second,
			setup: func(cb *CircuitBreaker) {
				cb.Execute(func() error { return errors.New("fail") })
			},
			fn:       func() error { return nil },
			wantErr:  ErrCircuitOpen,
			wantState: StateOpen,
		},
		{
			name:         "half_open_success_closes_circuit",
			maxFailures:  1,
			resetTimeout: 100 * time.Millisecond,
			setup: func(cb *CircuitBreaker) {
				cb.Execute(func() error { return errors.New("fail") })
				time.Sleep(150 * time.Millisecond)
			},
			fn:           func() error { return nil },
			wantState:    StateClosed,
			wantFailures: 0,
		},
		{
			name:         "half_open_failure_reopens_circuit",
			maxFailures:  1,
			resetTimeout: 100 * time.Millisecond,
			setup: func(cb *CircuitBreaker) {
				cb.Execute(func() error { return errors.New("fail") })
				time.Sleep(150 * time.Millisecond)
			},
			fn:       func() error { return errors.New("still failing") },
			wantErr:  errors.New("still failing"),
			wantState: StateOpen,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := NewCircuitBreaker(tt.maxFailures, tt.resetTimeout)
			
			if tt.setup != nil {
				tt.setup(cb)
			}

			err := cb.Execute(tt.fn)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else if err != nil {
				t.Errorf("Execute() unexpected error = %v", err)
			}

			if cb.state != tt.wantState {
				t.Errorf("Execute() state = %v, want %v", cb.state, tt.wantState)
			}

			if tt.wantFailures > 0 && cb.failures != tt.wantFailures {
				t.Errorf("Execute() failures = %v, want %v", cb.failures, tt.wantFailures)
			}
		})
	}
}

func TestExecute_TooManyRequests(t *testing.T) {
	cb := NewCircuitBreaker(1, 100*time.Millisecond)
	cb.Execute(func() error { return errors.New("fail") })
	time.Sleep(150 * time.Millisecond)

	// Manually set state to half-open and halfOpenRequests to 1
	cb.mu.Lock()
	cb.state = StateHalfOpen
	cb.halfOpenRequests = 1
	cb.mu.Unlock()

	err := cb.Execute(func() error { return nil })
	if !errors.Is(err, ErrTooManyRequests) {
		t.Errorf("Expected ErrTooManyRequests, got %v", err)
	}
}