package db

import "time"

// Connect creates attempts to connect with a db because when db is running in docker, it is not always possible to connect from the first time.
func Connect(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--

			continue
		}
		return nil
	}
	return
}
