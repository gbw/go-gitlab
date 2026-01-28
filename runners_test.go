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
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDisableRunner(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/runners/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Runners.DisableProjectRunner(1, 2, nil)
	require.NoError(t, err)
}

func TestListRunnersJobs(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/1/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, exampleListRunnerJobs)
	})

	opt := &ListRunnerJobsOptions{}

	jobs, _, err := client.Runners.ListRunnerJobs(1, opt)
	require.NoError(t, err)

	pipeline := JobPipeline{
		ID:        8777,
		ProjectID: 3252,
		Ref:       DefaultBranch,
		Sha:       "6c016b801a88f4bd31f927fc045b5c746a6f823e",
		Status:    "failed",
	}

	want := []*Job{
		{
			ID:             1,
			Status:         "failed",
			Stage:          "test",
			Name:           "run_tests",
			Ref:            DefaultBranch,
			CreatedAt:      Ptr(time.Date(2021, time.October, 22, 11, 59, 25, 201000000, time.UTC)),
			StartedAt:      Ptr(time.Date(2021, time.October, 22, 11, 59, 33, 660000000, time.UTC)),
			FinishedAt:     Ptr(time.Date(2021, time.October, 22, 15, 59, 25, 201000000, time.UTC)),
			Duration:       171.540594,
			QueuedDuration: 2.535766,
			User: &User{
				ID:          368,
				Name:        "John SMITH",
				Username:    "john.smith",
				AvatarURL:   "https://gitlab.example.com/uploads/-/system/user/avatar/368/avatar.png",
				State:       "blocked",
				WebURL:      "https://gitlab.example.com/john.smith",
				PublicEmail: "john.smith@example.com",
			},
			Commit: &Commit{
				ID:             "6c016b801a88f4bd31f927fc045b5c746a6f823e",
				ShortID:        "6c016b80",
				CreatedAt:      Ptr(time.Date(2018, time.March, 21, 14, 41, 0, 0, time.UTC)),
				ParentIDs:      []string{"6008b4902d40799ab11688e502d9f1f27f6d2e18"},
				Title:          "Update env for specific runner",
				Message:        "Update env for specific runner\n",
				AuthorName:     "John SMITH",
				AuthorEmail:    "john.smith@example.com",
				AuthoredDate:   Ptr(time.Date(2018, time.March, 21, 14, 41, 0, 0, time.UTC)),
				CommitterName:  "John SMITH",
				CommitterEmail: "john.smith@example.com",
				CommittedDate:  Ptr(time.Date(2018, time.March, 21, 14, 41, 0, 0, time.UTC)),
				WebURL:         "https://gitlab.example.com/awesome/packages/common/-/commit/6c016b801a88f4bd31f927fc045b5c746a6f823e",
			},
			Pipeline: pipeline,
			WebURL:   "https://gitlab.example.com/awesome/packages/common/-/jobs/14606",
			Project: &Project{
				ID:                3252,
				Description:       "Common nodejs paquet for producer",
				Name:              "common",
				NameWithNamespace: "awesome",
				Path:              "common",
				PathWithNamespace: "awesome",
				CreatedAt:         Ptr(time.Date(2018, time.February, 13, 9, 21, 48, 107000000, time.UTC)),
			},
		},
	}
	assert.Equal(t, want[0], jobs[0])
}

func TestRemoveRunner(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Runners.RemoveRunner(1, nil)
	require.NoError(t, err)
}

func TestUpdateRunnersDetails(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/6", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, exampleDetailResponse)
	})

	opt := &UpdateRunnerDetailsOptions{}

	details, _, err := client.Runners.UpdateRunnerDetails(6, opt, nil)
	require.NoError(t, err)

	projects := []RunnerDetailsProject{{
		ID:                1,
		Name:              "GitLab Community Edition",
		NameWithNamespace: "GitLab.org / GitLab Community Edition",
		Path:              "gitlab-ce",
		PathWithNamespace: "gitlab-org/gitlab-ce",
	}}

	want := &RunnerDetails{
		Active:         true,
		Description:    "test-1-20150125-test",
		ID:             6,
		IsShared:       false,
		RunnerType:     "project_type",
		ContactedAt:    Ptr(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
		Online:         true,
		Status:         "online",
		Token:          "205086a8e3b9a2b818ffac9b89d102",
		TagList:        []string{"ruby", "mysql"},
		RunUntagged:    true,
		AccessLevel:    "ref_protected",
		Projects:       projects,
		MaximumTimeout: 3600,
		Locked:         false,
	}
	assert.Equal(t, want, details)
}

func TestGetRunnerDetails(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/6", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, exampleDetailResponse)
	})

	details, _, err := client.Runners.GetRunnerDetails(6, nil)
	require.NoError(t, err)

	projects := []RunnerDetailsProject{{
		ID:                1,
		Name:              "GitLab Community Edition",
		NameWithNamespace: "GitLab.org / GitLab Community Edition",
		Path:              "gitlab-ce",
		PathWithNamespace: "gitlab-org/gitlab-ce",
	}}

	want := &RunnerDetails{
		Active:         true,
		Description:    "test-1-20150125-test",
		ID:             6,
		IsShared:       false,
		RunnerType:     "project_type",
		ContactedAt:    Ptr(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
		Online:         true,
		Status:         "online",
		Token:          "205086a8e3b9a2b818ffac9b89d102",
		TagList:        []string{"ruby", "mysql"},
		RunUntagged:    true,
		AccessLevel:    "ref_protected",
		Projects:       projects,
		MaximumTimeout: 3600,
		Locked:         false,
	}
	assert.Equal(t, want, details)
}

func TestRegisterNewRunner(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, exampleRegisterNewRunner)
	})

	opt := &RegisterNewRunnerOptions{}

	runner, resp, err := client.Runners.RegisterNewRunner(opt, nil)
	require.NoError(t, err)

	want := &Runner{
		ID:             12345,
		Token:          "6337ff461c94fd3fa32ba3b1ff4125",
		TokenExpiresAt: Ptr(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	assert.Equal(t, want, runner)
	assert.Equal(t, 201, resp.StatusCode)
}

// Similar to TestRegisterNewRunner but sends info struct and some extra other
// fields too.
func TestRegisterNewRunnerInfo(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"id": 53,
			"description": "some description",
			"active": true,
			"name": "some name",
			"online": true,
			"status": "online",
			"token": "1111122222333333444444",
			"token_expires_at": "2016-01-25T16:39:48.166Z"
		  }`)
	})

	opt := &RegisterNewRunnerOptions{
		Token:       Ptr("6337ff461c94fd3fa32ba3b1ff4125"),
		Description: Ptr("some description"),
		Info: &RegisterNewRunnerInfoOptions{
			Ptr("some name"),
			Ptr("13.7.0"),
			Ptr("943fc252"),
			Ptr("linux"),
			Ptr("amd64"),
		},
		Active:         Ptr(true),
		Locked:         Ptr(true),
		RunUntagged:    Ptr(false),
		MaximumTimeout: Ptr(int64(45)),
	}
	runner, resp, err := client.Runners.RegisterNewRunner(opt, nil)
	require.NoError(t, err)

	want := &Runner{
		ID:             53,
		Description:    "some description",
		Active:         true,
		Name:           "some name",
		Online:         true,
		Status:         "online",
		Token:          "1111122222333333444444",
		TokenExpiresAt: Ptr(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	assert.Equal(t, want, runner)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestDeleteRegisteredRunner(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	opt := &DeleteRegisteredRunnerOptions{}

	resp, err := client.Runners.DeleteRegisteredRunner(opt, nil)
	require.NoError(t, err)

	assert.Equal(t, 204, resp.StatusCode)
}

func TestDeleteRegisteredRunnerByID(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/11111", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	rid := int64(11111)

	resp, err := client.Runners.DeleteRegisteredRunnerByID(rid, nil)
	require.NoError(t, err)

	assert.Equal(t, 204, resp.StatusCode)
}

func TestVerifyRegisteredRunner(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/verify", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
	})

	opt := &VerifyRegisteredRunnerOptions{}

	resp, err := client.Runners.VerifyRegisteredRunner(opt, nil)
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestResetInstanceRunnerRegistrationToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/reset_registration_token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"token": "6337ff461c94fd3fa32ba3b1ff4125",
			"token_expires_at": "2016-01-25T16:39:48.166Z"
		}`)
	})

	token, resp, err := client.Runners.ResetInstanceRunnerRegistrationToken(nil)
	require.NoError(t, err)

	want := &RunnerRegistrationToken{
		Token:          Ptr("6337ff461c94fd3fa32ba3b1ff4125"),
		TokenExpiresAt: Ptr(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	assert.Equal(t, want, token)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestResetGroupRunnerRegistrationToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/foobar/runners/reset_registration_token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"token": "6337ff461c94fd3fa32ba3b1ff4125",
			"token_expires_at": "2016-01-25T16:39:48.166Z"
		}`)
	})

	token, resp, err := client.Runners.ResetGroupRunnerRegistrationToken("foobar", nil)
	require.NoError(t, err)

	want := &RunnerRegistrationToken{
		Token:          Ptr("6337ff461c94fd3fa32ba3b1ff4125"),
		TokenExpiresAt: Ptr(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	assert.Equal(t, want, token)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestResetProjectRunnerRegistrationToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/9/runners/reset_registration_token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"token": "6337ff461c94fd3fa32ba3b1ff4125",
			"token_expires_at": "2016-01-25T16:39:48.166Z"
		}`)
	})

	token, resp, err := client.Runners.ResetProjectRunnerRegistrationToken("9", nil)
	require.NoError(t, err)

	want := &RunnerRegistrationToken{
		Token:          Ptr("6337ff461c94fd3fa32ba3b1ff4125"),
		TokenExpiresAt: Ptr(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	assert.Equal(t, want, token)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestResetRunnerAuthenticationToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/42/reset_authentication_token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"token": "6337ff461c94fd3fa32ba3b1ff4125",
			"token_expires_at": "2016-01-25T16:39:48.166Z"
		}`)
	})

	token, resp, err := client.Runners.ResetRunnerAuthenticationToken(42, nil)
	require.NoError(t, err)

	want := &RunnerAuthenticationToken{
		Token:          Ptr("6337ff461c94fd3fa32ba3b1ff4125"),
		TokenExpiresAt: Ptr(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	assert.Equal(t, want, token)
	assert.Equal(t, 201, resp.StatusCode)
}
