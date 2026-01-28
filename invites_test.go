package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListGroupPendingInvites(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListPendingInvitationsOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
	}

	projects, _, err := client.Invites.ListPendingGroupInvitations("test", opt)
	require.NoError(t, err)

	want := []*PendingInvite{{ID: 1}, {ID: 2}}
	assert.Equal(t, want, projects)
}

func TestGroupInvites(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"status": "success"}`)
	})

	opt := &InvitesOptions{
		Email: Ptr("example@member.org"),
	}

	projects, _, err := client.Invites.GroupInvites("test", opt)
	require.NoError(t, err)

	want := &InvitesResult{Status: "success"}
	assert.Equal(t, want, projects)
}

func TestGroupInvitesError(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"status": "error","message": {"example@member.org": "Already invited"}}`)
	})

	opt := &InvitesOptions{
		Email: Ptr("example@member.org"),
	}

	projects, _, err := client.Invites.GroupInvites("test", opt)
	require.NoError(t, err)

	want := &InvitesResult{Status: "error", Message: map[string]string{"example@member.org": "Already invited"}}
	assert.Equal(t, want, projects)
}

func TestListProjectPendingInvites(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListPendingInvitationsOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
	}

	projects, _, err := client.Invites.ListPendingProjectInvitations("test", opt)
	require.NoError(t, err)

	want := []*PendingInvite{{ID: 1}, {ID: 2}}
	assert.Equal(t, want, projects)
}

func TestProjectInvites(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"status": "success"}`)
	})

	opt := &InvitesOptions{
		Email: Ptr("example@member.org"),
	}

	projects, _, err := client.Invites.ProjectInvites("test", opt)
	require.NoError(t, err)

	want := &InvitesResult{Status: "success"}
	assert.Equal(t, want, projects)
}

func TestProjectInvitesError(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"status": "error","message": {"example@member.org": "Already invited"}}`)
	})

	opt := &InvitesOptions{
		Email: Ptr("example@member.org"),
	}

	projects, _, err := client.Invites.ProjectInvites("test", opt)
	require.NoError(t, err)

	want := &InvitesResult{Status: "error", Message: map[string]string{"example@member.org": "Already invited"}}
	assert.Equal(t, want, projects)
}
