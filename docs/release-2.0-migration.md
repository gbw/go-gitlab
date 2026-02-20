# GitLab Go Client v2.0 Migration Guide

This document outlines the breaking changes introduced in GitLab Go Client v2.0 and provides guidance on how to migrate your code.

## Overview

Version 2.0 introduces breaking changes to improve API consistency, naming conventions, and parameter handling across the client library. This includes:

- Consistent return signatures for user moderation methods to return the "response" object along with errors
- New `Nullable[T]` generic type for explicit null handling, which allows users to differentiate explicit "null" vs empty, and make it easier to send explicit "null"
- New services for Work Items
- Removal of deprecated methods

## Update minimum required Go version

The client-go 2.0 major version upgrade aligns our supported Go versions to align to the Golang [Release Policy](https://go.dev/doc/devel/release#policy), and 
changes the minimum required Go version to 1.25 (up from 1.24). 

client-go 3.0 will release in roughly 6 months when Go version 1.27 releases, and will change the minimum required Go version to 1.26. 

With the future release of client-go 3.0, we will be deprecating the `gitlab.Ptr` function and aligning with the usage of `new()` that is now native to Go 1.26. We
encourage users who have already migrated to go 1.26 to use the `new()` function instead of `gitlab.Ptr`.

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

The Integrations often returned generic `Integration` objects, which prevented us from providing access to typesafe Properties. They now return typesafe integration
 structs.

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

### Merge Requests That Implement This Change
- [Refactor Jira Integration Settings](https://gitlab.com/gitlab-org/api/client-go/-/merge_requests/2678) by @HamzaHassanain

## Harbor Group Integration

The `GetGroupHarborSettings` method now returns a `*HarborIntegration` struct instead of the generic `*Integration` struct. This provides strongly typed access to Harbor-specific properties.

**Changes:**
- Return type: Changed from `*Integration` to `*HarborIntegration`
- Properties: The `Properties` field in `HarborIntegration` is now of type `HarborIntegrationProperties`, containing fields like `URL`, `ProjectName`, `Username`, etc.

```go
// Before (v1.x)
integration, _, err := client.Integrations.GetGroupHarborSettings(gid)
// integration.Properties was generic/untyped

// After (v2.0)
harborIntegration, _, err := client.Integrations.GetGroupHarborSettings(gid)
fmt.Println(harborIntegration.Properties.URL)
```

## User Moderation Methods Return Signature Changes

All user moderation methods now return `(*Response, error)` instead of just `error`. This provides access to HTTP response metadata including status codes and headers.

**Affected Methods:**
- `BlockUser`
- `UnblockUser`
- `BanUser`
- `UnbanUser`
- `DeactivateUser`
- `ActivateUser`
- `ApproveUser`
- `RejectUser`
- `DisableTwoFactor`

```go
// Before (v1.x)
err := client.Users.BlockUser(userID)
if err != nil {
    return err
}

// After (v2.0)
resp, err := client.Users.BlockUser(userID)
if err != nil {
    return err
}
```

## Personal Access Tokens - Removed Deprecated Method

The deprecated `RevokePersonalAccessToken` method has been removed. Use `RevokePersonalAccessTokenByID` instead.

**Changes:**
- Removed: `RevokePersonalAccessToken(token int64, options ...RequestOptionFunc) (*Response, error)`
- Use: `RevokePersonalAccessTokenByID(token int64, options ...RequestOptionFunc) (*Response, error)`

```go
// Before (v1.x)
resp, err := client.PersonalAccessTokens.RevokePersonalAccessToken(tokenID)

// After (v2.0)
resp, err := client.PersonalAccessTokens.RevokePersonalAccessTokenByID(tokenID)
```

## Nullable Type for Explicit Null Handling

A new generic `Nullable[T]` type has been introduced to handle fields that can be:
- Not set in the request
- Explicitly set to `null` in the request
- Explicitly set to a valid value in the request

This is particularly useful for optional fields where you need to distinguish between "not provided" and "explicitly set to null".

**Affected Fields:**
- `Label.Priority` - changed from `int64` to `Nullable[int64]`
- `CreateLabelOptions.Priority` - changed from `*int64` to `Nullable[int64]`
- `UpdateLabelOptions.Priority` - changed from `*int64` to `Nullable[int64]`
- `CreateGroupLabelOptions.Priority` - changed from `*int64` to `Nullable[int64]`
- `UpdateGroupLabelOptions.Priority` - changed from `*int64` to `Nullable[int64]`

```go
// Before (v1.x)
label, _, err := client.Labels.CreateLabel(projectID, &gitlab.CreateLabelOptions{
    Name:     gitlab.Ptr("bug"),
    Color:    gitlab.Ptr("#FF0000"),
    Priority: gitlab.Ptr(int64(10)),
})

// After (v2.0)
label, _, err := client.Labels.CreateLabel(projectID, &gitlab.CreateLabelOptions{
    Name:     gitlab.Ptr("bug"),
    Color:    gitlab.Ptr("#FF0000"),
    Priority: gitlab.NewNullableWithValue(int64(10)),
})

// To explicitly set null
opts := &gitlab.UpdateLabelOptions{
    Name:     gitlab.Ptr("bug"),
    Priority: gitlab.NewNullNullable[int64](), // Explicitly set to null
}

// To check if a value is set
if label.Priority.IsSpecified() {
    priority, err := label.Priority.Get()
    if err == nil {
        fmt.Printf("Priority: %d\n", priority)
    }
}
```

**Helper Functions:**
- `NewNullableWithValue[T](value T)` - Create a Nullable with a value
- `NewNullNullable[T]()` - Create a Nullable explicitly set to null
- `Nullable[T].Get()` - Retrieve the value (returns error if not set or null)
- `Nullable[T].MustGet()` - Retrieve the value (panics if not set or null)
- `Nullable[T].Set(value T)` - Set the value
- `Nullable[T].IsNull()` - Check if explicitly set to null
- `Nullable[T].IsSpecified()` - Check if a value was provided

## Time Type Changes

### BillableUserMembership.ExpiresAt

The `ExpiresAt` field in `BillableUserMembership` has been changed from `*time.Time` to `*ISOTime` for more accurate date-only representation.

**Changes:**
- Type: Changed from `*time.Time` to `*ISOTime`

```go
// Before (v1.x)
membership := &gitlab.BillableUserMembership{
    ExpiresAt: &time.Time{}, // Full timestamp
}

// After (v2.0)
membership := &gitlab.BillableUserMembership{
    ExpiresAt: &gitlab.ISOTime{}, // Date-only (YYYY-MM-DD)
}
```

**Note:** `ISOTime` is used for fields that only support year-month-day formatting, while `*time.Time` is used for full timestamps.

## API Update for PackageProtectionRule

The variable type for `MinimumAccessLevelForDelete` and `MinimumAccessLevelForPush` were previously `int64`, and have been changed to a `string` to align with the documentation. In addition, they use a new `Nullable[string]` type that allows explicit `null` values to be sent to the API, since `null` is a valid and intentional value for this API call.

**Changes:**
- Parameter type: Changed from `*int64` to `Nullable[ProtectionRuleAccessLevel]` ( string, nullable )

```go
// Before (v1.x)
rule, resp, err := client.ProtectedPackages.CreatePackageProtectionRules(1, &CreatePackageProtectionRulesOptions{
		PackageNamePattern:          Ptr("@my-scope/my-package-*"),
		PackageType:                 Ptr("npm"),
		MinimumAccessLevelForDelete: Ptr(int64(MaintainerPermissions)),
		MinimumAccessLevelForPush:   nil, // Ignored when sent to the API, preventing sending "null"
})

// After (v2.0):
rule, resp, err := client.ProtectedPackages.CreatePackageProtectionRules(1, &CreatePackageProtectionRulesOptions{
		PackageNamePattern:          Ptr("@my-scope/my-package-*"),
		PackageType:                 Ptr("npm"),
		MinimumAccessLevelForDelete: NewNullableWithValue(ProtectionRuleAccessLevelMaintainer),
		MinimumAccessLevelForPush:   NewNullNullable[ProtectionRuleAccessLevel](), // sends "null" to the API to reset the value to default
})
```


### Merge Requests That Implement This Change
- [Fix Package Protection Access Level Variable Type](https://gitlab.com/gitlab-org/api/client-go/-/merge_requests/2728) by @deepflame

