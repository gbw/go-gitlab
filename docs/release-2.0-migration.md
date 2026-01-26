# GitLab Go Client v2.0 Migration Guide

This document outlines the breaking changes introduced in GitLab Go Client v2.0 and provides guidance on how to migrate your code.

## Overview

Version 2.0 introduces breaking changes to improve API consistency, naming conventions, and parameter handling across the client library.


## Fix GetUser Function Parameter Naming

The `GetUser` function had inconsistent naming for its options parameter. The struct name `GetUsersOptions` (plural) 
was misleading since the function retrieves a single user, not multiple users. Additionally, the parameter has been 
changed from a value type to a pointer for consistency with other API methods.

**Changes:**
- Renamed: `GetUsersOptions` → `GetUserOptions`
- Parameter type: Changed from value to pointer (`*GetUserOptions`)


```go
// Before (v1.x)
user, _, err := client.Users.GetUser(1, GetUsersOptions{
    WithCustomAttributes: Ptr(true),
})
// After (v2.0):

user, _, err := client.Users.GetUser(1, &GetUserOptions{
    WithCustomAttributes: Ptr(true),
})
```

### Merge Requests That Implement This Change
- [Fix GetUser function parameter naming](https://gitlab.com/gitlab-org/api/client-go/-/merge_requests/2668#) by @seif-hatem

## Group Integrations

### Microsoft Teams

- `GetGroupMicrosoftTeamsNotifications` now returns `*MicrosoftTeamsIntegration` instead of `*Integration`.
- `SetGroupMicrosoftTeamsNotifications` now returns `*MicrosoftTeamsIntegration` instead of `*Integration`.

#### Merge Requests That Implement This Change
- [Refactor Microsoft Teams Group Integration](https://gitlab.com/gitlab-org/api/client-go/-/merge_requests/2679) by @HamzaHassanain

## Refactor Jira Integration Settings

The `GetGroupJiraSettings` and `SetUpGroupJira` methods now return a `*JiraIntegration` struct instead of the generic `*Integration` struct. This provides strongly typed access to Jira-specific properties.

**Changes:**
- Return type: Changed from `*Integration` to `*JiraIntegration`
- Properties: The `Properties` field in `JiraIntegration` is now of type `JiraIntegrationProperties`, containing fields like `URL`, `Username`, `Password`, etc., instead of being a `map[string]any` or similar generic accessible only via the API response JSON.

```go
// Before (v1.x)
integration, _, err := client.Integrations.GetGroupJiraSettings(gid)
// integration.Properties was generic/untyped

// After (v2.0)
jiraIntegration, _, err := client.Integrations.GetGroupJiraSettings(gid)
fmt.Println(jiraIntegration.Properties.URL)
```
