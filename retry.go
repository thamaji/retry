package retry

import (
	"time"
)

func Do[T any](s Strategy, f func() (T, error)) (T, error) {
	v, err := f()
	if err == nil {
		return v, nil
	}

	if s == nil {
		s = DefaultStrategy
	}

	for i := 0; i < s.Max(); i++ {
		time.Sleep(s.Delay(i))
		v, err = f()
		if err == nil {
			return v, nil
		}
	}

	return *new(T), err
}
