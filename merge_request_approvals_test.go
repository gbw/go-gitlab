//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApprovalRules(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/1/approval_rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "security",
				"rule_type": "regular",
				"eligible_approvers": [
					{
						"id": 5,
						"name": "John Doe",
						"username": "jdoe",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						"web_url": "http://localhost/jdoe"
					},
					{
						"id": 50,
						"name": "Group Member 1",
						"username": "group_member_1",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						"web_url": "http://localhost/group_member_1"
					}
				],
				"approvals_required": 3,
				"source_rule": null,
				"users": [
					{
						"id": 5,
						"name": "John Doe",
						"username": "jdoe",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						"web_url": "http://localhost/jdoe"
					}
				],
				"groups": [
					{
						"id": 5,
						"name": "group1",
						"path": "group1",
						"description": "",
						"visibility": "public",
						"lfs_enabled": false,
						"avatar_url": null,
						"web_url": "http://localhost/groups/group1",
						"request_access_enabled": false,
						"full_name": "group1",
						"full_path": "group1",
						"parent_id": null,
						"ldap_cn": null,
						"ldap_access": null
					}
				],
				"contains_hidden_groups": false,
				"section": "test"
			}
		]`)
	})

	approvals, resp, err := client.MergeRequestApprovals.GetApprovalRules(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := []*MergeRequestApprovalRule{
		{
			ID:       1,
			Name:     "security",
			RuleType: "regular",
			EligibleApprovers: []*BasicUser{
				{
					ID:        5,
					Name:      "John Doe",
					Username:  "jdoe",
					State:     "active",
					AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					WebURL:    "http://localhost/jdoe",
				},
				{
					ID:        50,
					Name:      "Group Member 1",
					Username:  "group_member_1",
					State:     "active",
					AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					WebURL:    "http://localhost/group_member_1",
				},
			},
			ApprovalsRequired: 3,
			Users: []*BasicUser{
				{
					ID:        5,
					Name:      "John Doe",
					Username:  "jdoe",
					State:     "active",
					AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					WebURL:    "http://localhost/jdoe",
				},
			},
			Groups: []*Group{
				{
					ID:                   5,
					Name:                 "group1",
					Path:                 "group1",
					Description:          "",
					Visibility:           PublicVisibility,
					LFSEnabled:           false,
					AvatarURL:            "",
					WebURL:               "http://localhost/groups/group1",
					RequestAccessEnabled: false,
					FullName:             "group1",
					FullPath:             "group1",
				},
			},
			Section: "test",
		},
	}
	assert.Equal(t, want, approvals)
}

func TestGetApprovalState(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/1/approval_state", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"approval_rules_overwritten": true,
			"rules": [
			{
				"id": 1,
				"name": "security",
				"rule_type": "regular",
				"eligible_approvers": [
					{
						"id": 5,
						"name": "John Doe",
						"username": "jdoe",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						"web_url": "http://localhost/jdoe"
					},
					{
						"id": 50,
						"name": "Group Member 1",
						"username": "group_member_1",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						"web_url": "http://localhost/group_member_1"
					}
				],
				"approvals_required": 3,
				"source_rule": null,
				"users": [
					{
						"id": 5,
						"name": "John Doe",
						"username": "jdoe",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						"web_url": "http://localhost/jdoe"
					}
				],
				"groups": [
					{
						"id": 5,
						"name": "group1",
						"path": "group1",
						"description": "",
						"visibility": "public",
						"lfs_enabled": false,
						"avatar_url": null,
						"web_url": "http://localhost/groups/group1",
						"request_access_enabled": false,
						"full_name": "group1",
						"full_path": "group1",
						"parent_id": null,
						"ldap_cn": null,
						"ldap_access": null
					}
				],
				"contains_hidden_groups": false,
				"section": "test",
				"approved_by": [
					{
						"id": 5,
						"name": "John Doe",
						"username": "jdoe",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						"web_url": "http://localhost/jdoe"
					}
				],
				"approved": false
			}
		]
		}`)
	})

	approvals, resp, err := client.MergeRequestApprovals.GetApprovalState(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := &MergeRequestApprovalState{
		ApprovalRulesOverwritten: true,
		Rules: []*MergeRequestApprovalRule{
			{
				ID:       1,
				Name:     "security",
				RuleType: "regular",
				EligibleApprovers: []*BasicUser{
					{
						ID:        5,
						Name:      "John Doe",
						Username:  "jdoe",
						State:     "active",
						AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						WebURL:    "http://localhost/jdoe",
					},
					{
						ID:        50,
						Name:      "Group Member 1",
						Username:  "group_member_1",
						State:     "active",
						AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						WebURL:    "http://localhost/group_member_1",
					},
				},
				ApprovalsRequired: 3,
				Users: []*BasicUser{
					{
						ID:        5,
						Name:      "John Doe",
						Username:  "jdoe",
						State:     "active",
						AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						WebURL:    "http://localhost/jdoe",
					},
				},
				Groups: []*Group{
					{
						ID:                   5,
						Name:                 "group1",
						Path:                 "group1",
						Description:          "",
						Visibility:           PublicVisibility,
						LFSEnabled:           false,
						AvatarURL:            "",
						WebURL:               "http://localhost/groups/group1",
						RequestAccessEnabled: false,
						FullName:             "group1",
						FullPath:             "group1",
					},
				},
				Section: "test",
				ApprovedBy: []*BasicUser{
					{
						ID:        5,
						Name:      "John Doe",
						Username:  "jdoe",
						State:     "active",
						AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						WebURL:    "http://localhost/jdoe",
					},
				},
				Approved: false,
			},
		},
	}
	assert.Equal(t, want, approvals)
}

func TestCreateApprovalRules(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/1/approval_rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "security",
			"rule_type": "regular",
			"eligible_approvers": [
				{
					"id": 5,
					"name": "John Doe",
					"username": "jdoe",
					"state": "active",
					"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					"web_url": "http://localhost/jdoe"
				},
				{
					"id": 50,
					"name": "Group Member 1",
					"username": "group_member_1",
					"state": "active",
					"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					"web_url": "http://localhost/group_member_1"
				}
			],
			"approvals_required": 3,
			"source_rule": null,
			"users": [
				{
					"id": 5,
					"name": "John Doe",
					"username": "jdoe",
					"state": "active",
					"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					"web_url": "http://localhost/jdoe"
				}
			],
			"groups": [
				{
					"id": 5,
					"name": "group1",
					"path": "group1",
					"description": "",
					"visibility": "public",
					"lfs_enabled": false,
					"avatar_url": null,
					"web_url": "http://localhost/groups/group1",
					"request_access_enabled": false,
					"full_name": "group1",
					"full_path": "group1",
					"parent_id": null,
					"ldap_cn": null,
					"ldap_access": null
				}
			],
			"contains_hidden_groups": false
		}`)
	})

	opt := &CreateMergeRequestApprovalRuleOptions{
		Name:              Ptr("security"),
		ApprovalsRequired: Ptr(3),
		UserIDs:           &[]int{5, 50},
		GroupIDs:          &[]int{5},
	}

	rule, resp, err := client.MergeRequestApprovals.CreateApprovalRule(1, 1, opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := &MergeRequestApprovalRule{
		ID:       1,
		Name:     "security",
		RuleType: "regular",
		EligibleApprovers: []*BasicUser{
			{
				ID:        5,
				Name:      "John Doe",
				Username:  "jdoe",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
				WebURL:    "http://localhost/jdoe",
			},
			{
				ID:        50,
				Name:      "Group Member 1",
				Username:  "group_member_1",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
				WebURL:    "http://localhost/group_member_1",
			},
		},
		ApprovalsRequired: 3,
		Users: []*BasicUser{
			{
				ID:        5,
				Name:      "John Doe",
				Username:  "jdoe",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
				WebURL:    "http://localhost/jdoe",
			},
		},
		Groups: []*Group{
			{
				ID:                   5,
				Name:                 "group1",
				Path:                 "group1",
				Description:          "",
				Visibility:           PublicVisibility,
				LFSEnabled:           false,
				AvatarURL:            "",
				WebURL:               "http://localhost/groups/group1",
				RequestAccessEnabled: false,
				FullName:             "group1",
				FullPath:             "group1",
			},
		},
	}
	assert.Equal(t, want, rule)
}
