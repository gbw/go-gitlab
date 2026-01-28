package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupListProtectedBranches(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/protected_branches", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
	{
		"id":1,
		"name":"master",
		"push_access_levels":[{
			"id":1,
			"access_level":40,
			"access_level_description":"Maintainers",
			"user_id":null,
			"group_id":null
		},{
			"id":2,
			"access_level":30,
			"access_level_description":"User name",
			"user_id":123,
			"group_id":null
		}],
		"merge_access_levels":[{
			"id":1,
			"access_level":40,
			"access_level_description":"Maintainers",
			"user_id":null,
			"group_id":null
		}],
		"code_owner_approval_required":false
	}
]`)
	})
	opt := &ListGroupProtectedBranchesOptions{}
	protectedBranches, resp, err := client.GroupProtectedBranches.ListProtectedBranches("1", opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	want := []*GroupProtectedBranch{
		{
			ID:   1,
			Name: "master",
			PushAccessLevels: []*GroupBranchAccessDescription{
				{
					ID:                     1,
					AccessLevel:            40,
					AccessLevelDescription: "Maintainers",
				},
				{
					ID:                     2,
					AccessLevel:            30,
					AccessLevelDescription: "User name",
					UserID:                 123,
				},
			},
			MergeAccessLevels: []*GroupBranchAccessDescription{
				{
					ID:                     1,
					AccessLevel:            40,
					AccessLevelDescription: "Maintainers",
				},
			},
			AllowForcePush:            false,
			CodeOwnerApprovalRequired: false,
		},
	}
	assert.Equal(t, want, protectedBranches)
}

func TestGroupGetProtectedBranch(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/protected_branches/main", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id":1,
			"name":"main",
			"push_access_levels":[{
				"id":1,
				"access_level":40,
				"access_level_description":"Maintainers",
				"user_id":null,
				"group_id":null
			}],
			"merge_access_levels":[{
				"id":1,
				"access_level":40,
				"access_level_description":"Maintainers",
				"user_id":null,
				"group_id":null
			}],
			"code_owner_approval_required":false
		}`)
	})
	protectedBranch, resp, err := client.GroupProtectedBranches.GetProtectedBranch(1, "main")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	want := &GroupProtectedBranch{
		ID:   1,
		Name: "main",
		PushAccessLevels: []*GroupBranchAccessDescription{
			{
				ID:                     1,
				AccessLevel:            40,
				AccessLevelDescription: "Maintainers",
			},
		},
		MergeAccessLevels: []*GroupBranchAccessDescription{
			{
				ID:                     1,
				AccessLevel:            40,
				AccessLevelDescription: "Maintainers",
			},
		},
		AllowForcePush:            false,
		CodeOwnerApprovalRequired: false,
	}
	assert.Equal(t, want, protectedBranch)
}

func TestGroupProtectRepositoryBranches(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/protected_branches", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
	{
		"id":1,
		"name":"master",
		"push_access_levels":[{
			"access_level":40,
			"access_level_description":"Maintainers"
		}],
		"merge_access_levels":[{
			"access_level":40,
			"access_level_description":"Maintainers"
		}],
		"allow_force_push":true,
		"code_owner_approval_required":true
	}`)
	})
	opt := &ProtectGroupRepositoryBranchesOptions{
		Name:                      Ptr("master"),
		PushAccessLevel:           Ptr(MaintainerPermissions),
		MergeAccessLevel:          Ptr(MaintainerPermissions),
		AllowForcePush:            Ptr(true),
		CodeOwnerApprovalRequired: Ptr(true),
	}
	protectedBranches, resp, err := client.GroupProtectedBranches.ProtectRepositoryBranches("1", opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	want := &GroupProtectedBranch{
		ID:   1,
		Name: "master",
		PushAccessLevels: []*GroupBranchAccessDescription{
			{
				AccessLevel:            40,
				AccessLevelDescription: "Maintainers",
			},
		},
		MergeAccessLevels: []*GroupBranchAccessDescription{
			{
				AccessLevel:            40,
				AccessLevelDescription: "Maintainers",
			},
		},
		AllowForcePush:            true,
		CodeOwnerApprovalRequired: true,
	}
	assert.Equal(t, want, protectedBranches)
}

func TestGroupUnprotectRepositoryBranches(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/protected_branches/main", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})
	resp, err := client.GroupProtectedBranches.UnprotectRepositoryBranches("1", "main")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestGroupUpdateProtectedBranch(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/protected_branches/master", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBodyJSON(t, r, map[string]bool{
			"code_owner_approval_required": true,
		})
		fmt.Fprintf(w, `{
			"name": "master",
			"code_owner_approval_required": true
		}`)
	})
	opt := &UpdateGroupProtectedBranchOptions{
		CodeOwnerApprovalRequired: Ptr(true),
	}
	protectedBranch, resp, err := client.GroupProtectedBranches.UpdateProtectedBranch("1", "master", opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := &GroupProtectedBranch{
		Name:                      "master",
		CodeOwnerApprovalRequired: true,
	}
	assert.Equal(t, want, protectedBranch)
}
