import pybreaker

circuit_breaker = pybreaker.CircuitBreaker(
    fail_max=3,
    reset_timeout=10 * 3.5
)