---
page_title: "Adding support for a new API"
---

# Adding support for a new API to `client-go`

This tutorial is to help new contributors out when adding support for a new API or endpoint.
It will walk through a step-by-step guide of adding support for a new API or endpoint.
This guide will assume that a development environment has already been set up by following the `Setting up your local development environment to contribute` section of the CONTRIBUTING.md documentation.

## Step 1: Understand the API from GitLab

For this guide, we use the [`branches API`](./branches.go) API as a step-by-step example.
This code aligns to the [Branches API](https://docs.gitlab.com/api/branches/) exposed by GitLab.

The documentation describes the URL, input parameters and JSON response for each endpoint.

Note, the documentation is not currently generated from the API code.
It may not match the real input parameters and response objects.

## Step 2: Create the API file

All APIs have a `.go` file in the top level directory of this repository.
The names match the API names.
In our case, the file is `branches.go`.

Add the package declaration:

```golang
package gitlab
```

## Step 3: Create the Endpoint functions

### Step 3a: Get endpoints

Usually, `Get` endpoints only need function parameters and no custom input structs.
They use IDs get a single entry from the API endpoint.
Project and group endpoints usually accept either a numeric ID or the namespace path.
These functions take an `any` type parameter like `pid` or `gid`.

Most of the time they need a custom struct for the decoded JSON response.
The fields in these structs can usually be plain fields (for example, `string`).
Use pointers only if you need to differentiate between "unset" and "empty string".
For example, `GetBranch` returns the custom `Branch` struct:

```golang
// Branch represents a GitLab branch.
//
// GitLab API docs: https://docs.gitlab.com/api/branches/
type Branch struct {
    Commit             *Commit `json:"commit"`
    Name               string  `json:"name"`
    Protected          bool    `json:"protected"`
    Merged             bool    `json:"merged"`
    Default            bool    `json:"default"`
    CanPush            bool    `json:"can_push"`
    DevelopersCanPush  bool    `json:"developers_can_push"`
    DevelopersCanMerge bool    `json:"developers_can_merge"`
    WebURL             string  `json:"web_url"`
}
```

Pass the URL path parameters and generic `RequestOptionFunc` optional parameter to the endpoint function.
For example, for `GetBranch`:

```golang
func (s *BranchesService) GetBranch(pid any, branch string, options ...RequestOptionFunc) (*Branch, *Response, error) {
    return do[*Branch](s.client,
        withMethod(http.MethodGet),
        withPath("projects/%s/repository/branches/%s", ProjectID{pid}, branch),
        withRequestOpts(options...),
    )
}
```

### Step 3b: `List` endpoints

`List` endpoints need a custom options struct.
This is for pagination fields and custom search options.
There is a reusable struct called `ListOptions` that encapsulates the pagination fields.
Many `List` endpoints then add custom search parameters in this struct.
All fields should be pointers, to distinguish between "unset" and "set to empty/zero/false".

> [!important]
> Always add an options struct for `List` functions, even if it does not add any fields bar `ListOptions`.
> This makes the API consistent and allows forward compatibility.

For example, for `ListBranches`:

```golang
// ListBranchesOptions represents the available ListBranches() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/branches/#list-repository-branches
type ListBranchesOptions struct {
    // Nest the default pagination options
    ListOptions

    // Add any custom search parameters for the endpoint
    Search *string `url:"search,omitempty" json:"search,omitempty"`
    Regex  *string `url:"regex,omitempty" json:"regex,omitempty"`
}
```

Pass the custom options struct into the function after any URL path parameters.
Make sure to include the generic `RequestOptionFunc` optional parameter.
Usually, `List` endpoints return a slice of the custom struct created for the `Get` endpoint.
For example, for `ListBranches`:

```golang
func (s *BranchesService) ListBranches(pid any, opts *ListBranchesOptions, options ...RequestOptionFunc) ([]*Branch, *Response, error) {
    return do[[]*Branch](s.client,
        withMethod(http.MethodGet),
        withPath("projects/%s/repository/branches", ProjectID{pid}),
        withAPIOpts(opts),
        withRequestOpts(options...),
    )
}
```

### Step 3c: `Create`/`Update` endpoints

As with `List` endpoints, `Create` and `Update` endpoints need custom options structs.
These are for query parameters or JSON request body fields.
All fields should be pointers, to distinguish between "unset" and "set to empty/zero/false".
For example, for `CreateBranch`:

```golang
// CreateBranchOptions represents the available CreateBranch() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/branches/#create-repository-branch
type CreateBranchOptions struct {
    Branch *string `url:"branch,omitempty" json:"branch,omitempty"`
    Ref    *string `url:"ref,omitempty" json:"ref,omitempty"`
}
```

Pass the custom options struct into the function after any URL path parameters.
Make sure to include the generic `RequestOptionFunc` optional parameter.
`Create`/`Update` endpoints return the created/updated object.
This is usually the same custom struct created for the `Get` endpoint.
For example, for `CreateBranch`:

```golang
func (s *BranchesService) CreateBranch(pid any, opt *CreateBranchOptions, options ...RequestOptionFunc) (*Branch, *Response, error) {
    return do[*Branch](s.client,
        withMethod(http.MethodPost),
        withPath("projects/%s/repository/branches", ProjectID{pid}),
        withAPIOpts(opt),
        withRequestOpts(options...),
    )
}
```

### Step 3d: `Delete` endpoints

`Delete` endpoints usually do not need custom structs.
They use IDs to delete a single entry from the API endpoint.

Pass the URL path parameters and generic `RequestOptionFunc` optional parameter to the endpoint function.
For example, for `DeleteBranch`:

```golang
func (s *BranchesService) DeleteBranch(pid any, branch string, options ...RequestOptionFunc) (*Response, error) {
    _, resp, err := do[none](s.client,
        withMethod(http.MethodDelete),
        withPath("projects/%s/repository/branches/%s", ProjectID{pid}, branch),
        withRequestOpts(options...),
    )
    return resp, err
}
```

## Step 4: Add the API Service Interface and Service struct

To make it easier for unit testing and mocking, each API has a service interface.
This declares all the functions in the API file.
It includes all the input parameters and response objects.

Place this at the top of the file.
Include the function definitions of any endpoint functions you have created.
For example, for the Branches API:

```golang
type (
    BranchesServiceInterface interface {
        // ListBranches gets a list of repository branches from a project, sorted by name alphabetically.
        //
        // GitLab API docs:
        // https://docs.gitlab.com/api/branches/#list-repository-branches
        ListBranches(pid any, opts *ListBranchesOptions, options ...RequestOptionFunc) ([]*Branch, *Response, error)

        // GetBranch gets a single project repository branch.
        //
        // GitLab API docs:
        // https://docs.gitlab.com/api/branches/#get-single-repository-branch
        GetBranch(pid any, branch string, options ...RequestOptionFunc) (*Branch, *Response, error)

        // CreateBranch creates branch from commit SHA or existing branch.
        //
        // GitLab API docs:
        // https://docs.gitlab.com/api/branches/#create-repository-branch
        CreateBranch(pid any, opt *CreateBranchOptions, options ...RequestOptionFunc) (*Branch, *Response, error)

        // DeleteBranch deletes an existing branch.
        //
        // GitLab API docs:
        // https://docs.gitlab.com/api/branches/#delete-repository-branch
        DeleteBranch(pid any, branch string, options ...RequestOptionFunc) (*Response, error)

        // DeleteMergedBranches deletes all branches that are merged into the project's default branch.
        //
        // GitLab API docs:
        // https://docs.gitlab.com/api/branches/#delete-merged-branches
        DeleteMergedBranches(pid any, options ...RequestOptionFunc) (*Response, error)
    }

    // BranchesService handles communication with the branch related methods
    // of the GitLab API.
    //
    // GitLab API docs: https://docs.gitlab.com/api/branches/
    BranchesService struct {
        client *Client
    }
)
```

To ensure the new service matches the interface, include the following line:

```golang
var _ BranchesServiceInterface = (*BranchesService)(nil)
```

## Step 5: Wire the new service into the gitlab package

All services need adding to `gitlab.go` in two places.

Add it as an attribute of the `Client` struct.
There is a long list of services included here.
Add any new services to the list in alphabetical order.
For example, for the Branches API:

```golang
type Client struct {
    ...
    Branches BranchesServiceInterface
    ...
}
```

Populate it on the `Client` struct instance in `NewAuthSourceClient`.
Again, there is a long list of services included here.
Add any new services to the list in alphabetical order.
For example, for the Branches API:

```golang
func NewAuthSourceClient(as AuthSource, options ...ClientOptionFunc) (*Client, error) {
    ...
    c.Branches = &BranchesService{client: c}
    ...
}
```

## Step 6: Create unit tests

Every endpoint function should have unit tests associated with it.
Unit tests are in a separate `go` file, using a standard naming convention, appending `_test` to the end of your API's file name.
For example, the `branches.go` endpoints have tests in `branches_test.go`.

These unit tests use `mux` to mock out the HTTP responses.
They apply the following pattern, for example with the test for `GetBranch`:

```golang
func TestGetBranch(t *testing.T) {
    // Always ensure the test runs in parallel
    t.Parallel()
    // Standard setup function gives you the mux server and a test client
    mux, client := setup(t)

    // Mock out the API request and response
    mux.HandleFunc("/api/v4/projects/1/repository/branches/master", func(w http.ResponseWriter, r *http.Request) {
        testMethod(t, r, http.MethodGet)
        // Creating test data JSON files is optional, you can also include the JSON inline.
        mustWriteHTTPResponse(t, w, "testdata/get_branch.json")
    })

    // Call the function being tested
    branch, resp, err := client.Branches.GetBranch(1, "master")

    // Assert the response is as expected
    assert.NoError(t, err)
    assert.NotNil(t, resp)

    authoredDate := time.Date(2012, time.June, 27, 5, 51, 39, 0, time.UTC)
    committedDate := time.Date(2012, time.June, 28, 3, 44, 20, 0, time.UTC)
    want := &Branch{
        Name:               DefaultBranch,
        Merged:             false,
        Protected:          true,
        Default:            true,
        DevelopersCanPush:  false,
        DevelopersCanMerge: false,
        CanPush:            true,
        Commit: &Commit{
            AuthorEmail:    "john@example.com",
            AuthorName:     exampleEventUserName,
            AuthoredDate:   &authoredDate,
            CommittedDate:  &committedDate,
            CommitterEmail: "john@example.com",
            CommitterName:  exampleEventUserName,
            ID:             "7b5c3cc8be40ee161ae89a06bba6229da1032a0c",
            ShortID:        "7b5c3cc",
            Title:          "add projects API",
            Message:        "add projects API",
            ParentIDs:      []string{"4ad91d3c1144c406e50c7b33bae684bd6837faf8"},
        },
    }

    assert.Equal(t, want, branch)
}
```

## Step 7: Create integration tests

Integration tests have the benefit of running against a real test GitLab instance.
This confirms that the code is able to interact with the API successfully.

The ability to write and run integration tests is relatively new.
Therefore, there are only a small amount of tests written.
These tests are all in the `gitlab_test` folder.

As with the unit tests, each supported API should have a test file with `_integration_test.go` appended to the API's file name.
For example, the `branches.go` endpoints would have tests in `branches_integration_test.go`.

The integration tests all apply the following pattern, for example with the test for `Projects.GetProjectHook`:

```golang
func Test_ProjectGetProjectHook_Integration(t *testing.T) {
    // Run in parallel unless the test uses shared resources (all tests run against the same GitLab instance)
    t.Parallel()

    // Create a real client for the test GitLab instance
    client := SetupIntegrationClient(t)

    // The following two functions are reusable helper functions defined in the utils_test.go file in the same directory as the tests
    project := CreateTestProject(t, client)
    hook, err := CreateTestProjectHook(t, project.ID, client)
    require.NoError(t, err, "Failed to create test hook")

    // Call the function being tested
    retrievedHook, _, err := client.Projects.GetProjectHook(project.ID, hook.ID)
    // Assert the results are as expected
    require.NoError(t, err, "Failed to get project hook")

    assert.Equal(t, hook.ID, retrievedHook.ID)
    assert.Equal(t, hook.URL, retrievedHook.URL)
    assert.True(t, retrievedHook.PushEvents)
}
```

### Running Integration Tests

The integration tests use Docker to run a community edition instance of GitLab.
To start this instance locally, run the `make testacc-up` command.
To stop this instance, run the `make testacc-down` command.
To run the integration tests, run the `make test-integration` command.

## Step 8: Validate code and generate mocks

There are various targets in the `Makefile` for validating your code changes as you go along.
You can run `make` (without arguments) to get a list of valid make targets.

Before you create a merge request, ensure you have run the `make reviewable` command.
This generates a mock version of the service.
Then it formats the code and tests it using a linter and by running the unit tests.
You can run each command yourself by using the other targets in the `Makefile`.

## Step 9: Create your merge request

We use semantic commits in this project.
Please make sure any commits you do have semantic commit prefixes.

Create your merge request and when you are happy, request a review from one of the maintainers!
Thank you for following this tutorial and helping contribute to this project.
