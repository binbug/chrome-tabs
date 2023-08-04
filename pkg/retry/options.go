package retry

// Function signature of retry if function
type RetryIfFunc func(error) bool

type OnRetryFunc func(n uint, err error)

type Config struct {
	attempts uint
	retryIf  RetryIfFunc
	onRetry  OnRetryFunc
}

func emptyOption(c *Config) {}

// Option represents an option for retry.
type Option func(*Config)

// Attempts sets the number of attempts.
func Attempts(attempts uint) Option {
	return func(c *Config) {
		c.attempts = attempts
	}
}

// OnRetry sets the callback function when retry.
func OnRetry(onRetry OnRetryFunc) Option {
	if onRetry == nil {
		return emptyOption
	}
	return func(c *Config) {
		c.onRetry = onRetry
	}
}

// RetryIf sets the function to determine whether to retry.
func RetryIf(retryIf RetryIfFunc) Option {
	if retryIf == nil {
		return emptyOption
	}
	return func(c *Config) {
		c.retryIf = retryIf
	}
}
