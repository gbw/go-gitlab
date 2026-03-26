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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSystemhookPush(t *testing.T) {
	t.Parallel()
	payload := loadFixture(t, "testdata/systemhooks/push.json")

	parsedEvent, err := ParseSystemhook(payload)
	require.NoError(t, err)

	event, ok := parsedEvent.(*PushSystemEvent)
	require.True(t, ok)
	assert.Equal(t, eventObjectKindPush, event.EventName)
}

func TestParseSystemhookTagPush(t *testing.T) {
	t.Parallel()
	payload := loadFixture(t, "testdata/systemhooks/tag_push.json")

	parsedEvent, err := ParseSystemhook(payload)
	require.NoError(t, err)

	event, ok := parsedEvent.(*TagPushSystemEvent)
	require.True(t, ok)
	assert.Equal(t, eventObjectKindTagPush, event.EventName)
}

func TestParseSystemhookMergeRequest(t *testing.T) {
	t.Parallel()
	payload := loadFixture(t, "testdata/systemhooks/merge_request.json")

	parsedEvent, err := ParseSystemhook(payload)
	require.NoError(t, err)

	event, ok := parsedEvent.(*MergeEvent)
	require.True(t, ok)
	assert.Equal(t, eventObjectKindMergeRequest, event.ObjectKind)
}

func TestParseSystemhookRepositoryUpdate(t *testing.T) {
	t.Parallel()
	payload := loadFixture(t, "testdata/systemhooks/repository_update.json")

	parsedEvent, err := ParseSystemhook(payload)
	require.NoError(t, err)

	event, ok := parsedEvent.(*RepositoryUpdateSystemEvent)
	require.True(t, ok)
	assert.Equal(t, "repository_update", event.EventName)
}

func TestParseSystemhookProject(t *testing.T) {
	t.Parallel()
	tests := []struct {
		event   string
		payload []byte
	}{
		{"project_create", loadFixture(t, "testdata/systemhooks/project_create.json")},
		{"project_update", loadFixture(t, "testdata/systemhooks/project_update.json")},
		{"project_destroy", loadFixture(t, "testdata/systemhooks/project_destroy.json")},
		{"project_transfer", loadFixture(t, "testdata/systemhooks/project_transfer.json")},
		{"project_rename", loadFixture(t, "testdata/systemhooks/project_rename.json")},
	}
	for _, tc := range tests {
		t.Run(tc.event, func(t *testing.T) {
			t.Parallel()

			parsedEvent, err := ParseSystemhook(tc.payload)
			require.NoError(t, err)
			event, ok := parsedEvent.(*ProjectSystemEvent)
			require.True(t, ok)
			assert.Equal(t, tc.event, event.EventName)
		})
	}
}

func TestParseSystemhookGroup(t *testing.T) {
	t.Parallel()
	tests := []struct {
		event   string
		payload []byte
	}{
		{"group_create", loadFixture(t, "testdata/systemhooks/group_create.json")},
		{"group_destroy", loadFixture(t, "testdata/systemhooks/group_destroy.json")},
		{"group_rename", loadFixture(t, "testdata/systemhooks/group_rename.json")},
	}
	for _, tc := range tests {
		t.Run(tc.event, func(t *testing.T) {
			t.Parallel()

			parsedEvent, err := ParseSystemhook(tc.payload)
			require.NoError(t, err)
			event, ok := parsedEvent.(*GroupSystemEvent)
			require.True(t, ok)
			assert.Equal(t, tc.event, event.EventName)
		})
	}
}

func TestParseSystemhookUser(t *testing.T) {
	t.Parallel()
	tests := []struct {
		event   string
		payload []byte
	}{
		{"user_create", loadFixture(t, "testdata/systemhooks/user_create.json")},
		{"user_destroy", loadFixture(t, "testdata/systemhooks/user_destroy.json")},
		{"user_rename", loadFixture(t, "testdata/systemhooks/user_rename.json")},
		{"user_failed_login", loadFixture(t, "testdata/systemhooks/user_failed_login.json")},
	}
	for _, tc := range tests {
		t.Run(tc.event, func(t *testing.T) {
			t.Parallel()

			parsedEvent, err := ParseSystemhook(tc.payload)
			require.NoError(t, err)
			event, ok := parsedEvent.(*UserSystemEvent)
			require.True(t, ok)
			assert.Equal(t, tc.event, event.EventName)
		})
	}
}

func TestParseSystemhookUserGroup(t *testing.T) {
	t.Parallel()
	tests := []struct {
		event   string
		payload []byte
	}{
		{"user_add_to_group", loadFixture(t, "testdata/systemhooks/user_add_to_group.json")},
		{"user_remove_from_group", loadFixture(t, "testdata/systemhooks/user_remove_from_group.json")},
		{"user_update_for_group", loadFixture(t, "testdata/systemhooks/user_update_for_group.json")},
	}
	for _, tc := range tests {
		t.Run(tc.event, func(t *testing.T) {
			t.Parallel()

			parsedEvent, err := ParseSystemhook(tc.payload)
			require.NoError(t, err)
			event, ok := parsedEvent.(*UserGroupSystemEvent)
			require.True(t, ok)
			assert.Equal(t, tc.event, event.EventName)
		})
	}
}

func TestParseSystemhookUserTeam(t *testing.T) {
	t.Parallel()
	tests := []struct {
		event   string
		payload []byte
	}{
		{"user_add_to_team", loadFixture(t, "testdata/systemhooks/user_add_to_team.json")},
		{"user_remove_from_team", loadFixture(t, "testdata/systemhooks/user_remove_from_team.json")},
		{"user_update_for_team", loadFixture(t, "testdata/systemhooks/user_update_for_team.json")},
	}
	for _, tc := range tests {
		t.Run(tc.event, func(t *testing.T) {
			t.Parallel()

			parsedEvent, err := ParseSystemhook(tc.payload)
			require.NoError(t, err)
			event, ok := parsedEvent.(*UserTeamSystemEvent)
			require.True(t, ok)
			assert.Equal(t, tc.event, event.EventName)
		})
	}
}

func TestParseHookSystemHook(t *testing.T) {
	t.Parallel()
	parsedEvent1, err := ParseHook("System Hook", loadFixture(t, "testdata/systemhooks/merge_request.json"))
	require.NoError(t, err)
	parsedEvent2, err := ParseSystemhook(loadFixture(t, "testdata/systemhooks/merge_request.json"))
	require.NoError(t, err)
	assert.Equal(t, parsedEvent1, parsedEvent2)
}
