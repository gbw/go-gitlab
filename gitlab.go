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

// Package gitlab implements a GitLab API client.
package gitlab

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"math"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-cleanhttp"

	"github.com/google/go-querystring/query"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"golang.org/x/oauth2"
	"golang.org/x/time/rate"
)

const (
	defaultBaseURL = "https://gitlab.com/"
	apiVersionPath = "api/v4/"
	userAgent      = "go-gitlab"

	headerRateLimit = "RateLimit-Limit"
	headerRateReset = "RateLimit-Reset"

	AccessTokenHeaderName = "PRIVATE-TOKEN"
	JobTokenHeaderName    = "JOB-TOKEN"
)

// AuthType represents an authentication type within GitLab.
//
// GitLab API docs: https://docs.gitlab.com/api/
type AuthType int

// List of available authentication types.
//
// GitLab API docs: https://docs.gitlab.com/api/
const (
	BasicAuth AuthType = iota
	JobToken
	OAuthToken
	PrivateToken
)

var ErrNotFound = errors.New("404 Not Found")

// A Client manages communication with the GitLab API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *retryablehttp.Client

	// Base URL for API requests. Defaults to the public GitLab API, but can be
	// set to a domain endpoint to use with a self hosted GitLab server. baseURL
	// should always be specified with a trailing slash.
	baseURL *url.URL

	// disableRetries is used to disable the default retry logic.
	disableRetries bool

	// configureLimiterOnce is used to make sure the limiter is configured exactly
	// once and block all other calls until the initial (one) call is done.
	configureLimiterOnce sync.Once

	// Limiter is used to limit API calls and prevent 429 responses.
	limiter RateLimiter

	// authSource is used to obtain authentication headers.
	authSource AuthSource

	// authSourceInit is used to ensure that AuthSources are initialized only
	// once.
	authSourceInit sync.Once

	// Default request options applied to every request.
	defaultRequestOptions []RequestOptionFunc

	// User agent used when communicating with the GitLab API.
	UserAgent string

	// GraphQL interface
	GraphQL GraphQLInterface

	// Services used for talking to different parts of the GitLab API.
	AccessRequests                   AccessRequestsServiceInterface
	AlertManagement                  AlertManagementServiceInterface
	Appearance                       AppearanceServiceInterface
	Applications                     ApplicationsServiceInterface
	ApplicationStatistics            ApplicationStatisticsServiceInterface
	AuditEvents                      AuditEventsServiceInterface
	Avatar                           AvatarRequestsServiceInterface
	AwardEmoji                       AwardEmojiServiceInterface
	Boards                           IssueBoardsServiceInterface
	Branches                         BranchesServiceInterface
	BroadcastMessage                 BroadcastMessagesServiceInterface
	BulkImports                      BulkImportsServiceInterface
	CIYMLTemplate                    CIYMLTemplatesServiceInterface
	ClusterAgents                    ClusterAgentsServiceInterface
	Commits                          CommitsServiceInterface
	ContainerRegistry                ContainerRegistryServiceInterface
	ContainerRegistryProtectionRules ContainerRegistryProtectionRulesServiceInterface
	CustomAttribute                  CustomAttributesServiceInterface
	DatabaseMigrations               DatabaseMigrationsServiceInterface
	Dependencies                     DependenciesServiceInterface
	DependencyListExport             DependencyListExportServiceInterface
	DependencyProxy                  DependencyProxyServiceInterface
	DeployKeys                       DeployKeysServiceInterface
	DeployTokens                     DeployTokensServiceInterface
	DeploymentMergeRequests          DeploymentMergeRequestsServiceInterface
	Deployments                      DeploymentsServiceInterface
	Discussions                      DiscussionsServiceInterface
	DockerfileTemplate               DockerfileTemplatesServiceInterface
	DORAMetrics                      DORAMetricsServiceInterface
	DraftNotes                       DraftNotesServiceInterface
	EnterpriseUsers                  EnterpriseUsersServiceInterface
	Environments                     EnvironmentsServiceInterface
	EpicIssues                       EpicIssuesServiceInterface
	Epics                            EpicsServiceInterface
	ErrorTracking                    ErrorTrackingServiceInterface
	Events                           EventsServiceInterface
	ExternalStatusChecks             ExternalStatusChecksServiceInterface
	FeatureFlagUserLists             FeatureFlagUserListsServiceInterface
	Features                         FeaturesServiceInterface
	FreezePeriods                    FreezePeriodsServiceInterface
	GenericPackages                  GenericPackagesServiceInterface
	GeoNodes                         GeoNodesServiceInterface
	GeoSites                         GeoSitesServiceInterface
	GitIgnoreTemplates               GitIgnoreTemplatesServiceInterface
	GroupAccessTokens                GroupAccessTokensServiceInterface
	GroupActivityAnalytics           GroupActivityAnalyticsServiceInterface
	GroupBadges                      GroupBadgesServiceInterface
	GroupCluster                     GroupClustersServiceInterface
	GroupEpicBoards                  GroupEpicBoardsServiceInterface
	GroupImportExport                GroupImportExportServiceInterface
	Integrations                     IntegrationsServiceInterface
	GroupIssueBoards                 GroupIssueBoardsServiceInterface
	GroupIterations                  GroupIterationsServiceInterface
	GroupLabels                      GroupLabelsServiceInterface
	GroupMarkdownUploads             GroupMarkdownUploadsServiceInterface
	GroupMembers                     GroupMembersServiceInterface
	GroupMilestones                  GroupMilestonesServiceInterface
	GroupProtectedEnvironments       GroupProtectedEnvironmentsServiceInterface
	GroupReleases                    GroupReleasesServiceInterface
	GroupRepositoryStorageMove       GroupRepositoryStorageMoveServiceInterface
	GroupSCIM                        GroupSCIMServiceInterface
	GroupSecuritySettings            GroupSecuritySettingsServiceInterface
	GroupSSHCertificates             GroupSSHCertificatesServiceInterface
	GroupVariables                   GroupVariablesServiceInterface
	GroupWikis                       GroupWikisServiceInterface
	Groups                           GroupsServiceInterface
	Import                           ImportServiceInterface
	InstanceCluster                  InstanceClustersServiceInterface
	InstanceVariables                InstanceVariablesServiceInterface
	Invites                          InvitesServiceInterface
	IssueLinks                       IssueLinksServiceInterface
	Issues                           IssuesServiceInterface
	IssuesStatistics                 IssuesStatisticsServiceInterface
	Jobs                             JobsServiceInterface
	JobTokenScope                    JobTokenScopeServiceInterface
	Keys                             KeysServiceInterface
	Labels                           LabelsServiceInterface
	License                          LicenseServiceInterface
	LicenseTemplates                 LicenseTemplatesServiceInterface
	Markdown                         MarkdownServiceInterface
	MemberRolesService               MemberRolesServiceInterface
	MergeRequestApprovals            MergeRequestApprovalsServiceInterface
	MergeRequestApprovalSettings     MergeRequestApprovalSettingsServiceInterface
	MergeRequests                    MergeRequestsServiceInterface
	MergeTrains                      MergeTrainsServiceInterface
	Metadata                         MetadataServiceInterface
	Milestones                       MilestonesServiceInterface
	Namespaces                       NamespacesServiceInterface
	Notes                            NotesServiceInterface
	NotificationSettings             NotificationSettingsServiceInterface
	Packages                         PackagesServiceInterface
	Pages                            PagesServiceInterface
	PagesDomains                     PagesDomainsServiceInterface
	PersonalAccessTokens             PersonalAccessTokensServiceInterface
	PipelineSchedules                PipelineSchedulesServiceInterface
	PipelineTriggers                 PipelineTriggersServiceInterface
	Pipelines                        PipelinesServiceInterface
	PlanLimits                       PlanLimitsServiceInterface
	ProjectAccessTokens              ProjectAccessTokensServiceInterface
	ProjectBadges                    ProjectBadgesServiceInterface
	ProjectCluster                   ProjectClustersServiceInterface
	ProjectFeatureFlags              ProjectFeatureFlagServiceInterface
	ProjectImportExport              ProjectImportExportServiceInterface
	ProjectIterations                ProjectIterationsServiceInterface
	ProjectMarkdownUploads           ProjectMarkdownUploadsServiceInterface
	ProjectMembers                   ProjectMembersServiceInterface
	ProjectMirrors                   ProjectMirrorServiceInterface
	ProjectRepositoryStorageMove     ProjectRepositoryStorageMoveServiceInterface
	ProjectSecuritySettings          ProjectSecuritySettingsServiceInterface
	ProjectSnippets                  ProjectSnippetsServiceInterface
	ProjectTemplates                 ProjectTemplatesServiceInterface
	ProjectVariables                 ProjectVariablesServiceInterface
	ProjectVulnerabilities           ProjectVulnerabilitiesServiceInterface
	Projects                         ProjectsServiceInterface
	ProtectedBranches                ProtectedBranchesServiceInterface
	ProtectedEnvironments            ProtectedEnvironmentsServiceInterface
	ProtectedTags                    ProtectedTagsServiceInterface
	ReleaseLinks                     ReleaseLinksServiceInterface
	Releases                         ReleasesServiceInterface
	Repositories                     RepositoriesServiceInterface
	RepositoryFiles                  RepositoryFilesServiceInterface
	RepositorySubmodules             RepositorySubmodulesServiceInterface
	ResourceGroup                    ResourceGroupServiceInterface
	ResourceIterationEvents          ResourceIterationEventsServiceInterface
	ResourceLabelEvents              ResourceLabelEventsServiceInterface
	ResourceMilestoneEvents          ResourceMilestoneEventsServiceInterface
	ResourceStateEvents              ResourceStateEventsServiceInterface
	ResourceWeightEvents             ResourceWeightEventsServiceInterface
	Runners                          RunnersServiceInterface
	Search                           SearchServiceInterface
	SecureFiles                      SecureFilesServiceInterface
	Services                         ServicesServiceInterface
	Settings                         SettingsServiceInterface
	Sidekiq                          SidekiqServiceInterface
	SnippetRepositoryStorageMove     SnippetRepositoryStorageMoveServiceInterface
	Snippets                         SnippetsServiceInterface
	SystemHooks                      SystemHooksServiceInterface
	Tags                             TagsServiceInterface
	TerraformStates                  TerraformStatesServiceInterface
	Todos                            TodosServiceInterface
	Topics                           TopicsServiceInterface
	UsageData                        UsageDataServiceInterface
	Users                            UsersServiceInterface
	Validate                         ValidateServiceInterface
	Version                          VersionServiceInterface
	Wikis                            WikisServiceInterface
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For keyset-based paginated result sets, the value must be `"keyset"`
	Pagination string `url:"pagination,omitempty" json:"pagination,omitempty"`
	// For offset-based and keyset-based paginated result sets, the number of results to include per page.
	PerPage int `url:"per_page,omitempty" json:"per_page,omitempty"`
	// For offset-based paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty" json:"page,omitempty"`
	// For keyset-based paginated result sets, tree record ID at which to fetch the next page.
	PageToken string `url:"page_token,omitempty" json:"page_token,omitempty"`
	// For keyset-based paginated result sets, name of the column by which to order
	OrderBy string `url:"order_by,omitempty" json:"order_by,omitempty"`
	// For keyset-based paginated result sets, sort order (`"asc"`` or `"desc"`)
	Sort string `url:"sort,omitempty" json:"sort,omitempty"`
}

// RateLimiter describes the interface that all (custom) rate limiters must implement.
type RateLimiter interface {
	Wait(context.Context) error
}

// NewClient returns a new GitLab API client. To use API methods which require
// authentication, provide a valid private or personal token.
func NewClient(token string, options ...ClientOptionFunc) (*Client, error) {
	as := AccessTokenAuthSource{Token: token}
	return NewAuthSourceClient(as, options...)
}

// NewBasicAuthClient returns a new GitLab API client using the OAuth 2.0 Resource Owner Password Credentials flow.
// The provided username and password are used to obtain an OAuth access token
// from GitLab's token endpoint on the first API request. The token is then
// cached, reused for subsequent requests, and refreshed when expired.
//
// The Resource Owner Password Credentials flow is only suitable for trusted,
// first-party applications and does not work for users who have two-factor
// authentication enabled.
//
// Note: This method uses OAuth tokens with Bearer authentication, not HTTP Basic Auth.
//
// Deprecated: GitLab recommends against using this authentication method.
func NewBasicAuthClient(username, password string, options ...ClientOptionFunc) (*Client, error) {
	as := &PasswordCredentialsAuthSource{
		Username: username,
		Password: password,
	}

	return NewAuthSourceClient(as, options...)
}

// NewJobClient returns a new GitLab API client. To use API methods which require
// authentication, provide a valid job token.
func NewJobClient(token string, options ...ClientOptionFunc) (*Client, error) {
	as := JobTokenAuthSource{Token: token}
	return NewAuthSourceClient(as, options...)
}

// NewOAuthClient returns a new GitLab API client using a static OAuth bearer token for authentication.
//
// Deprecated: use NewAuthSourceClient with a StaticTokenSource instead. For example:
//
//	ts := oauth2.StaticTokenSource(
//	    &oauth2.Token{AccessToken: "YOUR STATIC TOKEN"},
//	)
//	c, err := gitlab.NewAuthSourceClient(gitlab.OAuthTokenSource{ts})
func NewOAuthClient(token string, options ...ClientOptionFunc) (*Client, error) {
	as := OAuthTokenSource{
		TokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
	}

	return NewAuthSourceClient(as, options...)
}

// NewAuthSourceClient returns a new GitLab API client that uses the AuthSource for authentication.
func NewAuthSourceClient(as AuthSource, options ...ClientOptionFunc) (*Client, error) {
	c := &Client{
		UserAgent:  userAgent,
		authSource: as,
	}

	// Configure the HTTP client.
	c.client = &retryablehttp.Client{
		Backoff:      c.retryHTTPBackoff,
		CheckRetry:   c.retryHTTPCheck,
		ErrorHandler: retryablehttp.PassthroughErrorHandler,
		HTTPClient:   cleanhttp.DefaultPooledClient(),
		RetryWaitMin: 100 * time.Millisecond,
		RetryWaitMax: 400 * time.Millisecond,
		RetryMax:     5,
	}

	// Set the default base URL.
	c.setBaseURL(defaultBaseURL)

	// Apply any given client options.
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}

	// If no custom limiter was set using a client option, configure
	// the default rate limiter with values that implicitly disable
	// rate limiting until an initial HTTP call is done and we can
	// use the headers to try and properly configure the limiter.
	if c.limiter == nil {
		c.limiter = rate.NewLimiter(rate.Inf, 0)
	}

	// Create the internal timeStats service.
	timeStats := &timeStatsService{client: c}

	// GraphQL interface
	c.GraphQL = &GraphQL{client: c}

	// Create all the public services.
	c.AccessRequests = &AccessRequestsService{client: c}
	c.AlertManagement = &AlertManagementService{client: c}
	c.Appearance = &AppearanceService{client: c}
	c.Applications = &ApplicationsService{client: c}
	c.ApplicationStatistics = &ApplicationStatisticsService{client: c}
	c.AuditEvents = &AuditEventsService{client: c}
	c.Avatar = &AvatarRequestsService{client: c}
	c.AwardEmoji = &AwardEmojiService{client: c}
	c.Boards = &IssueBoardsService{client: c}
	c.Branches = &BranchesService{client: c}
	c.BroadcastMessage = &BroadcastMessagesService{client: c}
	c.BulkImports = &BulkImportsService{client: c}
	c.CIYMLTemplate = &CIYMLTemplatesService{client: c}
	c.ClusterAgents = &ClusterAgentsService{client: c}
	c.Commits = &CommitsService{client: c}
	c.ContainerRegistry = &ContainerRegistryService{client: c}
	c.ContainerRegistryProtectionRules = &ContainerRegistryProtectionRulesService{client: c}
	c.CustomAttribute = &CustomAttributesService{client: c}
	c.DatabaseMigrations = &DatabaseMigrationsService{client: c}
	c.Dependencies = &DependenciesService{client: c}
	c.DependencyListExport = &DependencyListExportService{client: c}
	c.DependencyProxy = &DependencyProxyService{client: c}
	c.DeployKeys = &DeployKeysService{client: c}
	c.DeployTokens = &DeployTokensService{client: c}
	c.DeploymentMergeRequests = &DeploymentMergeRequestsService{client: c}
	c.Deployments = &DeploymentsService{client: c}
	c.Discussions = &DiscussionsService{client: c}
	c.DockerfileTemplate = &DockerfileTemplatesService{client: c}
	c.DORAMetrics = &DORAMetricsService{client: c}
	c.DraftNotes = &DraftNotesService{client: c}
	c.EnterpriseUsers = &EnterpriseUsersService{client: c}
	c.Environments = &EnvironmentsService{client: c}
	c.EpicIssues = &EpicIssuesService{client: c}
	c.Epics = &EpicsService{client: c}
	c.ErrorTracking = &ErrorTrackingService{client: c}
	c.Events = &EventsService{client: c}
	c.ExternalStatusChecks = &ExternalStatusChecksService{client: c}
	c.FeatureFlagUserLists = &FeatureFlagUserListsService{client: c}
	c.Features = &FeaturesService{client: c}
	c.FreezePeriods = &FreezePeriodsService{client: c}
	c.GenericPackages = &GenericPackagesService{client: c}
	c.GeoNodes = &GeoNodesService{client: c}
	c.GeoSites = &GeoSitesService{client: c}
	c.GitIgnoreTemplates = &GitIgnoreTemplatesService{client: c}
	c.GroupAccessTokens = &GroupAccessTokensService{client: c}
	c.GroupActivityAnalytics = &GroupActivityAnalyticsService{client: c}
	c.GroupBadges = &GroupBadgesService{client: c}
	c.GroupCluster = &GroupClustersService{client: c}
	c.GroupEpicBoards = &GroupEpicBoardsService{client: c}
	c.GroupImportExport = &GroupImportExportService{client: c}
	c.Integrations = &IntegrationsService{client: c}
	c.GroupIssueBoards = &GroupIssueBoardsService{client: c}
	c.GroupIterations = &GroupIterationsService{client: c}
	c.GroupLabels = &GroupLabelsService{client: c}
	c.GroupMarkdownUploads = &GroupMarkdownUploadsService{client: c}
	c.GroupMembers = &GroupMembersService{client: c}
	c.GroupMilestones = &GroupMilestonesService{client: c}
	c.GroupProtectedEnvironments = &GroupProtectedEnvironmentsService{client: c}
	c.GroupReleases = &GroupReleasesService{client: c}
	c.GroupRepositoryStorageMove = &GroupRepositoryStorageMoveService{client: c}
	c.GroupSCIM = &GroupSCIMService{client: c}
	c.GroupSecuritySettings = &GroupSecuritySettingsService{client: c}
	c.GroupSSHCertificates = &GroupSSHCertificatesService{client: c}
	c.GroupVariables = &GroupVariablesService{client: c}
	c.GroupWikis = &GroupWikisService{client: c}
	c.Groups = &GroupsService{client: c}
	c.Import = &ImportService{client: c}
	c.InstanceCluster = &InstanceClustersService{client: c}
	c.InstanceVariables = &InstanceVariablesService{client: c}
	c.Invites = &InvitesService{client: c}
	c.IssueLinks = &IssueLinksService{client: c}
	c.Issues = &IssuesService{client: c, timeStats: timeStats}
	c.IssuesStatistics = &IssuesStatisticsService{client: c}
	c.Jobs = &JobsService{client: c}
	c.JobTokenScope = &JobTokenScopeService{client: c}
	c.Keys = &KeysService{client: c}
	c.Labels = &LabelsService{client: c}
	c.License = &LicenseService{client: c}
	c.LicenseTemplates = &LicenseTemplatesService{client: c}
	c.Markdown = &MarkdownService{client: c}
	c.MemberRolesService = &MemberRolesService{client: c}
	c.MergeRequestApprovals = &MergeRequestApprovalsService{client: c}
	c.MergeRequestApprovalSettings = &MergeRequestApprovalSettingsService{client: c}
	c.MergeRequests = &MergeRequestsService{client: c, timeStats: timeStats}
	c.MergeTrains = &MergeTrainsService{client: c}
	c.Metadata = &MetadataService{client: c}
	c.Milestones = &MilestonesService{client: c}
	c.Namespaces = &NamespacesService{client: c}
	c.Notes = &NotesService{client: c}
	c.NotificationSettings = &NotificationSettingsService{client: c}
	c.Packages = &PackagesService{client: c}
	c.Pages = &PagesService{client: c}
	c.PagesDomains = &PagesDomainsService{client: c}
	c.PersonalAccessTokens = &PersonalAccessTokensService{client: c}
	c.PipelineSchedules = &PipelineSchedulesService{client: c}
	c.PipelineTriggers = &PipelineTriggersService{client: c}
	c.Pipelines = &PipelinesService{client: c}
	c.PlanLimits = &PlanLimitsService{client: c}
	c.ProjectAccessTokens = &ProjectAccessTokensService{client: c}
	c.ProjectBadges = &ProjectBadgesService{client: c}
	c.ProjectCluster = &ProjectClustersService{client: c}
	c.ProjectFeatureFlags = &ProjectFeatureFlagService{client: c}
	c.ProjectImportExport = &ProjectImportExportService{client: c}
	c.ProjectIterations = &ProjectIterationsService{client: c}
	c.ProjectMarkdownUploads = &ProjectMarkdownUploadsService{client: c}
	c.ProjectMembers = &ProjectMembersService{client: c}
	c.ProjectMirrors = &ProjectMirrorService{client: c}
	c.ProjectRepositoryStorageMove = &ProjectRepositoryStorageMoveService{client: c}
	c.ProjectSecuritySettings = &ProjectSecuritySettingsService{client: c}
	c.ProjectSnippets = &ProjectSnippetsService{client: c}
	c.ProjectTemplates = &ProjectTemplatesService{client: c}
	c.ProjectVariables = &ProjectVariablesService{client: c}
	c.ProjectVulnerabilities = &ProjectVulnerabilitiesService{client: c}
	c.Projects = &ProjectsService{client: c}
	c.ProtectedBranches = &ProtectedBranchesService{client: c}
	c.ProtectedEnvironments = &ProtectedEnvironmentsService{client: c}
	c.ProtectedTags = &ProtectedTagsService{client: c}
	c.ReleaseLinks = &ReleaseLinksService{client: c}
	c.Releases = &ReleasesService{client: c}
	c.Repositories = &RepositoriesService{client: c}
	c.RepositoryFiles = &RepositoryFilesService{client: c}
	c.RepositorySubmodules = &RepositorySubmodulesService{client: c}
	c.ResourceGroup = &ResourceGroupService{client: c}
	c.ResourceIterationEvents = &ResourceIterationEventsService{client: c}
	c.ResourceLabelEvents = &ResourceLabelEventsService{client: c}
	c.ResourceMilestoneEvents = &ResourceMilestoneEventsService{client: c}
	c.ResourceStateEvents = &ResourceStateEventsService{client: c}
	c.ResourceWeightEvents = &ResourceWeightEventsService{client: c}
	c.Runners = &RunnersService{client: c}
	c.Search = &SearchService{client: c}
	c.SecureFiles = &SecureFilesService{client: c}
	c.Services = &ServicesService{client: c}
	c.Settings = &SettingsService{client: c}
	c.Sidekiq = &SidekiqService{client: c}
	c.Snippets = &SnippetsService{client: c}
	c.SnippetRepositoryStorageMove = &SnippetRepositoryStorageMoveService{client: c}
	c.SystemHooks = &SystemHooksService{client: c}
	c.Tags = &TagsService{client: c}
	c.TerraformStates = &TerraformStatesService{client: c}
	c.Todos = &TodosService{client: c}
	c.Topics = &TopicsService{client: c}
	c.UsageData = &UsageDataService{client: c}
	c.Users = &UsersService{client: c}
	c.Validate = &ValidateService{client: c}
	c.Version = &VersionService{client: c}
	c.Wikis = &WikisService{client: c}

	return c, nil
}

func (c *Client) HTTPClient() *http.Client {
	return c.client.HTTPClient
}

// retryHTTPCheck provides a callback for Client.CheckRetry which
// will retry both rate limit (429) and server (>= 500) errors.
func (c *Client) retryHTTPCheck(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	if err != nil {
		return false, err
	}
	if !c.disableRetries && (resp.StatusCode == 429 || resp.StatusCode >= 500) {
		return true, nil
	}
	return false, nil
}

// retryHTTPBackoff provides a generic callback for Client.Backoff which
// will pass through all calls based on the status code of the response.
func (c *Client) retryHTTPBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	// Use the rate limit backoff function when we are rate limited.
	if resp != nil && resp.StatusCode == 429 {
		return rateLimitBackoff(min, max, attemptNum, resp)
	}

	// Set custom duration's when we experience a service interruption.
	min = 700 * time.Millisecond
	max = 900 * time.Millisecond

	return retryablehttp.LinearJitterBackoff(min, max, attemptNum, resp)
}

// rateLimitBackoff provides a callback for Client.Backoff which will use the
// RateLimit-Reset header to determine the time to wait. We add some jitter
// to prevent a thundering herd.
//
// min and max are mainly used for bounding the jitter that will be added to
// the reset time retrieved from the headers. But if the final wait time is
// less then min, min will be used instead.
func rateLimitBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	// rnd is used to generate pseudo-random numbers.
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// First create some jitter bounded by the min and max durations.
	jitter := time.Duration(rnd.Float64() * float64(max-min))

	if resp != nil {
		if v := resp.Header.Get(headerRateReset); v != "" {
			if reset, _ := strconv.ParseInt(v, 10, 64); reset > 0 {
				// Only update min if the given time to wait is longer.
				if wait := time.Until(time.Unix(reset, 0)); wait > min {
					min = wait
				}
			}
		} else {
			// In case the RateLimit-Reset header is not set, back off an additional
			// 100% exponentially. With the default milliseconds being set to 100 for
			// `min`, this makes the 5th retry wait 3.2 seconds (3,200 ms) by default.
			min = time.Duration(float64(min) * math.Pow(2, float64(attemptNum)))
		}
	}

	return min + jitter
}

// configureLimiter configures the rate limiter.
func (c *Client) configureLimiter(ctx context.Context, headers http.Header) {
	if v := headers.Get(headerRateLimit); v != "" {
		if rateLimit, _ := strconv.ParseFloat(v, 64); rateLimit > 0 {
			// The rate limit is based on requests per minute, so for our limiter to
			// work correctly we divide the limit by 60 to get the limit per second.
			rateLimit /= 60

			// Configure the limit and burst using a split of 2/3 for the limit and
			// 1/3 for the burst. This enables clients to burst 1/3 of the allowed
			// calls before the limiter kicks in. The remaining calls will then be
			// spread out evenly using intervals of time.Second / limit which should
			// prevent hitting the rate limit.
			limit := rate.Limit(rateLimit * 0.66)
			burst := int(rateLimit * 0.33)

			// Need at least one allowed to burst or x/time will throw an error
			if burst == 0 {
				burst = 1
			}

			// Create a new limiter using the calculated values.
			c.limiter = rate.NewLimiter(limit, burst)

			// Call the limiter once as we have already made a request
			// to get the headers and the limiter is not aware of this.
			c.limiter.Wait(ctx)
		}
	}
}

// BaseURL return a copy of the baseURL.
func (c *Client) BaseURL() *url.URL {
	u := *c.baseURL
	return &u
}

// setBaseURL sets the base URL for API requests to a custom endpoint.
func (c *Client) setBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseURL.Path, apiVersionPath) {
		baseURL.Path += apiVersionPath
	}

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

// NewRequest creates a new API request. The method expects a relative URL
// path that will be resolved relative to the base URL of the Client.
// Relative URL paths should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included
// as the request body.
func (c *Client) NewRequest(method, path string, opt any, options []RequestOptionFunc) (*retryablehttp.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	return c.NewRequestToURL(method, &u, opt, options)
}

func (c *Client) NewRequestToURL(method string, u *url.URL, opt any, options []RequestOptionFunc) (*retryablehttp.Request, error) {
	if u.Scheme != c.baseURL.Scheme || u.Host != c.baseURL.Host {
		return nil, fmt.Errorf("client only allows requests to URLs matching the clients configured base URL. Got %q, base URL is %q", u.String(), c.baseURL.String())
	}

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	if c.UserAgent != "" {
		reqHeaders.Set("User-Agent", c.UserAgent)
	}

	var body any
	switch {
	case method == http.MethodPatch || method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		if opt != nil {
			b, err := json.Marshal(opt)
			if err != nil {
				return nil, err
			}
			body = b
		}
	case opt != nil:
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req, err := retryablehttp.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	for _, fn := range append(c.defaultRequestOptions, options...) {
		if fn == nil {
			continue
		}
		if err := fn(req); err != nil {
			return nil, err
		}
	}

	// Set the request specific headers.
	maps.Copy(req.Header, reqHeaders)

	return req, nil
}

// UploadRequest creates an API request for uploading a file. The method
// expects a relative URL path that will be resolved relative to the base
// URL of the Client. Relative URL paths should always be specified without
// a preceding slash. If specified, the value pointed to by body is JSON
// encoded and included as the request body.
func (c *Client) UploadRequest(method, path string, content io.Reader, filename string, uploadType UploadType, opt any, options []RequestOptionFunc) (*retryablehttp.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	if c.UserAgent != "" {
		reqHeaders.Set("User-Agent", c.UserAgent)
	}

	b := new(bytes.Buffer)
	w := multipart.NewWriter(b)

	fw, err := w.CreateFormFile(string(uploadType), filename)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(fw, content); err != nil {
		return nil, err
	}

	if opt != nil {
		fields, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		for name := range fields {
			if err = w.WriteField(name, fmt.Sprintf("%v", fields.Get(name))); err != nil {
				return nil, err
			}
		}
	}

	if err = w.Close(); err != nil {
		return nil, err
	}

	reqHeaders.Set("Content-Type", w.FormDataContentType())

	req, err := retryablehttp.NewRequest(method, u.String(), b)
	if err != nil {
		return nil, err
	}

	for _, fn := range append(c.defaultRequestOptions, options...) {
		if fn == nil {
			continue
		}
		if err := fn(req); err != nil {
			return nil, err
		}
	}

	// Set the request specific headers.
	maps.Copy(req.Header, reqHeaders)

	return req, nil
}

// Response is a GitLab API response. This wraps the standard http.Response
// returned from GitLab and provides convenient access to things like
// pagination links.
type Response struct {
	*http.Response

	// Fields used for offset-based pagination.
	TotalItems   int
	TotalPages   int
	ItemsPerPage int
	CurrentPage  int
	NextPage     int
	PreviousPage int

	// Fields used for keyset-based pagination.
	PreviousLink string
	NextLink     string
	FirstLink    string
	LastLink     string
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.populatePageValues()
	response.populateLinkValues()
	return response
}

const (
	// Headers used for offset-based pagination.
	xTotal      = "X-Total"
	xTotalPages = "X-Total-Pages"
	xPerPage    = "X-Per-Page"
	xPage       = "X-Page"
	xNextPage   = "X-Next-Page"
	xPrevPage   = "X-Prev-Page"

	// Headers used for keyset-based pagination.
	linkPrev  = "prev"
	linkNext  = "next"
	linkFirst = "first"
	linkLast  = "last"
)

// populatePageValues parses the HTTP Link response headers and populates the
// various pagination link values in the Response.
func (r *Response) populatePageValues() {
	if totalItems := r.Header.Get(xTotal); totalItems != "" {
		r.TotalItems, _ = strconv.Atoi(totalItems)
	}
	if totalPages := r.Header.Get(xTotalPages); totalPages != "" {
		r.TotalPages, _ = strconv.Atoi(totalPages)
	}
	if itemsPerPage := r.Header.Get(xPerPage); itemsPerPage != "" {
		r.ItemsPerPage, _ = strconv.Atoi(itemsPerPage)
	}
	if currentPage := r.Header.Get(xPage); currentPage != "" {
		r.CurrentPage, _ = strconv.Atoi(currentPage)
	}
	if nextPage := r.Header.Get(xNextPage); nextPage != "" {
		r.NextPage, _ = strconv.Atoi(nextPage)
	}
	if previousPage := r.Header.Get(xPrevPage); previousPage != "" {
		r.PreviousPage, _ = strconv.Atoi(previousPage)
	}
}

func (r *Response) populateLinkValues() {
	if link := r.Header.Get("Link"); link != "" {
		for _, link := range strings.Split(link, ",") {
			parts := strings.Split(link, ";")
			if len(parts) < 2 {
				continue
			}

			linkType := strings.Trim(strings.Split(parts[1], "=")[1], "\"")
			linkValue := strings.Trim(parts[0], "< >")

			switch linkType {
			case linkPrev:
				r.PreviousLink = linkValue
			case linkNext:
				r.NextLink = linkValue
			case linkFirst:
				r.FirstLink = linkValue
			case linkLast:
				r.LastLink = linkValue
			}
		}
	}
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *retryablehttp.Request, v any) (*Response, error) {
	// Wait will block until the limiter can obtain a new token.
	err := c.limiter.Wait(req.Context())
	if err != nil {
		return nil, err
	}

	c.authSourceInit.Do(func() {
		err = c.authSource.Init(req.Context(), c)
	})
	if err != nil {
		return nil, fmt.Errorf("initializing token source failed: %w", err)
	}

	authKey, authValue, err := c.authSource.Header(req.Context())
	if err != nil {
		return nil, err
	}

	if v := req.Header.Values(authKey); len(v) == 0 {
		req.Header.Set(authKey, authValue)
	}

	client := c.client

	if cr := checkRetryFromContext(req.Context()); cr != nil {
		// for avoid overwriting c.client. Use copy of c.client and apply checkRetry from request context
		client = c.newRetryableHTTPClientWithRetryCheck(cr)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	// If not yet configured, try to configure the rate limiter
	// using the response headers we just received. Fail silently
	// so the limiter will remain disabled in case of an error.
	c.configureLimiterOnce.Do(func() { c.configureLimiter(req.Context(), resp.Header) })

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further.
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return response, err
}

func (c *Client) endpoint() oauth2.Endpoint {
	baseURL := strings.TrimSuffix(c.baseURL.String(), apiVersionPath)

	return oauth2.Endpoint{
		AuthURL:       baseURL + "oauth/authorize",
		TokenURL:      baseURL + "oauth/token",
		DeviceAuthURL: baseURL + "oauth/authorize_device",
	}
}

// ErrInvalidIDType is returned when a function expecting an ID as either an integer
// or string receives a different type. This error commonly occurs when working with
// GitLab resources like groups and projects which support both numeric IDs and
// path-based string identifiers.
var ErrInvalidIDType = errors.New("the ID must be an int or a string")

// Helper function to accept and format both the project ID or name as project
// identifier for all API calls.
func parseID(id any) (string, error) {
	switch v := id.(type) {
	case int:
		return strconv.Itoa(v), nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("invalid ID type %#v, %w", id, ErrInvalidIDType)
	}
}

// PathEscape is a helper function to escape a project identifier.
func PathEscape(s string) string {
	return strings.ReplaceAll(url.PathEscape(s), ".", "%2E")
}

// An ErrorResponse reports one or more errors caused by an API request.
//
// GitLab API docs:
// https://docs.gitlab.com/api/rest/troubleshooting/
type ErrorResponse struct {
	Body     []byte
	Response *http.Response
	Message  string
}

func (e *ErrorResponse) Error() string {
	path := e.Response.Request.URL.RawPath
	if path == "" {
		path = e.Response.Request.URL.Path
	}
	url := fmt.Sprintf("%s://%s%s", e.Response.Request.URL.Scheme, e.Response.Request.URL.Host, path)

	if e.Message == "" {
		return fmt.Sprintf("%s %s: %d", e.Response.Request.Method, url, e.Response.StatusCode)
	} else {
		return fmt.Sprintf("%s %s: %d %s", e.Response.Request.Method, url, e.Response.StatusCode, e.Message)
	}
}

func (e *ErrorResponse) HasStatusCode(statusCode int) bool {
	return e != nil && e.Response != nil && e.Response.StatusCode == statusCode
}

// CheckResponse checks the API response for errors, and returns them if present.
func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	case 404:
		return ErrNotFound
	}

	errorResponse := &ErrorResponse{Response: r}

	data, err := io.ReadAll(r.Body)
	if err == nil && strings.TrimSpace(string(data)) != "" {
		errorResponse.Body = data

		var raw any
		if err := json.Unmarshal(data, &raw); err != nil {
			errorResponse.Message = fmt.Sprintf("failed to parse unknown error format: %s", data)
		} else {
			errorResponse.Message = parseError(raw)
		}
	}

	return errorResponse
}

// Format:
//
//	{
//	    "message": {
//	        "<property-name>": [
//	            "<error-message>",
//	            "<error-message>",
//	            ...
//	        ],
//	        "<embed-entity>": {
//	            "<property-name>": [
//	                "<error-message>",
//	                "<error-message>",
//	                ...
//	            ],
//	        }
//	    },
//	    "error": "<error-message>"
//	}
func parseError(raw any) string {
	switch raw := raw.(type) {
	case string:
		return raw

	case []any:
		var errs []string
		for _, v := range raw {
			errs = append(errs, parseError(v))
		}
		return fmt.Sprintf("[%s]", strings.Join(errs, ", "))

	case map[string]any:
		var errs []string
		for k, v := range raw {
			errs = append(errs, fmt.Sprintf("{%s: %s}", k, parseError(v)))
		}
		sort.Strings(errs)
		return strings.Join(errs, ", ")

	default:
		return fmt.Sprintf("failed to parse unexpected error type: %T", raw)
	}
}

func HasStatusCode(err error, statusCode int) bool {
	var errResponse *ErrorResponse
	if !errors.As(err, &errResponse) {
		return false
	}

	return errResponse.HasStatusCode(statusCode)
}

// newRetryableHTTPClientWithRetryCheck returns a `retryablehttp.Client` clone of itself with the given CheckRetry function
func (c *Client) newRetryableHTTPClientWithRetryCheck(cr retryablehttp.CheckRetry) *retryablehttp.Client {
	return &retryablehttp.Client{
		HTTPClient:     c.client.HTTPClient,
		Logger:         c.client.Logger,
		RetryWaitMin:   c.client.RetryWaitMin,
		RetryWaitMax:   c.client.RetryWaitMax,
		RetryMax:       c.client.RetryMax,
		RequestLogHook: c.client.RequestLogHook,
		CheckRetry:     cr,
		Backoff:        c.client.Backoff,
		ErrorHandler:   c.client.ErrorHandler,
		PrepareRetry:   c.client.PrepareRetry,
	}
}

// AuthSource is used to obtain access tokens.
type AuthSource interface {
	// Init is called once before making any requests.
	// If the token source needs access to client to initialize itself, it should do so here.
	Init(context.Context, *Client) error

	// Header returns an authentication header. When no error is returned, the
	// key and value should never be empty.
	Header(ctx context.Context) (key, value string, err error)
}

// OAuthTokenSource wraps an oauth2.TokenSource to implement the AuthSource interface.
type OAuthTokenSource struct {
	TokenSource oauth2.TokenSource
}

func (OAuthTokenSource) Init(context.Context, *Client) error {
	return nil
}

func (as OAuthTokenSource) Header(_ context.Context) (string, string, error) {
	t, err := as.TokenSource.Token()
	if err != nil {
		return "", "", err
	}

	return "Authorization", "Bearer " + t.AccessToken, nil
}

// JobTokenAuthSource used as an AuthSource for CI Job Tokens
type JobTokenAuthSource struct {
	Token string
}

func (JobTokenAuthSource) Init(context.Context, *Client) error {
	return nil
}

func (s JobTokenAuthSource) Header(_ context.Context) (string, string, error) {
	return JobTokenHeaderName, s.Token, nil
}

// AccessTokenAuthSource used as an AuthSource for various access tokens, like Personal-, Project- and Group- Access Tokens.
// Can be used for all tokens that authorize with the Private-Token header.
type AccessTokenAuthSource struct {
	Token string
}

func (AccessTokenAuthSource) Init(context.Context, *Client) error {
	return nil
}

func (s AccessTokenAuthSource) Header(_ context.Context) (string, string, error) {
	return AccessTokenHeaderName, s.Token, nil
}

// PasswordCredentialsAuthSource implements the AuthSource interface for the OAuth 2.0
// resource owner password credentials flow.
type PasswordCredentialsAuthSource struct {
	Username string
	Password string

	AuthSource
}

func (as *PasswordCredentialsAuthSource) Init(ctx context.Context, client *Client) error {
	ctx = context.WithValue(ctx, oauth2.HTTPClient, client.client.HTTPClient)

	config := &oauth2.Config{
		Endpoint: client.endpoint(),
	}

	pct, err := config.PasswordCredentialsToken(ctx, as.Username, as.Password)
	if err != nil {
		return fmt.Errorf("PasswordCredentialsToken(%q, ******): %w", as.Username, err)
	}

	as.AuthSource = OAuthTokenSource{
		config.TokenSource(ctx, pct),
	}

	return nil
}
