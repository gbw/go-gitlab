//
// Copyright 2021, Igor Varavko
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
	"fmt"
	"net/http"
	"time"
)

type PipelineSource string

// PipelineSource is the source of a pipeline.
// GitLab API docs: https://docs.gitlab.com/ci/jobs/job_rules/#ci_pipeline_source-predefined-variable
const (
	PipelineSourceAPI                         PipelineSource = "api"
	PipelineSourceChat                        PipelineSource = "chat"
	PipelineSourceExternal                    PipelineSource = "external"
	PipelineSourceExternalPullRequestEvent    PipelineSource = "external_pull_request_event"
	PipelineSourceMergeRequestEvent           PipelineSource = "merge_request_event"
	PipelineSourceOndemandDastScan            PipelineSource = "ondemand_dast_scan"
	PipelineSourceOndemandDastValidation      PipelineSource = "ondemand_dast_validation"
	PipelineSourceParentPipeline              PipelineSource = "parent_pipeline"
	PipelineSourcePipeline                    PipelineSource = "pipeline"
	PipelineSourcePush                        PipelineSource = "push"
	PipelineSourceSchedule                    PipelineSource = "schedule"
	PipelineSourceSecurityOrchestrationPolicy PipelineSource = "security_orchestration_policy"
	PipelineSourceTrigger                     PipelineSource = "trigger"
	PipelineSourceWeb                         PipelineSource = "web"
	PipelineSourceWebIDE                      PipelineSource = "webide"
)

type (
	PipelinesServiceInterface interface {
		ListProjectPipelines(pid any, opt *ListProjectPipelinesOptions, options ...RequestOptionFunc) ([]*PipelineInfo, *Response, error)
		GetPipeline(pid any, pipeline int, options ...RequestOptionFunc) (*Pipeline, *Response, error)
		GetPipelineVariables(pid any, pipeline int, options ...RequestOptionFunc) ([]*PipelineVariable, *Response, error)
		GetPipelineTestReport(pid any, pipeline int, options ...RequestOptionFunc) (*PipelineTestReport, *Response, error)
		GetLatestPipeline(pid any, opt *GetLatestPipelineOptions, options ...RequestOptionFunc) (*Pipeline, *Response, error)
		CreatePipeline(pid any, opt *CreatePipelineOptions, options ...RequestOptionFunc) (*Pipeline, *Response, error)
		RetryPipelineBuild(pid any, pipeline int, options ...RequestOptionFunc) (*Pipeline, *Response, error)
		CancelPipelineBuild(pid any, pipeline int, options ...RequestOptionFunc) (*Pipeline, *Response, error)
		DeletePipeline(pid any, pipeline int, options ...RequestOptionFunc) (*Response, error)
		UpdatePipelineMetadata(pid any, pipeline int, opt *UpdatePipelineMetadataOptions, options ...RequestOptionFunc) (*Pipeline, *Response, error)
	}

	// PipelinesService handles communication with the repositories related
	// methods of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/pipelines/
	PipelinesService struct {
		client *Client
	}
)

var _ PipelinesServiceInterface = (*PipelinesService)(nil)

// PipelineVariable represents a pipeline variable.
//
// GitLab API docs: https://docs.gitlab.com/api/pipelines/
type PipelineVariable struct {
	Key          string            `json:"key"`
	Value        string            `json:"value"`
	VariableType VariableTypeValue `json:"variable_type"`
}

// PipelineInput represents a pipeline input.
//
// GitLab API docs: https://docs.gitlab.com/api/pipelines/
type PipelineInput struct {
	Name  string `json:"key"`
	Value any    `json:"value"`
}

// Pipeline represents a GitLab pipeline.
//
// GitLab API docs: https://docs.gitlab.com/api/pipelines/
type Pipeline struct {
	ID             int             `json:"id"`
	IID            int             `json:"iid"`
	ProjectID      int             `json:"project_id"`
	Status         string          `json:"status"`
	Source         PipelineSource  `json:"source"`
	Ref            string          `json:"ref"`
	Name           string          `json:"name"`
	SHA            string          `json:"sha"`
	BeforeSHA      string          `json:"before_sha"`
	Tag            bool            `json:"tag"`
	YamlErrors     string          `json:"yaml_errors"`
	User           *BasicUser      `json:"user"`
	UpdatedAt      *time.Time      `json:"updated_at"`
	CreatedAt      *time.Time      `json:"created_at"`
	StartedAt      *time.Time      `json:"started_at"`
	FinishedAt     *time.Time      `json:"finished_at"`
	CommittedAt    *time.Time      `json:"committed_at"`
	Duration       int             `json:"duration"`
	QueuedDuration int             `json:"queued_duration"`
	Coverage       string          `json:"coverage"`
	WebURL         string          `json:"web_url"`
	DetailedStatus *DetailedStatus `json:"detailed_status"`
}

// DetailedStatus contains detailed information about the status of a pipeline.
type DetailedStatus struct {
	Icon         string `json:"icon"`
	Text         string `json:"text"`
	Label        string `json:"label"`
	Group        string `json:"group"`
	Tooltip      string `json:"tooltip"`
	HasDetails   bool   `json:"has_details"`
	DetailsPath  string `json:"details_path"`
	Illustration struct {
		Image string `json:"image"`
	} `json:"illustration"`
	Favicon string `json:"favicon"`
}

func (p Pipeline) String() string {
	return Stringify(p)
}

// PipelineTestReport contains a detailed report of a test run.
type PipelineTestReport struct {
	TotalTime    float64               `json:"total_time"`
	TotalCount   int                   `json:"total_count"`
	SuccessCount int                   `json:"success_count"`
	FailedCount  int                   `json:"failed_count"`
	SkippedCount int                   `json:"skipped_count"`
	ErrorCount   int                   `json:"error_count"`
	TestSuites   []*PipelineTestSuites `json:"test_suites"`
}

// PipelineTestSuites contains test suites results.
type PipelineTestSuites struct {
	Name         string               `json:"name"`
	TotalTime    float64              `json:"total_time"`
	TotalCount   int                  `json:"total_count"`
	SuccessCount int                  `json:"success_count"`
	FailedCount  int                  `json:"failed_count"`
	SkippedCount int                  `json:"skipped_count"`
	ErrorCount   int                  `json:"error_count"`
	TestCases    []*PipelineTestCases `json:"test_cases"`
}

// PipelineTestCases contains test cases details.
type PipelineTestCases struct {
	Status         string          `json:"status"`
	Name           string          `json:"name"`
	Classname      string          `json:"classname"`
	File           string          `json:"file"`
	ExecutionTime  float64         `json:"execution_time"`
	SystemOutput   any             `json:"system_output"`
	StackTrace     string          `json:"stack_trace"`
	AttachmentURL  string          `json:"attachment_url"`
	RecentFailures *RecentFailures `json:"recent_failures"`
}

// RecentFailures contains failures count for the project's default branch.
type RecentFailures struct {
	Count      int    `json:"count"`
	BaseBranch string `json:"base_branch"`
}

func (p PipelineTestReport) String() string {
	return Stringify(p)
}

// PipelineInfo shows the basic entities of a pipeline, mostly used as fields
// on other assets, like Commit.
type PipelineInfo struct {
	ID        int        `json:"id"`
	IID       int        `json:"iid"`
	ProjectID int        `json:"project_id"`
	Status    string     `json:"status"`
	Source    string     `json:"source"`
	Ref       string     `json:"ref"`
	SHA       string     `json:"sha"`
	Name      string     `json:"name"`
	WebURL    string     `json:"web_url"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}

func (p PipelineInfo) String() string {
	return Stringify(p)
}

// ListProjectPipelinesOptions represents the available ListProjectPipelines()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#list-project-pipelines
type ListProjectPipelinesOptions struct {
	ListOptions
	Scope         *string          `url:"scope,omitempty" json:"scope,omitempty"`
	Status        *BuildStateValue `url:"status,omitempty" json:"status,omitempty"`
	Source        *string          `url:"source,omitempty" json:"source,omitempty"`
	Ref           *string          `url:"ref,omitempty" json:"ref,omitempty"`
	SHA           *string          `url:"sha,omitempty" json:"sha,omitempty"`
	YamlErrors    *bool            `url:"yaml_errors,omitempty" json:"yaml_errors,omitempty"`
	Name          *string          `url:"name,omitempty" json:"name,omitempty"`
	Username      *string          `url:"username,omitempty" json:"username,omitempty"`
	UpdatedAfter  *time.Time       `url:"updated_after,omitempty" json:"updated_after,omitempty"`
	UpdatedBefore *time.Time       `url:"updated_before,omitempty" json:"updated_before,omitempty"`
	OrderBy       *string          `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort          *string          `url:"sort,omitempty" json:"sort,omitempty"`
	CreatedAfter  *time.Time       `url:"created_after,omitempty" json:"created_after,omitempty"`
	CreatedBefore *time.Time       `url:"created_before,omitempty" json:"created_before,omitempty"`
}

// ListProjectPipelines gets a list of project pipelines.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#list-project-pipelines
func (s *PipelinesService) ListProjectPipelines(pid any, opt *ListProjectPipelinesOptions, options ...RequestOptionFunc) ([]*PipelineInfo, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pipelines", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*PipelineInfo
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// GetPipeline gets a single project pipeline.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#get-a-single-pipeline
func (s *PipelinesService) GetPipeline(pid any, pipeline int, options ...RequestOptionFunc) (*Pipeline, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pipelines/%d", PathEscape(project), pipeline)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Pipeline)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// GetPipelineVariables gets the variables of a single project pipeline.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#get-variables-of-a-pipeline
func (s *PipelinesService) GetPipelineVariables(pid any, pipeline int, options ...RequestOptionFunc) ([]*PipelineVariable, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pipelines/%d/variables", PathEscape(project), pipeline)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*PipelineVariable
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// GetPipelineTestReport gets the test report of a single project pipeline.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#get-a-pipelines-test-report
func (s *PipelinesService) GetPipelineTestReport(pid any, pipeline int, options ...RequestOptionFunc) (*PipelineTestReport, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pipelines/%d/test_report", PathEscape(project), pipeline)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(PipelineTestReport)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// GetLatestPipelineOptions represents the available GetLatestPipeline() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#get-the-latest-pipeline
type GetLatestPipelineOptions struct {
	Ref *string `url:"ref,omitempty" json:"ref,omitempty"`
}

// GetLatestPipeline gets the latest pipeline for a specific ref in a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#get-the-latest-pipeline
func (s *PipelinesService) GetLatestPipeline(pid any, opt *GetLatestPipelineOptions, options ...RequestOptionFunc) (*Pipeline, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pipelines/latest", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Pipeline)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// CreatePipelineOptions represents the available CreatePipeline() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#create-a-new-pipeline
type CreatePipelineOptions struct {
	Ref       *string                     `url:"ref" json:"ref"`
	Variables *[]*PipelineVariableOptions `url:"variables,omitempty" json:"variables,omitempty"`
}

// PipelineVariableOptions represents a pipeline variable option.
//
// GitLab API docs: https://docs.gitlab.com/api/pipelines/#create-a-new-pipeline
type PipelineVariableOptions struct {
	Key          *string            `url:"key,omitempty" json:"key,omitempty"`
	Value        *string            `url:"value,omitempty" json:"value,omitempty"`
	VariableType *VariableTypeValue `url:"variable_type,omitempty" json:"variable_type,omitempty"`
}

// CreatePipeline creates a new project pipeline.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#create-a-new-pipeline
func (s *PipelinesService) CreatePipeline(pid any, opt *CreatePipelineOptions, options ...RequestOptionFunc) (*Pipeline, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pipeline", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Pipeline)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// RetryPipelineBuild retries failed builds in a pipeline.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#retry-jobs-in-a-pipeline
func (s *PipelinesService) RetryPipelineBuild(pid any, pipeline int, options ...RequestOptionFunc) (*Pipeline, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pipelines/%d/retry", PathEscape(project), pipeline)

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Pipeline)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// CancelPipelineBuild cancels a pipeline builds.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#cancel-a-pipelines-jobs
func (s *PipelinesService) CancelPipelineBuild(pid any, pipeline int, options ...RequestOptionFunc) (*Pipeline, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pipelines/%d/cancel", PathEscape(project), pipeline)

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Pipeline)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// DeletePipeline deletes an existing pipeline.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#delete-a-pipeline
func (s *PipelinesService) DeletePipeline(pid any, pipeline int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/pipelines/%d", PathEscape(project), pipeline)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UpdatePipelineMetadataOptions represents the available UpdatePipelineMetadata()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#update-pipeline-metadata
type UpdatePipelineMetadataOptions struct {
	Name *string `url:"name,omitempty" json:"name,omitempty"`
}

// UpdatePipelineMetadata You can update the metadata of a pipeline. The metadata
// contains the name of the pipeline.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipelines/#update-pipeline-metadata
func (s *PipelinesService) UpdatePipelineMetadata(pid any, pipeline int, opt *UpdatePipelineMetadataOptions, options ...RequestOptionFunc) (*Pipeline, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pipelines/%d/metadata", PathEscape(project), pipeline)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Pipeline)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}
