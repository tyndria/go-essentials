package exercises

import (
	"fmt"
	"time"
)

func ContentTypesOverChannel(urls []string) {
	ch := make(chan string)

	for _, url := range urls {
		go func() {
			contentType, error := ContentType(url)
			if error != nil {
				ch <- fmt.Sprintf("Url: %v, error: %v", url, error)
			} else {
				ch <- fmt.Sprintf("Url: %v, type: %v", url, contentType)
			}
		}()
		// Add Sleep to ensure order with some probability just for tests
		time.Sleep(1000 * time.Millisecond)
	}

	for range urls {
		out := <-ch
		fmt.Println(out)
	}
}
