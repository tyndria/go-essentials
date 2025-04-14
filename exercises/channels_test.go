package exercises

func ExampleContentTypesOverChannel() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://pkg.go.dev",
	}
	ContentTypesOverChannel(urls)
	// Output:
	// Url: https://golang.org, type: text/html; charset=utf-8
	// Url: https://api.github.com, type: application/json; charset=utf-8
	// Url: https://pkg.go.dev, type: text/html; charset=utf-8
}
