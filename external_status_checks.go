package gitlab

import (
	"fmt"
	"net/http"
	"time"
)

type (
	// ExternalStatusChecksServiceInterface defines all the API methods for the ExternalStatusChecksService
	ExternalStatusChecksServiceInterface interface {
		ListMergeStatusChecks(pid any, mr int, opt *ListOptions, options ...RequestOptionFunc) ([]*MergeStatusCheck, *Response, error)
		SetExternalStatusCheckStatus(pid any, mergeRequest int, opt *SetExternalStatusCheckStatusOptions, options ...RequestOptionFunc) (*Response, error)
		ListProjectStatusChecks(pid any, opt *ListOptions, options ...RequestOptionFunc) ([]*ProjectStatusCheck, *Response, error)
		CreateExternalStatusCheck(pid any, opt *CreateExternalStatusCheckOptions, options ...RequestOptionFunc) (*Response, error)
		DeleteExternalStatusCheck(pid any, check int, options ...RequestOptionFunc) (*Response, error)
		UpdateExternalStatusCheck(pid any, check int, opt *UpdateExternalStatusCheckOptions, options ...RequestOptionFunc) (*Response, error)
		RetryFailedStatusCheckForAMergeRequest(pid any, mergeRequest int, externalStatusCheck int, options ...RequestOptionFunc) (*Response, error)
	}

	// ExternalStatusChecksService handles communication with the external
	// status check related methods of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/status_checks/
	ExternalStatusChecksService struct {
		client *Client
	}
)

var _ ExternalStatusChecksServiceInterface = (*ExternalStatusChecksService)(nil)

type MergeStatusCheck struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ExternalURL string `json:"external_url"`
	Status      string `json:"status"`
}

type ProjectStatusCheck struct {
	ID                int                          `json:"id"`
	Name              string                       `json:"name"`
	ProjectID         int                          `json:"project_id"`
	ExternalURL       string                       `json:"external_url"`
	ProtectedBranches []StatusCheckProtectedBranch `json:"protected_branches"`
}

type StatusCheckProtectedBranch struct {
	ID                        int        `json:"id"`
	ProjectID                 int        `json:"project_id"`
	Name                      string     `json:"name"`
	CreatedAt                 *time.Time `json:"created_at"`
	UpdatedAt                 *time.Time `json:"updated_at"`
	CodeOwnerApprovalRequired bool       `json:"code_owner_approval_required"`
}

// ListMergeStatusChecks lists the external status checks that apply to it
// and their status for a single merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/api/status_checks/#list-status-checks-for-a-merge-request
func (s *ExternalStatusChecksService) ListMergeStatusChecks(pid any, mr int, opt *ListOptions, options ...RequestOptionFunc) ([]*MergeStatusCheck, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/status_checks", PathEscape(project), mr)

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var mscs []*MergeStatusCheck
	resp, err := s.client.Do(req, &mscs)
	if err != nil {
		return nil, resp, err
	}

	return mscs, resp, nil
}

// SetExternalStatusCheckStatusOptions represents the available
// SetExternalStatusCheckStatus() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/status_checks/#set-status-of-an-external-status-check
type SetExternalStatusCheckStatusOptions struct {
	SHA                   *string `url:"sha,omitempty" json:"sha,omitempty"`
	ExternalStatusCheckID *int    `url:"external_status_check_id,omitempty" json:"external_status_check_id,omitempty"`
	Status                *string `url:"status,omitempty" json:"status,omitempty"`
}

// SetExternalStatusCheckStatus sets the status of an external status check.
//
// Gitlab API docs:
// https://docs.gitlab.com/api/status_checks/#set-status-of-an-external-status-check
func (s *ExternalStatusChecksService) SetExternalStatusCheckStatus(pid any, mergeRequest int, opt *SetExternalStatusCheckStatusOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/status_check_responses", PathEscape(project), mergeRequest)

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListProjectStatusChecks lists the project external status checks.
//
// GitLab API docs:
// https://docs.gitlab.com/api/status_checks/#get-project-external-status-check-services
func (s *ExternalStatusChecksService) ListProjectStatusChecks(pid any, opt *ListOptions, options ...RequestOptionFunc) ([]*ProjectStatusCheck, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/external_status_checks", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var pscs []*ProjectStatusCheck
	resp, err := s.client.Do(req, &pscs)
	if err != nil {
		return nil, resp, err
	}

	return pscs, resp, nil
}

// CreateExternalStatusCheckOptions represents the available
// CreateExternalStatusCheck() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/status_checks/#create-external-status-check-service
type CreateExternalStatusCheckOptions struct {
	Name               *string `url:"name,omitempty" json:"name,omitempty"`
	ExternalURL        *string `url:"external_url,omitempty" json:"external_url,omitempty"`
	ProtectedBranchIDs *[]int  `url:"protected_branch_ids,omitempty" json:"protected_branch_ids,omitempty"`
}

// CreateExternalStatusCheck creates an external status check.
//
// Gitlab API docs:
// https://docs.gitlab.com/api/status_checks/#create-external-status-check-service
func (s *ExternalStatusChecksService) CreateExternalStatusCheck(pid any, opt *CreateExternalStatusCheckOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/external_status_checks", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DeleteExternalStatusCheck deletes an external status check.
//
// Gitlab API docs:
// https://docs.gitlab.com/api/status_checks/#delete-external-status-check-service
func (s *ExternalStatusChecksService) DeleteExternalStatusCheck(pid any, check int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/external_status_checks/%d", PathEscape(project), check)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UpdateExternalStatusCheckOptions represents the available
// UpdateExternalStatusCheck() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/status_checks/#update-external-status-check-service
type UpdateExternalStatusCheckOptions struct {
	Name               *string `url:"name,omitempty" json:"name,omitempty"`
	ExternalURL        *string `url:"external_url,omitempty" json:"external_url,omitempty"`
	ProtectedBranchIDs *[]int  `url:"protected_branch_ids,omitempty" json:"protected_branch_ids,omitempty"`
}

// UpdateExternalStatusCheck updates an external status check.
//
// Gitlab API docs:
// https://docs.gitlab.com/api/status_checks/#update-external-status-check-service
func (s *ExternalStatusChecksService) UpdateExternalStatusCheck(pid any, check int, opt *UpdateExternalStatusCheckOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/external_status_checks/%d", PathEscape(project), check)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// RetryFailedStatusCheckForAMergeRequest retries the specified failed external status check.
//
// Gitlab API docs:
// https://docs.gitlab.com/api/status_checks/#retry-failed-status-check-for-a-merge-request
func (s *ExternalStatusChecksService) RetryFailedStatusCheckForAMergeRequest(pid any, mergeRequest int, externalStatusCheck int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/status_checks/%d/retry", PathEscape(project), mergeRequest, externalStatusCheck)

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
