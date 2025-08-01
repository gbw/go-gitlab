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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type (
	// ProjectsServiceInterface handles communication with the repositories related methods
	// of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/projects/
	ProjectsServiceInterface interface {
		ListProjects(opt *ListProjectsOptions, options ...RequestOptionFunc) ([]*Project, *Response, error)
		ListUserProjects(uid any, opt *ListProjectsOptions, options ...RequestOptionFunc) ([]*Project, *Response, error)
		ListUserContributedProjects(uid any, opt *ListProjectsOptions, options ...RequestOptionFunc) ([]*Project, *Response, error)
		ListUserStarredProjects(uid any, opt *ListProjectsOptions, options ...RequestOptionFunc) ([]*Project, *Response, error)
		ListProjectsUsers(pid any, opt *ListProjectUserOptions, options ...RequestOptionFunc) ([]*ProjectUser, *Response, error)
		ListProjectsGroups(pid any, opt *ListProjectGroupOptions, options ...RequestOptionFunc) ([]*ProjectGroup, *Response, error)
		GetProjectLanguages(pid any, options ...RequestOptionFunc) (*ProjectLanguages, *Response, error)
		GetProject(pid any, opt *GetProjectOptions, options ...RequestOptionFunc) (*Project, *Response, error)
		CreateProject(opt *CreateProjectOptions, options ...RequestOptionFunc) (*Project, *Response, error)
		CreateProjectForUser(user int, opt *CreateProjectForUserOptions, options ...RequestOptionFunc) (*Project, *Response, error)
		EditProject(pid any, opt *EditProjectOptions, options ...RequestOptionFunc) (*Project, *Response, error)
		ForkProject(pid any, opt *ForkProjectOptions, options ...RequestOptionFunc) (*Project, *Response, error)
		StarProject(pid any, options ...RequestOptionFunc) (*Project, *Response, error)
		ListProjectsInvitedGroups(pid any, opt *ListProjectInvidedGroupOptions, options ...RequestOptionFunc) ([]*ProjectGroup, *Response, error)
		UnstarProject(pid any, options ...RequestOptionFunc) (*Project, *Response, error)
		ArchiveProject(pid any, options ...RequestOptionFunc) (*Project, *Response, error)
		UnarchiveProject(pid any, options ...RequestOptionFunc) (*Project, *Response, error)
		RestoreProject(pid any, options ...RequestOptionFunc) (*Project, *Response, error)
		DeleteProject(pid any, opt *DeleteProjectOptions, options ...RequestOptionFunc) (*Response, error)
		ShareProjectWithGroup(pid any, opt *ShareWithGroupOptions, options ...RequestOptionFunc) (*Response, error)
		DeleteSharedProjectFromGroup(pid any, groupID int, options ...RequestOptionFunc) (*Response, error)
		ListProjectHooks(pid any, opt *ListProjectHooksOptions, options ...RequestOptionFunc) ([]*ProjectHook, *Response, error)
		GetProjectHook(pid any, hook int, options ...RequestOptionFunc) (*ProjectHook, *Response, error)
		AddProjectHook(pid any, opt *AddProjectHookOptions, options ...RequestOptionFunc) (*ProjectHook, *Response, error)
		EditProjectHook(pid any, hook int, opt *EditProjectHookOptions, options ...RequestOptionFunc) (*ProjectHook, *Response, error)
		DeleteProjectHook(pid any, hook int, options ...RequestOptionFunc) (*Response, error)
		TriggerTestProjectHook(pid any, hook int, event ProjectHookEvent, options ...RequestOptionFunc) (*Response, error)
		SetProjectCustomHeader(pid any, hook int, key string, opt *SetHookCustomHeaderOptions, options ...RequestOptionFunc) (*Response, error)
		DeleteProjectCustomHeader(pid any, hook int, key string, options ...RequestOptionFunc) (*Response, error)
		SetProjectWebhookURLVariable(pid any, hook int, key string, opt *SetProjectWebhookURLVariableOptions, options ...RequestOptionFunc) (*Response, error)
		DeleteProjectWebhookURLVariable(pid any, hook int, key string, options ...RequestOptionFunc) (*Response, error)
		CreateProjectForkRelation(pid any, fork int, options ...RequestOptionFunc) (*ProjectForkRelation, *Response, error)
		DeleteProjectForkRelation(pid any, options ...RequestOptionFunc) (*Response, error)
		UploadAvatar(pid any, avatar io.Reader, filename string, options ...RequestOptionFunc) (*Project, *Response, error)
		DownloadAvatar(pid any, options ...RequestOptionFunc) (*bytes.Reader, *Response, error)
		ListProjectForks(pid any, opt *ListProjectsOptions, options ...RequestOptionFunc) ([]*Project, *Response, error)
		GetProjectPushRules(pid any, options ...RequestOptionFunc) (*ProjectPushRules, *Response, error)
		AddProjectPushRule(pid any, opt *AddProjectPushRuleOptions, options ...RequestOptionFunc) (*ProjectPushRules, *Response, error)
		EditProjectPushRule(pid any, opt *EditProjectPushRuleOptions, options ...RequestOptionFunc) (*ProjectPushRules, *Response, error)
		DeleteProjectPushRule(pid any, options ...RequestOptionFunc) (*Response, error)
		GetApprovalConfiguration(pid any, options ...RequestOptionFunc) (*ProjectApprovals, *Response, error)
		ChangeApprovalConfiguration(pid any, opt *ChangeApprovalConfigurationOptions, options ...RequestOptionFunc) (*ProjectApprovals, *Response, error)
		GetProjectApprovalRules(pid any, opt *GetProjectApprovalRulesListsOptions, options ...RequestOptionFunc) ([]*ProjectApprovalRule, *Response, error)
		GetProjectApprovalRule(pid any, ruleID int, options ...RequestOptionFunc) (*ProjectApprovalRule, *Response, error)
		CreateProjectApprovalRule(pid any, opt *CreateProjectLevelRuleOptions, options ...RequestOptionFunc) (*ProjectApprovalRule, *Response, error)
		UpdateProjectApprovalRule(pid any, approvalRule int, opt *UpdateProjectLevelRuleOptions, options ...RequestOptionFunc) (*ProjectApprovalRule, *Response, error)
		DeleteProjectApprovalRule(pid any, approvalRule int, options ...RequestOptionFunc) (*Response, error)
		GetProjectPullMirrorDetails(pid any, options ...RequestOptionFunc) (*ProjectPullMirrorDetails, *Response, error)
		ConfigureProjectPullMirror(pid any, opt *ConfigureProjectPullMirrorOptions, options ...RequestOptionFunc) (*ProjectPullMirrorDetails, *Response, error)
		StartMirroringProject(pid any, options ...RequestOptionFunc) (*Response, error)
		TransferProject(pid any, opt *TransferProjectOptions, options ...RequestOptionFunc) (*Project, *Response, error)
		StartHousekeepingProject(pid any, options ...RequestOptionFunc) (*Response, error)
		GetRepositoryStorage(pid any, options ...RequestOptionFunc) (*ProjectReposityStorage, *Response, error)
	}

	// ProjectsService handles communication with the repositories related methods
	// of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/projects/
	ProjectsService struct {
		client *Client
	}
)

var _ ProjectsServiceInterface = (*ProjectsService)(nil)

// Project represents a GitLab project.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/
type Project struct {
	ID                                        int                        `json:"id"`
	Description                               string                     `json:"description"`
	DefaultBranch                             string                     `json:"default_branch"`
	Visibility                                VisibilityValue            `json:"visibility"`
	SSHURLToRepo                              string                     `json:"ssh_url_to_repo"`
	HTTPURLToRepo                             string                     `json:"http_url_to_repo"`
	WebURL                                    string                     `json:"web_url"`
	ReadmeURL                                 string                     `json:"readme_url"`
	Topics                                    []string                   `json:"topics"`
	Owner                                     *User                      `json:"owner"`
	Name                                      string                     `json:"name"`
	NameWithNamespace                         string                     `json:"name_with_namespace"`
	Path                                      string                     `json:"path"`
	PathWithNamespace                         string                     `json:"path_with_namespace"`
	OpenIssuesCount                           int                        `json:"open_issues_count"`
	ResolveOutdatedDiffDiscussions            bool                       `json:"resolve_outdated_diff_discussions"`
	ContainerExpirationPolicy                 *ContainerExpirationPolicy `json:"container_expiration_policy,omitempty"`
	ContainerRegistryAccessLevel              AccessControlValue         `json:"container_registry_access_level"`
	ContainerRegistryImagePrefix              string                     `json:"container_registry_image_prefix,omitempty"`
	CreatedAt                                 *time.Time                 `json:"created_at,omitempty"`
	UpdatedAt                                 *time.Time                 `json:"updated_at,omitempty"`
	LastActivityAt                            *time.Time                 `json:"last_activity_at,omitempty"`
	CreatorID                                 int                        `json:"creator_id"`
	Namespace                                 *ProjectNamespace          `json:"namespace"`
	Permissions                               *Permissions               `json:"permissions"`
	MarkedForDeletionOn                       *ISOTime                   `json:"marked_for_deletion_on"`
	EmptyRepo                                 bool                       `json:"empty_repo"`
	Archived                                  bool                       `json:"archived"`
	AvatarURL                                 string                     `json:"avatar_url"`
	LicenseURL                                string                     `json:"license_url"`
	License                                   *ProjectLicense            `json:"license"`
	SharedRunnersEnabled                      bool                       `json:"shared_runners_enabled"`
	GroupRunnersEnabled                       bool                       `json:"group_runners_enabled"`
	RunnerTokenExpirationInterval             int                        `json:"runner_token_expiration_interval"`
	ForksCount                                int                        `json:"forks_count"`
	StarCount                                 int                        `json:"star_count"`
	RunnersToken                              string                     `json:"runners_token"`
	AllowMergeOnSkippedPipeline               bool                       `json:"allow_merge_on_skipped_pipeline"`
	AllowPipelineTriggerApproveDeployment     bool                       `json:"allow_pipeline_trigger_approve_deployment"`
	OnlyAllowMergeIfPipelineSucceeds          bool                       `json:"only_allow_merge_if_pipeline_succeeds"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool                       `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool                       `json:"remove_source_branch_after_merge"`
	PreventMergeWithoutJiraIssue              bool                       `json:"prevent_merge_without_jira_issue"`
	PrintingMergeRequestLinkEnabled           bool                       `json:"printing_merge_request_link_enabled"`
	LFSEnabled                                bool                       `json:"lfs_enabled"`
	RepositoryStorage                         string                     `json:"repository_storage"`
	RequestAccessEnabled                      bool                       `json:"request_access_enabled"`
	MergeMethod                               MergeMethodValue           `json:"merge_method"`
	CanCreateMergeRequestIn                   bool                       `json:"can_create_merge_request_in"`
	ForkedFromProject                         *ForkParent                `json:"forked_from_project"`
	Mirror                                    bool                       `json:"mirror"`
	MirrorUserID                              int                        `json:"mirror_user_id"`
	MirrorTriggerBuilds                       bool                       `json:"mirror_trigger_builds"`
	OnlyMirrorProtectedBranches               bool                       `json:"only_mirror_protected_branches"`
	MirrorOverwritesDivergedBranches          bool                       `json:"mirror_overwrites_diverged_branches"`
	PackagesEnabled                           bool                       `json:"packages_enabled"`
	ServiceDeskEnabled                        bool                       `json:"service_desk_enabled"`
	ServiceDeskAddress                        string                     `json:"service_desk_address"`
	IssuesAccessLevel                         AccessControlValue         `json:"issues_access_level"`
	ReleasesAccessLevel                       AccessControlValue         `json:"releases_access_level,omitempty"`
	RepositoryAccessLevel                     AccessControlValue         `json:"repository_access_level"`
	MergeRequestsAccessLevel                  AccessControlValue         `json:"merge_requests_access_level"`
	ForkingAccessLevel                        AccessControlValue         `json:"forking_access_level"`
	WikiAccessLevel                           AccessControlValue         `json:"wiki_access_level"`
	BuildsAccessLevel                         AccessControlValue         `json:"builds_access_level"`
	SnippetsAccessLevel                       AccessControlValue         `json:"snippets_access_level"`
	PagesAccessLevel                          AccessControlValue         `json:"pages_access_level"`
	OperationsAccessLevel                     AccessControlValue         `json:"operations_access_level"`
	AnalyticsAccessLevel                      AccessControlValue         `json:"analytics_access_level"`
	EnvironmentsAccessLevel                   AccessControlValue         `json:"environments_access_level"`
	FeatureFlagsAccessLevel                   AccessControlValue         `json:"feature_flags_access_level"`
	InfrastructureAccessLevel                 AccessControlValue         `json:"infrastructure_access_level"`
	MonitorAccessLevel                        AccessControlValue         `json:"monitor_access_level"`
	AutocloseReferencedIssues                 bool                       `json:"autoclose_referenced_issues"`
	SuggestionCommitMessage                   string                     `json:"suggestion_commit_message"`
	SquashOption                              SquashOptionValue          `json:"squash_option"`
	EnforceAuthChecksOnUploads                bool                       `json:"enforce_auth_checks_on_uploads,omitempty"`
	SharedWithGroups                          []struct {
		GroupID          int    `json:"group_id"`
		GroupName        string `json:"group_name"`
		GroupFullPath    string `json:"group_full_path"`
		GroupAccessLevel int    `json:"group_access_level"`
	} `json:"shared_with_groups"`
	Statistics                               *Statistics                                 `json:"statistics"`
	Links                                    *Links                                      `json:"_links,omitempty"`
	ImportURL                                string                                      `json:"import_url"`
	ImportType                               string                                      `json:"import_type"`
	ImportStatus                             string                                      `json:"import_status"`
	ImportError                              string                                      `json:"import_error"`
	CIDefaultGitDepth                        int                                         `json:"ci_default_git_depth"`
	CIDeletePipelinesInSeconds               int                                         `json:"ci_delete_pipelines_in_seconds,omitempty"`
	CIForwardDeploymentEnabled               bool                                        `json:"ci_forward_deployment_enabled"`
	CIForwardDeploymentRollbackAllowed       bool                                        `json:"ci_forward_deployment_rollback_allowed"`
	CIPushRepositoryForJobTokenAllowed       bool                                        `json:"ci_push_repository_for_job_token_allowed"`
	CIIdTokenSubClaimComponents              []string                                    `json:"ci_id_token_sub_claim_components"`
	CISeperateCache                          bool                                        `json:"ci_separated_caches"`
	CIJobTokenScopeEnabled                   bool                                        `json:"ci_job_token_scope_enabled"`
	CIOptInJWT                               bool                                        `json:"ci_opt_in_jwt"`
	CIAllowForkPipelinesToRunInParentProject bool                                        `json:"ci_allow_fork_pipelines_to_run_in_parent_project"`
	CIRestrictPipelineCancellationRole       AccessControlValue                          `json:"ci_restrict_pipeline_cancellation_role"`
	PublicJobs                               bool                                        `json:"public_jobs"`
	BuildTimeout                             int                                         `json:"build_timeout"`
	AutoCancelPendingPipelines               string                                      `json:"auto_cancel_pending_pipelines"`
	CIConfigPath                             string                                      `json:"ci_config_path"`
	CustomAttributes                         []*CustomAttribute                          `json:"custom_attributes"`
	ComplianceFrameworks                     []string                                    `json:"compliance_frameworks"`
	BuildCoverageRegex                       string                                      `json:"build_coverage_regex"`
	IssuesTemplate                           string                                      `json:"issues_template"`
	MergeRequestsTemplate                    string                                      `json:"merge_requests_template"`
	IssueBranchTemplate                      string                                      `json:"issue_branch_template"`
	KeepLatestArtifact                       bool                                        `json:"keep_latest_artifact"`
	MergePipelinesEnabled                    bool                                        `json:"merge_pipelines_enabled"`
	MergeTrainsEnabled                       bool                                        `json:"merge_trains_enabled"`
	MergeTrainsSkipTrainAllowed              bool                                        `json:"merge_trains_skip_train_allowed"`
	CIPipelineVariablesMinimumOverrideRole   CIPipelineVariablesMinimumOverrideRoleValue `json:"ci_pipeline_variables_minimum_override_role"`
	MergeCommitTemplate                      string                                      `json:"merge_commit_template"`
	SquashCommitTemplate                     string                                      `json:"squash_commit_template"`
	AutoDevopsDeployStrategy                 string                                      `json:"auto_devops_deploy_strategy"`
	AutoDevopsEnabled                        bool                                        `json:"auto_devops_enabled"`
	BuildGitStrategy                         string                                      `json:"build_git_strategy"`
	EmailsEnabled                            bool                                        `json:"emails_enabled"`
	ExternalAuthorizationClassificationLabel string                                      `json:"external_authorization_classification_label"`
	RequirementsEnabled                      bool                                        `json:"requirements_enabled"`
	RequirementsAccessLevel                  AccessControlValue                          `json:"requirements_access_level"`
	SecurityAndComplianceEnabled             bool                                        `json:"security_and_compliance_enabled"`
	SecurityAndComplianceAccessLevel         AccessControlValue                          `json:"security_and_compliance_access_level"`
	MergeRequestDefaultTargetSelf            bool                                        `json:"mr_default_target_self"`
	ModelExperimentsAccessLevel              AccessControlValue                          `json:"model_experiments_access_level"`
	ModelRegistryAccessLevel                 AccessControlValue                          `json:"model_registry_access_level"`
	PreReceiveSecretDetectionEnabled         bool                                        `json:"pre_receive_secret_detection_enabled"`
	AutoDuoCodeReviewEnabled                 bool                                        `json:"auto_duo_code_review_enabled"`

	// Deprecated: use Topics instead
	TagList []string `json:"tag_list"`
	// Deprecated: use IssuesAccessLevel instead
	IssuesEnabled bool `json:"issues_enabled"`
	// Deprecated: use MergeRequestsAccessLevel instead
	MergeRequestsEnabled bool `json:"merge_requests_enabled"`
	// Deprecated: use Merge Request Approvals API instead
	ApprovalsBeforeMerge int `json:"approvals_before_merge"`
	// Deprecated: use BuildsAccessLevel instead
	JobsEnabled bool `json:"jobs_enabled"`
	// Deprecated: use WikiAccessLevel instead
	WikiEnabled bool `json:"wiki_enabled"`
	// Deprecated: use SnippetsAccessLevel instead
	SnippetsEnabled bool `json:"snippets_enabled"`
	// Deprecated: use ContainerRegistryAccessLevel instead
	ContainerRegistryEnabled bool `json:"container_registry_enabled"`
	// Deprecated: use MarkedForDeletionOn instead
	MarkedForDeletionAt *ISOTime `json:"marked_for_deletion_at"`
	// Deprecated: use CIPipelineVariablesMinimumOverrideRole instead
	RestrictUserDefinedVariables bool `json:"restrict_user_defined_variables"`
	// Deprecated: Use EmailsEnabled instead
	EmailsDisabled bool `json:"emails_disabled"`
	// Deprecated: This parameter has been renamed to PublicJobs in GitLab 9.0.
	PublicBuilds bool `json:"public_builds"`
}

// BasicProject included in other service responses (such as todos).
type BasicProject struct {
	ID                int        `json:"id"`
	Description       string     `json:"description"`
	Name              string     `json:"name"`
	NameWithNamespace string     `json:"name_with_namespace"`
	Path              string     `json:"path"`
	PathWithNamespace string     `json:"path_with_namespace"`
	CreatedAt         *time.Time `json:"created_at"`
}

// ContainerExpirationPolicy represents the container expiration policy.
type ContainerExpirationPolicy struct {
	Cadence         string     `json:"cadence"`
	KeepN           int        `json:"keep_n"`
	OlderThan       string     `json:"older_than"`
	NameRegexDelete string     `json:"name_regex_delete"`
	NameRegexKeep   string     `json:"name_regex_keep"`
	Enabled         bool       `json:"enabled"`
	NextRunAt       *time.Time `json:"next_run_at"`

	// Deprecated: use NameRegexDelete instead
	NameRegex string `json:"name_regex"`
}

// ForkParent represents the parent project when this is a fork.
type ForkParent struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	NameWithNamespace string `json:"name_with_namespace"`
	Path              string `json:"path"`
	PathWithNamespace string `json:"path_with_namespace"`
	HTTPURLToRepo     string `json:"http_url_to_repo"`
	WebURL            string `json:"web_url"`
	RepositoryStorage string `json:"repository_storage"`
}

// GroupAccess represents group access.
type GroupAccess struct {
	AccessLevel       AccessLevelValue       `json:"access_level"`
	NotificationLevel NotificationLevelValue `json:"notification_level"`
}

// Links represents a project web links for self, issues, merge_requests,
// repo_branches, labels, events, members.
type Links struct {
	Self          string `json:"self"`
	Issues        string `json:"issues"`
	MergeRequests string `json:"merge_requests"`
	RepoBranches  string `json:"repo_branches"`
	Labels        string `json:"labels"`
	Events        string `json:"events"`
	Members       string `json:"members"`
	ClusterAgents string `json:"cluster_agents"`
}

// Permissions represents permissions.
type Permissions struct {
	ProjectAccess *ProjectAccess `json:"project_access"`
	GroupAccess   *GroupAccess   `json:"group_access"`
}

// ProjectAccess represents project access.
type ProjectAccess struct {
	AccessLevel       AccessLevelValue       `json:"access_level"`
	NotificationLevel NotificationLevelValue `json:"notification_level"`
}

// ProjectLicense represent the license for a project.
type ProjectLicense struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	HTMLURL   string `json:"html_url"`
	SourceURL string `json:"source_url"`
}

// ProjectNamespace represents a project namespace.
type ProjectNamespace struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Kind      string `json:"kind"`
	FullPath  string `json:"full_path"`
	ParentID  int    `json:"parent_id"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

// Repository represents a repository.
type Repository struct {
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	WebURL            string          `json:"web_url"`
	AvatarURL         string          `json:"avatar_url"`
	GitSSHURL         string          `json:"git_ssh_url"`
	GitHTTPURL        string          `json:"git_http_url"`
	Namespace         string          `json:"namespace"`
	Visibility        VisibilityValue `json:"visibility"`
	PathWithNamespace string          `json:"path_with_namespace"`
	DefaultBranch     string          `json:"default_branch"`
	Homepage          string          `json:"homepage"`
	URL               string          `json:"url"`
	SSHURL            string          `json:"ssh_url"`
	HTTPURL           string          `json:"http_url"`
}

// Statistics represents a statistics record for a group or project.
type Statistics struct {
	CommitCount           int64 `json:"commit_count"`
	StorageSize           int64 `json:"storage_size"`
	RepositorySize        int64 `json:"repository_size"`
	WikiSize              int64 `json:"wiki_size"`
	LFSObjectsSize        int64 `json:"lfs_objects_size"`
	JobArtifactsSize      int64 `json:"job_artifacts_size"`
	PipelineArtifactsSize int64 `json:"pipeline_artifacts_size"`
	PackagesSize          int64 `json:"packages_size"`
	SnippetsSize          int64 `json:"snippets_size"`
	UploadsSize           int64 `json:"uploads_size"`
	ContainerRegistrySize int64 `json:"container_registry_size"`
}

func (s Project) String() string {
	return Stringify(s)
}

// ProjectApprovalRule represents a GitLab project approval rule.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#get-all-approval-rules-for-project
type ProjectApprovalRule struct {
	ID                            int                `json:"id"`
	Name                          string             `json:"name"`
	RuleType                      string             `json:"rule_type"`
	ReportType                    string             `json:"report_type"`
	EligibleApprovers             []*BasicUser       `json:"eligible_approvers"`
	ApprovalsRequired             int                `json:"approvals_required"`
	Users                         []*BasicUser       `json:"users"`
	Groups                        []*Group           `json:"groups"`
	ContainsHiddenGroups          bool               `json:"contains_hidden_groups"`
	ProtectedBranches             []*ProtectedBranch `json:"protected_branches"`
	AppliesToAllProtectedBranches bool               `json:"applies_to_all_protected_branches"`
}

func (s ProjectApprovalRule) String() string {
	return Stringify(s)
}

// ListProjectsOptions represents the available ListProjects() options.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#list-all-projects
type ListProjectsOptions struct {
	ListOptions
	Active                   *bool             `url:"active,omitempty" json:"active,omitempty"`
	Archived                 *bool             `url:"archived,omitempty" json:"archived,omitempty"`
	IDAfter                  *int              `url:"id_after,omitempty" json:"id_after,omitempty"`
	IDBefore                 *int              `url:"id_before,omitempty" json:"id_before,omitempty"`
	Imported                 *bool             `url:"imported,omitempty" json:"imported,omitempty"`
	IncludeHidden            *bool             `url:"include_hidden,omitempty" json:"include_hidden,omitempty"`
	IncludePendingDelete     *bool             `url:"include_pending_delete,omitempty" json:"include_pending_delete,omitempty"`
	LastActivityAfter        *time.Time        `url:"last_activity_after,omitempty" json:"last_activity_after,omitempty"`
	LastActivityBefore       *time.Time        `url:"last_activity_before,omitempty" json:"last_activity_before,omitempty"`
	Membership               *bool             `url:"membership,omitempty" json:"membership,omitempty"`
	MinAccessLevel           *AccessLevelValue `url:"min_access_level,omitempty" json:"min_access_level,omitempty"`
	OrderBy                  *string           `url:"order_by,omitempty" json:"order_by,omitempty"`
	Owned                    *bool             `url:"owned,omitempty" json:"owned,omitempty"`
	RepositoryChecksumFailed *bool             `url:"repository_checksum_failed,omitempty" json:"repository_checksum_failed,omitempty"`
	RepositoryStorage        *string           `url:"repository_storage,omitempty" json:"repository_storage,omitempty"`
	Search                   *string           `url:"search,omitempty" json:"search,omitempty"`
	SearchNamespaces         *bool             `url:"search_namespaces,omitempty" json:"search_namespaces,omitempty"`
	Simple                   *bool             `url:"simple,omitempty" json:"simple,omitempty"`
	Sort                     *string           `url:"sort,omitempty" json:"sort,omitempty"`
	Starred                  *bool             `url:"starred,omitempty" json:"starred,omitempty"`
	Statistics               *bool             `url:"statistics,omitempty" json:"statistics,omitempty"`
	Topic                    *string           `url:"topic,omitempty" json:"topic,omitempty"`
	Visibility               *VisibilityValue  `url:"visibility,omitempty" json:"visibility,omitempty"`
	WikiChecksumFailed       *bool             `url:"wiki_checksum_failed,omitempty" json:"wiki_checksum_failed,omitempty"`
	WithCustomAttributes     *bool             `url:"with_custom_attributes,omitempty" json:"with_custom_attributes,omitempty"`
	WithIssuesEnabled        *bool             `url:"with_issues_enabled,omitempty" json:"with_issues_enabled,omitempty"`
	WithMergeRequestsEnabled *bool             `url:"with_merge_requests_enabled,omitempty" json:"with_merge_requests_enabled,omitempty"`
	WithProgrammingLanguage  *string           `url:"with_programming_language,omitempty" json:"with_programming_language,omitempty"`
}

// ListProjects gets a list of projects accessible by the authenticated user.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#list-all-projects
func (s *ProjectsService) ListProjects(opt *ListProjectsOptions, options ...RequestOptionFunc) ([]*Project, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "projects", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*Project
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// ListUserProjects gets a list of projects for the given user.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#list-a-users-projects
func (s *ProjectsService) ListUserProjects(uid any, opt *ListProjectsOptions, options ...RequestOptionFunc) ([]*Project, *Response, error) {
	user, err := parseID(uid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("users/%s/projects", user)

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*Project
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// ListUserContributedProjects gets a list of visible projects a given user has contributed to.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#list-projects-a-user-has-contributed-to
func (s *ProjectsService) ListUserContributedProjects(uid any, opt *ListProjectsOptions, options ...RequestOptionFunc) ([]*Project, *Response, error) {
	user, err := parseID(uid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("users/%s/contributed_projects", user)

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*Project
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// ListUserStarredProjects gets a list of projects starred by the given user.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_starring/#list-projects-starred-by-a-user
func (s *ProjectsService) ListUserStarredProjects(uid any, opt *ListProjectsOptions, options ...RequestOptionFunc) ([]*Project, *Response, error) {
	user, err := parseID(uid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("users/%s/starred_projects", user)

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*Project
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// ProjectUser represents a GitLab project user.
type ProjectUser struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

// ListProjectUserOptions represents the available ListProjectsUsers() options.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#list-users
type ListProjectUserOptions struct {
	ListOptions
	Search *string `url:"search,omitempty" json:"search,omitempty"`
}

// ListProjectsUsers gets a list of users for the given project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#list-users
func (s *ProjectsService) ListProjectsUsers(pid any, opt *ListProjectUserOptions, options ...RequestOptionFunc) ([]*ProjectUser, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/users", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*ProjectUser
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// ProjectGroup represents a GitLab project group.
// GitLab API docs: https://docs.gitlab.com/api/projects/#list-groups
type ProjectGroup struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
	FullName  string `json:"full_name"`
	FullPath  string `json:"full_path"`
}

// ListProjectGroupOptions represents the available ListProjectsGroups() options.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#list-groups
type ListProjectGroupOptions struct {
	ListOptions
	Search               *string           `url:"search,omitempty" json:"search,omitempty"`
	SharedMinAccessLevel *AccessLevelValue `url:"shared_min_access_level,omitempty" json:"shared_min_access_level,omitempty"`
	SharedVisiableOnly   *bool             `url:"shared_visible_only,omitempty" json:"shared_visible_only,omitempty"`
	SkipGroups           *[]int            `url:"skip_groups,omitempty" json:"skip_groups,omitempty"`
	WithShared           *bool             `url:"with_shared,omitempty" json:"with_shared,omitempty"`
}

// ListProjectsGroups gets a list of groups for the given project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#list-groups
func (s *ProjectsService) ListProjectsGroups(pid any, opt *ListProjectGroupOptions, options ...RequestOptionFunc) ([]*ProjectGroup, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/groups", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*ProjectGroup
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// ProjectLanguages is a map of strings because the response is arbitrary
//
// Gitlab API docs:
// https://docs.gitlab.com/api/projects/#list-programming-languages-used
type ProjectLanguages map[string]float32

// GetProjectLanguages gets a list of languages used by the project
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#list-programming-languages-used
func (s *ProjectsService) GetProjectLanguages(pid any, options ...RequestOptionFunc) (*ProjectLanguages, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/languages", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(ProjectLanguages)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// GetProjectOptions represents the available GetProject() options.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#get-a-single-project
type GetProjectOptions struct {
	License              *bool `url:"license,omitempty" json:"license,omitempty"`
	Statistics           *bool `url:"statistics,omitempty" json:"statistics,omitempty"`
	WithCustomAttributes *bool `url:"with_custom_attributes,omitempty" json:"with_custom_attributes,omitempty"`
}

// GetProject gets a specific project, identified by project ID or
// NAMESPACE/PROJECT_NAME, which is owned by the authenticated user.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#get-a-single-project
func (s *ProjectsService) GetProject(pid any, opt *GetProjectOptions, options ...RequestOptionFunc) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// CreateProjectOptions represents the available CreateProject() options.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#create-a-project
type CreateProjectOptions struct {
	AllowMergeOnSkippedPipeline               *bool                                `url:"allow_merge_on_skipped_pipeline,omitempty" json:"allow_merge_on_skipped_pipeline,omitempty"`
	OnlyAllowMergeIfAllStatusChecksPassed     *bool                                `url:"only_allow_merge_if_all_status_checks_passed,omitempty" json:"only_allow_merge_if_all_status_checks_passed,omitempty"`
	AnalyticsAccessLevel                      *AccessControlValue                  `url:"analytics_access_level,omitempty" json:"analytics_access_level,omitempty"`
	AutoCancelPendingPipelines                *string                              `url:"auto_cancel_pending_pipelines,omitempty" json:"auto_cancel_pending_pipelines,omitempty"`
	AutoDevopsDeployStrategy                  *string                              `url:"auto_devops_deploy_strategy,omitempty" json:"auto_devops_deploy_strategy,omitempty"`
	AutoDevopsEnabled                         *bool                                `url:"auto_devops_enabled,omitempty" json:"auto_devops_enabled,omitempty"`
	AutocloseReferencedIssues                 *bool                                `url:"autoclose_referenced_issues,omitempty" json:"autoclose_referenced_issues,omitempty"`
	Avatar                                    *ProjectAvatar                       `url:"-" json:"-"`
	BuildCoverageRegex                        *string                              `url:"build_coverage_regex,omitempty" json:"build_coverage_regex,omitempty"`
	BuildGitStrategy                          *string                              `url:"build_git_strategy,omitempty" json:"build_git_strategy,omitempty"`
	BuildTimeout                              *int                                 `url:"build_timeout,omitempty" json:"build_timeout,omitempty"`
	BuildsAccessLevel                         *AccessControlValue                  `url:"builds_access_level,omitempty" json:"builds_access_level,omitempty"`
	CIConfigPath                              *string                              `url:"ci_config_path,omitempty" json:"ci_config_path,omitempty"`
	ContainerExpirationPolicyAttributes       *ContainerExpirationPolicyAttributes `url:"container_expiration_policy_attributes,omitempty" json:"container_expiration_policy_attributes,omitempty"`
	ContainerRegistryAccessLevel              *AccessControlValue                  `url:"container_registry_access_level,omitempty" json:"container_registry_access_level,omitempty"`
	DefaultBranch                             *string                              `url:"default_branch,omitempty" json:"default_branch,omitempty"`
	Description                               *string                              `url:"description,omitempty" json:"description,omitempty"`
	EmailsEnabled                             *bool                                `url:"emails_enabled,omitempty" json:"emails_enabled,omitempty"`
	EnforceAuthChecksOnUploads                *bool                                `url:"enforce_auth_checks_on_uploads,omitempty" json:"enforce_auth_checks_on_uploads,omitempty"`
	ExternalAuthorizationClassificationLabel  *string                              `url:"external_authorization_classification_label,omitempty" json:"external_authorization_classification_label,omitempty"`
	ForkingAccessLevel                        *AccessControlValue                  `url:"forking_access_level,omitempty" json:"forking_access_level,omitempty"`
	GroupWithProjectTemplatesID               *int                                 `url:"group_with_project_templates_id,omitempty" json:"group_with_project_templates_id,omitempty"`
	ImportURL                                 *string                              `url:"import_url,omitempty" json:"import_url,omitempty"`
	InitializeWithReadme                      *bool                                `url:"initialize_with_readme,omitempty" json:"initialize_with_readme,omitempty"`
	IssuesAccessLevel                         *AccessControlValue                  `url:"issues_access_level,omitempty" json:"issues_access_level,omitempty"`
	IssueBranchTemplate                       *string                              `url:"issue_branch_template,omitempty" json:"issue_branch_template,omitempty"`
	LFSEnabled                                *bool                                `url:"lfs_enabled,omitempty" json:"lfs_enabled,omitempty"`
	MergeCommitTemplate                       *string                              `url:"merge_commit_template,omitempty" json:"merge_commit_template,omitempty"`
	MergeMethod                               *MergeMethodValue                    `url:"merge_method,omitempty" json:"merge_method,omitempty"`
	MergePipelinesEnabled                     *bool                                `url:"merge_pipelines_enabled,omitempty" json:"merge_pipelines_enabled,omitempty"`
	MergeRequestsAccessLevel                  *AccessControlValue                  `url:"merge_requests_access_level,omitempty" json:"merge_requests_access_level,omitempty"`
	MergeTrainsEnabled                        *bool                                `url:"merge_trains_enabled,omitempty" json:"merge_trains_enabled,omitempty"`
	MergeTrainsSkipTrainAllowed               *bool                                `url:"merge_trains_skip_train_allowed,omitempty" json:"merge_trains_skip_train_allowed,omitempty"`
	Mirror                                    *bool                                `url:"mirror,omitempty" json:"mirror,omitempty"`
	MirrorTriggerBuilds                       *bool                                `url:"mirror_trigger_builds,omitempty" json:"mirror_trigger_builds,omitempty"`
	ModelExperimentsAccessLevel               *AccessControlValue                  `url:"model_experiments_access_level,omitempty" json:"model_experiments_access_level,omitempty"`
	ModelRegistryAccessLevel                  *AccessControlValue                  `url:"model_registry_access_level,omitempty" json:"model_registry_access_level,omitempty"`
	Name                                      *string                              `url:"name,omitempty" json:"name,omitempty"`
	NamespaceID                               *int                                 `url:"namespace_id,omitempty" json:"namespace_id,omitempty"`
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool                                `url:"only_allow_merge_if_all_discussions_are_resolved,omitempty" json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	OnlyAllowMergeIfPipelineSucceeds          *bool                                `url:"only_allow_merge_if_pipeline_succeeds,omitempty" json:"only_allow_merge_if_pipeline_succeeds,omitempty"`
	OperationsAccessLevel                     *AccessControlValue                  `url:"operations_access_level,omitempty" json:"operations_access_level,omitempty"`
	PackagesEnabled                           *bool                                `url:"packages_enabled,omitempty" json:"packages_enabled,omitempty"`
	PagesAccessLevel                          *AccessControlValue                  `url:"pages_access_level,omitempty" json:"pages_access_level,omitempty"`
	Path                                      *string                              `url:"path,omitempty" json:"path,omitempty"`
	ReleasesAccessLevel                       *AccessControlValue                  `url:"releases_access_level,omitempty" json:"releases_access_level,omitempty"`
	EnvironmentsAccessLevel                   *AccessControlValue                  `url:"environments_access_level,omitempty" json:"environments_access_level,omitempty"`
	FeatureFlagsAccessLevel                   *AccessControlValue                  `url:"feature_flags_access_level,omitempty" json:"feature_flags_access_level,omitempty"`
	InfrastructureAccessLevel                 *AccessControlValue                  `url:"infrastructure_access_level,omitempty" json:"infrastructure_access_level,omitempty"`
	MonitorAccessLevel                        *AccessControlValue                  `url:"monitor_access_level,omitempty" json:"monitor_access_level,omitempty"`
	RemoveSourceBranchAfterMerge              *bool                                `url:"remove_source_branch_after_merge,omitempty" json:"remove_source_branch_after_merge,omitempty"`
	PrintingMergeRequestLinkEnabled           *bool                                `url:"printing_merge_request_link_enabled,omitempty" json:"printing_merge_request_link_enabled,omitempty"`
	RepositoryAccessLevel                     *AccessControlValue                  `url:"repository_access_level,omitempty" json:"repository_access_level,omitempty"`
	RepositoryStorage                         *string                              `url:"repository_storage,omitempty" json:"repository_storage,omitempty"`
	RequestAccessEnabled                      *bool                                `url:"request_access_enabled,omitempty" json:"request_access_enabled,omitempty"`
	RequirementsAccessLevel                   *AccessControlValue                  `url:"requirements_access_level,omitempty" json:"requirements_access_level,omitempty"`
	ResolveOutdatedDiffDiscussions            *bool                                `url:"resolve_outdated_diff_discussions,omitempty" json:"resolve_outdated_diff_discussions,omitempty"`
	SecurityAndComplianceAccessLevel          *AccessControlValue                  `url:"security_and_compliance_access_level,omitempty" json:"security_and_compliance_access_level,omitempty"`
	SharedRunnersEnabled                      *bool                                `url:"shared_runners_enabled,omitempty" json:"shared_runners_enabled,omitempty"`
	GroupRunnersEnabled                       *bool                                `url:"group_runners_enabled,omitempty" json:"group_runners_enabled,omitempty"`
	ShowDefaultAwardEmojis                    *bool                                `url:"show_default_award_emojis,omitempty" json:"show_default_award_emojis,omitempty"`
	SnippetsAccessLevel                       *AccessControlValue                  `url:"snippets_access_level,omitempty" json:"snippets_access_level,omitempty"`
	SquashCommitTemplate                      *string                              `url:"squash_commit_template,omitempty" json:"squash_commit_template,omitempty"`
	SquashOption                              *SquashOptionValue                   `url:"squash_option,omitempty" json:"squash_option,omitempty"`
	SuggestionCommitMessage                   *string                              `url:"suggestion_commit_message,omitempty" json:"suggestion_commit_message,omitempty"`
	TemplateName                              *string                              `url:"template_name,omitempty" json:"template_name,omitempty"`
	TemplateProjectID                         *int                                 `url:"template_project_id,omitempty" json:"template_project_id,omitempty"`
	Topics                                    *[]string                            `url:"topics,omitempty" json:"topics,omitempty"`
	UseCustomTemplate                         *bool                                `url:"use_custom_template,omitempty" json:"use_custom_template,omitempty"`
	Visibility                                *VisibilityValue                     `url:"visibility,omitempty" json:"visibility,omitempty"`
	WikiAccessLevel                           *AccessControlValue                  `url:"wiki_access_level,omitempty" json:"wiki_access_level,omitempty"`

	// Deprecated: use Merge Request Approvals API instead
	ApprovalsBeforeMerge *int `url:"approvals_before_merge,omitempty" json:"approvals_before_merge,omitempty"`
	// Deprecated: use PublicJobs instead
	PublicBuilds *bool `url:"public_builds,omitempty" json:"public_builds,omitempty"`
	// Deprecated: No longer supported in recent versions.
	CIForwardDeploymentEnabled *bool `url:"ci_forward_deployment_enabled,omitempty" json:"ci_forward_deployment_enabled,omitempty"`
	// Deprecated: Use ContainerRegistryAccessLevel instead.
	ContainerRegistryEnabled *bool `url:"container_registry_enabled,omitempty" json:"container_registry_enabled,omitempty"`
	// Deprecated: Use EmailsEnabled instead
	EmailsDisabled *bool `url:"emails_disabled,omitempty" json:"emails_disabled,omitempty"`
	// Deprecated: Use IssuesAccessLevel instead.
	IssuesEnabled *bool `url:"issues_enabled,omitempty" json:"issues_enabled,omitempty"`
	// Deprecated: No longer supported in recent versions.
	IssuesTemplate *string `url:"issues_template,omitempty" json:"issues_template,omitempty"`
	// Deprecated: Use BuildsAccessLevel instead.
	JobsEnabled *bool `url:"jobs_enabled,omitempty" json:"jobs_enabled,omitempty"`
	// Deprecated: Use MergeRequestsAccessLevel instead.
	MergeRequestsEnabled *bool `url:"merge_requests_enabled,omitempty" json:"merge_requests_enabled,omitempty"`
	// Deprecated: No longer supported in recent versions.
	MergeRequestsTemplate *string `url:"merge_requests_template,omitempty" json:"merge_requests_template,omitempty"`
	// Deprecated: No longer supported in recent versions.
	ServiceDeskEnabled *bool `url:"service_desk_enabled,omitempty" json:"service_desk_enabled,omitempty"`
	// Deprecated: Use SnippetsAccessLevel instead.
	SnippetsEnabled *bool `url:"snippets_enabled,omitempty" json:"snippets_enabled,omitempty"`
	// Deprecated: Use Topics instead. (Deprecated in GitLab 14.0)
	TagList *[]string `url:"tag_list,omitempty" json:"tag_list,omitempty"`
	// Deprecated: Use WikiAccessLevel instead.
	WikiEnabled *bool `url:"wiki_enabled,omitempty" json:"wiki_enabled,omitempty"`
}

// ContainerExpirationPolicyAttributes represents the available container
// expiration policy attributes.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#create-a-project
type ContainerExpirationPolicyAttributes struct {
	Cadence         *string `url:"cadence,omitempty" json:"cadence,omitempty"`
	KeepN           *int    `url:"keep_n,omitempty" json:"keep_n,omitempty"`
	OlderThan       *string `url:"older_than,omitempty" json:"older_than,omitempty"`
	NameRegexDelete *string `url:"name_regex_delete,omitempty" json:"name_regex_delete,omitempty"`
	NameRegexKeep   *string `url:"name_regex_keep,omitempty" json:"name_regex_keep,omitempty"`
	Enabled         *bool   `url:"enabled,omitempty" json:"enabled,omitempty"`

	// Deprecated: Is replaced by NameRegexDelete and is internally hardwired to its value.
	NameRegex *string `url:"name_regex,omitempty" json:"name_regex,omitempty"`
}

// ProjectAvatar represents a GitLab project avatar.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#create-a-project
type ProjectAvatar struct {
	Filename string
	Image    io.Reader
}

// MarshalJSON implements the json.Marshaler interface.
func (a *ProjectAvatar) MarshalJSON() ([]byte, error) {
	if a.Filename == "" && a.Image == nil {
		return []byte(`""`), nil
	}
	type alias ProjectAvatar
	return json.Marshal((*alias)(a))
}

// CreateProject creates a new project owned by the authenticated user.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#create-a-project
func (s *ProjectsService) CreateProject(opt *CreateProjectOptions, options ...RequestOptionFunc) (*Project, *Response, error) {
	if opt.ContainerExpirationPolicyAttributes != nil {
		// This is needed to satisfy the API. Should be deleted
		// when NameRegex is removed (it's now deprecated).
		opt.ContainerExpirationPolicyAttributes.NameRegex = opt.ContainerExpirationPolicyAttributes.NameRegexDelete
	}

	var err error
	var req *retryablehttp.Request

	if opt.Avatar == nil {
		req, err = s.client.NewRequest(http.MethodPost, "projects", opt, options)
	} else {
		req, err = s.client.UploadRequest(
			http.MethodPost,
			"projects",
			opt.Avatar.Image,
			opt.Avatar.Filename,
			UploadAvatar,
			opt,
			options,
		)
	}
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// CreateProjectForUserOptions represents the available CreateProjectForUser()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#create-a-project-for-a-user
type CreateProjectForUserOptions CreateProjectOptions

// CreateProjectForUser creates a new project owned by the specified user.
// Available only for admins.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#create-a-project-for-a-user
func (s *ProjectsService) CreateProjectForUser(user int, opt *CreateProjectForUserOptions, options ...RequestOptionFunc) (*Project, *Response, error) {
	if opt.ContainerExpirationPolicyAttributes != nil {
		// This is needed to satisfy the API. Should be deleted
		// when NameRegex is removed (it's now deprecated).
		opt.ContainerExpirationPolicyAttributes.NameRegex = opt.ContainerExpirationPolicyAttributes.NameRegexDelete
	}

	var err error
	var req *retryablehttp.Request
	u := fmt.Sprintf("projects/user/%d", user)

	if opt.Avatar == nil {
		req, err = s.client.NewRequest(http.MethodPost, u, opt, options)
	} else {
		req, err = s.client.UploadRequest(
			http.MethodPost,
			u,
			opt.Avatar.Image,
			opt.Avatar.Filename,
			UploadAvatar,
			opt,
			options,
		)
	}
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// EditProjectOptions represents the available EditProject() options.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#edit-a-project
type EditProjectOptions struct {
	AllowMergeOnSkippedPipeline               *bool                                        `url:"allow_merge_on_skipped_pipeline,omitempty" json:"allow_merge_on_skipped_pipeline,omitempty"`
	AllowPipelineTriggerApproveDeployment     *bool                                        `url:"allow_pipeline_trigger_approve_deployment,omitempty" json:"allow_pipeline_trigger_approve_deployment,omitempty"`
	OnlyAllowMergeIfAllStatusChecksPassed     *bool                                        `url:"only_allow_merge_if_all_status_checks_passed,omitempty" json:"only_allow_merge_if_all_status_checks_passed,omitempty"`
	AnalyticsAccessLevel                      *AccessControlValue                          `url:"analytics_access_level,omitempty" json:"analytics_access_level,omitempty"`
	AutoCancelPendingPipelines                *string                                      `url:"auto_cancel_pending_pipelines,omitempty" json:"auto_cancel_pending_pipelines,omitempty"`
	AutoDevopsDeployStrategy                  *string                                      `url:"auto_devops_deploy_strategy,omitempty" json:"auto_devops_deploy_strategy,omitempty"`
	AutoDevopsEnabled                         *bool                                        `url:"auto_devops_enabled,omitempty" json:"auto_devops_enabled,omitempty"`
	AutoDuoCodeReviewEnabled                  *bool                                        `url:"auto_duo_code_review_enabled,omitempty" json:"auto_duo_code_review_enabled,omitempty"`
	AutocloseReferencedIssues                 *bool                                        `url:"autoclose_referenced_issues,omitempty" json:"autoclose_referenced_issues,omitempty"`
	Avatar                                    *ProjectAvatar                               `url:"-" json:"avatar,omitempty"`
	BuildCoverageRegex                        *string                                      `url:"build_coverage_regex,omitempty" json:"build_coverage_regex,omitempty"`
	BuildGitStrategy                          *string                                      `url:"build_git_strategy,omitempty" json:"build_git_strategy,omitempty"`
	BuildTimeout                              *int                                         `url:"build_timeout,omitempty" json:"build_timeout,omitempty"`
	BuildsAccessLevel                         *AccessControlValue                          `url:"builds_access_level,omitempty" json:"builds_access_level,omitempty"`
	CIConfigPath                              *string                                      `url:"ci_config_path,omitempty" json:"ci_config_path,omitempty"`
	CIDefaultGitDepth                         *int                                         `url:"ci_default_git_depth,omitempty" json:"ci_default_git_depth,omitempty"`
	CIDeletePipelinesInSeconds                *int                                         `url:"ci_delete_pipelines_in_seconds,omitempty" json:"ci_delete_pipelines_in_seconds,omitempty"`
	CIForwardDeploymentEnabled                *bool                                        `url:"ci_forward_deployment_enabled,omitempty" json:"ci_forward_deployment_enabled,omitempty"`
	CIForwardDeploymentRollbackAllowed        *bool                                        `url:"ci_forward_deployment_rollback_allowed,omitempty" json:"ci_forward_deployment_rollback_allowed,omitempty"`
	CIPushRepositoryForJobTokenAllowed        *bool                                        `url:"ci_push_repository_for_job_token_allowed,omitempty" json:"ci_push_repository_for_job_token_allowed,omitempty"`
	CIIdTokenSubClaimComponents               *[]string                                    `url:"ci_id_token_sub_claim_components,omitempty" json:"ci_id_token_sub_claim_components,omitempty"`
	CISeperateCache                           *bool                                        `url:"ci_separated_caches,omitempty" json:"ci_separated_caches,omitempty"`
	CIRestrictPipelineCancellationRole        *AccessControlValue                          `url:"ci_restrict_pipeline_cancellation_role,omitempty" json:"ci_restrict_pipeline_cancellation_role,omitempty"`
	CIPipelineVariablesMinimumOverrideRole    *CIPipelineVariablesMinimumOverrideRoleValue `url:"ci_pipeline_variables_minimum_override_role,omitempty" json:"ci_pipeline_variables_minimum_override_role,omitempty"`
	ContainerExpirationPolicyAttributes       *ContainerExpirationPolicyAttributes         `url:"container_expiration_policy_attributes,omitempty" json:"container_expiration_policy_attributes,omitempty"`
	ContainerRegistryAccessLevel              *AccessControlValue                          `url:"container_registry_access_level,omitempty" json:"container_registry_access_level,omitempty"`
	DefaultBranch                             *string                                      `url:"default_branch,omitempty" json:"default_branch,omitempty"`
	Description                               *string                                      `url:"description,omitempty" json:"description,omitempty"`
	EmailsEnabled                             *bool                                        `url:"emails_enabled,omitempty" json:"emails_enabled,omitempty"`
	EnforceAuthChecksOnUploads                *bool                                        `url:"enforce_auth_checks_on_uploads,omitempty" json:"enforce_auth_checks_on_uploads,omitempty"`
	ExternalAuthorizationClassificationLabel  *string                                      `url:"external_authorization_classification_label,omitempty" json:"external_authorization_classification_label,omitempty"`
	ForkingAccessLevel                        *AccessControlValue                          `url:"forking_access_level,omitempty" json:"forking_access_level,omitempty"`
	ImportURL                                 *string                                      `url:"import_url,omitempty" json:"import_url,omitempty"`
	IssuesAccessLevel                         *AccessControlValue                          `url:"issues_access_level,omitempty" json:"issues_access_level,omitempty"`
	IssueBranchTemplate                       *string                                      `url:"issue_branch_template,omitempty" json:"issue_branch_template,omitempty"`
	IssuesTemplate                            *string                                      `url:"issues_template,omitempty" json:"issues_template,omitempty"`
	KeepLatestArtifact                        *bool                                        `url:"keep_latest_artifact,omitempty" json:"keep_latest_artifact,omitempty"`
	LFSEnabled                                *bool                                        `url:"lfs_enabled,omitempty" json:"lfs_enabled,omitempty"`
	MergeCommitTemplate                       *string                                      `url:"merge_commit_template,omitempty" json:"merge_commit_template,omitempty"`
	MergeRequestDefaultTargetSelf             *bool                                        `url:"mr_default_target_self,omitempty" json:"mr_default_target_self,omitempty"`
	MergeMethod                               *MergeMethodValue                            `url:"merge_method,omitempty" json:"merge_method,omitempty"`
	MergePipelinesEnabled                     *bool                                        `url:"merge_pipelines_enabled,omitempty" json:"merge_pipelines_enabled,omitempty"`
	MergeRequestsAccessLevel                  *AccessControlValue                          `url:"merge_requests_access_level,omitempty" json:"merge_requests_access_level,omitempty"`
	MergeRequestsTemplate                     *string                                      `url:"merge_requests_template,omitempty" json:"merge_requests_template,omitempty"`
	MergeTrainsEnabled                        *bool                                        `url:"merge_trains_enabled,omitempty" json:"merge_trains_enabled,omitempty"`
	MergeTrainsSkipTrainAllowed               *bool                                        `url:"merge_trains_skip_train_allowed,omitempty" json:"merge_trains_skip_train_allowed,omitempty"`
	Mirror                                    *bool                                        `url:"mirror,omitempty" json:"mirror,omitempty"`
	MirrorBranchRegex                         *string                                      `url:"mirror_branch_regex,omitempty" json:"mirror_branch_regex,omitempty"`
	MirrorOverwritesDivergedBranches          *bool                                        `url:"mirror_overwrites_diverged_branches,omitempty" json:"mirror_overwrites_diverged_branches,omitempty"`
	MirrorTriggerBuilds                       *bool                                        `url:"mirror_trigger_builds,omitempty" json:"mirror_trigger_builds,omitempty"`
	MirrorUserID                              *int                                         `url:"mirror_user_id,omitempty" json:"mirror_user_id,omitempty"`
	ModelExperimentsAccessLevel               *AccessControlValue                          `url:"model_experiments_access_level,omitempty" json:"model_experiments_access_level,omitempty"`
	ModelRegistryAccessLevel                  *AccessControlValue                          `url:"model_registry_access_level,omitempty" json:"model_registry_access_level,omitempty"`
	Name                                      *string                                      `url:"name,omitempty" json:"name,omitempty"`
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool                                        `url:"only_allow_merge_if_all_discussions_are_resolved,omitempty" json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	OnlyAllowMergeIfPipelineSucceeds          *bool                                        `url:"only_allow_merge_if_pipeline_succeeds,omitempty" json:"only_allow_merge_if_pipeline_succeeds,omitempty"`
	OnlyMirrorProtectedBranches               *bool                                        `url:"only_mirror_protected_branches,omitempty" json:"only_mirror_protected_branches,omitempty"`
	OperationsAccessLevel                     *AccessControlValue                          `url:"operations_access_level,omitempty" json:"operations_access_level,omitempty"`
	PackagesEnabled                           *bool                                        `url:"packages_enabled,omitempty" json:"packages_enabled,omitempty"`
	PagesAccessLevel                          *AccessControlValue                          `url:"pages_access_level,omitempty" json:"pages_access_level,omitempty"`
	Path                                      *string                                      `url:"path,omitempty" json:"path,omitempty"`
	PublicJobs                                *bool                                        `url:"public_jobs,omitempty" json:"public_jobs,omitempty"`
	ReleasesAccessLevel                       *AccessControlValue                          `url:"releases_access_level,omitempty" json:"releases_access_level,omitempty"`
	EnvironmentsAccessLevel                   *AccessControlValue                          `url:"environments_access_level,omitempty" json:"environments_access_level,omitempty"`
	FeatureFlagsAccessLevel                   *AccessControlValue                          `url:"feature_flags_access_level,omitempty" json:"feature_flags_access_level,omitempty"`
	InfrastructureAccessLevel                 *AccessControlValue                          `url:"infrastructure_access_level,omitempty" json:"infrastructure_access_level,omitempty"`
	MonitorAccessLevel                        *AccessControlValue                          `url:"monitor_access_level,omitempty" json:"monitor_access_level,omitempty"`
	RemoveSourceBranchAfterMerge              *bool                                        `url:"remove_source_branch_after_merge,omitempty" json:"remove_source_branch_after_merge,omitempty"`
	PreventMergeWithoutJiraIssue              *bool                                        `url:"prevent_merge_without_jira_issue,omitempty" json:"prevent_merge_without_jira_issue,omitempty"`
	PrintingMergeRequestLinkEnabled           *bool                                        `url:"printing_merge_request_link_enabled,omitempty" json:"printing_merge_request_link_enabled,omitempty"`
	RepositoryAccessLevel                     *AccessControlValue                          `url:"repository_access_level,omitempty" json:"repository_access_level,omitempty"`
	RepositoryStorage                         *string                                      `url:"repository_storage,omitempty" json:"repository_storage,omitempty"`
	RequestAccessEnabled                      *bool                                        `url:"request_access_enabled,omitempty" json:"request_access_enabled,omitempty"`
	RequirementsAccessLevel                   *AccessControlValue                          `url:"requirements_access_level,omitempty" json:"requirements_access_level,omitempty"`
	ResolveOutdatedDiffDiscussions            *bool                                        `url:"resolve_outdated_diff_discussions,omitempty" json:"resolve_outdated_diff_discussions,omitempty"`
	SecurityAndComplianceAccessLevel          *AccessControlValue                          `url:"security_and_compliance_access_level,omitempty" json:"security_and_compliance_access_level,omitempty"`
	ServiceDeskEnabled                        *bool                                        `url:"service_desk_enabled,omitempty" json:"service_desk_enabled,omitempty"`
	SharedRunnersEnabled                      *bool                                        `url:"shared_runners_enabled,omitempty" json:"shared_runners_enabled,omitempty"`
	GroupRunnersEnabled                       *bool                                        `url:"group_runners_enabled,omitempty" json:"group_runners_enabled,omitempty"`
	ShowDefaultAwardEmojis                    *bool                                        `url:"show_default_award_emojis,omitempty" json:"show_default_award_emojis,omitempty"`
	SnippetsAccessLevel                       *AccessControlValue                          `url:"snippets_access_level,omitempty" json:"snippets_access_level,omitempty"`
	SquashCommitTemplate                      *string                                      `url:"squash_commit_template,omitempty" json:"squash_commit_template,omitempty"`
	SquashOption                              *SquashOptionValue                           `url:"squash_option,omitempty" json:"squash_option,omitempty"`
	SuggestionCommitMessage                   *string                                      `url:"suggestion_commit_message,omitempty" json:"suggestion_commit_message,omitempty"`
	Topics                                    *[]string                                    `url:"topics,omitempty" json:"topics,omitempty"`
	Visibility                                *VisibilityValue                             `url:"visibility,omitempty" json:"visibility,omitempty"`
	WikiAccessLevel                           *AccessControlValue                          `url:"wiki_access_level,omitempty" json:"wiki_access_level,omitempty"`

	// Deprecated: use Merge Request Approvals API instead
	ApprovalsBeforeMerge *int `url:"approvals_before_merge,omitempty" json:"approvals_before_merge,omitempty"`
	// Deprecated: use PublicJobs instead
	PublicBuilds *bool `url:"public_builds,omitempty" json:"public_builds,omitempty"`
	// Deprecated: use CIPipelineVariablesMinimumOverrideRole instead
	RestrictUserDefinedVariables *bool `url:"restrict_user_defined_variables,omitempty" json:"restrict_user_defined_variables,omitempty"`
	// Deprecated: Use ContainerRegistryAccessLevel instead.
	ContainerRegistryEnabled *bool `url:"container_registry_enabled,omitempty" json:"container_registry_enabled,omitempty"`
	// Deprecated: Use EmailsEnabled instead
	EmailsDisabled *bool `url:"emails_disabled,omitempty" json:"emails_disabled,omitempty"`
	// Deprecated: Use IssuesAccessLevel instead.
	IssuesEnabled *bool `url:"issues_enabled,omitempty" json:"issues_enabled,omitempty"`
	// Deprecated: Use BuildsAccessLevel instead.
	JobsEnabled *bool `url:"jobs_enabled,omitempty" json:"jobs_enabled,omitempty"`
	// Deprecated: Use MergeRequestsAccessLevel instead.
	MergeRequestsEnabled *bool `url:"merge_requests_enabled,omitempty" json:"merge_requests_enabled,omitempty"`
	// Deprecated: Use SnippetsAccessLevel instead.
	SnippetsEnabled *bool `url:"snippets_enabled,omitempty" json:"snippets_enabled,omitempty"`
	// Deprecated: Use Topics instead. (Deprecated in GitLab 14.0)
	TagList *[]string `url:"tag_list,omitempty" json:"tag_list,omitempty"`
	// Deprecated: Use WikiAccessLevel instead.
	WikiEnabled *bool `url:"wiki_enabled,omitempty" json:"wiki_enabled,omitempty"`
}

// EditProject updates an existing project.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#edit-a-project
func (s *ProjectsService) EditProject(pid any, opt *EditProjectOptions, options ...RequestOptionFunc) (*Project, *Response, error) {
	if opt.ContainerExpirationPolicyAttributes != nil {
		// This is needed to satisfy the API. Should be deleted
		// when NameRegex is removed (it's now deprecated).
		opt.ContainerExpirationPolicyAttributes.NameRegex = opt.ContainerExpirationPolicyAttributes.NameRegexDelete
	}

	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s", PathEscape(project))

	var req *retryablehttp.Request

	if opt.Avatar == nil || (opt.Avatar.Filename == "" && opt.Avatar.Image == nil) {
		req, err = s.client.NewRequest(http.MethodPut, u, opt, options)
	} else {
		req, err = s.client.UploadRequest(
			http.MethodPut,
			u,
			opt.Avatar.Image,
			opt.Avatar.Filename,
			UploadAvatar,
			opt,
			options,
		)
	}
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// ForkProjectOptions represents the available ForkProject() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_forks/#fork-a-project
type ForkProjectOptions struct {
	Branches                      *string          `url:"branches,omitempty" json:"branches,omitempty"`
	Description                   *string          `url:"description,omitempty" json:"description,omitempty"`
	MergeRequestDefaultTargetSelf *bool            `url:"mr_default_target_self,omitempty" json:"mr_default_target_self,omitempty"`
	Name                          *string          `url:"name,omitempty" json:"name,omitempty"`
	NamespaceID                   *int             `url:"namespace_id,omitempty" json:"namespace_id,omitempty"`
	NamespacePath                 *string          `url:"namespace_path,omitempty" json:"namespace_path,omitempty"`
	Path                          *string          `url:"path,omitempty" json:"path,omitempty"`
	Visibility                    *VisibilityValue `url:"visibility,omitempty" json:"visibility,omitempty"`

	// Deprecated: This parameter has been split into NamespaceID and NamespacePath.
	Namespace *string `url:"namespace,omitempty" json:"namespace,omitempty"`
}

// ForkProject forks a project into the user namespace of the authenticated
// user.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_forks/#fork-a-project
func (s *ProjectsService) ForkProject(pid any, opt *ForkProjectOptions, options ...RequestOptionFunc) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/fork", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// StarProject stars a given the project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_starring/#star-a-project
func (s *ProjectsService) StarProject(pid any, options ...RequestOptionFunc) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/star", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// ListProjectInvitedGroupOptions represents the available
// ListProjectsInvitedGroups() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#list-a-projects-invited-groups
type ListProjectInvitedGroupOptions struct {
	ListOptions
	Search               *string           `url:"search,omitempty" json:"search,omitempty"`
	MinAccessLevel       *AccessLevelValue `url:"min_access_level,omitempty" json:"min_access_level,omitempty"`
	Relation             *[]string         `url:"relation,omitempty" json:"relation,omitempty"`
	WithCustomAttributes *bool             `url:"with_custom_attributes,omitempty" json:"with_custom_attributes,omitempty"`
}

// ListProjectInvidedGroupOptions is kept for backwards compatibility.
//
// Deprecated: use ListProjectInvitedGroupOptions instead. The ListProjectInvidedGroupOptions type will be removed in the next release.
type ListProjectInvidedGroupOptions = ListProjectInvitedGroupOptions

// ListProjectsInvitedGroups lists invited groups of a project
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#list-a-projects-invited-groups
func (s *ProjectsService) ListProjectsInvitedGroups(pid any, opt *ListProjectInvidedGroupOptions, options ...RequestOptionFunc) ([]*ProjectGroup, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/invited_groups", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var pg []*ProjectGroup
	resp, err := s.client.Do(req, &pg)
	if err != nil {
		return nil, resp, err
	}

	return pg, resp, nil
}

// UnstarProject unstars a given project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_starring/#unstar-a-project
func (s *ProjectsService) UnstarProject(pid any, options ...RequestOptionFunc) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/unstar", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// ArchiveProject archives the project if the user is either admin or the
// project owner of this project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#archive-a-project
func (s *ProjectsService) ArchiveProject(pid any, options ...RequestOptionFunc) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/archive", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// UnarchiveProject unarchives the project if the user is either admin or
// the project owner of this project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#unarchive-a-project
func (s *ProjectsService) UnarchiveProject(pid any, options ...RequestOptionFunc) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/unarchive", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// RestoreProject restores a project that is marked for deletion.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#restore-a-project-marked-for-deletion
func (s *ProjectsService) RestoreProject(pid any, options ...RequestOptionFunc) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/restore", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// DeleteProjectOptions represents the available DeleteProject() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#delete-a-project
type DeleteProjectOptions struct {
	FullPath          *string `url:"full_path" json:"full_path"`
	PermanentlyRemove *bool   `url:"permanently_remove" json:"permanently_remove"`
}

// DeleteProject removes a project including all associated resources
// (issues, merge requests etc.)
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#delete-a-project
func (s *ProjectsService) DeleteProject(pid any, opt *DeleteProjectOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ShareWithGroupOptions represents the available SharedWithGroup() options.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#share-a-project-with-a-group
type ShareWithGroupOptions struct {
	ExpiresAt   *string           `url:"expires_at" json:"expires_at"`
	GroupAccess *AccessLevelValue `url:"group_access" json:"group_access"`
	GroupID     *int              `url:"group_id" json:"group_id"`
}

// ShareProjectWithGroup allows to share a project with a group.
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#share-a-project-with-a-group
func (s *ProjectsService) ShareProjectWithGroup(pid any, opt *ShareWithGroupOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/share", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DeleteSharedProjectFromGroup allows to unshare a project from a group.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#delete-a-shared-project-link-in-a-group
func (s *ProjectsService) DeleteSharedProjectFromGroup(pid any, groupID int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/share/%d", PathEscape(project), groupID)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// HookCustomHeader represents a project or group hook custom header
// Note: "Key" is returned from the Get operation, but "Value" is not
// The List operation doesn't return any headers at all for Projects,
// but does return headers for Groups
type HookCustomHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// HookURLVariable represents a project or group hook URL variable
type HookURLVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ProjectHook represents a project hook.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#list-webhooks-for-a-project
type ProjectHook struct {
	ID                        int                 `json:"id"`
	URL                       string              `json:"url"`
	Name                      string              `json:"name"`
	Description               string              `json:"description"`
	ProjectID                 int                 `json:"project_id"`
	PushEvents                bool                `json:"push_events"`
	PushEventsBranchFilter    string              `json:"push_events_branch_filter"`
	IssuesEvents              bool                `json:"issues_events"`
	ConfidentialIssuesEvents  bool                `json:"confidential_issues_events"`
	MergeRequestsEvents       bool                `json:"merge_requests_events"`
	TagPushEvents             bool                `json:"tag_push_events"`
	NoteEvents                bool                `json:"note_events"`
	ConfidentialNoteEvents    bool                `json:"confidential_note_events"`
	JobEvents                 bool                `json:"job_events"`
	PipelineEvents            bool                `json:"pipeline_events"`
	WikiPageEvents            bool                `json:"wiki_page_events"`
	DeploymentEvents          bool                `json:"deployment_events"`
	ReleasesEvents            bool                `json:"releases_events"`
	MilestoneEvents           bool                `json:"milestone_events"`
	FeatureFlagEvents         bool                `json:"feature_flag_events"`
	EnableSSLVerification     bool                `json:"enable_ssl_verification"`
	RepositoryUpdateEvents    bool                `json:"repository_update_events"`
	AlertStatus               string              `json:"alert_status"`
	DisabledUntil             *time.Time          `json:"disabled_until"`
	URLVariables              []HookURLVariable   `json:"url_variables"`
	CreatedAt                 *time.Time          `json:"created_at"`
	ResourceAccessTokenEvents bool                `json:"resource_access_token_events"`
	CustomWebhookTemplate     string              `json:"custom_webhook_template"`
	CustomHeaders             []*HookCustomHeader `json:"custom_headers"`
}

// ListProjectHooksOptions represents the available ListProjectHooks() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#list-webhooks-for-a-project
type ListProjectHooksOptions ListOptions

// ListProjectHooks gets a list of project hooks.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#list-webhooks-for-a-project
func (s *ProjectsService) ListProjectHooks(pid any, opt *ListProjectHooksOptions, options ...RequestOptionFunc) ([]*ProjectHook, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/hooks", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var ph []*ProjectHook
	resp, err := s.client.Do(req, &ph)
	if err != nil {
		return nil, resp, err
	}

	return ph, resp, nil
}

// GetProjectHook gets a specific hook for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#get-a-project-webhook
func (s *ProjectsService) GetProjectHook(pid any, hook int, options ...RequestOptionFunc) (*ProjectHook, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/hooks/%d", PathEscape(project), hook)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	ph := new(ProjectHook)
	resp, err := s.client.Do(req, ph)
	if err != nil {
		return nil, resp, err
	}

	return ph, resp, nil
}

// AddProjectHookOptions represents the available AddProjectHook() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#add-a-webhook-to-a-project
type AddProjectHookOptions struct {
	Name                      *string              `url:"name,omitempty" json:"name,omitempty"`
	Description               *string              `url:"description,omitempty" json:"description,omitempty"`
	ConfidentialIssuesEvents  *bool                `url:"confidential_issues_events,omitempty" json:"confidential_issues_events,omitempty"`
	ConfidentialNoteEvents    *bool                `url:"confidential_note_events,omitempty" json:"confidential_note_events,omitempty"`
	DeploymentEvents          *bool                `url:"deployment_events,omitempty" json:"deployment_events,omitempty"`
	EnableSSLVerification     *bool                `url:"enable_ssl_verification,omitempty" json:"enable_ssl_verification,omitempty"`
	IssuesEvents              *bool                `url:"issues_events,omitempty" json:"issues_events,omitempty"`
	JobEvents                 *bool                `url:"job_events,omitempty" json:"job_events,omitempty"`
	MergeRequestsEvents       *bool                `url:"merge_requests_events,omitempty" json:"merge_requests_events,omitempty"`
	NoteEvents                *bool                `url:"note_events,omitempty" json:"note_events,omitempty"`
	PipelineEvents            *bool                `url:"pipeline_events,omitempty" json:"pipeline_events,omitempty"`
	PushEvents                *bool                `url:"push_events,omitempty" json:"push_events,omitempty"`
	PushEventsBranchFilter    *string              `url:"push_events_branch_filter,omitempty" json:"push_events_branch_filter,omitempty"`
	ReleasesEvents            *bool                `url:"releases_events,omitempty" json:"releases_events,omitempty"`
	TagPushEvents             *bool                `url:"tag_push_events,omitempty" json:"tag_push_events,omitempty"`
	Token                     *string              `url:"token,omitempty" json:"token,omitempty"`
	URL                       *string              `url:"url,omitempty" json:"url,omitempty"`
	WikiPageEvents            *bool                `url:"wiki_page_events,omitempty" json:"wiki_page_events,omitempty"`
	ResourceAccessTokenEvents *bool                `url:"resource_access_token_events,omitempty" json:"resource_access_token_events,omitempty"`
	CustomWebhookTemplate     *string              `url:"custom_webhook_template,omitempty" json:"custom_webhook_template,omitempty"`
	CustomHeaders             *[]*HookCustomHeader `url:"custom_headers,omitempty" json:"custom_headers,omitempty"`
}

// AddProjectHook adds a hook to a specified project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#add-a-webhook-to-a-project
func (s *ProjectsService) AddProjectHook(pid any, opt *AddProjectHookOptions, options ...RequestOptionFunc) (*ProjectHook, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/hooks", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ph := new(ProjectHook)
	resp, err := s.client.Do(req, ph)
	if err != nil {
		return nil, resp, err
	}

	return ph, resp, nil
}

// EditProjectHookOptions represents the available EditProjectHook() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#edit-a-project-webhook
type EditProjectHookOptions struct {
	Name                      *string              `url:"name,omitempty" json:"name,omitempty"`
	Description               *string              `url:"description,omitempty" json:"description,omitempty"`
	ConfidentialIssuesEvents  *bool                `url:"confidential_issues_events,omitempty" json:"confidential_issues_events,omitempty"`
	ConfidentialNoteEvents    *bool                `url:"confidential_note_events,omitempty" json:"confidential_note_events,omitempty"`
	DeploymentEvents          *bool                `url:"deployment_events,omitempty" json:"deployment_events,omitempty"`
	EnableSSLVerification     *bool                `url:"enable_ssl_verification,omitempty" json:"enable_ssl_verification,omitempty"`
	IssuesEvents              *bool                `url:"issues_events,omitempty" json:"issues_events,omitempty"`
	JobEvents                 *bool                `url:"job_events,omitempty" json:"job_events,omitempty"`
	MergeRequestsEvents       *bool                `url:"merge_requests_events,omitempty" json:"merge_requests_events,omitempty"`
	NoteEvents                *bool                `url:"note_events,omitempty" json:"note_events,omitempty"`
	PipelineEvents            *bool                `url:"pipeline_events,omitempty" json:"pipeline_events,omitempty"`
	PushEvents                *bool                `url:"push_events,omitempty" json:"push_events,omitempty"`
	PushEventsBranchFilter    *string              `url:"push_events_branch_filter,omitempty" json:"push_events_branch_filter,omitempty"`
	ReleasesEvents            *bool                `url:"releases_events,omitempty" json:"releases_events,omitempty"`
	TagPushEvents             *bool                `url:"tag_push_events,omitempty" json:"tag_push_events,omitempty"`
	Token                     *string              `url:"token,omitempty" json:"token,omitempty"`
	URL                       *string              `url:"url,omitempty" json:"url,omitempty"`
	WikiPageEvents            *bool                `url:"wiki_page_events,omitempty" json:"wiki_page_events,omitempty"`
	ResourceAccessTokenEvents *bool                `url:"resource_access_token_events,omitempty" json:"resource_access_token_events,omitempty"`
	CustomWebhookTemplate     *string              `url:"custom_webhook_template,omitempty" json:"custom_webhook_template,omitempty"`
	CustomHeaders             *[]*HookCustomHeader `url:"custom_headers,omitempty" json:"custom_headers,omitempty"`
}

// EditProjectHook edits a hook for a specified project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#edit-a-project-webhook
func (s *ProjectsService) EditProjectHook(pid any, hook int, opt *EditProjectHookOptions, options ...RequestOptionFunc) (*ProjectHook, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/hooks/%d", PathEscape(project), hook)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ph := new(ProjectHook)
	resp, err := s.client.Do(req, ph)
	if err != nil {
		return nil, resp, err
	}

	return ph, resp, nil
}

// DeleteProjectHook removes a hook from a project. This is an idempotent
// method and can be called multiple times. Either the hook is available or not.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#delete-project-webhook
func (s *ProjectsService) DeleteProjectHook(pid any, hook int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/hooks/%d", PathEscape(project), hook)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// TriggerTestProjectHook Trigger a test hook for a specified project.
//
// In GitLab 17.0 and later, this endpoint has a special rate limit.
// In GitLab 17.0 the rate was three requests per minute for each project hook.
// In GitLab 17.1 this was changed to five requests per minute for each project
// and authenticated user.
//
// To disable this limit on self-managed GitLab and GitLab Dedicated,
// an administrator can disable the feature flag named web_hook_test_api_endpoint_rate_limit.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#trigger-a-test-project-webhook
func (s *ProjectsService) TriggerTestProjectHook(pid any, hook int, event ProjectHookEvent, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/hooks/%d/test/%s", PathEscape(project), hook, string(event))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// SetHookCustomHeaderOptions represents the available SetProjectCustomHeader()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#set-a-custom-header
type SetHookCustomHeaderOptions struct {
	Value *string `json:"value,omitempty"`
}

// SetProjectCustomHeader creates or updates a project custom webhook header.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#set-a-custom-header
func (s *ProjectsService) SetProjectCustomHeader(pid any, hook int, key string, opt *SetHookCustomHeaderOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/hooks/%d/custom_headers/%s", PathEscape(project), hook, key)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DeleteProjectCustomHeader deletes a project custom webhook header.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#delete-a-custom-header
func (s *ProjectsService) DeleteProjectCustomHeader(pid any, hook int, key string, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/hooks/%d/custom_headers/%s", PathEscape(project), hook, key)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// SetProjectWebhookURLVariableOptions represents the available
// SetProjectWebhookURLVariable() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#set-a-url-variable
type SetProjectWebhookURLVariableOptions struct {
	Value *string `json:"value,omitempty"`
}

// SetProjectWebhookURLVariable creates or updates a project webhook URL variable.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#set-a-url-variable
func (s *ProjectsService) SetProjectWebhookURLVariable(pid any, hook int, key string, opt *SetProjectWebhookURLVariableOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/hooks/%d/url_variables/%s", PathEscape(project), hook, PathEscape(key))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DeleteProjectWebhookURLVariable deletes a project webhook URL variable.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_webhooks/#delete-a-url-variable
func (s *ProjectsService) DeleteProjectWebhookURLVariable(pid any, hook int, key string, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/hooks/%d/url_variables/%s", PathEscape(project), hook, PathEscape(key))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ProjectForkRelation represents a project fork relationship.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_forks/#create-a-fork-relationship-between-projects
type ProjectForkRelation struct {
	ID                  int        `json:"id"`
	ForkedToProjectID   int        `json:"forked_to_project_id"`
	ForkedFromProjectID int        `json:"forked_from_project_id"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

// CreateProjectForkRelation creates a forked from/to relation between
// existing projects.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_forks/#create-a-fork-relationship-between-projects
func (s *ProjectsService) CreateProjectForkRelation(pid any, fork int, options ...RequestOptionFunc) (*ProjectForkRelation, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/fork/%d", PathEscape(project), fork)

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	pfr := new(ProjectForkRelation)
	resp, err := s.client.Do(req, pfr)
	if err != nil {
		return nil, resp, err
	}

	return pfr, resp, nil
}

// DeleteProjectForkRelation deletes an existing forked from relationship.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_forks/#delete-a-fork-relationship-between-projects
func (s *ProjectsService) DeleteProjectForkRelation(pid any, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/fork", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UploadAvatar uploads an avatar.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#upload-a-project-avatar
func (s *ProjectsService) UploadAvatar(pid any, avatar io.Reader, filename string, options ...RequestOptionFunc) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s", PathEscape(project))

	req, err := s.client.UploadRequest(
		http.MethodPut,
		u,
		avatar,
		filename,
		UploadAvatar,
		nil,
		options,
	)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// DownloadAvatar downloads an avatar.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#download-a-project-avatar
func (s *ProjectsService) DownloadAvatar(pid any, options ...RequestOptionFunc) (*bytes.Reader, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/avatar", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	avatar := new(bytes.Buffer)
	resp, err := s.client.Do(req, avatar)
	if err != nil {
		return nil, resp, err
	}

	return bytes.NewReader(avatar.Bytes()), resp, err
}

// ListProjectForks gets a list of project forks.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_forks/#list-forks-of-a-project
func (s *ProjectsService) ListProjectForks(pid any, opt *ListProjectsOptions, options ...RequestOptionFunc) ([]*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/forks", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var forks []*Project
	resp, err := s.client.Do(req, &forks)
	if err != nil {
		return nil, resp, err
	}

	return forks, resp, nil
}

// ProjectPushRules represents a project push rule.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_push_rules/
type ProjectPushRules struct {
	ID                         int        `json:"id"`
	ProjectID                  int        `json:"project_id"`
	CommitMessageRegex         string     `json:"commit_message_regex"`
	CommitMessageNegativeRegex string     `json:"commit_message_negative_regex"`
	BranchNameRegex            string     `json:"branch_name_regex"`
	DenyDeleteTag              bool       `json:"deny_delete_tag"`
	CreatedAt                  *time.Time `json:"created_at"`
	MemberCheck                bool       `json:"member_check"`
	PreventSecrets             bool       `json:"prevent_secrets"`
	AuthorEmailRegex           string     `json:"author_email_regex"`
	FileNameRegex              string     `json:"file_name_regex"`
	MaxFileSize                int        `json:"max_file_size"`
	CommitCommitterCheck       bool       `json:"commit_committer_check"`
	CommitCommitterNameCheck   bool       `json:"commit_committer_name_check"`
	RejectUnsignedCommits      bool       `json:"reject_unsigned_commits"`
	RejectNonDCOCommits        bool       `json:"reject_non_dco_commits"`
}

// GetProjectPushRules gets the push rules of a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_push_rules/#get-project-push-rules
func (s *ProjectsService) GetProjectPushRules(pid any, options ...RequestOptionFunc) (*ProjectPushRules, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/push_rule", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	ppr := new(ProjectPushRules)
	resp, err := s.client.Do(req, ppr)
	if err != nil {
		return nil, resp, err
	}

	return ppr, resp, nil
}

// AddProjectPushRuleOptions represents the available AddProjectPushRule()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_push_rules/#add-a-project-push-rule
type AddProjectPushRuleOptions struct {
	AuthorEmailRegex           *string `url:"author_email_regex,omitempty" json:"author_email_regex,omitempty"`
	BranchNameRegex            *string `url:"branch_name_regex,omitempty" json:"branch_name_regex,omitempty"`
	CommitCommitterCheck       *bool   `url:"commit_committer_check,omitempty" json:"commit_committer_check,omitempty"`
	CommitCommitterNameCheck   *bool   `url:"commit_committer_name_check,omitempty" json:"commit_committer_name_check,omitempty"`
	CommitMessageNegativeRegex *string `url:"commit_message_negative_regex,omitempty" json:"commit_message_negative_regex,omitempty"`
	CommitMessageRegex         *string `url:"commit_message_regex,omitempty" json:"commit_message_regex,omitempty"`
	DenyDeleteTag              *bool   `url:"deny_delete_tag,omitempty" json:"deny_delete_tag,omitempty"`
	FileNameRegex              *string `url:"file_name_regex,omitempty" json:"file_name_regex,omitempty"`
	MaxFileSize                *int    `url:"max_file_size,omitempty" json:"max_file_size,omitempty"`
	MemberCheck                *bool   `url:"member_check,omitempty" json:"member_check,omitempty"`
	PreventSecrets             *bool   `url:"prevent_secrets,omitempty" json:"prevent_secrets,omitempty"`
	RejectUnsignedCommits      *bool   `url:"reject_unsigned_commits,omitempty" json:"reject_unsigned_commits,omitempty"`
	RejectNonDCOCommits        *bool   `url:"reject_non_dco_commits,omitempty" json:"reject_non_dco_commits,omitempty"`
}

// AddProjectPushRule adds a push rule to a specified project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_push_rules/#add-a-project-push-rule
func (s *ProjectsService) AddProjectPushRule(pid any, opt *AddProjectPushRuleOptions, options ...RequestOptionFunc) (*ProjectPushRules, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/push_rule", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ppr := new(ProjectPushRules)
	resp, err := s.client.Do(req, ppr)
	if err != nil {
		return nil, resp, err
	}

	return ppr, resp, nil
}

// EditProjectPushRuleOptions represents the available EditProjectPushRule()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_push_rules/#edit-project-push-rule
type EditProjectPushRuleOptions struct {
	AuthorEmailRegex           *string `url:"author_email_regex,omitempty" json:"author_email_regex,omitempty"`
	BranchNameRegex            *string `url:"branch_name_regex,omitempty" json:"branch_name_regex,omitempty"`
	CommitCommitterCheck       *bool   `url:"commit_committer_check,omitempty" json:"commit_committer_check,omitempty"`
	CommitCommitterNameCheck   *bool   `url:"commit_committer_name_check,omitempty" json:"commit_committer_name_check,omitempty"`
	CommitMessageNegativeRegex *string `url:"commit_message_negative_regex,omitempty" json:"commit_message_negative_regex,omitempty"`
	CommitMessageRegex         *string `url:"commit_message_regex,omitempty" json:"commit_message_regex,omitempty"`
	DenyDeleteTag              *bool   `url:"deny_delete_tag,omitempty" json:"deny_delete_tag,omitempty"`
	FileNameRegex              *string `url:"file_name_regex,omitempty" json:"file_name_regex,omitempty"`
	MaxFileSize                *int    `url:"max_file_size,omitempty" json:"max_file_size,omitempty"`
	MemberCheck                *bool   `url:"member_check,omitempty" json:"member_check,omitempty"`
	PreventSecrets             *bool   `url:"prevent_secrets,omitempty" json:"prevent_secrets,omitempty"`
	RejectUnsignedCommits      *bool   `url:"reject_unsigned_commits,omitempty" json:"reject_unsigned_commits,omitempty"`
	RejectNonDCOCommits        *bool   `url:"reject_non_dco_commits,omitempty" json:"reject_non_dco_commits,omitempty"`
}

// EditProjectPushRule edits a push rule for a specified project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_push_rules/#edit-project-push-rule
func (s *ProjectsService) EditProjectPushRule(pid any, opt *EditProjectPushRuleOptions, options ...RequestOptionFunc) (*ProjectPushRules, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/push_rule", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ppr := new(ProjectPushRules)
	resp, err := s.client.Do(req, ppr)
	if err != nil {
		return nil, resp, err
	}

	return ppr, resp, nil
}

// DeleteProjectPushRule removes a push rule from a project. This is an
// idempotent method and can be called multiple times. Either the push rule is
// available or not.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_push_rules/#delete-project-push-rule
func (s *ProjectsService) DeleteProjectPushRule(pid any, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/push_rule", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ProjectApprovals represents GitLab project level merge request approvals.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#project-approval-rules
type ProjectApprovals struct {
	Approvers                                 []*MergeRequestApproverUser  `json:"approvers"`
	ApproverGroups                            []*MergeRequestApproverGroup `json:"approver_groups"`
	ResetApprovalsOnPush                      bool                         `json:"reset_approvals_on_push"`
	DisableOverridingApproversPerMergeRequest bool                         `json:"disable_overriding_approvers_per_merge_request"`
	MergeRequestsAuthorApproval               bool                         `json:"merge_requests_author_approval"`
	MergeRequestsDisableCommittersApproval    bool                         `json:"merge_requests_disable_committers_approval"`
	RequirePasswordToApprove                  bool                         `json:"require_password_to_approve"`
	SelectiveCodeOwnerRemovals                bool                         `json:"selective_code_owner_removals,omitempty"`

	// Deprecated: use Merge Request Approvals API instead
	ApprovalsBeforeMerge int `json:"approvals_before_merge"`
}

// GetApprovalConfiguration get the approval configuration for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#project-approval-rules
func (s *ProjectsService) GetApprovalConfiguration(pid any, options ...RequestOptionFunc) (*ProjectApprovals, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/approvals", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	pa := new(ProjectApprovals)
	resp, err := s.client.Do(req, pa)
	if err != nil {
		return nil, resp, err
	}

	return pa, resp, nil
}

// ChangeApprovalConfigurationOptions represents the available
// ApprovalConfiguration() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#change-configuration
type ChangeApprovalConfigurationOptions struct {
	DisableOverridingApproversPerMergeRequest *bool `url:"disable_overriding_approvers_per_merge_request,omitempty" json:"disable_overriding_approvers_per_merge_request,omitempty"`
	MergeRequestsAuthorApproval               *bool `url:"merge_requests_author_approval,omitempty" json:"merge_requests_author_approval,omitempty"`
	MergeRequestsDisableCommittersApproval    *bool `url:"merge_requests_disable_committers_approval,omitempty" json:"merge_requests_disable_committers_approval,omitempty"`
	RequirePasswordToApprove                  *bool `url:"require_password_to_approve,omitempty" json:"require_password_to_approve,omitempty"`
	ResetApprovalsOnPush                      *bool `url:"reset_approvals_on_push,omitempty" json:"reset_approvals_on_push,omitempty"`
	SelectiveCodeOwnerRemovals                *bool `url:"selective_code_owner_removals,omitempty" json:"selective_code_owner_removals,omitempty"`

	// Deprecated: use Merge Request Approvals API instead
	ApprovalsBeforeMerge *int `url:"approvals_before_merge,omitempty" json:"approvals_before_merge,omitempty"`
}

// ChangeApprovalConfiguration updates the approval configuration for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#change-configuration
func (s *ProjectsService) ChangeApprovalConfiguration(pid any, opt *ChangeApprovalConfigurationOptions, options ...RequestOptionFunc) (*ProjectApprovals, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/approvals", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pa := new(ProjectApprovals)
	resp, err := s.client.Do(req, pa)
	if err != nil {
		return nil, resp, err
	}

	return pa, resp, nil
}

// GetProjectApprovalRulesListsOptions represents the available
// GetProjectApprovalRules() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#get-all-approval-rules-for-project
type GetProjectApprovalRulesListsOptions ListOptions

// GetProjectApprovalRules looks up the list of project level approver rules.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#get-all-approval-rules-for-project
func (s *ProjectsService) GetProjectApprovalRules(pid any, opt *GetProjectApprovalRulesListsOptions, options ...RequestOptionFunc) ([]*ProjectApprovalRule, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/approval_rules", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var par []*ProjectApprovalRule
	resp, err := s.client.Do(req, &par)
	if err != nil {
		return nil, resp, err
	}

	return par, resp, nil
}

// GetProjectApprovalRule gets the project level approvers.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#get-single-approval-rule-for-project
func (s *ProjectsService) GetProjectApprovalRule(pid any, ruleID int, options ...RequestOptionFunc) (*ProjectApprovalRule, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/approval_rules/%d", PathEscape(project), ruleID)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	par := new(ProjectApprovalRule)
	resp, err := s.client.Do(req, &par)
	if err != nil {
		return nil, resp, err
	}

	return par, resp, nil
}

// CreateProjectLevelRuleOptions represents the available CreateProjectApprovalRule()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#create-project-approval-rule
type CreateProjectLevelRuleOptions struct {
	Name                          *string   `url:"name,omitempty" json:"name,omitempty"`
	ApprovalsRequired             *int      `url:"approvals_required,omitempty" json:"approvals_required,omitempty"`
	ReportType                    *string   `url:"report_type,omitempty" json:"report_type,omitempty"`
	RuleType                      *string   `url:"rule_type,omitempty" json:"rule_type,omitempty"`
	UserIDs                       *[]int    `url:"user_ids,omitempty" json:"user_ids,omitempty"`
	GroupIDs                      *[]int    `url:"group_ids,omitempty" json:"group_ids,omitempty"`
	ProtectedBranchIDs            *[]int    `url:"protected_branch_ids,omitempty" json:"protected_branch_ids,omitempty"`
	AppliesToAllProtectedBranches *bool     `url:"applies_to_all_protected_branches,omitempty" json:"applies_to_all_protected_branches,omitempty"`
	Usernames                     *[]string `url:"usernames,omitempty" json:"usernames,omitempty"`
}

// CreateProjectApprovalRule creates a new project-level approval rule.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#create-project-approval-rule
func (s *ProjectsService) CreateProjectApprovalRule(pid any, opt *CreateProjectLevelRuleOptions, options ...RequestOptionFunc) (*ProjectApprovalRule, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/approval_rules", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	par := new(ProjectApprovalRule)
	resp, err := s.client.Do(req, &par)
	if err != nil {
		return nil, resp, err
	}

	return par, resp, nil
}

// UpdateProjectLevelRuleOptions represents the available UpdateProjectApprovalRule()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#update-project-approval-rule
type UpdateProjectLevelRuleOptions struct {
	Name                          *string   `url:"name,omitempty" json:"name,omitempty"`
	ApprovalsRequired             *int      `url:"approvals_required,omitempty" json:"approvals_required,omitempty"`
	UserIDs                       *[]int    `url:"user_ids,omitempty" json:"user_ids,omitempty"`
	GroupIDs                      *[]int    `url:"group_ids,omitempty" json:"group_ids,omitempty"`
	ProtectedBranchIDs            *[]int    `url:"protected_branch_ids,omitempty" json:"protected_branch_ids,omitempty"`
	AppliesToAllProtectedBranches *bool     `url:"applies_to_all_protected_branches,omitempty" json:"applies_to_all_protected_branches,omitempty"`
	Usernames                     *[]string `url:"usernames,omitempty" json:"usernames,omitempty"`
}

// UpdateProjectApprovalRule updates an existing approval rule with new options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#update-project-approval-rule
func (s *ProjectsService) UpdateProjectApprovalRule(pid any, approvalRule int, opt *UpdateProjectLevelRuleOptions, options ...RequestOptionFunc) (*ProjectApprovalRule, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/approval_rules/%d", PathEscape(project), approvalRule)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	par := new(ProjectApprovalRule)
	resp, err := s.client.Do(req, &par)
	if err != nil {
		return nil, resp, err
	}

	return par, resp, nil
}

// DeleteProjectApprovalRule deletes a project-level approval rule.
//
// GitLab API docs:
// https://docs.gitlab.com/api/merge_request_approvals/#delete-project-approval-rule
func (s *ProjectsService) DeleteProjectApprovalRule(pid any, approvalRule int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/approval_rules/%d", PathEscape(project), approvalRule)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ProjectPullMirrorDetails represent the details of the configuration pull
// mirror and its update status.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_pull_mirroring/
type ProjectPullMirrorDetails struct {
	ID                     int        `json:"id"`
	LastError              string     `json:"last_error"`
	LastSuccessfulUpdateAt *time.Time `json:"last_successful_update_at"`
	LastUpdateAt           *time.Time `json:"last_update_at"`
	LastUpdateStartedAt    *time.Time `json:"last_update_started_at"`
	UpdateStatus           string     `json:"update_status"`
	URL                    string     `json:"url"`
}

// GetProjectPullMirrorDetails returns the pull mirror details.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_pull_mirroring/#get-a-projects-pull-mirror-details
func (s *ProjectsService) GetProjectPullMirrorDetails(pid any, options ...RequestOptionFunc) (*ProjectPullMirrorDetails, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/mirror/pull", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	pmd := new(ProjectPullMirrorDetails)
	resp, err := s.client.Do(req, pmd)
	if err != nil {
		return nil, resp, err
	}

	return pmd, resp, nil
}

// ConfigureProjectPullMirrorOptions represents the available ConfigureProjectPullMirror() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_pull_mirroring/#configure-pull-mirroring-for-a-project
type ConfigureProjectPullMirrorOptions struct {
	Enabled                          *bool   `url:"enabled,omitempty" json:"enabled,omitempty"`
	URL                              *string `url:"url,omitempty" json:"url,omitempty"`
	AuthUser                         *string `url:"auth_user,omitempty" json:"auth_user,omitempty"`
	AuthPassword                     *string `url:"auth_password,omitempty" json:"auth_password,omitempty"`
	MirrorTriggerBuilds              *bool   `url:"mirror_trigger_builds,omitempty" json:"mirror_trigger_builds,omitempty"`
	OnlyMirrorProtectedBranches      *bool   `url:"only_mirror_protected_branches,omitempty" json:"only_mirror_protected_branches,omitempty"`
	MirrorOverwritesDivergedBranches *bool   `url:"mirror_overwrites_diverged_branches,omitempty" json:"mirror_overwrites_diverged_branches,omitempty"`
	MirrorBranchRegex                *string `url:"mirror_branch_regex,omitempty" json:"mirror_branch_regex,omitempty"`
}

// ConfigureProjectPullMirror configures pull mirroring settings.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_pull_mirroring/#configure-pull-mirroring-for-a-project
func (s *ProjectsService) ConfigureProjectPullMirror(pid any, opt *ConfigureProjectPullMirrorOptions, options ...RequestOptionFunc) (*ProjectPullMirrorDetails, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/mirror/pull", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pmd := new(ProjectPullMirrorDetails)
	resp, err := s.client.Do(req, pmd)
	if err != nil {
		return nil, resp, err
	}

	return pmd, resp, nil
}

// StartMirroringProject start the pull mirroring process for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/project_pull_mirroring/#start-the-pull-mirroring-process-for-a-project
func (s *ProjectsService) StartMirroringProject(pid any, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/mirror/pull", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// TransferProjectOptions represents the available TransferProject() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#transfer-a-project-to-a-new-namespace
type TransferProjectOptions struct {
	Namespace any `url:"namespace,omitempty" json:"namespace,omitempty"`
}

// TransferProject transfer a project into the specified namespace
//
// GitLab API docs: https://docs.gitlab.com/api/projects/#transfer-a-project-to-a-new-namespace
func (s *ProjectsService) TransferProject(pid any, opt *TransferProjectOptions, options ...RequestOptionFunc) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/transfer", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// StartHousekeepingProject start the Housekeeping task for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#start-the-housekeeping-task-for-a-project
func (s *ProjectsService) StartHousekeepingProject(pid any, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/housekeeping", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ProjectRepositoryStorage represents the repository storage information for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#get-the-path-to-repository-storage
type ProjectRepositoryStorage struct {
	ProjectID         int        `json:"project_id"`
	DiskPath          string     `json:"disk_path"`
	CreatedAt         *time.Time `json:"created_at"`
	RepositoryStorage string     `json:"repository_storage"`
}

// ProjectReposityStorage is kept for backwards compatibility.
//
// Deprecated: use ProjectRepositoryStorage instead. The ProjectReposityStorage type will be removed in the next release.
type ProjectReposityStorage = ProjectRepositoryStorage

// GetRepositoryStorage Get the path to repository storage.
//
// GitLab API docs:
// https://docs.gitlab.com/api/projects/#get-the-path-to-repository-storage
func (s *ProjectsService) GetRepositoryStorage(pid any, options ...RequestOptionFunc) (*ProjectReposityStorage, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/storage", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	prs := new(ProjectReposityStorage)
	resp, err := s.client.Do(req, prs)
	if err != nil {
		return nil, resp, err
	}

	return prs, resp, nil
}
