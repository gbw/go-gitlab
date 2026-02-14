package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetGroupMattermostIntegration(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/integrations/mattermost", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"title": "Mattermost",
			"slug": "mattermost",
			"created_at": "2023-01-01T00:00:00.000Z",
			"updated_at": "2023-01-02T00:00:00.000Z",
			"active": true,
			"notify_only_broken_pipelines": true,
			"branches_to_be_notified": "default",
			"push_events": true,
			"issues_events": true,
			"confidential_issues_events": true,
			"merge_requests_events": true,
			"tag_push_events": true,
			"note_events": true,
			"confidential_note_events": true,
			"pipeline_events": true,
			"wiki_page_events": true,
			"deployment_events": true,
			"alert_events": true,
			"vulnerability_events": true,
			"labels_to_be_notified": "label1,label2",
			"labels_to_be_notified_behavior": "match_all",
			"notify_only_default_branch": true,
			"properties": {
				"webhook": "http://mattermost.example.com/hooks/xxx",
				"channel": "#alerts",
				"username": "GitLab",
				"push_channel": "push",
				"issue_channel": "issue",
				"confidential_issue_channel": "confidential_issue",
				"merge_request_channel": "merge_request",
				"note_channel": "note",
				"confidential_note_channel": "confidential_note",
				"tag_push_channel": "tag_push",
				"pipeline_channel": "pipeline_channel",
				"wiki_page_channel": "wiki_page",
				"deployment_channel": "deployment",
				"alert_channel": "alert",
				"vulnerability_channel": "vulnerability"
			}
		}`)
	})

	gmi, resp, err := client.Integrations.GetGroupMattermostIntegration(1)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	createdAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00.000Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2023-01-02T00:00:00.000Z")

	want := &GroupMattermostIntegration{
		Integration: Integration{
			ID:                       1,
			Title:                    "Mattermost",
			Slug:                     "mattermost",
			CreatedAt:                &createdAt,
			UpdatedAt:                &updatedAt,
			Active:                   true,
			PushEvents:               true,
			IssuesEvents:             true,
			ConfidentialIssuesEvents: true,
			MergeRequestsEvents:      true,
			TagPushEvents:            true,
			NoteEvents:               true,
			ConfidentialNoteEvents:   true,
			PipelineEvents:           true,
			WikiPageEvents:           true,
			DeploymentEvents:         true,
			AlertEvents:              true,
			VulnerabilityEvents:      true,
		},
		NotifyOnlyBrokenPipelines:  true,
		BranchesToBeNotified:       "default",
		LabelsToBeNotified:         "label1,label2",
		LabelsToBeNotifiedBehavior: "match_all",
		NotifyOnlyDefaultBranch:    true,
		Properties: &GroupMattermostIntegrationProperties{
			WebHook:                  "http://mattermost.example.com/hooks/xxx",
			Username:                 "GitLab",
			Channel:                  "#alerts",
			PushChannel:              "push",
			IssueChannel:             "issue",
			ConfidentialIssueChannel: "confidential_issue",
			MergeRequestChannel:      "merge_request",
			NoteChannel:              "note",
			ConfidentialNoteChannel:  "confidential_note",
			TagPushChannel:           "tag_push",
			PipelineChannel:          "pipeline_channel",
			WikiPageChannel:          "wiki_page",
			DeploymentChannel:        "deployment",
			AlertChannel:             "alert",
			VulnerabilityChannel:     "vulnerability",
		},
	}

	assert.Equal(t, want, gmi)
}

func TestSetGroupMattermostIntegration(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/integrations/mattermost", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{
			"id": 1,
			"title": "Mattermost",
			"slug": "mattermost",
			"created_at": "2023-01-01T00:00:00.000Z",
			"updated_at": "2023-01-02T00:00:00.000Z",
			"active": true,
			"push_events": true,
			"properties": {
				"webhook": "http://mattermost.example.com/hooks/xxx",
				"channel": "#alerts",
				"username": "GitLab"
			}
		}`)
	})

	opt := &GroupMattermostIntegrationOptions{
		WebHook: Ptr("http://mattermost.example.com/hooks/xxx"),
		Channel: Ptr("#alerts"),
	}

	gmi, resp, err := client.Integrations.SetGroupMattermostIntegration(1, opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	createdAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00.000Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2023-01-02T00:00:00.000Z")

	want := &GroupMattermostIntegration{
		Integration: Integration{
			ID:         1,
			Title:      "Mattermost",
			Slug:       "mattermost",
			CreatedAt:  &createdAt,
			UpdatedAt:  &updatedAt,
			Active:     true,
			PushEvents: true,
		},
		Properties: &GroupMattermostIntegrationProperties{
			WebHook:  "http://mattermost.example.com/hooks/xxx",
			Channel:  "#alerts",
			Username: "GitLab",
		},
	}

	assert.Equal(t, want, gmi)
}

func TestDeleteGroupMattermostIntegration(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/integrations/mattermost", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Integrations.DeleteGroupMattermostIntegration(1)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestGetGroupMattermostSlashCommandsIntegration(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/integrations/mattermost-slash-commands", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"title": "Mattermost Slash Commands",
			"slug": "mattermost-slash-commands",
			"created_at": "2023-01-01T00:00:00.000Z",
			"updated_at": "2023-01-02T00:00:00.000Z",
			"token": "secret"
		}`)
	})

	integration, resp, err := client.Integrations.GetGroupMattermostSlashCommandsIntegration(1)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	createdAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00.000Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2023-01-02T00:00:00.000Z")

	want := &GroupMattermostSlashCommandsIntegration{
		ID:        1,
		Title:     "Mattermost Slash Commands",
		Slug:      "mattermost-slash-commands",
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Token:     "secret",
	}

	assert.Equal(t, want, integration)
}

func TestSetGroupMattermostSlashCommandsIntegration(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/integrations/mattermost-slash-commands", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testBodyJSON(t, r, map[string]string{"token": "new-token"})
		fmt.Fprint(w, `{
			"id": 1,
			"title": "Mattermost Slash Commands",
			"slug": "mattermost-slash-commands",
			"created_at": "2023-01-01T00:00:00.000Z",
			"updated_at": "2023-01-02T00:00:00.000Z",
			"token": "new-token"
		}`)
	})

	opt := &GroupMattermostSlashCommandsIntegrationOptions{
		Token: Ptr("new-token"),
	}

	integration, resp, err := client.Integrations.SetGroupMattermostSlashCommandsIntegration(1, opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	createdAt, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00.000Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2023-01-02T00:00:00.000Z")

	want := &GroupMattermostSlashCommandsIntegration{
		ID:        1,
		Title:     "Mattermost Slash Commands",
		Slug:      "mattermost-slash-commands",
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Token:     "new-token",
	}

	assert.Equal(t, want, integration)
}

func TestDeleteGroupMattermostSlashCommandsIntegration(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/integrations/mattermost-slash-commands", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Integrations.DeleteGroupMattermostSlashCommandsIntegration(1)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
