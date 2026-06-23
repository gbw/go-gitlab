//go:build integration

package gitlab_test

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gitlab "gitlab.com/gitlab-org/api/client-go/v2"
)

// Integration tests for the Groups API.
// These tests require a GitLab instance running on localhost:8095.
// They also require a valid admin token in GITLAB_TOKEN environment variable.

func Test_GroupsGetGroup_MergeSettings_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// WHEN retrieving the group
	retrievedGroup, _, err := client.Groups.GetGroup(group.ID, nil)
	require.NoError(t, err, "Failed to get group")

	// THEN the merge settings fields should be present
	assert.NotNil(t, retrievedGroup)
	// The fields should have default values (false for most merge restrictions)
	assert.False(t, retrievedGroup.OnlyAllowMergeIfPipelineSucceeds)
	assert.False(t, retrievedGroup.AllowMergeOnSkippedPipeline)
	assert.False(t, retrievedGroup.OnlyAllowMergeIfAllDiscussionsAreResolved)
}

func Test_GroupsUpdateGroup_MergeSettings_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// WHEN updating the group with merge settings enabled
	updatedGroup, _, err := client.Groups.UpdateGroup(group.ID, &gitlab.UpdateGroupOptions{
		OnlyAllowMergeIfPipelineSucceeds:          gitlab.Ptr(true),
		AllowMergeOnSkippedPipeline:               gitlab.Ptr(true),
		OnlyAllowMergeIfAllDiscussionsAreResolved: gitlab.Ptr(true),
	})
	require.NoError(t, err, "Failed to update group")

	// THEN the merge settings should be updated
	assert.True(t, updatedGroup.OnlyAllowMergeIfPipelineSucceeds)
	assert.True(t, updatedGroup.AllowMergeOnSkippedPipeline)
	assert.True(t, updatedGroup.OnlyAllowMergeIfAllDiscussionsAreResolved)

	// AND WHEN retrieving the group again
	retrievedGroup, _, err := client.Groups.GetGroup(group.ID, nil)
	require.NoError(t, err, "Failed to get group")

	// THEN the merge settings should persist
	assert.True(t, retrievedGroup.OnlyAllowMergeIfPipelineSucceeds)
	assert.True(t, retrievedGroup.AllowMergeOnSkippedPipeline)
	assert.True(t, retrievedGroup.OnlyAllowMergeIfAllDiscussionsAreResolved)
}

func Test_GroupsUpdateGroup_MergeSettings_Disable_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group with merge settings enabled
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// Enable merge settings first
	_, _, err := client.Groups.UpdateGroup(group.ID, &gitlab.UpdateGroupOptions{
		OnlyAllowMergeIfPipelineSucceeds:          gitlab.Ptr(true),
		AllowMergeOnSkippedPipeline:               gitlab.Ptr(true),
		OnlyAllowMergeIfAllDiscussionsAreResolved: gitlab.Ptr(true),
	})
	require.NoError(t, err, "Failed to enable merge settings")

	// WHEN updating the group to disable merge settings
	updatedGroup, _, err := client.Groups.UpdateGroup(group.ID, &gitlab.UpdateGroupOptions{
		OnlyAllowMergeIfPipelineSucceeds:          gitlab.Ptr(false),
		AllowMergeOnSkippedPipeline:               gitlab.Ptr(false),
		OnlyAllowMergeIfAllDiscussionsAreResolved: gitlab.Ptr(false),
	})
	require.NoError(t, err, "Failed to update group")

	// THEN the merge settings should be disabled
	assert.False(t, updatedGroup.OnlyAllowMergeIfPipelineSucceeds)
	assert.False(t, updatedGroup.AllowMergeOnSkippedPipeline)
	assert.False(t, updatedGroup.OnlyAllowMergeIfAllDiscussionsAreResolved)
}

func Test_GroupsMaxArtifactsSize_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// WHEN updating the group to set MaxArtifactsSize to 100 MB
	updatedGroup, _, err := client.Groups.UpdateGroup(group.ID, &gitlab.UpdateGroupOptions{
		MaxArtifactsSize: gitlab.Ptr(int64(100)), // 100 MB
	})
	require.NoError(t, err, "Failed to update group MaxArtifactsSize")

	// THEN the setting should be reflected in the update response
	assert.Equal(t, int64(100), updatedGroup.MaxArtifactsSize)

	// AND WHEN retrieving the group again
	retrievedGroup, _, err := client.Groups.GetGroup(group.ID, nil)
	require.NoError(t, err, "Failed to retrieve group after update")

	// THEN MaxArtifactsSize should persist
	assert.Equal(t, int64(100), retrievedGroup.MaxArtifactsSize)
}

func Test_GroupProtectedBranches_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// Define branch name
	branchName := "main"

	// WHEN protecting a branch
	protectedBranch, _, err := client.GroupProtectedBranches.ProtectRepositoryBranches(group.ID, &gitlab.ProtectGroupRepositoryBranchesOptions{
		Name:             gitlab.Ptr(branchName),
		PushAccessLevel:  gitlab.Ptr(gitlab.MaintainerPermissions),
		MergeAccessLevel: gitlab.Ptr(gitlab.MaintainerPermissions),
	})
	require.NoError(t, err, "Failed to protect branch")

	// THEN the branch should be protected
	assert.Equal(t, branchName, protectedBranch.Name)

	// WHEN listing protected branches
	branches, _, err := client.GroupProtectedBranches.ListProtectedBranches(group.ID, nil)
	require.NoError(t, err, "Failed to list protected branches")

	// THEN the protected branch should be in the list
	found := false
	for _, b := range branches {
		if b.Name == branchName {
			found = true
			break
		}
	}
	assert.True(t, found, "Protected branch not found in list")

	// WHEN getting the protected branch
	gotBranch, _, err := client.GroupProtectedBranches.GetProtectedBranch(group.ID, branchName)
	require.NoError(t, err, "Failed to get protected branch")

	// THEN it should match
	assert.Equal(t, branchName, gotBranch.Name)

	// WHEN updating the protected branch
	updatedBranch, _, err := client.GroupProtectedBranches.UpdateProtectedBranch(group.ID, branchName, &gitlab.UpdateGroupProtectedBranchOptions{
		AllowForcePush: gitlab.Ptr(true),
	})
	require.NoError(t, err, "Failed to update protected branch")

	// THEN the update should be reflected
	assert.True(t, updatedBranch.AllowForcePush)

	// WHEN unprotecting the branch
	_, err = client.GroupProtectedBranches.UnprotectRepositoryBranches(group.ID, branchName)
	require.NoError(t, err, "Failed to unprotect branch")

	// THEN getting the branch should fail (404)
	_, resp, err := client.GroupProtectedBranches.GetProtectedBranch(group.ID, branchName)
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode)
}

func Test_GroupsCreateGroup_CodeOwnerApprovalRequired_Integration(t *testing.T) {
	// GIVEN a GitLab client
	client := SetupIntegrationClient(t)

	// WHEN creating a group with CodeOwnerApprovalRequired enabled
	group, _, err := client.Groups.CreateGroup(&gitlab.CreateGroupOptions{
		Name: gitlab.Ptr("code-owner-create-test"),
		Path: gitlab.Ptr("code-owner-create-test"),
		DefaultBranchProtectionDefaults: &gitlab.DefaultBranchProtectionDefaultsOptions{
			CodeOwnerApprovalRequired: gitlab.Ptr(true),
		},
	})
	require.NoError(t, err, "Failed to create group")

	// THEN the setting should be enabled
	require.NotNil(t, group.DefaultBranchProtectionDefaults)
	assert.True(t, group.DefaultBranchProtectionDefaults.CodeOwnerApprovalRequired)

	// AND WHEN retrieving the group
	retrievedGroup, _, err := client.Groups.GetGroup(group.ID, nil)
	require.NoError(t, err)

	// THEN it should persist
	require.NotNil(t, retrievedGroup.DefaultBranchProtectionDefaults)
	assert.True(t, retrievedGroup.DefaultBranchProtectionDefaults.CodeOwnerApprovalRequired)
}

func Test_GroupsUpdateGroup_CodeOwnerApprovalRequired_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// WHEN updating CodeOwnerApprovalRequired to true
	updatedGroup, _, err := client.Groups.UpdateGroup(group.ID, &gitlab.UpdateGroupOptions{
		DefaultBranchProtectionDefaults: &gitlab.DefaultBranchProtectionDefaultsOptions{
			CodeOwnerApprovalRequired: gitlab.Ptr(true),
		},
	})
	require.NoError(t, err)

	// THEN the update response should reflect the change
	require.NotNil(t, updatedGroup.DefaultBranchProtectionDefaults)
	assert.True(t, updatedGroup.DefaultBranchProtectionDefaults.CodeOwnerApprovalRequired)

	// AND WHEN retrieving the group again
	retrievedGroup, _, err := client.Groups.GetGroup(group.ID, nil)
	require.NoError(t, err)

	// THEN the setting should persist
	require.NotNil(t, retrievedGroup.DefaultBranchProtectionDefaults)
	assert.True(t, retrievedGroup.DefaultBranchProtectionDefaults.CodeOwnerApprovalRequired)
}

// Test_GroupsCreateGroup_AllFields_Integration iterates over every field of
// CreateGroupOptions via reflection and verifies the GitLab create-group
// endpoint accepts it. Each pointer field becomes its own subtest, so any
// field added to CreateGroupOptions in the future is exercised automatically.
//
// For fields whose JSON tag matches a field on the Group response struct,
// the test also verifies the value round-trips. For everything else, the
// assertion is that create did not return an error — GitLab returns 400 on
// some invalid fields, but does silently ignore others, so this is a
// necessary-but-not-sufficient signal.
//
// Adding a new pointer-to-{bool,int64,string,[]string,[]int64} field needs
// no change here. Fields of other types (named string enums, nested structs)
// surface as t.Fatal with a message — add an override to exercise them.
func Test_GroupsCreateGroup_AllFields_Integration(t *testing.T) {
	client := SetupIntegrationClient(t)

	// Fields skipped by this test, keyed by JSON name with a reason.
	skip := map[string]string{
		"name":                               "always set as required setup",
		"path":                               "always set as required setup",
		"visibility":                         "always set as required setup",
		"emails_disabled":                    "deprecated",
		"default_branch_protection":          "deprecated",
		"parent_id":                          "would make the group a subgroup, changing the fixture",
		"organization_id":                    "env-specific — targets a specific organization",
		"default_branch_protection_defaults": "nested struct; covered by CodeOwnerApprovalRequired test",
		"enabled_git_access_protocol":        "not returned in create response; tested via update",

		// Duo isn't yet enabled on the integration test instance — see
		// https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/merge_requests/3086
		// for the in-flight work to wire it up. Re-enable when that lands.
		"duo_availability":            "Duo not yet enabled in CI",
		"duo_features_enabled":        "Duo not yet enabled in CI",
		"lock_duo_features_enabled":   "Duo not yet enabled in CI",
		"experiment_features_enabled": "Duo not yet enabled in CI",
	}

	skipPremium := map[string]struct{}{
		"ip_restriction_ranges":      {},
		"allowed_email_domains_list": {},
	}

	// Per-field value overrides. Use when the default value of the
	// corresponding Group response field would match an auto-generated value
	// (so the round-trip assertion couldn't distinguish "applied" from
	// "silently ignored"), or when the field type needs a specific value.
	overrides := map[string]any{
		// Enum-typed fields: pick a non-default value so the round-trip check
		// distinguishes "applied" from "silently ignored".
		"project_creation_level":  gitlab.MaintainerProjectCreation,            // default "developer"
		"subgroup_creation_level": gitlab.MaintainerSubGroupCreationLevelValue, // default "owner"
		"wiki_access_level":       gitlab.DisabledAccessControl,                // default "enabled"
		"default_branch":          "main",                                      // valid branch name

		// These fields require valid formats even on licensed instances.
		"ip_restriction_ranges":      "192.168.0.0/24",
		"allowed_email_domains_list": "example.com",
	}

	// Build a json-tag → Group struct field index for round-trip checks.
	groupByJSON := indexByJSONTag(reflect.TypeOf(gitlab.Group{}))
	suffix := time.Now().UnixNano()

	optsT := reflect.TypeOf(gitlab.CreateGroupOptions{})
	for i := 0; i < optsT.NumField(); i++ {
		f := optsT.Field(i)
		urlTag := strings.Split(f.Tag.Get("url"), ",")[0]
		if urlTag == "" || urlTag == "-" {
			// Either no API field (e.g. Avatar with url:"-") or no tag at all.
			continue
		}
		if reason, ok := skip[urlTag]; ok {
			t.Logf("skipping %q: %s", urlTag, reason)
			continue
		}

		t.Run(urlTag, func(t *testing.T) {
			if _, ok := skipPremium[urlTag]; ok {
				SkipIfNotLicensed(t, client)
			}

			fieldVal := getFieldVal(t, f, urlTag, overrides)

			// Build options: required setup + the single field under test.
			name := fmt.Sprintf("test-%s-%d", f.Name, suffix)
			opts := &gitlab.CreateGroupOptions{
				Name:       gitlab.Ptr(name),
				Path:       gitlab.Ptr(name),
				Visibility: gitlab.Ptr(gitlab.PublicVisibility),
			}
			reflect.ValueOf(opts).Elem().FieldByName(f.Name).Set(fieldVal)

			// Create + cleanup.
			group, _, err := client.Groups.CreateGroup(opts, gitlab.WithContext(t.Context()))
			require.NoErrorf(t, err, "create rejected field %q", urlTag)
			t.Cleanup(func() {
				_, _ = client.Groups.DeleteGroup(group.ID, nil, gitlab.WithContext(context.Background()))
			})

			// Round-trip the value if Group exposes it under the same JSON tag.
			gf, hasResponseField := groupByJSON[urlTag]
			if !hasResponseField {
				t.Logf("field %q has no matching Group response field; only verified create did not error", urlTag)
				return
			}
			got := reflect.ValueOf(group).Elem().FieldByIndex(gf.Index).Interface()
			want := fieldVal.Elem().Interface()
			assert.Equalf(t, want, got, "field %q not persisted on Group response", urlTag)
		})
	}
}

// getFieldVal returns the pointer-typed reflect.Value to assign to `f`.
// If `overrides` names this field, the override value is wrapped in a pointer
// of the field's type. Otherwise, it falls back to autoTestValue. Fields whose
// type isn't handled by either path fail the test with a fix-it message.
func getFieldVal(t *testing.T, f reflect.StructField, urlTag string, overrides map[string]any) reflect.Value {
	t.Helper()
	if ov, ok := overrides[urlTag]; ok {
		target := f.Type.Elem()
		ovV := reflect.ValueOf(ov)
		require.Truef(t, ovV.Type().ConvertibleTo(target),
			"override for %q has incompatible type %s (need %s)", urlTag, ovV.Type(), target)
		p := reflect.New(target)
		p.Elem().Set(ovV.Convert(target))
		return p
	}
	v, ok := autoTestValue(f.Type)
	if !ok {
		t.Fatalf("no auto-generated test value for type %s; add an entry to `overrides` for %q to cover this field", f.Type, urlTag)
	}
	return v
}

// indexByJSONTag returns a map from the first segment of each field's json
// struct tag to the field, skipping fields with no tag or `json:"-"`.
func indexByJSONTag(t reflect.Type) map[string]reflect.StructField {
	out := map[string]reflect.StructField{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := strings.Split(f.Tag.Get("json"), ",")[0]
		if tag != "" && tag != "-" {
			out[tag] = f
		}
	}
	return out
}

// Test_GroupsUpdateGroup_AllFields_Integration is the update-endpoint analogue
// of Test_GroupsCreateGroup_AllFields_Integration: iterates over every field
// of UpdateGroupOptions via reflection, calls UpdateGroup on a fresh fixture
// group, and verifies the result. Each pointer field becomes its own subtest,
// so any new field added to UpdateGroupOptions is exercised automatically.
func Test_GroupsUpdateGroup_AllFields_Integration(t *testing.T) {
	client := SetupIntegrationClient(t)

	skip := map[string]string{
		"path":                                 "must be a URL-safe slug; not exercised by the generic *string path",
		"emails_disabled":                      "deprecated",
		"default_branch_protection":            "deprecated",
		"default_branch_protection_defaults":   "nested struct; covered by CodeOwnerApprovalRequired test",
		"step_up_auth_required_oauth_provider": "needs a configured OAuth provider",
		// Duo isn't yet enabled on the integration test instance — see
		// https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/merge_requests/3086
		// for the in-flight work to wire it up. Re-enable when that lands.
		"duo_availability":            "Duo not yet enabled in CI",
		"duo_features_enabled":        "Duo not yet enabled in CI",
		"lock_duo_features_enabled":   "Duo not yet enabled in CI",
		"experiment_features_enabled": "Duo not yet enabled in CI",
	}

	// prerequisiteUpdates maps a field under test to a setup function run
	// before the actual field update. A non-nil return value overrides the
	// field value for that subtest (use when the value must be computed from
	//  the setup state, e.g. a freshly created project ID).
	prerequisiteUpdates := map[string]func(*testing.T, *gitlab.Client, *gitlab.Group) any{
		"allow_merge_on_skipped_pipeline": func(t *testing.T, c *gitlab.Client, g *gitlab.Group) any {
			t.Helper()
			_, _, err := c.Groups.UpdateGroup(g.ID, &gitlab.UpdateGroupOptions{
				OnlyAllowMergeIfPipelineSucceeds: gitlab.Ptr(true),
			}, gitlab.WithContext(t.Context()))
			require.NoErrorf(t, err, "prerequisite: enable pipeline success requirement on group %d", g.ID)
			return nil
		},
		"file_template_project_id": func(t *testing.T, c *gitlab.Client, g *gitlab.Group) any {
			t.Helper()
			// The template project must live inside the group's namespace.
			project := CreateTestProjectWithOptions(t, c, &gitlab.CreateProjectOptions{
				Name:        gitlab.Ptr(fmt.Sprintf("template-project-%d", time.Now().UnixNano())),
				Visibility:  gitlab.Ptr(gitlab.PublicVisibility),
				NamespaceID: gitlab.Ptr(g.ID),
			})
			return project.ID
		},
	}

	skipPremium := map[string]struct{}{
		"ip_restriction_ranges":      {},
		"allowed_email_domains_list": {},
		"file_template_project_id":   {},
	}

	overrides := map[string]any{
		// CreateTestGroup builds a public group; flip to detect the update applied.
		"visibility": gitlab.PrivateVisibility,

		// Enum-typed fields: pick a non-default value so the round-trip check
		// distinguishes "applied" from "silently ignored".
		"project_creation_level":      gitlab.MaintainerProjectCreation,
		"subgroup_creation_level":     gitlab.MaintainerSubGroupCreationLevelValue,
		"wiki_access_level":           gitlab.DisabledAccessControl,
		"shared_runners_setting":      gitlab.DisabledAndUnoverridableSharedRunnersSettingValue,
		"enabled_git_access_protocol": gitlab.EnabledGitAccessProtocolSSH, // default "all"
		"default_branch":              "main",                             // valid branch name

		// These fields require valid formats even on licensed instances.
		"ip_restriction_ranges":      "192.168.0.0/24",
		"allowed_email_domains_list": "example.com",
	}

	groupByJSON := indexByJSONTag(reflect.TypeOf(gitlab.Group{}))

	optsT := reflect.TypeOf(gitlab.UpdateGroupOptions{})
	for i := 0; i < optsT.NumField(); i++ {
		f := optsT.Field(i)
		urlTag := strings.Split(f.Tag.Get("url"), ",")[0]
		if urlTag == "" || urlTag == "-" {
			continue
		}
		if reason, ok := skip[urlTag]; ok {
			t.Logf("skipping %q: %s", urlTag, reason)
			continue
		}

		t.Run(urlTag, func(t *testing.T) {
			if _, ok := skipPremium[urlTag]; ok {
				SkipIfNotLicensed(t, client)
			}

			// Fresh fixture per subtest so update state doesn't leak between subtests.
			group := CreateTestGroup(t, client)

			subtestOverrides := overrides
			if prereq, ok := prerequisiteUpdates[urlTag]; ok {
				if v := prereq(t, client, group); v != nil {
					m := make(map[string]any, len(overrides)+1)
					for k, val := range overrides {
						m[k] = val
					}
					m[urlTag] = v
					subtestOverrides = m
				}
			}

			fieldVal := getFieldVal(t, f, urlTag, subtestOverrides)
			opts := &gitlab.UpdateGroupOptions{}
			reflect.ValueOf(opts).Elem().FieldByName(f.Name).Set(fieldVal)

			updated, _, err := client.Groups.UpdateGroup(group.ID, opts, gitlab.WithContext(t.Context()))
			require.NoErrorf(t, err, "update rejected field %q", urlTag)

			gf, hasResponseField := groupByJSON[urlTag]
			if !hasResponseField {
				t.Logf("field %q has no matching Group response field; only verified update did not error", urlTag)
				return
			}
			got := reflect.ValueOf(updated).Elem().FieldByIndex(gf.Index).Interface()
			want := fieldVal.Elem().Interface()
			assert.Equalf(t, want, got, "field %q not persisted on Group response", urlTag)
		})
	}
}

// autoTestValue returns a pointer-typed reflect.Value with a non-zero value
// of the given pointer type, for the simple types we know how to handle.
// Returns (zero, false) for types that need a manual override.
func autoTestValue(t reflect.Type) (reflect.Value, bool) {
	if t.Kind() != reflect.Ptr {
		return reflect.Value{}, false
	}
	elem := t.Elem()
	p := reflect.New(elem)
	switch elem.Kind() {
	case reflect.Bool:
		p.Elem().SetBool(true)
		return p, true
	case reflect.String:
		// Only handle the plain `string` type — named string types (enums)
		// need a known-good value and must use an override.
		if elem == reflect.TypeOf("") {
			p.Elem().SetString("integration test value")
			return p, true
		}
		return reflect.Value{}, false
	case reflect.Int, reflect.Int64:
		p.Elem().SetInt(100)
		return p, true
	case reflect.Slice:
		switch elem.Elem().Kind() {
		case reflect.String:
			p.Elem().Set(reflect.ValueOf([]string{"root"}))
			return p, true
		case reflect.Int64:
			p.Elem().Set(reflect.ValueOf([]int64{1}))
			return p, true
		}
	}
	return reflect.Value{}, false
}

func Test_GroupsArchive_Unarchive_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// WHEN the group is archived
	_, err := client.Groups.ArchiveGroup(group.ID)
	require.NoError(t, err)

	// THEN archiving again should fail
	resp, err := client.Groups.ArchiveGroup(group.ID)
	assert.Error(t, err)
	assert.Equal(t, 422, resp.StatusCode)

	// AND WHEN the group is unarchived
	_, err = client.Groups.UnarchiveGroup(group.ID)
	require.NoError(t, err)

	// THEN unarchiving again should fail
	resp, err = client.Groups.UnarchiveGroup(group.ID)
	assert.Error(t, err)
	assert.Equal(t, 422, resp.StatusCode)
}

func Test_GroupsListSAMLUsers_Integration(t *testing.T) {
	// GIVEN a licensed GitLab instance and a test group
	client := SetupIntegrationClient(t)
	SkipIfNotLicensed(t, client)
	group := CreateTestGroup(t, client)

	// WHEN ListSAMLUsers is called on a group with no SAML users provisioned
	users, _, err := client.Groups.ListSAMLUsers(group.ID, &gitlab.ListSAMLUsersOptions{})

	// THEN the call succeeds and returns an empty slice
	require.NoError(t, err)
	assert.NotNil(t, users)
}

func Test_GroupsSyncGroupWithLDAP_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// WHEN SyncGroupWithLDAP is called
	resp, err := client.Groups.SyncGroupWithLDAP(group.ID)
	// THEN the call either succeeds or returns 404 (no LDAP configured)
	if err != nil {
		assert.Equal(t, 404, resp.StatusCode, "unexpected error: %v", err)
	}
}
