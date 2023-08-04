package retry

import "errors"

func Do[T any](f func() (T, error), opts ...Option) (T, error) {
	var err error
	var result T
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}

	for i := 0; i <= int(config.attempts); i++ {
		result, err = f()
		if err == nil {
			return result, nil
		}

		if config.retryIf != nil && !config.retryIf(err) {
			return result, err
		}

		if config.onRetry != nil {
			config.onRetry(uint(i), err)
		}
	}

	err = errors.New("retry failed")
	return result, err
}
