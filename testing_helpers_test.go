package gitlab_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// setupBasicAuthMock creates a mock server for BasicAuth example
func setupBasicAuthMock() (*gitlab.Client, *httptest.Server) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.URL.Path {
		case "/oauth/token":
			fmt.Fprint(w, `{"access_token":"mock-token","token_type":"Bearer"}`)
		case "/api/v4/projects":
			fmt.Fprint(w, `[{"id":1,"name":"project1"},{"id":2,"name":"project2"}]`)
		default:
			http.NotFound(w, r)
		}
	})

	server := httptest.NewServer(handler)
	client, _ := gitlab.NewBasicAuthClient("your-username", "your-password", gitlab.WithBaseURL(server.URL))
	return client, server
}

// setupPaginationMock creates a mock server for pagination examples
func setupPaginationMock(useKeyset bool) (*gitlab.Client, *httptest.Server) {
	callCount := 0
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		if callCount == 1 {
			if useKeyset {
				w.Header().Set("Link", `<next>; rel="next"`)
			} else {
				w.Header().Set("X-Next-Page", "2")
			}
			fmt.Fprint(w, `[{"id":1,"name":"project1"}]`)
		} else {
			fmt.Fprint(w, `[{"id":2,"name":"project2"}]`)
		}
	})

	server := httptest.NewServer(handler)
	client, _ := gitlab.NewClient("token", gitlab.WithBaseURL(server.URL))
	return client, server
}

// setupFileUploadMock creates a mock server for file upload example
func setupFileUploadMock() (*gitlab.Client, *httptest.Server) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"file_path":"README.md","branch":"main"}`)
	})

	server := httptest.NewServer(handler)
	client, _ := gitlab.NewClient("token", gitlab.WithBaseURL(server.URL))
	return client, server
}
