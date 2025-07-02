//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
)

var timeLayout = "2006-01-02T15:04:05Z07:00"

// setup sets up a test HTTP server along with a gitlab.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup(t *testing.T) (*http.ServeMux, *Client) {
	// mux is the HTTP request multiplexer used with the test server.
	mux := http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	// client is the Gitlab client being tested.
	client, err := NewClient("",
		WithBaseURL(server.URL),
		// Disable backoff to speed up tests that expect errors.
		WithCustomBackoff(func(_, _ time.Duration, _ int, _ *http.Response) time.Duration {
			return 0
		}),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	return mux, client
}

func testURL(t *testing.T, r *http.Request, want string) {
	if got := r.RequestURI; got != want {
		t.Errorf("Request url: %+v, want %s", got, want)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %s, want %s", got, want)
	}
}

// Tests that a given form attribute has a value in a form request. Useful
// for testing file upload API requests.
func testFormBody(t *testing.T, r *http.Request, key string, want string) {
	if got := r.FormValue(key); got != want {
		t.Errorf("Request body for key %s got: %s, want %s", key, got, want)
	}
}

// testBodyJSON tests that the JSON request body is what we expect. The want
// argument is typically either a struct, a map[string]string, or a
// map[string]any, though other types are handled as well.
//
// Calls t.Fatal if decoding the request body fails, failing the test
// immediately.
//
// When the request body is not equal to "want", the error is reported but the
// test is allowed to continue. You can use the return value to end the test on
// error: returns true if the decoded body is identical to want, false
// otherwise.
func testBodyJSON[T any](t *testing.T, r *http.Request, want T) bool {
	var got T

	if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
		t.Fatalf("Failed to decode JSON from request body: %v", err)
	}

	return assert.Equal(t, want, got)
}

// testParam checks whether the given request contains the expected parameter and whether the parameter has the expected value.
func testParam(t *testing.T, r *http.Request, key, value string) {
	require.True(t, r.URL.Query().Has(key), "Request does not contain the %q parameter", key)
	assert.Equal(t, 1, len(r.URL.Query()[key]), "Request contains multiple %q parameters when only one is expected", key)
	require.Equal(t, value, r.URL.Query().Get(key))
}

func mustWriteHTTPResponse(t *testing.T, w io.Writer, fixturePath string) {
	f, err := os.Open(fixturePath)
	if err != nil {
		t.Fatalf("error opening fixture file: %v", err)
	}
	defer f.Close()

	if _, err = io.Copy(w, f); err != nil {
		t.Fatalf("error writing response: %v", err)
	}
}

// mustWriteJSONResponse writes a JSON response to w.
// It uses t.Fatal to stop the test and report an error if encoding the response fails.
// This helper is useful when implementing handlers in unit tests.
func mustWriteJSONResponse(t *testing.T, w io.Writer, response any) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		t.Fatalf("Failed to write response: %v", err)
	}
}

// mustWriteErrorResponse writes an error response to w in a format that CheckResponse can parse.
// It uses t.Fatal to stop the test and report an error if encoding the response fails.
// This is useful when testing error conditions.
func mustWriteErrorResponse(t *testing.T, w io.Writer, err error) {
	mustWriteJSONResponse(t, w, map[string]any{
		"error": err.Error(),
	})
}

func errorOption(*retryablehttp.Request) error {
	return errors.New("RequestOptionFunc returns an error")
}

func TestNewClient(t *testing.T) {
	t.Parallel()

	t.Run("Default Configuration", func(t *testing.T) {
		t.Parallel()
		c, err := NewClient("")
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		expectedBaseURL := defaultBaseURL + apiVersionPath

		if c.BaseURL().String() != expectedBaseURL {
			t.Errorf("NewClient BaseURL is %s, want %s", c.BaseURL().String(), expectedBaseURL)
		}
		if c.UserAgent != userAgent {
			t.Errorf("NewClient UserAgent is %s, want %s", c.UserAgent, userAgent)
		}
	})

	t.Run("Custom UserAgent", func(t *testing.T) {
		t.Parallel()
		c, err := NewClient("", WithUserAgent("any-custom-user-agent"))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		expectedBaseURL := defaultBaseURL + apiVersionPath

		if c.BaseURL().String() != expectedBaseURL {
			t.Errorf("NewClient BaseURL is %s, want %s", c.BaseURL().String(), expectedBaseURL)
		}
		if c.UserAgent != "any-custom-user-agent" {
			t.Errorf("NewClient UserAgent is %s, want any-custom-user-agent", c.UserAgent)
		}
	})

	t.Run("Custom Base URL", func(t *testing.T) {
		t.Parallel()
		customURL := "https://custom.gitlab.com/api/v4"
		c, err := NewClient("", WithBaseURL(customURL))
		require.NoError(t, err)
		require.NotNil(t, c, "Client is nil")

		// The client will append a trailing slash to the base URL
		expectedURL := customURL + "/"
		require.Equal(t, expectedURL, c.BaseURL().String(), "BaseURL mismatch")
	})

	t.Run("Invalid Base URL", func(t *testing.T) {
		t.Parallel()
		_, err := NewClient("", WithBaseURL(":invalid:"))
		require.Error(t, err)
	})
}

func TestCheckResponse(t *testing.T) {
	t.Parallel()
	c, err := NewClient("")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req, err := c.NewRequest(http.MethodGet, "test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp := &http.Response{
		Request:    req.Request,
		StatusCode: http.StatusBadRequest,
		Body: io.NopCloser(strings.NewReader(`
		{
			"message": {
				"prop1": [
					"message 1",
					"message 2"
				],
				"prop2":[
					"message 3"
				],
				"embed1": {
					"prop3": [
						"msg 1",
						"msg2"
					]
				},
				"embed2": {
					"prop4": [
						"some msg"
					]
				}
			},
			"error": "message 1"
		}`)),
	}

	errResp := CheckResponse(resp)
	if errResp == nil {
		t.Fatal("Expected error response.")
	}

	want := "GET https://gitlab.com/api/v4/test: 400 {error: message 1}, {message: {embed1: {prop3: [msg 1, msg2]}}, {embed2: {prop4: [some msg]}}, {prop1: [message 1, message 2]}, {prop2: [message 3]}}"

	if errResp.Error() != want {
		t.Errorf("Expected error: %s, got %s", want, errResp.Error())
	}
}

func TestCheckResponseOnUnknownErrorFormat(t *testing.T) {
	t.Parallel()
	c, err := NewClient("")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req, err := c.NewRequest(http.MethodGet, "test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp := &http.Response{
		Request:    req.Request,
		StatusCode: http.StatusBadRequest,
		Body:       io.NopCloser(strings.NewReader("some error message but not JSON")),
	}

	errResp := CheckResponse(resp)
	if errResp == nil {
		t.Fatal("Expected error response.")
	}

	want := "GET https://gitlab.com/api/v4/test: 400 failed to parse unknown error format: some error message but not JSON"

	if errResp.Error() != want {
		t.Errorf("Expected error: %s, got %s", want, errResp.Error())
	}
}

func TestCheckResponseOnHeadRequestError(t *testing.T) {
	t.Parallel()
	c, err := NewClient("")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req, err := c.NewRequest(http.MethodHead, "test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp := &http.Response{
		Request:    req.Request,
		StatusCode: http.StatusNotFound,
		Body:       nil,
	}

	errResp := CheckResponse(resp)
	if errResp == nil {
		t.Fatal("Expected error response.")
	}

	want := "404 Not Found"

	if errResp.Error() != want {
		t.Errorf("Expected error: %s, got %s", want, errResp.Error())
	}
}

func TestRequestWithContext(t *testing.T) {
	t.Parallel()
	c, err := NewClient("")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	req, err := c.NewRequest(http.MethodGet, "test", nil, []RequestOptionFunc{WithContext(ctx)})
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	defer cancel()

	if req.Context() != ctx {
		t.Fatal("Context was not set correctly")
	}
}

func loadFixture(t *testing.T, filePath string) []byte {
	t.Helper()
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}

	return content
}

func TestPathEscape(t *testing.T) {
	t.Parallel()
	want := "diaspora%2Fdiaspora"
	got := PathEscape("diaspora/diaspora")
	if want != got {
		t.Errorf("Expected: %s, got %s", want, got)
	}
}

func TestPaginationPopulatePageValuesEmpty(t *testing.T) {
	t.Parallel()
	wantPageHeaders := map[string]int{
		xTotal:      0,
		xTotalPages: 0,
		xPerPage:    0,
		xPage:       0,
		xNextPage:   0,
		xPrevPage:   0,
	}
	wantLinkHeaders := map[string]string{
		linkPrev:  "",
		linkNext:  "",
		linkFirst: "",
		linkLast:  "",
	}

	r := newResponse(&http.Response{
		Header: http.Header{},
	})

	gotPageHeaders := map[string]int{
		xTotal:      r.TotalItems,
		xTotalPages: r.TotalPages,
		xPerPage:    r.ItemsPerPage,
		xPage:       r.CurrentPage,
		xNextPage:   r.NextPage,
		xPrevPage:   r.PreviousPage,
	}
	for k, v := range wantPageHeaders {
		if v != gotPageHeaders[k] {
			t.Errorf("For %s, expected %d, got %d", k, v, gotPageHeaders[k])
		}
	}

	gotLinkHeaders := map[string]string{
		linkPrev:  r.PreviousLink,
		linkNext:  r.NextLink,
		linkFirst: r.FirstLink,
		linkLast:  r.LastLink,
	}
	for k, v := range wantLinkHeaders {
		if v != gotLinkHeaders[k] {
			t.Errorf("For %s, expected %s, got %s", k, v, gotLinkHeaders[k])
		}
	}
}

func TestPaginationPopulatePageValuesOffset(t *testing.T) {
	t.Parallel()
	wantPageHeaders := map[string]int{
		xTotal:      100,
		xTotalPages: 5,
		xPerPage:    20,
		xPage:       2,
		xNextPage:   3,
		xPrevPage:   1,
	}
	wantLinkHeaders := map[string]string{
		linkPrev:  "https://gitlab.example.com/api/v4/projects/8/issues/8/notes?page=1&per_page=3",
		linkNext:  "https://gitlab.example.com/api/v4/projects/8/issues/8/notes?page=3&per_page=3",
		linkFirst: "https://gitlab.example.com/api/v4/projects/8/issues/8/notes?page=1&per_page=3",
		linkLast:  "https://gitlab.example.com/api/v4/projects/8/issues/8/notes?page=3&per_page=3",
	}

	h := http.Header{}
	for k, v := range wantPageHeaders {
		h.Add(k, fmt.Sprint(v))
	}
	var linkHeaderComponents []string
	for k, v := range wantLinkHeaders {
		if v != "" {
			linkHeaderComponents = append(linkHeaderComponents, fmt.Sprintf("<%s>; rel=\"%s\"", v, k))
		}
	}
	h.Add("Link", strings.Join(linkHeaderComponents, ", "))

	r := newResponse(&http.Response{
		Header: h,
	})

	gotPageHeaders := map[string]int{
		xTotal:      r.TotalItems,
		xTotalPages: r.TotalPages,
		xPerPage:    r.ItemsPerPage,
		xPage:       r.CurrentPage,
		xNextPage:   r.NextPage,
		xPrevPage:   r.PreviousPage,
	}
	for k, v := range wantPageHeaders {
		if v != gotPageHeaders[k] {
			t.Errorf("For %s, expected %d, got %d", k, v, gotPageHeaders[k])
		}
	}

	gotLinkHeaders := map[string]string{
		linkPrev:  r.PreviousLink,
		linkNext:  r.NextLink,
		linkFirst: r.FirstLink,
		linkLast:  r.LastLink,
	}
	for k, v := range wantLinkHeaders {
		if v != gotLinkHeaders[k] {
			t.Errorf("For %s, expected %s, got %s", k, v, gotLinkHeaders[k])
		}
	}
}

func TestPaginationPopulatePageValuesKeyset(t *testing.T) {
	t.Parallel()
	wantPageHeaders := map[string]int{
		xTotal:      0,
		xTotalPages: 0,
		xPerPage:    0,
		xPage:       0,
		xNextPage:   0,
		xPrevPage:   0,
	}
	wantLinkHeaders := map[string]string{
		linkPrev:  "",
		linkFirst: "",
		linkLast:  "",
	}

	h := http.Header{}
	for k, v := range wantPageHeaders {
		h.Add(k, fmt.Sprint(v))
	}
	var linkHeaderComponents []string
	for k, v := range wantLinkHeaders {
		if v != "" {
			linkHeaderComponents = append(linkHeaderComponents, fmt.Sprintf("<%s>; rel=\"%s\"", v, k))
		}
	}
	h.Add("Link", strings.Join(linkHeaderComponents, ", "))

	r := newResponse(&http.Response{
		Header: h,
	})

	gotPageHeaders := map[string]int{
		xTotal:      r.TotalItems,
		xTotalPages: r.TotalPages,
		xPerPage:    r.ItemsPerPage,
		xPage:       r.CurrentPage,
		xNextPage:   r.NextPage,
		xPrevPage:   r.PreviousPage,
	}
	for k, v := range wantPageHeaders {
		if v != gotPageHeaders[k] {
			t.Errorf("For %s, expected %d, got %d", k, v, gotPageHeaders[k])
		}
	}
}

func TestNewRetryableHTTPClientWithRetryCheck(t *testing.T) {
	t.Parallel()

	_, client := setup(t)

	httpClient := &http.Client{}
	logger := struct{}{}
	retryWaitMin := 10 * time.Second
	retryWaitMax := 20 * time.Second
	retryMax := 30
	requestLogHook := retryablehttp.RequestLogHook(func(logger retryablehttp.Logger, request *http.Request, i int) {
	})
	checkRetry := retryablehttp.CheckRetry(func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		return false, nil
	})
	backoff := retryablehttp.Backoff(func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		return time.Second
	})
	errorHandler := retryablehttp.ErrorHandler(func(resp *http.Response, err error, numTries int) (*http.Response, error) {
		return nil, nil
	})
	prepareRetry := retryablehttp.PrepareRetry(func(req *http.Request) error {
		return nil
	})

	newCheckRetry := retryablehttp.CheckRetry(func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		return false, nil
	})

	client.client.HTTPClient = httpClient
	client.client.Logger = logger
	client.client.RetryWaitMin = retryWaitMin
	client.client.RetryWaitMax = retryWaitMax
	client.client.RetryMax = retryMax
	client.client.RequestLogHook = requestLogHook
	client.client.CheckRetry = checkRetry
	client.client.Backoff = backoff
	client.client.ErrorHandler = errorHandler
	client.client.PrepareRetry = prepareRetry

	actual := client.newRetryableHTTPClientWithRetryCheck(newCheckRetry)

	assert.Equal(t, httpClient, actual.HTTPClient)
	assert.Equal(t, logger, actual.Logger)
	assert.Equal(t, retryWaitMin, actual.RetryWaitMin)
	assert.Equal(t, retryWaitMax, actual.RetryWaitMax)
	assert.Equal(t, retryMax, actual.RetryMax)
	assert.Equal(t, reflect.ValueOf(requestLogHook).Pointer(), reflect.ValueOf(actual.RequestLogHook).Pointer())
	assert.Equal(t, reflect.ValueOf(newCheckRetry).Pointer(), reflect.ValueOf(actual.CheckRetry).Pointer())
	assert.Equal(t, reflect.ValueOf(backoff).Pointer(), reflect.ValueOf(actual.Backoff).Pointer())
	assert.Equal(t, reflect.ValueOf(errorHandler).Pointer(), reflect.ValueOf(actual.ErrorHandler).Pointer())
	assert.Equal(t, reflect.ValueOf(prepareRetry).Pointer(), reflect.ValueOf(actual.PrepareRetry).Pointer())
}

func TestExponentialBackoffLogic(t *testing.T) {
	t.Parallel()
	// Can't use the default `setup` because it disabled the backoff
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)
	client, err := NewClient("",
		WithBaseURL(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create a method that returns 429
	mux.HandleFunc("/api/v4/projects/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusTooManyRequests)
	})

	// Measure the time at the start of the test
	start := time.Now()

	// Send a request (which will get a bunch of 429s)
	// None of the responses matter, so ignore them all
	_, resp, _ := client.Projects.GetProject(1, nil)
	end := time.Now()

	// The test should run for _at least_ 3,200 milliseconds
	duration := float64(end.Sub(start))
	if duration < float64(3200*time.Millisecond) {
		t.Fatal("Wait was shorter than expected. Expected a minimum of 5 retries taking 3200 milliseconds, got:", duration)
	}
	if resp.StatusCode != 429 {
		t.Fatal("Expected to get a 429 code given the server is hard-coded to return this. Received instead:", resp.StatusCode)
	}
}

func TestErrorResponsePreservesURLEncoding(t *testing.T) {
	t.Parallel()

	projectID := "group/subgroup"
	fileName := "path/file.txt"

	expectedEscapedPath := "/api/v4/projects/group%2Fsubgroup/repository/files/path%2Ffile%2Etxt"

	escapedProjectID := PathEscape(projectID)
	escapedFileName := PathEscape(fileName)
	escapedPath := fmt.Sprintf("/api/v4/projects/%s/repository/files/%s",
		escapedProjectID, escapedFileName)

	require.Equal(t, expectedEscapedPath, escapedPath)

	fullURL := "https://gitlab.com" + expectedEscapedPath
	req, _ := http.NewRequest("GET", fullURL, nil)
	resp := &http.Response{
		Request:    req,
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader(`{"message":"Not Found"}`)),
	}

	errorResponse := &ErrorResponse{Response: resp, Message: "Not Found"}

	require.ErrorContains(t, errorResponse, expectedEscapedPath)

	unescapedPath := fmt.Sprintf("/api/v4/projects/%s/repository/files/%s", projectID, fileName)
	assert.NotContains(t, errorResponse.Error(), unescapedPath)
}

func TestNewClient_auth(t *testing.T) {
	t.Parallel()

	const token = "glpat-0123456789abcdefg"

	handler := func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("PRIVATE-TOKEN"), token; got != want {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Authorization = %q, want %q", got, want)
			return
		}

		fmt.Fprint(w, "[]")
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	t.Cleanup(server.Close)

	client, err := NewClient(token,
		WithBaseURL(server.URL),
		WithHTTPClient(server.Client()),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	projects, resp, err := client.Projects.ListProjects(&ListProjectsOptions{})
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, []*Project{}, projects)
}

func TestNewJobClient_auth(t *testing.T) {
	t.Parallel()

	const token = "glcbt-0123456789abcdefg"

	handler := func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("JOB-TOKEN"), token; got != want {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Authorization = %q, want %q", got, want)
			return
		}

		fmt.Fprint(w, "[]")
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	t.Cleanup(server.Close)

	client, err := NewJobClient(token,
		WithBaseURL(server.URL),
		WithHTTPClient(server.Client()),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	projects, resp, err := client.Projects.ListProjects(&ListProjectsOptions{})
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, []*Project{}, projects)
}

func TestNewBasicAuthClient_auth(t *testing.T) {
	t.Parallel()

	const (
		username = "test-username"
		password = "test-p4ssw0rd"
		token    = "test-token"
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "r.ParseForm: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if got, want := r.Form.Get("grant_type"), "password"; got != want {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "grant_type = %q, want %q", got, want)
			return
		}

		if gotUsername, gotPassword := r.Form.Get("username"), r.Form.Get("password"); gotUsername != username || gotPassword != password {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "username is %q, want %q", gotUsername, username)
			fmt.Fprintf(w, "password is %q, want %q", gotPassword, password)
			return
		}

		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		fmt.Fprint(w, url.Values{
			"access_token": {token},
			"token_type":   {"bearer"},
			"expires_in":   {"1800"},
		}.Encode())
	})
	mux.HandleFunc("/api/v4/projects", func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("Authorization"), "Bearer "+token; got != want {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Authorization = %q, want %q", got, want)
			return
		}

		fmt.Fprint(w, "[]")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unexpected %s request to %s", r.Method, r.URL.String())
	})

	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client, err := NewBasicAuthClient(username, password,
		WithBaseURL(server.URL),
		WithHTTPClient(server.Client()),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	projects, resp, err := client.Projects.ListProjects(&ListProjectsOptions{})
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, []*Project{}, projects)
}

func TestNewAuthSourceClient(t *testing.T) {
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

	client, err := NewAuthSourceClient(OAuthTokenSource{ts},
		WithBaseURL(server.URL),
		WithHTTPClient(server.Client()),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	projects, resp, err := client.Projects.ListProjects(&ListProjectsOptions{})
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, []*Project{}, projects)
}

// Interface implementation checks.
var (
	_ AuthSource = OAuthTokenSource{}
	_ AuthSource = JobTokenAuthSource{}
	_ AuthSource = AccessTokenAuthSource{}
	_ AuthSource = (*passwordCredentialsAuthSource)(nil)
)
