// This file is generate from scripts/generate_testing_client.sh
package testing

import (
	"go.uber.org/mock/gomock"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type testClientMocks struct {
	MockGraphQL                          *MockGraphQLInterface
	MockAccessRequests                   *MockAccessRequestsServiceInterface
	MockAlertManagement                  *MockAlertManagementServiceInterface
	MockAppearance                       *MockAppearanceServiceInterface
	MockApplications                     *MockApplicationsServiceInterface
	MockApplicationStatistics            *MockApplicationStatisticsServiceInterface
	MockAuditEvents                      *MockAuditEventsServiceInterface
	MockAvatar                           *MockAvatarRequestsServiceInterface
	MockAwardEmoji                       *MockAwardEmojiServiceInterface
	MockBoards                           *MockIssueBoardsServiceInterface
	MockBranches                         *MockBranchesServiceInterface
	MockBroadcastMessage                 *MockBroadcastMessagesServiceInterface
	MockBulkImports                      *MockBulkImportsServiceInterface
	MockCIYMLTemplate                    *MockCIYMLTemplatesServiceInterface
	MockClusterAgents                    *MockClusterAgentsServiceInterface
	MockCommits                          *MockCommitsServiceInterface
	MockContainerRegistry                *MockContainerRegistryServiceInterface
	MockContainerRegistryProtectionRules *MockContainerRegistryProtectionRulesServiceInterface
	MockCustomAttribute                  *MockCustomAttributesServiceInterface
	MockDatabaseMigrations               *MockDatabaseMigrationsServiceInterface
	MockDependencies                     *MockDependenciesServiceInterface
	MockDependencyListExport             *MockDependencyListExportServiceInterface
	MockDependencyProxy                  *MockDependencyProxyServiceInterface
	MockDeployKeys                       *MockDeployKeysServiceInterface
	MockDeployTokens                     *MockDeployTokensServiceInterface
	MockDeploymentMergeRequests          *MockDeploymentMergeRequestsServiceInterface
	MockDeployments                      *MockDeploymentsServiceInterface
	MockDiscussions                      *MockDiscussionsServiceInterface
	MockDockerfileTemplate               *MockDockerfileTemplatesServiceInterface
	MockDORAMetrics                      *MockDORAMetricsServiceInterface
	MockDraftNotes                       *MockDraftNotesServiceInterface
	MockEnterpriseUsers                  *MockEnterpriseUsersServiceInterface
	MockEnvironments                     *MockEnvironmentsServiceInterface
	MockEpicIssues                       *MockEpicIssuesServiceInterface
	MockEpics                            *MockEpicsServiceInterface
	MockErrorTracking                    *MockErrorTrackingServiceInterface
	MockEvents                           *MockEventsServiceInterface
	MockExternalStatusChecks             *MockExternalStatusChecksServiceInterface
	MockFeatureFlagUserLists             *MockFeatureFlagUserListsServiceInterface
	MockFeatures                         *MockFeaturesServiceInterface
	MockFreezePeriods                    *MockFreezePeriodsServiceInterface
	MockGenericPackages                  *MockGenericPackagesServiceInterface
	MockGeoNodes                         *MockGeoNodesServiceInterface
	MockGeoSites                         *MockGeoSitesServiceInterface
	MockGitIgnoreTemplates               *MockGitIgnoreTemplatesServiceInterface
	MockGroupAccessTokens                *MockGroupAccessTokensServiceInterface
	MockGroupActivityAnalytics           *MockGroupActivityAnalyticsServiceInterface
	MockGroupBadges                      *MockGroupBadgesServiceInterface
	MockGroupCluster                     *MockGroupClustersServiceInterface
	MockGroupEpicBoards                  *MockGroupEpicBoardsServiceInterface
	MockGroupImportExport                *MockGroupImportExportServiceInterface
	MockGroupIssueBoards                 *MockGroupIssueBoardsServiceInterface
	MockGroupIterations                  *MockGroupIterationsServiceInterface
	MockGroupLabels                      *MockGroupLabelsServiceInterface
	MockGroupMarkdownUploads             *MockGroupMarkdownUploadsServiceInterface
	MockGroupMembers                     *MockGroupMembersServiceInterface
	MockGroupMilestones                  *MockGroupMilestonesServiceInterface
	MockGroupProtectedEnvironments       *MockGroupProtectedEnvironmentsServiceInterface
	MockGroupReleases                    *MockGroupReleasesServiceInterface
	MockGroupRepositoryStorageMove       *MockGroupRepositoryStorageMoveServiceInterface
	MockGroupSCIM                        *MockGroupSCIMServiceInterface
	MockGroupSecuritySettings            *MockGroupSecuritySettingsServiceInterface
	MockGroupSSHCertificates             *MockGroupSSHCertificatesServiceInterface
	MockGroupVariables                   *MockGroupVariablesServiceInterface
	MockGroupWikis                       *MockGroupWikisServiceInterface
	MockGroups                           *MockGroupsServiceInterface
	MockImport                           *MockImportServiceInterface
	MockInstanceCluster                  *MockInstanceClustersServiceInterface
	MockInstanceVariables                *MockInstanceVariablesServiceInterface
	MockInvites                          *MockInvitesServiceInterface
	MockIssueLinks                       *MockIssueLinksServiceInterface
	MockIssues                           *MockIssuesServiceInterface
	MockIssuesStatistics                 *MockIssuesStatisticsServiceInterface
	MockJobs                             *MockJobsServiceInterface
	MockJobTokenScope                    *MockJobTokenScopeServiceInterface
	MockKeys                             *MockKeysServiceInterface
	MockLabels                           *MockLabelsServiceInterface
	MockLicense                          *MockLicenseServiceInterface
	MockLicenseTemplates                 *MockLicenseTemplatesServiceInterface
	MockMarkdown                         *MockMarkdownServiceInterface
	MockMemberRolesService               *MockMemberRolesServiceInterface
	MockMergeRequestApprovals            *MockMergeRequestApprovalsServiceInterface
	MockMergeRequestApprovalSettings     *MockMergeRequestApprovalSettingsServiceInterface
	MockMergeRequests                    *MockMergeRequestsServiceInterface
	MockMergeTrains                      *MockMergeTrainsServiceInterface
	MockMetadata                         *MockMetadataServiceInterface
	MockMilestones                       *MockMilestonesServiceInterface
	MockNamespaces                       *MockNamespacesServiceInterface
	MockNotes                            *MockNotesServiceInterface
	MockNotificationSettings             *MockNotificationSettingsServiceInterface
	MockPackages                         *MockPackagesServiceInterface
	MockPages                            *MockPagesServiceInterface
	MockPagesDomains                     *MockPagesDomainsServiceInterface
	MockPersonalAccessTokens             *MockPersonalAccessTokensServiceInterface
	MockPipelineSchedules                *MockPipelineSchedulesServiceInterface
	MockPipelineTriggers                 *MockPipelineTriggersServiceInterface
	MockPipelines                        *MockPipelinesServiceInterface
	MockPlanLimits                       *MockPlanLimitsServiceInterface
	MockProjectAccessTokens              *MockProjectAccessTokensServiceInterface
	MockProjectBadges                    *MockProjectBadgesServiceInterface
	MockProjectCluster                   *MockProjectClustersServiceInterface
	MockProjectFeatureFlags              *MockProjectFeatureFlagServiceInterface
	MockProjectImportExport              *MockProjectImportExportServiceInterface
	MockProjectIterations                *MockProjectIterationsServiceInterface
	MockProjectMarkdownUploads           *MockProjectMarkdownUploadsServiceInterface
	MockProjectMembers                   *MockProjectMembersServiceInterface
	MockProjectMirrors                   *MockProjectMirrorServiceInterface
	MockProjectRepositoryStorageMove     *MockProjectRepositoryStorageMoveServiceInterface
	MockProjectSecuritySettings          *MockProjectSecuritySettingsServiceInterface
	MockProjectSnippets                  *MockProjectSnippetsServiceInterface
	MockProjectTemplates                 *MockProjectTemplatesServiceInterface
	MockProjectVariables                 *MockProjectVariablesServiceInterface
	MockProjectVulnerabilities           *MockProjectVulnerabilitiesServiceInterface
	MockProjects                         *MockProjectsServiceInterface
	MockProtectedBranches                *MockProtectedBranchesServiceInterface
	MockProtectedEnvironments            *MockProtectedEnvironmentsServiceInterface
	MockProtectedTags                    *MockProtectedTagsServiceInterface
	MockReleaseLinks                     *MockReleaseLinksServiceInterface
	MockReleases                         *MockReleasesServiceInterface
	MockRepositories                     *MockRepositoriesServiceInterface
	MockRepositoryFiles                  *MockRepositoryFilesServiceInterface
	MockRepositorySubmodules             *MockRepositorySubmodulesServiceInterface
	MockResourceGroup                    *MockResourceGroupServiceInterface
	MockResourceIterationEvents          *MockResourceIterationEventsServiceInterface
	MockResourceLabelEvents              *MockResourceLabelEventsServiceInterface
	MockResourceMilestoneEvents          *MockResourceMilestoneEventsServiceInterface
	MockResourceStateEvents              *MockResourceStateEventsServiceInterface
	MockResourceWeightEvents             *MockResourceWeightEventsServiceInterface
	MockRunners                          *MockRunnersServiceInterface
	MockSearch                           *MockSearchServiceInterface
	MockSecureFiles                      *MockSecureFilesServiceInterface
	MockServices                         *MockServicesServiceInterface
	MockSettings                         *MockSettingsServiceInterface
	MockSidekiq                          *MockSidekiqServiceInterface
	MockSnippetRepositoryStorageMove     *MockSnippetRepositoryStorageMoveServiceInterface
	MockSnippets                         *MockSnippetsServiceInterface
	MockSystemHooks                      *MockSystemHooksServiceInterface
	MockTags                             *MockTagsServiceInterface
	MockTerraformStates                  *MockTerraformStatesServiceInterface
	MockTodos                            *MockTodosServiceInterface
	MockTopics                           *MockTopicsServiceInterface
	MockUsageData                        *MockUsageDataServiceInterface
	MockUsers                            *MockUsersServiceInterface
	MockValidate                         *MockValidateServiceInterface
	MockVersion                          *MockVersionServiceInterface
	MockWikis                            *MockWikisServiceInterface
}

func newTestClientWithCtrl(ctrl *gomock.Controller) *TestClient {
	mockGraphQL := NewMockGraphQLInterface(ctrl)
	mockAccessRequests := NewMockAccessRequestsServiceInterface(ctrl)
	mockAlertManagement := NewMockAlertManagementServiceInterface(ctrl)
	mockAppearance := NewMockAppearanceServiceInterface(ctrl)
	mockApplications := NewMockApplicationsServiceInterface(ctrl)
	mockApplicationStatistics := NewMockApplicationStatisticsServiceInterface(ctrl)
	mockAuditEvents := NewMockAuditEventsServiceInterface(ctrl)
	mockAvatar := NewMockAvatarRequestsServiceInterface(ctrl)
	mockAwardEmoji := NewMockAwardEmojiServiceInterface(ctrl)
	mockBoards := NewMockIssueBoardsServiceInterface(ctrl)
	mockBranches := NewMockBranchesServiceInterface(ctrl)
	mockBroadcastMessage := NewMockBroadcastMessagesServiceInterface(ctrl)
	mockBulkImports := NewMockBulkImportsServiceInterface(ctrl)
	mockCIYMLTemplate := NewMockCIYMLTemplatesServiceInterface(ctrl)
	mockClusterAgents := NewMockClusterAgentsServiceInterface(ctrl)
	mockCommits := NewMockCommitsServiceInterface(ctrl)
	mockContainerRegistry := NewMockContainerRegistryServiceInterface(ctrl)
	mockContainerRegistryProtectionRules := NewMockContainerRegistryProtectionRulesServiceInterface(ctrl)
	mockCustomAttribute := NewMockCustomAttributesServiceInterface(ctrl)
	mockDatabaseMigrations := NewMockDatabaseMigrationsServiceInterface(ctrl)
	mockDependencies := NewMockDependenciesServiceInterface(ctrl)
	mockDependencyListExport := NewMockDependencyListExportServiceInterface(ctrl)
	mockDependencyProxy := NewMockDependencyProxyServiceInterface(ctrl)
	mockDeployKeys := NewMockDeployKeysServiceInterface(ctrl)
	mockDeployTokens := NewMockDeployTokensServiceInterface(ctrl)
	mockDeploymentMergeRequests := NewMockDeploymentMergeRequestsServiceInterface(ctrl)
	mockDeployments := NewMockDeploymentsServiceInterface(ctrl)
	mockDiscussions := NewMockDiscussionsServiceInterface(ctrl)
	mockDockerfileTemplate := NewMockDockerfileTemplatesServiceInterface(ctrl)
	mockDORAMetrics := NewMockDORAMetricsServiceInterface(ctrl)
	mockDraftNotes := NewMockDraftNotesServiceInterface(ctrl)
	mockEnterpriseUsers := NewMockEnterpriseUsersServiceInterface(ctrl)
	mockEnvironments := NewMockEnvironmentsServiceInterface(ctrl)
	mockEpicIssues := NewMockEpicIssuesServiceInterface(ctrl)
	mockEpics := NewMockEpicsServiceInterface(ctrl)
	mockErrorTracking := NewMockErrorTrackingServiceInterface(ctrl)
	mockEvents := NewMockEventsServiceInterface(ctrl)
	mockExternalStatusChecks := NewMockExternalStatusChecksServiceInterface(ctrl)
	mockFeatureFlagUserLists := NewMockFeatureFlagUserListsServiceInterface(ctrl)
	mockFeatures := NewMockFeaturesServiceInterface(ctrl)
	mockFreezePeriods := NewMockFreezePeriodsServiceInterface(ctrl)
	mockGenericPackages := NewMockGenericPackagesServiceInterface(ctrl)
	mockGeoNodes := NewMockGeoNodesServiceInterface(ctrl)
	mockGeoSites := NewMockGeoSitesServiceInterface(ctrl)
	mockGitIgnoreTemplates := NewMockGitIgnoreTemplatesServiceInterface(ctrl)
	mockGroupAccessTokens := NewMockGroupAccessTokensServiceInterface(ctrl)
	mockGroupActivityAnalytics := NewMockGroupActivityAnalyticsServiceInterface(ctrl)
	mockGroupBadges := NewMockGroupBadgesServiceInterface(ctrl)
	mockGroupCluster := NewMockGroupClustersServiceInterface(ctrl)
	mockGroupEpicBoards := NewMockGroupEpicBoardsServiceInterface(ctrl)
	mockGroupImportExport := NewMockGroupImportExportServiceInterface(ctrl)
	mockGroupIssueBoards := NewMockGroupIssueBoardsServiceInterface(ctrl)
	mockGroupIterations := NewMockGroupIterationsServiceInterface(ctrl)
	mockGroupLabels := NewMockGroupLabelsServiceInterface(ctrl)
	mockGroupMarkdownUploads := NewMockGroupMarkdownUploadsServiceInterface(ctrl)
	mockGroupMembers := NewMockGroupMembersServiceInterface(ctrl)
	mockGroupMilestones := NewMockGroupMilestonesServiceInterface(ctrl)
	mockGroupProtectedEnvironments := NewMockGroupProtectedEnvironmentsServiceInterface(ctrl)
	mockGroupReleases := NewMockGroupReleasesServiceInterface(ctrl)
	mockGroupRepositoryStorageMove := NewMockGroupRepositoryStorageMoveServiceInterface(ctrl)
	mockGroupSCIM := NewMockGroupSCIMServiceInterface(ctrl)
	mockGroupSecuritySettings := NewMockGroupSecuritySettingsServiceInterface(ctrl)
	mockGroupSSHCertificates := NewMockGroupSSHCertificatesServiceInterface(ctrl)
	mockGroupVariables := NewMockGroupVariablesServiceInterface(ctrl)
	mockGroupWikis := NewMockGroupWikisServiceInterface(ctrl)
	mockGroups := NewMockGroupsServiceInterface(ctrl)
	mockImport := NewMockImportServiceInterface(ctrl)
	mockInstanceCluster := NewMockInstanceClustersServiceInterface(ctrl)
	mockInstanceVariables := NewMockInstanceVariablesServiceInterface(ctrl)
	mockInvites := NewMockInvitesServiceInterface(ctrl)
	mockIssueLinks := NewMockIssueLinksServiceInterface(ctrl)
	mockIssues := NewMockIssuesServiceInterface(ctrl)
	mockIssuesStatistics := NewMockIssuesStatisticsServiceInterface(ctrl)
	mockJobs := NewMockJobsServiceInterface(ctrl)
	mockJobTokenScope := NewMockJobTokenScopeServiceInterface(ctrl)
	mockKeys := NewMockKeysServiceInterface(ctrl)
	mockLabels := NewMockLabelsServiceInterface(ctrl)
	mockLicense := NewMockLicenseServiceInterface(ctrl)
	mockLicenseTemplates := NewMockLicenseTemplatesServiceInterface(ctrl)
	mockMarkdown := NewMockMarkdownServiceInterface(ctrl)
	mockMemberRolesService := NewMockMemberRolesServiceInterface(ctrl)
	mockMergeRequestApprovals := NewMockMergeRequestApprovalsServiceInterface(ctrl)
	mockMergeRequestApprovalSettings := NewMockMergeRequestApprovalSettingsServiceInterface(ctrl)
	mockMergeRequests := NewMockMergeRequestsServiceInterface(ctrl)
	mockMergeTrains := NewMockMergeTrainsServiceInterface(ctrl)
	mockMetadata := NewMockMetadataServiceInterface(ctrl)
	mockMilestones := NewMockMilestonesServiceInterface(ctrl)
	mockNamespaces := NewMockNamespacesServiceInterface(ctrl)
	mockNotes := NewMockNotesServiceInterface(ctrl)
	mockNotificationSettings := NewMockNotificationSettingsServiceInterface(ctrl)
	mockPackages := NewMockPackagesServiceInterface(ctrl)
	mockPages := NewMockPagesServiceInterface(ctrl)
	mockPagesDomains := NewMockPagesDomainsServiceInterface(ctrl)
	mockPersonalAccessTokens := NewMockPersonalAccessTokensServiceInterface(ctrl)
	mockPipelineSchedules := NewMockPipelineSchedulesServiceInterface(ctrl)
	mockPipelineTriggers := NewMockPipelineTriggersServiceInterface(ctrl)
	mockPipelines := NewMockPipelinesServiceInterface(ctrl)
	mockPlanLimits := NewMockPlanLimitsServiceInterface(ctrl)
	mockProjectAccessTokens := NewMockProjectAccessTokensServiceInterface(ctrl)
	mockProjectBadges := NewMockProjectBadgesServiceInterface(ctrl)
	mockProjectCluster := NewMockProjectClustersServiceInterface(ctrl)
	mockProjectFeatureFlags := NewMockProjectFeatureFlagServiceInterface(ctrl)
	mockProjectImportExport := NewMockProjectImportExportServiceInterface(ctrl)
	mockProjectIterations := NewMockProjectIterationsServiceInterface(ctrl)
	mockProjectMarkdownUploads := NewMockProjectMarkdownUploadsServiceInterface(ctrl)
	mockProjectMembers := NewMockProjectMembersServiceInterface(ctrl)
	mockProjectMirrors := NewMockProjectMirrorServiceInterface(ctrl)
	mockProjectRepositoryStorageMove := NewMockProjectRepositoryStorageMoveServiceInterface(ctrl)
	mockProjectSecuritySettings := NewMockProjectSecuritySettingsServiceInterface(ctrl)
	mockProjectSnippets := NewMockProjectSnippetsServiceInterface(ctrl)
	mockProjectTemplates := NewMockProjectTemplatesServiceInterface(ctrl)
	mockProjectVariables := NewMockProjectVariablesServiceInterface(ctrl)
	mockProjectVulnerabilities := NewMockProjectVulnerabilitiesServiceInterface(ctrl)
	mockProjects := NewMockProjectsServiceInterface(ctrl)
	mockProtectedBranches := NewMockProtectedBranchesServiceInterface(ctrl)
	mockProtectedEnvironments := NewMockProtectedEnvironmentsServiceInterface(ctrl)
	mockProtectedTags := NewMockProtectedTagsServiceInterface(ctrl)
	mockReleaseLinks := NewMockReleaseLinksServiceInterface(ctrl)
	mockReleases := NewMockReleasesServiceInterface(ctrl)
	mockRepositories := NewMockRepositoriesServiceInterface(ctrl)
	mockRepositoryFiles := NewMockRepositoryFilesServiceInterface(ctrl)
	mockRepositorySubmodules := NewMockRepositorySubmodulesServiceInterface(ctrl)
	mockResourceGroup := NewMockResourceGroupServiceInterface(ctrl)
	mockResourceIterationEvents := NewMockResourceIterationEventsServiceInterface(ctrl)
	mockResourceLabelEvents := NewMockResourceLabelEventsServiceInterface(ctrl)
	mockResourceMilestoneEvents := NewMockResourceMilestoneEventsServiceInterface(ctrl)
	mockResourceStateEvents := NewMockResourceStateEventsServiceInterface(ctrl)
	mockResourceWeightEvents := NewMockResourceWeightEventsServiceInterface(ctrl)
	mockRunners := NewMockRunnersServiceInterface(ctrl)
	mockSearch := NewMockSearchServiceInterface(ctrl)
	mockSecureFiles := NewMockSecureFilesServiceInterface(ctrl)
	mockServices := NewMockServicesServiceInterface(ctrl)
	mockSettings := NewMockSettingsServiceInterface(ctrl)
	mockSidekiq := NewMockSidekiqServiceInterface(ctrl)
	mockSnippetRepositoryStorageMove := NewMockSnippetRepositoryStorageMoveServiceInterface(ctrl)
	mockSnippets := NewMockSnippetsServiceInterface(ctrl)
	mockSystemHooks := NewMockSystemHooksServiceInterface(ctrl)
	mockTags := NewMockTagsServiceInterface(ctrl)
	mockTerraformStates := NewMockTerraformStatesServiceInterface(ctrl)
	mockTodos := NewMockTodosServiceInterface(ctrl)
	mockTopics := NewMockTopicsServiceInterface(ctrl)
	mockUsageData := NewMockUsageDataServiceInterface(ctrl)
	mockUsers := NewMockUsersServiceInterface(ctrl)
	mockValidate := NewMockValidateServiceInterface(ctrl)
	mockVersion := NewMockVersionServiceInterface(ctrl)
	mockWikis := NewMockWikisServiceInterface(ctrl)

	return &TestClient{
		Client: &gitlab.Client{
			GraphQL:                          mockGraphQL,
			AccessRequests:                   mockAccessRequests,
			AlertManagement:                  mockAlertManagement,
			Appearance:                       mockAppearance,
			Applications:                     mockApplications,
			ApplicationStatistics:            mockApplicationStatistics,
			AuditEvents:                      mockAuditEvents,
			Avatar:                           mockAvatar,
			AwardEmoji:                       mockAwardEmoji,
			Boards:                           mockBoards,
			Branches:                         mockBranches,
			BroadcastMessage:                 mockBroadcastMessage,
			BulkImports:                      mockBulkImports,
			CIYMLTemplate:                    mockCIYMLTemplate,
			ClusterAgents:                    mockClusterAgents,
			Commits:                          mockCommits,
			ContainerRegistry:                mockContainerRegistry,
			ContainerRegistryProtectionRules: mockContainerRegistryProtectionRules,
			CustomAttribute:                  mockCustomAttribute,
			DatabaseMigrations:               mockDatabaseMigrations,
			Dependencies:                     mockDependencies,
			DependencyListExport:             mockDependencyListExport,
			DependencyProxy:                  mockDependencyProxy,
			DeployKeys:                       mockDeployKeys,
			DeployTokens:                     mockDeployTokens,
			DeploymentMergeRequests:          mockDeploymentMergeRequests,
			Deployments:                      mockDeployments,
			Discussions:                      mockDiscussions,
			DockerfileTemplate:               mockDockerfileTemplate,
			DORAMetrics:                      mockDORAMetrics,
			DraftNotes:                       mockDraftNotes,
			EnterpriseUsers:                  mockEnterpriseUsers,
			Environments:                     mockEnvironments,
			EpicIssues:                       mockEpicIssues,
			Epics:                            mockEpics,
			ErrorTracking:                    mockErrorTracking,
			Events:                           mockEvents,
			ExternalStatusChecks:             mockExternalStatusChecks,
			FeatureFlagUserLists:             mockFeatureFlagUserLists,
			Features:                         mockFeatures,
			FreezePeriods:                    mockFreezePeriods,
			GenericPackages:                  mockGenericPackages,
			GeoNodes:                         mockGeoNodes,
			GeoSites:                         mockGeoSites,
			GitIgnoreTemplates:               mockGitIgnoreTemplates,
			GroupAccessTokens:                mockGroupAccessTokens,
			GroupActivityAnalytics:           mockGroupActivityAnalytics,
			GroupBadges:                      mockGroupBadges,
			GroupCluster:                     mockGroupCluster,
			GroupEpicBoards:                  mockGroupEpicBoards,
			GroupImportExport:                mockGroupImportExport,
			GroupIssueBoards:                 mockGroupIssueBoards,
			GroupIterations:                  mockGroupIterations,
			GroupLabels:                      mockGroupLabels,
			GroupMarkdownUploads:             mockGroupMarkdownUploads,
			GroupMembers:                     mockGroupMembers,
			GroupMilestones:                  mockGroupMilestones,
			GroupProtectedEnvironments:       mockGroupProtectedEnvironments,
			GroupReleases:                    mockGroupReleases,
			GroupRepositoryStorageMove:       mockGroupRepositoryStorageMove,
			GroupSCIM:                        mockGroupSCIM,
			GroupSecuritySettings:            mockGroupSecuritySettings,
			GroupSSHCertificates:             mockGroupSSHCertificates,
			GroupVariables:                   mockGroupVariables,
			GroupWikis:                       mockGroupWikis,
			Groups:                           mockGroups,
			Import:                           mockImport,
			InstanceCluster:                  mockInstanceCluster,
			InstanceVariables:                mockInstanceVariables,
			Invites:                          mockInvites,
			IssueLinks:                       mockIssueLinks,
			Issues:                           mockIssues,
			IssuesStatistics:                 mockIssuesStatistics,
			Jobs:                             mockJobs,
			JobTokenScope:                    mockJobTokenScope,
			Keys:                             mockKeys,
			Labels:                           mockLabels,
			License:                          mockLicense,
			LicenseTemplates:                 mockLicenseTemplates,
			Markdown:                         mockMarkdown,
			MemberRolesService:               mockMemberRolesService,
			MergeRequestApprovals:            mockMergeRequestApprovals,
			MergeRequestApprovalSettings:     mockMergeRequestApprovalSettings,
			MergeRequests:                    mockMergeRequests,
			MergeTrains:                      mockMergeTrains,
			Metadata:                         mockMetadata,
			Milestones:                       mockMilestones,
			Namespaces:                       mockNamespaces,
			Notes:                            mockNotes,
			NotificationSettings:             mockNotificationSettings,
			Packages:                         mockPackages,
			Pages:                            mockPages,
			PagesDomains:                     mockPagesDomains,
			PersonalAccessTokens:             mockPersonalAccessTokens,
			PipelineSchedules:                mockPipelineSchedules,
			PipelineTriggers:                 mockPipelineTriggers,
			Pipelines:                        mockPipelines,
			PlanLimits:                       mockPlanLimits,
			ProjectAccessTokens:              mockProjectAccessTokens,
			ProjectBadges:                    mockProjectBadges,
			ProjectCluster:                   mockProjectCluster,
			ProjectFeatureFlags:              mockProjectFeatureFlags,
			ProjectImportExport:              mockProjectImportExport,
			ProjectIterations:                mockProjectIterations,
			ProjectMarkdownUploads:           mockProjectMarkdownUploads,
			ProjectMembers:                   mockProjectMembers,
			ProjectMirrors:                   mockProjectMirrors,
			ProjectRepositoryStorageMove:     mockProjectRepositoryStorageMove,
			ProjectSecuritySettings:          mockProjectSecuritySettings,
			ProjectSnippets:                  mockProjectSnippets,
			ProjectTemplates:                 mockProjectTemplates,
			ProjectVariables:                 mockProjectVariables,
			ProjectVulnerabilities:           mockProjectVulnerabilities,
			Projects:                         mockProjects,
			ProtectedBranches:                mockProtectedBranches,
			ProtectedEnvironments:            mockProtectedEnvironments,
			ProtectedTags:                    mockProtectedTags,
			ReleaseLinks:                     mockReleaseLinks,
			Releases:                         mockReleases,
			Repositories:                     mockRepositories,
			RepositoryFiles:                  mockRepositoryFiles,
			RepositorySubmodules:             mockRepositorySubmodules,
			ResourceGroup:                    mockResourceGroup,
			ResourceIterationEvents:          mockResourceIterationEvents,
			ResourceLabelEvents:              mockResourceLabelEvents,
			ResourceMilestoneEvents:          mockResourceMilestoneEvents,
			ResourceStateEvents:              mockResourceStateEvents,
			ResourceWeightEvents:             mockResourceWeightEvents,
			Runners:                          mockRunners,
			Search:                           mockSearch,
			SecureFiles:                      mockSecureFiles,
			Services:                         mockServices,
			Settings:                         mockSettings,
			Sidekiq:                          mockSidekiq,
			SnippetRepositoryStorageMove:     mockSnippetRepositoryStorageMove,
			Snippets:                         mockSnippets,
			SystemHooks:                      mockSystemHooks,
			Tags:                             mockTags,
			TerraformStates:                  mockTerraformStates,
			Todos:                            mockTodos,
			Topics:                           mockTopics,
			UsageData:                        mockUsageData,
			Users:                            mockUsers,
			Validate:                         mockValidate,
			Version:                          mockVersion,
			Wikis:                            mockWikis,
		},
		testClientMocks: &testClientMocks{
			MockGraphQL:                          mockGraphQL,
			MockAccessRequests:                   mockAccessRequests,
			MockAlertManagement:                  mockAlertManagement,
			MockAppearance:                       mockAppearance,
			MockApplications:                     mockApplications,
			MockApplicationStatistics:            mockApplicationStatistics,
			MockAuditEvents:                      mockAuditEvents,
			MockAvatar:                           mockAvatar,
			MockAwardEmoji:                       mockAwardEmoji,
			MockBoards:                           mockBoards,
			MockBranches:                         mockBranches,
			MockBroadcastMessage:                 mockBroadcastMessage,
			MockBulkImports:                      mockBulkImports,
			MockCIYMLTemplate:                    mockCIYMLTemplate,
			MockClusterAgents:                    mockClusterAgents,
			MockCommits:                          mockCommits,
			MockContainerRegistry:                mockContainerRegistry,
			MockContainerRegistryProtectionRules: mockContainerRegistryProtectionRules,
			MockCustomAttribute:                  mockCustomAttribute,
			MockDatabaseMigrations:               mockDatabaseMigrations,
			MockDependencies:                     mockDependencies,
			MockDependencyListExport:             mockDependencyListExport,
			MockDependencyProxy:                  mockDependencyProxy,
			MockDeployKeys:                       mockDeployKeys,
			MockDeployTokens:                     mockDeployTokens,
			MockDeploymentMergeRequests:          mockDeploymentMergeRequests,
			MockDeployments:                      mockDeployments,
			MockDiscussions:                      mockDiscussions,
			MockDockerfileTemplate:               mockDockerfileTemplate,
			MockDORAMetrics:                      mockDORAMetrics,
			MockDraftNotes:                       mockDraftNotes,
			MockEnterpriseUsers:                  mockEnterpriseUsers,
			MockEnvironments:                     mockEnvironments,
			MockEpicIssues:                       mockEpicIssues,
			MockEpics:                            mockEpics,
			MockErrorTracking:                    mockErrorTracking,
			MockEvents:                           mockEvents,
			MockExternalStatusChecks:             mockExternalStatusChecks,
			MockFeatureFlagUserLists:             mockFeatureFlagUserLists,
			MockFeatures:                         mockFeatures,
			MockFreezePeriods:                    mockFreezePeriods,
			MockGenericPackages:                  mockGenericPackages,
			MockGeoNodes:                         mockGeoNodes,
			MockGeoSites:                         mockGeoSites,
			MockGitIgnoreTemplates:               mockGitIgnoreTemplates,
			MockGroupAccessTokens:                mockGroupAccessTokens,
			MockGroupActivityAnalytics:           mockGroupActivityAnalytics,
			MockGroupBadges:                      mockGroupBadges,
			MockGroupCluster:                     mockGroupCluster,
			MockGroupEpicBoards:                  mockGroupEpicBoards,
			MockGroupImportExport:                mockGroupImportExport,
			MockGroupIssueBoards:                 mockGroupIssueBoards,
			MockGroupIterations:                  mockGroupIterations,
			MockGroupLabels:                      mockGroupLabels,
			MockGroupMarkdownUploads:             mockGroupMarkdownUploads,
			MockGroupMembers:                     mockGroupMembers,
			MockGroupMilestones:                  mockGroupMilestones,
			MockGroupProtectedEnvironments:       mockGroupProtectedEnvironments,
			MockGroupReleases:                    mockGroupReleases,
			MockGroupRepositoryStorageMove:       mockGroupRepositoryStorageMove,
			MockGroupSCIM:                        mockGroupSCIM,
			MockGroupSecuritySettings:            mockGroupSecuritySettings,
			MockGroupSSHCertificates:             mockGroupSSHCertificates,
			MockGroupVariables:                   mockGroupVariables,
			MockGroupWikis:                       mockGroupWikis,
			MockGroups:                           mockGroups,
			MockImport:                           mockImport,
			MockInstanceCluster:                  mockInstanceCluster,
			MockInstanceVariables:                mockInstanceVariables,
			MockInvites:                          mockInvites,
			MockIssueLinks:                       mockIssueLinks,
			MockIssues:                           mockIssues,
			MockIssuesStatistics:                 mockIssuesStatistics,
			MockJobs:                             mockJobs,
			MockJobTokenScope:                    mockJobTokenScope,
			MockKeys:                             mockKeys,
			MockLabels:                           mockLabels,
			MockLicense:                          mockLicense,
			MockLicenseTemplates:                 mockLicenseTemplates,
			MockMarkdown:                         mockMarkdown,
			MockMemberRolesService:               mockMemberRolesService,
			MockMergeRequestApprovals:            mockMergeRequestApprovals,
			MockMergeRequestApprovalSettings:     mockMergeRequestApprovalSettings,
			MockMergeRequests:                    mockMergeRequests,
			MockMergeTrains:                      mockMergeTrains,
			MockMetadata:                         mockMetadata,
			MockMilestones:                       mockMilestones,
			MockNamespaces:                       mockNamespaces,
			MockNotes:                            mockNotes,
			MockNotificationSettings:             mockNotificationSettings,
			MockPackages:                         mockPackages,
			MockPages:                            mockPages,
			MockPagesDomains:                     mockPagesDomains,
			MockPersonalAccessTokens:             mockPersonalAccessTokens,
			MockPipelineSchedules:                mockPipelineSchedules,
			MockPipelineTriggers:                 mockPipelineTriggers,
			MockPipelines:                        mockPipelines,
			MockPlanLimits:                       mockPlanLimits,
			MockProjectAccessTokens:              mockProjectAccessTokens,
			MockProjectBadges:                    mockProjectBadges,
			MockProjectCluster:                   mockProjectCluster,
			MockProjectFeatureFlags:              mockProjectFeatureFlags,
			MockProjectImportExport:              mockProjectImportExport,
			MockProjectIterations:                mockProjectIterations,
			MockProjectMarkdownUploads:           mockProjectMarkdownUploads,
			MockProjectMembers:                   mockProjectMembers,
			MockProjectMirrors:                   mockProjectMirrors,
			MockProjectRepositoryStorageMove:     mockProjectRepositoryStorageMove,
			MockProjectSecuritySettings:          mockProjectSecuritySettings,
			MockProjectSnippets:                  mockProjectSnippets,
			MockProjectTemplates:                 mockProjectTemplates,
			MockProjectVariables:                 mockProjectVariables,
			MockProjectVulnerabilities:           mockProjectVulnerabilities,
			MockProjects:                         mockProjects,
			MockProtectedBranches:                mockProtectedBranches,
			MockProtectedEnvironments:            mockProtectedEnvironments,
			MockProtectedTags:                    mockProtectedTags,
			MockReleaseLinks:                     mockReleaseLinks,
			MockReleases:                         mockReleases,
			MockRepositories:                     mockRepositories,
			MockRepositoryFiles:                  mockRepositoryFiles,
			MockRepositorySubmodules:             mockRepositorySubmodules,
			MockResourceGroup:                    mockResourceGroup,
			MockResourceIterationEvents:          mockResourceIterationEvents,
			MockResourceLabelEvents:              mockResourceLabelEvents,
			MockResourceMilestoneEvents:          mockResourceMilestoneEvents,
			MockResourceStateEvents:              mockResourceStateEvents,
			MockResourceWeightEvents:             mockResourceWeightEvents,
			MockRunners:                          mockRunners,
			MockSearch:                           mockSearch,
			MockSecureFiles:                      mockSecureFiles,
			MockServices:                         mockServices,
			MockSettings:                         mockSettings,
			MockSidekiq:                          mockSidekiq,
			MockSnippetRepositoryStorageMove:     mockSnippetRepositoryStorageMove,
			MockSnippets:                         mockSnippets,
			MockSystemHooks:                      mockSystemHooks,
			MockTags:                             mockTags,
			MockTerraformStates:                  mockTerraformStates,
			MockTodos:                            mockTodos,
			MockTopics:                           mockTopics,
			MockUsageData:                        mockUsageData,
			MockUsers:                            mockUsers,
			MockValidate:                         mockValidate,
			MockVersion:                          mockVersion,
			MockWikis:                            mockWikis,
		},
	}
}
