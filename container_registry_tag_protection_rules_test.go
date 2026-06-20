package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListContainerRegistryTagProtectionRules(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/7/registry/protection/tag/rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[ 
		{
			"id": 1,
			"project_id": 7,
			"tag_name_pattern": "v*-release",
			"minimum_access_level_for_push": "maintainer",
			"minimum_access_level_for_delete": "owner"
		},
		{
			"id": 2,
			"project_id": 7,
			"tag_name_pattern": "latest",
			"minimum_access_level_for_push": "owner",
			"minimum_access_level_for_delete": "admin"
		}
	]`)
	})

	want := []*ContainerRegistryTagProtectionRule{
		{
			ID:                          1,
			ProjectID:                   7,
			TagNamePattern:              "v*-release",
			MinimumAccessLevelForPush:   ProtectionRuleAccessLevelMaintainer,
			MinimumAccessLevelForDelete: ProtectionRuleAccessLevelOwner,
		},
		{
			ID:                          2,
			ProjectID:                   7,
			TagNamePattern:              "latest",
			MinimumAccessLevelForPush:   ProtectionRuleAccessLevelOwner,
			MinimumAccessLevelForDelete: ProtectionRuleAccessLevelAdmin,
		},
	}

	rules, resp, err := client.ContainerRegistryTagProtectionRules.ListContainerRegistryTagProtectionRules(7)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, rules)
}

func TestCreateContainerRegistryTagProtectionRule(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/7/registry/protection/tag/rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBodyJSON(t, r, map[string]string{
			"tag_name_pattern":                "v*-release",
			"minimum_access_level_for_push":   "maintainer",
			"minimum_access_level_for_delete": "owner",
		})
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{
			"id": 1,
			"project_id": 7,
			"tag_name_pattern": "v*-release",
			"minimum_access_level_for_push": "maintainer",
			"minimum_access_level_for_delete": "owner"
		}`)
	})

	want := &ContainerRegistryTagProtectionRule{
		ID:                          1,
		ProjectID:                   7,
		TagNamePattern:              "v*-release",
		MinimumAccessLevelForPush:   ProtectionRuleAccessLevelMaintainer,
		MinimumAccessLevelForDelete: ProtectionRuleAccessLevelOwner,
	}

	rule, resp, err := client.ContainerRegistryTagProtectionRules.CreateContainerRegistryTagProtectionRule(7, &CreateContainerRegistryTagProtectionRuleOptions{
		TagNamePattern:              Ptr("v*-release"),
		MinimumAccessLevelForPush:   Ptr(ProtectionRuleAccessLevelMaintainer),
		MinimumAccessLevelForDelete: Ptr(ProtectionRuleAccessLevelOwner),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, want, rule)
}

func TestUpdateContainerRegistryTagProtectionRule(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/7/registry/protection/tag/rules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBodyJSON(t, r, map[string]string{"tag_name_pattern": "v*-stable", "minimum_access_level_for_delete": "maintainer"})
		fmt.Fprintf(w, `{
			"id": 1,
			"project_id": 7,
			"tag_name_pattern": "v*-stable",
			"minimum_access_level_for_push": "maintainer",
			"minimum_access_level_for_delete": "maintainer"
		}`)
	})

	want := &ContainerRegistryTagProtectionRule{
		ID:                          1,
		ProjectID:                   7,
		TagNamePattern:              "v*-stable",
		MinimumAccessLevelForPush:   ProtectionRuleAccessLevelMaintainer,
		MinimumAccessLevelForDelete: ProtectionRuleAccessLevelMaintainer,
	}

	rule, resp, err := client.ContainerRegistryTagProtectionRules.UpdateContainerRegistryTagProtectionRule(7, 1, &UpdateContainerRegistryTagProtectionRuleOptions{
		TagNamePattern:              Ptr("v*-stable"),
		MinimumAccessLevelForDelete: Ptr(ProtectionRuleAccessLevelMaintainer),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, rule)
}

func TestDeleteContainerRegistryTagProtectionRule(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/7/registry/protection/tag/rules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.ContainerRegistryTagProtectionRules.DeleteContainerRegistryTagProtectionRule(7, 1)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
