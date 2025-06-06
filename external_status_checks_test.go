package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListMergeStatusChecks(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/1/status_checks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, exampleStatusChecks)
	})

	statusChecks, _, err := client.ExternalStatusChecks.ListMergeStatusChecks(1, 1, nil)
	if err != nil {
		t.Fatalf("ExternalStatusChecks.ListMergeStatusChecks returns an error: %v", err)
	}

	expectedStatusChecks := []*MergeStatusCheck{
		{
			ID:          2,
			Name:        "Rule 1",
			ExternalURL: "https://gitlab.com/test-endpoint",
			Status:      "approved",
		},
		{
			ID:          1,
			Name:        "Rule 2",
			ExternalURL: "https://gitlab.com/test-endpoint-2",
			Status:      "pending",
		},
	}

	assert.Equal(t, expectedStatusChecks, statusChecks)
}

func TestListProjectStatusChecks(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/external_status_checks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, exampleProjectStatusChecks)
	})

	projectStatusChecks, _, err := client.ExternalStatusChecks.ListProjectStatusChecks(1, nil)
	if err != nil {
		t.Fatalf("ExternalStatusChecks.ListProjectStatusChecks returns an error: %v", err)
	}

	expectedProjectStatusChecks := []*ProjectStatusCheck{
		{
			ID:          1,
			Name:        "Compliance Check",
			ProjectID:   6,
			ExternalURL: "https://gitlab.com/example/test.json",
			ProtectedBranches: []StatusCheckProtectedBranch{
				{
					ID:                        14,
					ProjectID:                 6,
					Name:                      "master",
					CreatedAt:                 mustParseTime("2020-10-12T14:04:50.787Z"),
					UpdatedAt:                 mustParseTime("2020-10-12T14:04:50.787Z"),
					CodeOwnerApprovalRequired: false,
				},
			},
		},
	}

	assert.Equal(t, expectedProjectStatusChecks, projectStatusChecks)
}

func TestRetryFailedStatusCheckForAMergeRequest(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/2/status_checks/3/retry", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"message": "202 Accepted"}`)
	})

	resp, err := client.ExternalStatusChecks.RetryFailedStatusCheckForAMergeRequest(1, 2, 3)
	if err != nil {
		t.Fatalf("ExternalStatusChecks.RetryFailedStatusCheckForAMergeRequest returns an error: %v", err)
	}

	assert.NotNil(t, resp)
}
