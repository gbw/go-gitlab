// This file is generate from scripts/generate_mock_api.sh
package testing

//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=access_requests_mock.go -package=testing gitlab.com/gitlab-org/api/client-go AccessRequestsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=appearance_mock.go -package=testing gitlab.com/gitlab-org/api/client-go AppearanceServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=applications_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ApplicationsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=audit_events_mock.go -package=testing gitlab.com/gitlab-org/api/client-go AuditEventsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=avatar_mock.go -package=testing gitlab.com/gitlab-org/api/client-go AvatarRequestsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=award_emojis_mock.go -package=testing gitlab.com/gitlab-org/api/client-go AwardEmojiServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=boards_mock.go -package=testing gitlab.com/gitlab-org/api/client-go IssueBoardsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=branches_mock.go -package=testing gitlab.com/gitlab-org/api/client-go BranchesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=broadcast_messages_mock.go -package=testing gitlab.com/gitlab-org/api/client-go BroadcastMessagesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=bulk_imports_mock.go -package=testing gitlab.com/gitlab-org/api/client-go BulkImportsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=ci_yml_templates_mock.go -package=testing gitlab.com/gitlab-org/api/client-go CIYMLTemplatesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=cluster_agents_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ClusterAgentsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=commits_mock.go -package=testing gitlab.com/gitlab-org/api/client-go CommitsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=container_registry_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ContainerRegistryServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=container_registry_protection_rules_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ContainerRegistryProtectionRulesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=custom_attributes_mock.go -package=testing gitlab.com/gitlab-org/api/client-go CustomAttributesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=database_migrations_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DatabaseMigrationsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=dependency_list_export_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DependencyListExportServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=deploy_keys_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DeployKeysServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=deploy_tokens_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DeployTokensServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=deployments_merge_requests_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DeploymentMergeRequestsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=deployments_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DeploymentsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=discussions_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DiscussionsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=dockerfile_templates_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DockerfileTemplatesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=dora_metrics_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DORAMetricsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=draft_notes_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DraftNotesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=environments_mock.go -package=testing gitlab.com/gitlab-org/api/client-go EnvironmentsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=epic_issues_mock.go -package=testing gitlab.com/gitlab-org/api/client-go EpicIssuesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=epics_mock.go -package=testing gitlab.com/gitlab-org/api/client-go EpicsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=error_tracking_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ErrorTrackingServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=events_mock.go -package=testing gitlab.com/gitlab-org/api/client-go EventsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=external_status_checks_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ExternalStatusChecksServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=feature_flags_mock.go -package=testing gitlab.com/gitlab-org/api/client-go FeaturesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=freeze_periods_mock.go -package=testing gitlab.com/gitlab-org/api/client-go FreezePeriodsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=generic_packages_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GenericPackagesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=geo_nodes_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GeoNodesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=geo_sites_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GeoSitesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=gitignore_templates_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GitIgnoreTemplatesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=graphql_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GraphQLInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_access_tokens_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupAccessTokensServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_badges_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupBadgesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_boards_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupIssueBoardsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_clusters_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupClustersServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_epic_boards_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupEpicBoardsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_import_export_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupImportExportServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_iterations_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupIterationsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_labels_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupLabelsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_members_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupMembersServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_milestones_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupMilestonesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_protected_environments_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupProtectedEnvironmentsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_releases_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupReleasesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_repository_storage_move_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupRepositoryStorageMoveServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_scim_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupSCIMServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_security_settings_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupSecuritySettingsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_ssh_certificates_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupSSHCertificatesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_variables_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupVariablesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=group_wikis_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupWikisServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=groups_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=import_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ImportServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=instance_clusters_mock.go -package=testing gitlab.com/gitlab-org/api/client-go InstanceClustersServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=instance_variables_mock.go -package=testing gitlab.com/gitlab-org/api/client-go InstanceVariablesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=invites_mock.go -package=testing gitlab.com/gitlab-org/api/client-go InvitesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=issue_links_mock.go -package=testing gitlab.com/gitlab-org/api/client-go IssueLinksServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=issues_mock.go -package=testing gitlab.com/gitlab-org/api/client-go IssuesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=issues_statistics_mock.go -package=testing gitlab.com/gitlab-org/api/client-go IssuesStatisticsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=job_token_scope_mock.go -package=testing gitlab.com/gitlab-org/api/client-go JobTokenScopeServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=jobs_mock.go -package=testing gitlab.com/gitlab-org/api/client-go JobsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=keys_mock.go -package=testing gitlab.com/gitlab-org/api/client-go KeysServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=labels_mock.go -package=testing gitlab.com/gitlab-org/api/client-go LabelsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=license_mock.go -package=testing gitlab.com/gitlab-org/api/client-go LicenseServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=license_templates_mock.go -package=testing gitlab.com/gitlab-org/api/client-go LicenseTemplatesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=markdown_mock.go -package=testing gitlab.com/gitlab-org/api/client-go MarkdownServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=member_roles_mock.go -package=testing gitlab.com/gitlab-org/api/client-go MemberRolesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=merge_request_approval_settings_mock.go -package=testing gitlab.com/gitlab-org/api/client-go MergeRequestApprovalSettingsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=merge_request_approvals_mock.go -package=testing gitlab.com/gitlab-org/api/client-go MergeRequestApprovalsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=merge_requests_mock.go -package=testing gitlab.com/gitlab-org/api/client-go MergeRequestsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=merge_trains_mock.go -package=testing gitlab.com/gitlab-org/api/client-go MergeTrainsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=metadata_mock.go -package=testing gitlab.com/gitlab-org/api/client-go MetadataServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=milestones_mock.go -package=testing gitlab.com/gitlab-org/api/client-go MilestonesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=namespaces_mock.go -package=testing gitlab.com/gitlab-org/api/client-go NamespacesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=notes_mock.go -package=testing gitlab.com/gitlab-org/api/client-go NotesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=notifications_mock.go -package=testing gitlab.com/gitlab-org/api/client-go NotificationSettingsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=packages_mock.go -package=testing gitlab.com/gitlab-org/api/client-go PackagesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=pages_domains_mock.go -package=testing gitlab.com/gitlab-org/api/client-go PagesDomainsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=pages_mock.go -package=testing gitlab.com/gitlab-org/api/client-go PagesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=personal_access_tokens_mock.go -package=testing gitlab.com/gitlab-org/api/client-go PersonalAccessTokensServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=pipeline_schedules_mock.go -package=testing gitlab.com/gitlab-org/api/client-go PipelineSchedulesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=pipeline_triggers_mock.go -package=testing gitlab.com/gitlab-org/api/client-go PipelineTriggersServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=pipelines_mock.go -package=testing gitlab.com/gitlab-org/api/client-go PipelinesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=plan_limits_mock.go -package=testing gitlab.com/gitlab-org/api/client-go PlanLimitsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_access_tokens_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectAccessTokensServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_badges_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectBadgesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_clusters_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectClustersServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_feature_flags_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectFeatureFlagServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_import_export_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectImportExportServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_iterations_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectIterationsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_managed_licenses_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ManagedLicensesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_markdown_uploads_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectMarkdownUploadsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_members_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectMembersServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_mirror_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectMirrorServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_repository_storage_move_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectRepositoryStorageMoveServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_security_settings_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectSecuritySettingsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_snippets_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectSnippetsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_templates_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectTemplatesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_variables_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectVariablesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=project_vulnerabilities_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectVulnerabilitiesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=projects_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=protected_branches_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProtectedBranchesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=protected_environments_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProtectedEnvironmentsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=protected_tags_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProtectedTagsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=releaselinks_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ReleaseLinksServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=releases_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ReleasesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=repositories_mock.go -package=testing gitlab.com/gitlab-org/api/client-go RepositoriesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=repository_files_mock.go -package=testing gitlab.com/gitlab-org/api/client-go RepositoryFilesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=repository_submodules_mock.go -package=testing gitlab.com/gitlab-org/api/client-go RepositorySubmodulesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=resource_group_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ResourceGroupServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=resource_iteration_events_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ResourceIterationEventsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=resource_label_events_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ResourceLabelEventsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=resource_milestone_events_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ResourceMilestoneEventsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=resource_state_events_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ResourceStateEventsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=resource_weight_events_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ResourceWeightEventsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=runners_mock.go -package=testing gitlab.com/gitlab-org/api/client-go RunnersServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=search_mock.go -package=testing gitlab.com/gitlab-org/api/client-go SearchServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=secure_files_mock.go -package=testing gitlab.com/gitlab-org/api/client-go SecureFilesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=services_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ServicesServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=settings_mock.go -package=testing gitlab.com/gitlab-org/api/client-go SettingsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=sidekiq_metrics_mock.go -package=testing gitlab.com/gitlab-org/api/client-go SidekiqServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=snippet_repository_storage_move_mock.go -package=testing gitlab.com/gitlab-org/api/client-go SnippetRepositoryStorageMoveServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=snippets_mock.go -package=testing gitlab.com/gitlab-org/api/client-go SnippetsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=system_hooks_mock.go -package=testing gitlab.com/gitlab-org/api/client-go SystemHooksServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=tags_mock.go -package=testing gitlab.com/gitlab-org/api/client-go TagsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=todos_mock.go -package=testing gitlab.com/gitlab-org/api/client-go TodosServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=topics_mock.go -package=testing gitlab.com/gitlab-org/api/client-go TopicsServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=usage_data_mock.go -package=testing gitlab.com/gitlab-org/api/client-go UsageDataServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=users_mock.go -package=testing gitlab.com/gitlab-org/api/client-go UsersServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=validate_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ValidateServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=version_mock.go -package=testing gitlab.com/gitlab-org/api/client-go VersionServiceInterface
//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -typed -destination=wikis_mock.go -package=testing gitlab.com/gitlab-org/api/client-go WikisServiceInterface
