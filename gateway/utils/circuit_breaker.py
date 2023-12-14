from circuitbreaker import CircuitBreaker
import time

class CircuitBreaker:
    def __init__(self, max_failures, reset_timeout, max_re_routes):
        self.max_failures = max_failures
        self.reset_timeout = reset_timeout
        self.max_re_routes = max_re_routes
        self.failures = 0
        self.re_routes = 0
        self.last_failure_time = None
        self.state = "CLOSED"

    def allow_request(self):
        if self.state == "CLOSED":
            return True
        elif self.state == "OPEN":
            current_time = time.time()
            if current_time - self.last_failure_time >= self.reset_timeout:
                self.state = "HALF-OPEN"
                return True
            else:
                return False
        elif self.state == "HALF-OPEN":
            return self.failures < self.max_failures and self.re_routes < self.max_re_routes

    def record_success(self):
        if self.state == "HALF-OPEN":
            self.reset_circuit()

    def record_failure(self):
        self.failures += 1
        if self.failures >= self.max_failures:
            self.open_circuit()

    def record_re_route(self):
        if self.state == "HALF-OPEN":
            self.re_routes += 1
            if self.re_routes >= self.max_re_routes:
                self.open_circuit()

    def open_circuit(self):
        self.state = "OPEN"
        self.last_failure_time = time.time()

    def reset_circuit(self):
        self.state = "CLOSED"
        self.failures = 0
        self.re_routes = 0
        self.last_failure_time = None

circuit_breakers = {}

def create_circuit_breaker():
    max_failures = 3
    reset_timeout = 10 * 3.5
    max_re_routes = 2
    return CircuitBreaker(max_failures=max_failures, reset_timeout=reset_timeout, max_re_routes=max_re_routes)

def get_circuit_breaker(service_key):
    if service_key not in circuit_breakers:
        circuit_breakers[service_key] = create_circuit_breaker()
    return circuit_breakers[service_key]
