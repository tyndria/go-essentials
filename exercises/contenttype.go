package exercises

import (
    "net/http"
	"fmt"
)

func ContentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	value := resp.Header.Get("Content-Type")
	if value == "" {
		return "", fmt.Errorf("Content-Type is empty for the url %v", url)
	}

	return value, nil
}