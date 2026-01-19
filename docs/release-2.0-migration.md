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