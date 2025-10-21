package gitlab

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test structs for testing
type TestUser struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TestProject struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func TestDoRequestSuccess(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a mock server that returns a test user from testdata
	path := "/api/v4/users/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_user.json")
	})

	// WHEN doRequest is called with valid parameters
	user, resp, err := doRequest[*TestUser](
		client,
		http.MethodGet,
		"users/1",
		nil,
	)

	// THEN the request should succeed and return the expected user
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, "John Smith", user.Name)
}

func TestDoRequestPOSTWithBody(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a mock server that creates and returns a project
	path := "/api/v4/projects"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		var reqBody TestProject
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)
		require.Equal(t, "New Project", reqBody.Name)

		w.WriteHeader(201)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"id": 1, "name": "New Project", "description": "Test project"}`))
	})

	// GIVEN a project creation request body
	requestBody := &TestProject{
		Name:        "New Project",
		Description: "Test project",
	}

	// WHEN doRequest is called with POST method and body
	project, resp, err := doRequest[*TestProject](
		client,
		http.MethodPost,
		"projects",
		requestBody,
	)

	// THEN the project should be created successfully
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "New Project", project.Name)
}

func TestDoRequestErrorResponse(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a mock server that returns a 404 error
	path := "/api/v4/users/999"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(404)
		w.Write([]byte(`{"message": "Not found"}`))
	})

	// WHEN doRequest is called for a non-existent user
	user, resp, err := doRequest[*TestUser](
		client,
		http.MethodGet,
		"users/999",
		nil,
		nil,
	)

	// THEN the request should fail with a 404 error
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Nil(t, user)
}

func TestDoRequestSliceSuccess(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a mock server that returns a list of users from testdata
	path := "/api/v4/users"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_users.json")
	})

	// WHEN doRequestSlice is called to fetch users
	users, resp, err := doRequestSlice[TestUser](
		client,
		http.MethodGet,
		"users",
		nil,
		nil,
	)

	// THEN the request should succeed and return all users
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, users, 3)

	expectedUsers := []TestUser{
		{ID: 1, Name: "Example User 1"},
		{ID: 2, Name: "Example User 2"},
		{ID: 3, Name: "Example User 3"},
	}
	for i, user := range users {
		assert.Equal(t, expectedUsers[i].ID, user.ID)
		assert.Equal(t, expectedUsers[i].Name, user.Name)
	}
}

func TestDoRequestSliceEmptySlice(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a mock server that returns an empty array
	path := "/api/v4/users"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
	})

	// WHEN doRequestSlice is called on an endpoint with no data
	users, resp, err := doRequestSlice[TestUser](
		client,
		http.MethodGet,
		"users",
		nil,
		nil,
	)

	// THEN the request should succeed and return an empty slice
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Empty(t, users)
}

func TestDoRequestSliceErrorResponse(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a mock server that returns a 500 error
	path := "/api/v4/users"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(500)
		w.Write([]byte(`{"message": "Internal server error"}`))
	})

	// WHEN doRequestSlice is called on a failing endpoint
	users, resp, err := doRequestSlice[TestUser](
		client,
		http.MethodGet,
		"users",
		nil,
		nil,
	)

	// THEN the request should fail with a 500 error
	assert.Error(t, err)
	assert.Equal(t, 500, resp.StatusCode)
	assert.Nil(t, users)
}

func TestDoRequestVoidSuccessDELETE(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a mock server that accepts DELETE requests
	path := "/api/v4/users/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(204) // No Content
	})

	// WHEN doRequestVoid is called with DELETE method
	resp, err := doRequestVoid(
		client,
		http.MethodDelete,
		"users/1",
		nil,
		nil,
	)

	// THEN the request should succeed with no content
	assert.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestDoRequestVoidSuccessPUT(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a mock server that accepts PUT requests with body validation
	path := "/api/v4/merge_requests/1/approve"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)

		var reqBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.NoError(t, err)
		require.Equal(t, "approve", reqBody["action"])

		w.WriteHeader(200)
	})

	// GIVEN an approval request body
	requestBody := map[string]string{"action": "approve"}

	// WHEN doRequestVoid is called with PUT method and body
	resp, err := doRequestVoid(
		client,
		http.MethodPut,
		"merge_requests/1/approve",
		requestBody,
		nil,
	)

	// THEN the request should succeed
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestDoRequestVoidErrorResponse(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a mock server that returns a 403 forbidden error
	path := "/api/v4/users/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(403)
		w.Write([]byte(`{"message": "Forbidden"}`))
	})

	// WHEN doRequestVoid is called on a forbidden operation
	resp, err := doRequestVoid(
		client,
		http.MethodDelete,
		"users/1",
		nil,
		nil,
	)

	// THEN the request should fail with a 403 error
	assert.Error(t, err)
	assert.Equal(t, 403, resp.StatusCode)
}

func TestRequestHandlerWithOptions(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a mock server that validates query parameters and headers
	// from options
	path := "/api/v4/users"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		// Check that query parameters from options are included
		assert.Equal(t, "2", r.URL.Query().Get("page"))
		assert.Equal(t, "test-value", r.Header.Get("X-Test-Header"))

		mustWriteHTTPResponse(t, w, "testdata/list_users_public_email.json")
	})

	// GIVEN request options with pagination and custom header
	options := []RequestOptionFunc{
		WithOffsetPaginationParameters(2),
		WithHeader("X-Test-Header", "test-value"),
	}

	// WHEN doRequestSlice is called with request options
	users, resp, err := doRequestSlice[TestUser](
		client,
		http.MethodGet,
		"users",
		nil,
		options...,
	)

	// THEN the request should succeed with options applied
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, users, 1)
}
