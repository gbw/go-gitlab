package gitlab_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	gitlab "gitlab.com/gitlab-org/api/client-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

// ExampleWithTokenSource demonstrates how to use the Proof Key for Code
// Exchange (PKCE) flow to acquire an OAuth token, wrap the OAuth token in a
// token source and use that to authenticate API requests.
func ExampleWithTokenSource() {
	ctx := context.Background()

	config := oauth2.Config{
		ClientID: "your-client-id",
		Endpoint: endpoints.GitLab, // for gitlab.com
		Scopes:   []string{"api"},
	}

	// Generate a random code verifier
	verifier := oauth2.GenerateVerifier()

	// Create the authorization URL with PKCE parameters
	authCodeURL := config.AuthCodeURL("state", oauth2.S256ChallengeOption(verifier))

	// At this point, the user would visit authCodeURL in their browser
	// and authorize the application, receiving a code in return.
	fmt.Println("Visit the URL for the auth dialog:", authCodeURL)

	// After authorization, the user would receive a code
	var authCode string
	fmt.Print("Enter the authorization code: ")
	fmt.Scan(&authCode)

	// Exchange the authorization code for a token using the code verifier
	token, err := config.Exchange(ctx, authCode, oauth2.VerifierOption(verifier))
	if err != nil {
		panic(err)
	}

	// Wrap the token in a token source to refresh it when needed
	ts := config.TokenSource(ctx, token)

	// Create a client with the token
	client, err := gitlab.NewOAuthClient("", gitlab.WithTokenSource(ts))
	if err != nil {
		panic(err)
	}

	// Use the client to make API requests
	user, _, err := client.Users.CurrentUser(gitlab.WithContext(ctx))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s!\n", user.Name)
}

func TestWithTokenSource(t *testing.T) {
	t.Parallel()

	token := &oauth2.Token{
		AccessToken: "0123456789abcdefg",
	}
	ts := oauth2.StaticTokenSource(token)

	handler := func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("Authorization"), "Bearer 0123456789abcdefg"; got != want {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Authorization = %q, want %q", got, want)
			return
		}

		fmt.Fprint(w, "[]")
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	t.Cleanup(server.Close)

	client, err := gitlab.NewOAuthClient("unused",
		gitlab.WithBaseURL(server.URL),
		gitlab.WithHTTPClient(server.Client()),
		gitlab.WithTokenSource(ts),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	projects, resp, err := client.Projects.ListProjects(&gitlab.ListProjectsOptions{})
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, []*gitlab.Project{}, projects)
}
