package checker

import (
	"net/http"
)

// Health Health
func Health(url string) error {
	_, err := http.Get(url)
	if err != nil {
		return err
	}
	return nil
}
