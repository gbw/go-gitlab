// This file is generate from scripts/generate_service_interface_map.sh
package gitlab

var serviceMap = map[any]any{
	&AccessRequestsService{}:                   (*AccessRequestsServiceInterface)(nil),
	&AppearanceService{}:                       (*AppearanceServiceInterface)(nil),
	&ApplicationStatisticsService{}:            (*ApplicationStatisticsServiceInterface)(nil),
	&ApplicationsService{}:                     (*ApplicationsServiceInterface)(nil),
	&AuditEventsService{}:                      (*AuditEventsServiceInterface)(nil),
	&AvatarRequestsService{}:                   (*AvatarRequestsServiceInterface)(nil),
	&AwardEmojiService{}:                       (*AwardEmojiServiceInterface)(nil),
	&BranchesService{}:                         (*BranchesServiceInterface)(nil),
	&BroadcastMessagesService{}:                (*BroadcastMessagesServiceInterface)(nil),
	&BulkImportsService{}:                      (*BulkImportsServiceInterface)(nil),
	&CIYMLTemplatesService{}:                   (*CIYMLTemplatesServiceInterface)(nil),
	&ClusterAgentsService{}:                    (*ClusterAgentsServiceInterface)(nil),
	&CommitsService{}:                          (*CommitsServiceInterface)(nil),
	&ContainerRegistryProtectionRulesService{}: (*ContainerRegistryProtectionRulesServiceInterface)(nil),
	&ContainerRegistryService{}:                (*ContainerRegistryServiceInterface)(nil),
	&CustomAttributesService{}:                 (*CustomAttributesServiceInterface)(nil),
	&DORAMetricsService{}:                      (*DORAMetricsServiceInterface)(nil),
	&DependencyListExportService{}:             (*DependencyListExportServiceInterface)(nil),
	&DependencyProxyService{}:                  (*DependencyProxyServiceInterface)(nil),
	&DeployKeysService{}:                       (*DeployKeysServiceInterface)(nil),
	&DeployTokensService{}:                     (*DeployTokensServiceInterface)(nil),
	&DeploymentMergeRequestsService{}:          (*DeploymentMergeRequestsServiceInterface)(nil),
	&DeploymentsService{}:                      (*DeploymentsServiceInterface)(nil),
	&DiscussionsService{}:                      (*DiscussionsServiceInterface)(nil),
	&DockerfileTemplatesService{}:              (*DockerfileTemplatesServiceInterface)(nil),
	&DraftNotesService{}:                       (*DraftNotesServiceInterface)(nil),
	&EnterpriseUsersService{}:                  (*EnterpriseUsersServiceInterface)(nil),
	&EnvironmentsService{}:                     (*EnvironmentsServiceInterface)(nil),
	&EpicIssuesService{}:                       (*EpicIssuesServiceInterface)(nil),
	&EpicsService{}:                            (*EpicsServiceInterface)(nil),
	&ErrorTrackingService{}:                    (*ErrorTrackingServiceInterface)(nil),
	&EventsService{}:                           (*EventsServiceInterface)(nil),
	&ExternalStatusChecksService{}:             (*ExternalStatusChecksServiceInterface)(nil),
	&FeatureFlagUserListsService{}:             (*FeatureFlagUserListsServiceInterface)(nil),
	&FeaturesService{}:                         (*FeaturesServiceInterface)(nil),
	&FreezePeriodsService{}:                    (*FreezePeriodsServiceInterface)(nil),
	&GenericPackagesService{}:                  (*GenericPackagesServiceInterface)(nil),
	&GeoNodesService{}:                         (*GeoNodesServiceInterface)(nil),
	&GeoSitesService{}:                         (*GeoSitesServiceInterface)(nil),
	&GitIgnoreTemplatesService{}:               (*GitIgnoreTemplatesServiceInterface)(nil),
	&GroupAccessTokensService{}:                (*GroupAccessTokensServiceInterface)(nil),
	&GroupActivityAnalyticsService{}:           (*GroupActivityAnalyticsServiceInterface)(nil),
	&GroupBadgesService{}:                      (*GroupBadgesServiceInterface)(nil),
	&GroupClustersService{}:                    (*GroupClustersServiceInterface)(nil),
	&GroupEpicBoardsService{}:                  (*GroupEpicBoardsServiceInterface)(nil),
	&GroupImportExportService{}:                (*GroupImportExportServiceInterface)(nil),
	&GroupIssueBoardsService{}:                 (*GroupIssueBoardsServiceInterface)(nil),
	&GroupIterationsService{}:                  (*GroupIterationsServiceInterface)(nil),
	&GroupLabelsService{}:                      (*GroupLabelsServiceInterface)(nil),
	&GroupMembersService{}:                     (*GroupMembersServiceInterface)(nil),
	&GroupMilestonesService{}:                  (*GroupMilestonesServiceInterface)(nil),
	&GroupProtectedEnvironmentsService{}:       (*GroupProtectedEnvironmentsServiceInterface)(nil),
	&GroupReleasesService{}:                    (*GroupReleasesServiceInterface)(nil),
	&GroupRepositoryStorageMoveService{}:       (*GroupRepositoryStorageMoveServiceInterface)(nil),
	&GroupSCIMService{}:                        (*GroupSCIMServiceInterface)(nil),
	&GroupSSHCertificatesService{}:             (*GroupSSHCertificatesServiceInterface)(nil),
	&GroupSecuritySettingsService{}:            (*GroupSecuritySettingsServiceInterface)(nil),
	&GroupVariablesService{}:                   (*GroupVariablesServiceInterface)(nil),
	&GroupWikisService{}:                       (*GroupWikisServiceInterface)(nil),
	&GroupsService{}:                           (*GroupsServiceInterface)(nil),
	&ImportService{}:                           (*ImportServiceInterface)(nil),
	&InstanceClustersService{}:                 (*InstanceClustersServiceInterface)(nil),
	&InstanceVariablesService{}:                (*InstanceVariablesServiceInterface)(nil),
	&InvitesService{}:                          (*InvitesServiceInterface)(nil),
	&IssueBoardsService{}:                      (*IssueBoardsServiceInterface)(nil),
	&IssueLinksService{}:                       (*IssueLinksServiceInterface)(nil),
	&IssuesService{}:                           (*IssuesServiceInterface)(nil),
	&IssuesStatisticsService{}:                 (*IssuesStatisticsServiceInterface)(nil),
	&JobTokenScopeService{}:                    (*JobTokenScopeServiceInterface)(nil),
	&JobsService{}:                             (*JobsServiceInterface)(nil),
	&KeysService{}:                             (*KeysServiceInterface)(nil),
	&LabelsService{}:                           (*LabelsServiceInterface)(nil),
	&LicenseService{}:                          (*LicenseServiceInterface)(nil),
	&LicenseTemplatesService{}:                 (*LicenseTemplatesServiceInterface)(nil),
	&ManagedLicensesService{}:                  (*ManagedLicensesServiceInterface)(nil),
	&MarkdownService{}:                         (*MarkdownServiceInterface)(nil),
	&MemberRolesService{}:                      (*MemberRolesServiceInterface)(nil),
	&MergeRequestApprovalSettingsService{}:     (*MergeRequestApprovalSettingsServiceInterface)(nil),
	&MergeRequestApprovalsService{}:            (*MergeRequestApprovalsServiceInterface)(nil),
	&MergeRequestsService{}:                    (*MergeRequestsServiceInterface)(nil),
	&MergeTrainsService{}:                      (*MergeTrainsServiceInterface)(nil),
	&MetadataService{}:                         (*MetadataServiceInterface)(nil),
	&MilestonesService{}:                       (*MilestonesServiceInterface)(nil),
	&NamespacesService{}:                       (*NamespacesServiceInterface)(nil),
	&NotesService{}:                            (*NotesServiceInterface)(nil),
	&NotificationSettingsService{}:             (*NotificationSettingsServiceInterface)(nil),
	&PackagesService{}:                         (*PackagesServiceInterface)(nil),
	&PagesDomainsService{}:                     (*PagesDomainsServiceInterface)(nil),
	&PagesService{}:                            (*PagesServiceInterface)(nil),
	&PersonalAccessTokensService{}:             (*PersonalAccessTokensServiceInterface)(nil),
	&PipelineSchedulesService{}:                (*PipelineSchedulesServiceInterface)(nil),
	&PipelineTriggersService{}:                 (*PipelineTriggersServiceInterface)(nil),
	&PipelinesService{}:                        (*PipelinesServiceInterface)(nil),
	&PlanLimitsService{}:                       (*PlanLimitsServiceInterface)(nil),
	&ProjectAccessTokensService{}:              (*ProjectAccessTokensServiceInterface)(nil),
	&ProjectBadgesService{}:                    (*ProjectBadgesServiceInterface)(nil),
	&ProjectClustersService{}:                  (*ProjectClustersServiceInterface)(nil),
	&ProjectFeatureFlagService{}:               (*ProjectFeatureFlagServiceInterface)(nil),
	&ProjectImportExportService{}:              (*ProjectImportExportServiceInterface)(nil),
	&ProjectIterationsService{}:                (*ProjectIterationsServiceInterface)(nil),
	&ProjectMarkdownUploadsService{}:           (*ProjectMarkdownUploadsServiceInterface)(nil),
	&ProjectMembersService{}:                   (*ProjectMembersServiceInterface)(nil),
	&ProjectMirrorService{}:                    (*ProjectMirrorServiceInterface)(nil),
	&ProjectRepositoryStorageMoveService{}:     (*ProjectRepositoryStorageMoveServiceInterface)(nil),
	&ProjectSecuritySettingsService{}:          (*ProjectSecuritySettingsServiceInterface)(nil),
	&ProjectSnippetsService{}:                  (*ProjectSnippetsServiceInterface)(nil),
	&ProjectTemplatesService{}:                 (*ProjectTemplatesServiceInterface)(nil),
	&ProjectVariablesService{}:                 (*ProjectVariablesServiceInterface)(nil),
	&ProjectVulnerabilitiesService{}:           (*ProjectVulnerabilitiesServiceInterface)(nil),
	&ProjectsService{}:                         (*ProjectsServiceInterface)(nil),
	&ProtectedBranchesService{}:                (*ProtectedBranchesServiceInterface)(nil),
	&ProtectedEnvironmentsService{}:            (*ProtectedEnvironmentsServiceInterface)(nil),
	&ProtectedTagsService{}:                    (*ProtectedTagsServiceInterface)(nil),
	&ReleaseLinksService{}:                     (*ReleaseLinksServiceInterface)(nil),
	&ReleasesService{}:                         (*ReleasesServiceInterface)(nil),
	&RepositoriesService{}:                     (*RepositoriesServiceInterface)(nil),
	&RepositoryFilesService{}:                  (*RepositoryFilesServiceInterface)(nil),
	&RepositorySubmodulesService{}:             (*RepositorySubmodulesServiceInterface)(nil),
	&ResourceGroupService{}:                    (*ResourceGroupServiceInterface)(nil),
	&ResourceIterationEventsService{}:          (*ResourceIterationEventsServiceInterface)(nil),
	&ResourceLabelEventsService{}:              (*ResourceLabelEventsServiceInterface)(nil),
	&ResourceMilestoneEventsService{}:          (*ResourceMilestoneEventsServiceInterface)(nil),
	&ResourceStateEventsService{}:              (*ResourceStateEventsServiceInterface)(nil),
	&ResourceWeightEventsService{}:             (*ResourceWeightEventsServiceInterface)(nil),
	&RunnersService{}:                          (*RunnersServiceInterface)(nil),
	&SearchService{}:                           (*SearchServiceInterface)(nil),
	&SecureFilesService{}:                      (*SecureFilesServiceInterface)(nil),
	&ServicesService{}:                         (*ServicesServiceInterface)(nil),
	&SettingsService{}:                         (*SettingsServiceInterface)(nil),
	&SidekiqService{}:                          (*SidekiqServiceInterface)(nil),
	&SnippetRepositoryStorageMoveService{}:     (*SnippetRepositoryStorageMoveServiceInterface)(nil),
	&SnippetsService{}:                         (*SnippetsServiceInterface)(nil),
	&SystemHooksService{}:                      (*SystemHooksServiceInterface)(nil),
	&TagsService{}:                             (*TagsServiceInterface)(nil),
	&TodosService{}:                            (*TodosServiceInterface)(nil),
	&TopicsService{}:                           (*TopicsServiceInterface)(nil),
	&UsageDataService{}:                        (*UsageDataServiceInterface)(nil),
	&UsersService{}:                            (*UsersServiceInterface)(nil),
	&ValidateService{}:                         (*ValidateServiceInterface)(nil),
	&VersionService{}:                          (*VersionServiceInterface)(nil),
	&WikisService{}:                            (*WikisServiceInterface)(nil),
}
