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
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/users/1"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_user.json")
	})

	user, _, err := client.Users.GetUser(1, GetUsersOptions{})
	require.NoError(t, err)

	want := &User{
		ID:           1,
		Username:     "john_smith",
		Name:         "John Smith",
		State:        "active",
		WebURL:       "http://localhost:3000/john_smith",
		CreatedAt:    Ptr(time.Date(2012, time.May, 23, 8, 0o0, 58, 0, time.UTC)),
		Bio:          "Bio of John Smith",
		Location:     "USA",
		PublicEmail:  "john@example.com",
		Skype:        "john_smith",
		Linkedin:     "john_smith",
		Twitter:      "john_smith",
		WebsiteURL:   "john_smith.example.com",
		Organization: "Smith Inc",
		JobTitle:     "Operations Specialist",
		AvatarURL:    "http://localhost:3000/uploads/user/avatar/1/cd8.jpeg",
	}
	require.Equal(t, want, user)
}

func TestGetUserAdmin(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/users/1"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_user_admin.json")
	})

	user, _, err := client.Users.GetUser(1, GetUsersOptions{})
	require.NoError(t, err)

	lastActivityOn := ISOTime(time.Date(2012, time.May, 23, 0, 0, 0, 0, time.UTC))
	currentSignInIP := net.ParseIP("8.8.8.8")
	lastSignInIP := net.ParseIP("2001:db8::68")

	want := &User{
		ID:               1,
		Username:         "john_smith",
		Email:            "john@example.com",
		Name:             "John Smith",
		State:            "active",
		WebURL:           "http://localhost:3000/john_smith",
		CreatedAt:        Ptr(time.Date(2012, time.May, 23, 8, 0, 58, 0, time.UTC)),
		Bio:              "Bio of John Smith",
		Location:         "USA",
		PublicEmail:      "john@example.com",
		Skype:            "john_smith",
		Linkedin:         "john_smith",
		Twitter:          "john_smith",
		WebsiteURL:       "john_smith.example.com",
		Organization:     "Smith Inc",
		JobTitle:         "Operations Specialist",
		ThemeID:          1,
		LastActivityOn:   &lastActivityOn,
		ColorSchemeID:    2,
		IsAdmin:          true,
		IsAuditor:        true,
		AvatarURL:        "http://localhost:3000/uploads/user/avatar/1/index.jpg",
		CanCreateGroup:   true,
		CanCreateProject: true,
		ProjectsLimit:    100,
		CurrentSignInAt:  Ptr(time.Date(2012, time.June, 2, 6, 36, 55, 0, time.UTC)),
		CurrentSignInIP:  &currentSignInIP,
		LastSignInAt:     Ptr(time.Date(2012, time.June, 1, 11, 41, 1, 0, time.UTC)),
		LastSignInIP:     &lastSignInIP,
		ConfirmedAt:      Ptr(time.Date(2012, time.May, 23, 9, 0o5, 22, 0, time.UTC)),
		TwoFactorEnabled: true,
		Note:             "DMCA Request: 2018-11-05 | DMCA Violation | Abuse | https://gitlab.zendesk.com/agent/tickets/123",
		Identities:       []*UserIdentity{{Provider: "github", ExternUID: "2435223452345"}},
		NamespaceID:      42,
	}
	require.Equal(t, want, user)
}

func TestCreatedBy(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/users/2"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_user_bot.json")
	})

	user, _, err := client.Users.GetUser(2, GetUsersOptions{})
	require.NoError(t, err)

	lastActivityOn := ISOTime(time.Date(2012, time.May, 23, 0, 0, 0, 0, time.UTC))

	want := &User{
		ID:        2,
		Username:  "project_1_bot_3cca1d029554e372cf8f39df76bf507d",
		Email:     "project_1_bot_3cca1d029554e372cf8f39df76bf507d@example.com",
		Name:      "John Bot",
		State:     "active",
		WebURL:    "http://localhost:3000/project_1_bot_3cca1d029554e372cf8f39df76bf507d",
		CreatedAt: Ptr(time.Date(2012, time.May, 23, 8, 0o0, 58, 0, time.UTC)),
		Bot:       true,
		// Bio:          "Bio of John Smith",
		// Location:     "USA",
		// PublicEmail:  "john@example.com",
		// Skype:        "john_smith",
		// Linkedin:     "john_smith",
		// Twitter:      "john_smith",
		// WebsiteURL:   "john_smith.example.com",
		// Organization: "Smith Inc",
		// JobTitle:     "Operations Specialist",
		ThemeID:        3,
		LastActivityOn: &lastActivityOn,
		ColorSchemeID:  1,
		IsAdmin:        false,
		AvatarURL:      "http://localhost:3000/uploads/user/avatar/2/index.jpg",
		ConfirmedAt:    Ptr(time.Date(2012, time.May, 23, 8, 0o0, 58, 0, time.UTC)),
		Identities:     []*UserIdentity{},
		NamespaceID:    4,
		Locked:         false,
		CreatedBy: &BasicUser{
			ID:        1,
			Username:  "john_smith",
			Name:      "John Smith",
			State:     "active",
			Locked:    false,
			WebURL:    "http://localhost:3000/john_smith",
			AvatarURL: "http://localhost:3000/uploads/user/avatar/1/cd8.jpeg",
		},
	}
	require.Equal(t, want, user)
}

func TestBlockUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.BlockUser(1)
	if err != nil {
		t.Errorf("Users.BlockUser returned error: %v", err)
	}
}

func TestBlockUser_UserNotFound(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.BlockUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.BlockUser error.\nExpected: %+v\nGot: %+v", ErrUserNotFound, err)
	}
}

func TestBlockUser_BlockPrevented(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.BlockUser(1)
	if !errors.Is(err, ErrUserBlockPrevented) {
		t.Errorf("Users.BlockUser error.\nExpected: %+v\nGot: %+v", ErrUserBlockPrevented, err)
	}
}

func TestBlockUser_UnknownError(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.BlockUser(1)
	if err.Error() != want {
		t.Errorf("Users.BlockUser error.\nExpected: %s\nGot: %v", want, err)
	}
}

func TestUnblockUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.UnblockUser(1)
	if err != nil {
		t.Errorf("Users.UnblockUser returned error: %v", err)
	}
}

func TestUnblockUser_UserNotFound(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.UnblockUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.UnblockUser error.\nExpected: %v\nGot: %v", ErrUserNotFound, err)
	}
}

func TestUnblockUser_UnblockPrevented(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.UnblockUser(1)
	if !errors.Is(err, ErrUserUnblockPrevented) {
		t.Errorf("Users.UnblockUser error.\nExpected: %v\nGot: %v", ErrUserUnblockPrevented, err)
	}
}

func TestUnblockUser_UnknownError(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.UnblockUser(1)
	if err.Error() != want {
		t.Errorf("Users.UnblockUser error.\nExpected: %s\n\tGot: %v", want, err)
	}
}

func TestBanUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.BlockUser(1)
	if err != nil {
		t.Errorf("Users.BlockUser returned error: %v", err)
	}
}

func TestBanUser_UserNotFound(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/ban", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.BanUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.BanUser error.\nExpected: %+v\nGot: %+v", ErrUserNotFound, err)
	}
}

func TestBanUser_UnknownError(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/ban", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.BanUser(1)
	if err.Error() != want {
		t.Errorf("Users.BanUSer error.\nExpected: %s\nGot: %v", want, err)
	}
}

func TestUnbanUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/unban", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.UnbanUser(1)
	if err != nil {
		t.Errorf("Users.UnbanUser returned error: %v", err)
	}
}

func TestUnbanUser_UserNotFound(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/unban", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.UnbanUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.UnbanUser error.\nExpected: %v\nGot: %v", ErrUserNotFound, err)
	}
}

func TestUnbanUser_UnknownError(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/unban", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.UnbanUser(1)
	if err.Error() != want {
		t.Errorf("Users.UnbanUser error.\nExpected: %s\n\tGot: %v", want, err)
	}
}

func TestDeactivateUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/deactivate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.DeactivateUser(1)
	if err != nil {
		t.Errorf("Users.DeactivateUser returned error: %v", err)
	}
}

func TestDeactivateUser_UserNotFound(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/deactivate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.DeactivateUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.DeactivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserNotFound, err)
	}
}

func TestDeactivateUser_DeactivatePrevented(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/deactivate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.DeactivateUser(1)
	if !errors.Is(err, ErrUserDeactivatePrevented) {
		t.Errorf("Users.DeactivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserDeactivatePrevented, err)
	}
}

func TestActivateUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/activate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.ActivateUser(1)
	if err != nil {
		t.Errorf("Users.ActivateUser returned error: %v", err)
	}
}

func TestActivateUser_ActivatePrevented(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/activate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.ActivateUser(1)
	if !errors.Is(err, ErrUserActivatePrevented) {
		t.Errorf("Users.ActivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserActivatePrevented, err)
	}
}

func TestActivateUser_UserNotFound(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/activate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.ActivateUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.ActivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserNotFound, err)
	}
}

func TestApproveUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/approve", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.ApproveUser(1)
	if err != nil {
		t.Errorf("Users.ApproveUser returned error: %v", err)
	}
}

func TestApproveUser_UserNotFound(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/approve", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.ApproveUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.ApproveUser error.\nExpected: %v\nGot: %v", ErrUserNotFound, err)
	}
}

func TestApproveUser_ApprovePrevented(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/approve", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.ApproveUser(1)
	if !errors.Is(err, ErrUserApprovePrevented) {
		t.Errorf("Users.ApproveUser error.\nExpected: %v\nGot: %v", ErrUserApprovePrevented, err)
	}
}

func TestApproveUser_UnknownError(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/approve", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.ApproveUser(1)
	if err.Error() != want {
		t.Errorf("Users.ApproveUser error.\nExpected: %s\n\tGot: %v", want, err)
	}
}

func TestRejectUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/reject", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
	})

	err := client.Users.RejectUser(1)
	if err != nil {
		t.Errorf("Users.RejectUser returned error: %v", err)
	}
}

func TestRejectUser_UserNotFound(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/reject", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.RejectUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.RejectUser error.\nExpected: %v\nGot: %v", ErrUserNotFound, err)
	}
}

func TestRejectUser_RejectPrevented(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/reject", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.RejectUser(1)
	if !errors.Is(err, ErrUserRejectPrevented) {
		t.Errorf("Users.RejectUser error.\nExpected: %v\nGot: %v", ErrUserRejectPrevented, err)
	}
}

func TestRejectUser_Conflict(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/reject", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusConflict)
	})

	err := client.Users.RejectUser(1)
	if !errors.Is(err, ErrUserConflict) {
		t.Errorf("Users.RejectUser error.\nExpected: %v\nGot: %v", ErrUserConflict, err)
	}
}

func TestRejectUser_UnknownError(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/reject", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.RejectUser(1)
	if err.Error() != want {
		t.Errorf("Users.RejectUser error.\nExpected: %s\n\tGot: %v", want, err)
	}
}

func TestGetMemberships(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/memberships", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_user_memberships.json")
	})

	opt := new(GetUserMembershipOptions)

	memberships, _, err := client.Users.GetUserMemberships(1, opt)
	require.NoError(t, err)

	want := []*UserMembership{{SourceID: 1, SourceName: "Project one", SourceType: "Project", AccessLevel: 20}, {SourceID: 3, SourceName: "Group three", SourceType: "Namespace", AccessLevel: 20}}
	assert.Equal(t, want, memberships)
}

func TestGetUserAssociationsCount(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/users/1/associations_count"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_user_associations_count.json")
	})

	userAssociationsCount, _, err := client.Users.GetUserAssociationsCount(1)
	require.NoError(t, err)

	want := &UserAssociationsCount{
		GroupsCount:        1,
		ProjectsCount:      2,
		IssuesCount:        3,
		MergeRequestsCount: 4,
	}
	require.Equal(t, want, userAssociationsCount)
}

func TestGetSingleSSHKeyForUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/users/1/keys/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		{
			"id": 1,
			"title": "Public key",
			"key": "ssh-rsa AAAA...",
			"created_at": "2014-08-01T14:47:39.080Z",
			"usage_type": "auth"
		}`)
	})

	sshKey, _, err := client.Users.GetSSHKeyForUser(1, 1)
	if err != nil {
		t.Errorf("Users.GetSSHKeyForUser returned an error: %v", err)
	}

	wantCreatedAt := time.Date(2014, 8, 1, 14, 47, 39, 80000000, time.UTC)

	want := &SSHKey{
		ID:        1,
		Title:     "Public key",
		Key:       "ssh-rsa AAAA...",
		UsageType: "auth",
		CreatedAt: &wantCreatedAt,
	}

	if !reflect.DeepEqual(want, sshKey) {
		t.Errorf("Users.GetSSHKeyForUser returned %+v, want %+v", sshKey, want)
	}
}

func TestDisableUser2FA(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%susers/1/disable_two_factor", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.Users.DisableTwoFactor(1)
	if err != nil {
		t.Errorf("Users.DisableTwoFactor returned error: %v", err)
	}
}

func TestCreateUserRunner(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := fmt.Sprintf("/%suser/runners", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`
    {
      "id": 1234,
      "token": "glrt-1234567890ABCD",
      "token_expires_at":null
    }`))
	})

	createRunnerOpts := &CreateUserRunnerOptions{
		ProjectID:  Ptr(1),
		RunnerType: Ptr("project_type"),
	}

	response, _, err := client.Users.CreateUserRunner(createRunnerOpts)
	if err != nil {
		t.Errorf("Users.CreateUserRunner returned an error: %v", err)
	}

	require.Equal(t, 1234, response.ID)
	require.Equal(t, "glrt-1234567890ABCD", response.Token)
	require.Equal(t, (*time.Time)(nil), response.TokenExpiresAt)
}

func TestCreatePersonalAccessTokenForCurrentUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/user/personal_access_tokens"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/post_user_personal_access_tokens.json")
	})

	scopes := []string{"k8s_proxy"}
	expiresAt := ISOTime(time.Date(2020, time.October, 15, 0, 0, 0, 0, time.UTC))
	user, _, err := client.Users.CreatePersonalAccessTokenForCurrentUser(&CreatePersonalAccessTokenForCurrentUserOptions{
		Name:      String("mytoken"),
		Scopes:    &scopes,
		ExpiresAt: &expiresAt,
	})
	require.NoError(t, err)

	createdAt := time.Date(2020, time.October, 14, 11, 58, 53, 526000000, time.UTC)
	want := &PersonalAccessToken{
		ID:          3,
		Name:        "mytoken",
		Description: "Describe mytoken",
		Revoked:     false,
		CreatedAt:   &createdAt,
		Scopes:      scopes,
		UserID:      42,
		Active:      true,
		ExpiresAt:   &expiresAt,
		Token:       "glpat-aaaaaaaa-bbbbbbbbb",
	}
	require.Equal(t, want, user)
}

func TestCreateServiceAccountUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/service_accounts"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			t.Fatalf("Users.CreateServiceAccountUser request content-type %+v want application/json;", r.Header.Get("Content-Type"))
		}
		if r.ContentLength == -1 {
			t.Fatalf("Users.CreateServiceAccountUser request content-length is -1")
		}
		mustWriteHTTPResponse(t, w, "testdata/create_service_account_user.json")
	})

	user, _, err := client.Users.CreateServiceAccountUser(&CreateServiceAccountUserOptions{
		Name:     Ptr("Test Service Account"),
		Username: Ptr("serviceaccount"),
		Email:    Ptr("serviceaccount@test.com"),
	})
	require.NoError(t, err)

	want := &User{
		ID:        999,
		Username:  "serviceaccount",
		Name:      "Test Service Account",
		Email:     "serviceaccount@test.com",
		State:     "active",
		Locked:    false,
		AvatarURL: "http://localhost:3000/uploads/user/avatar/999/cd8.jpeg",
		WebURL:    "http://localhost:3000/serviceaccount",
	}
	require.Equal(t, want, user)
}

func TestCreateUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/users"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			t.Fatalf("Users.CreateUser request content-type %+v want application/json;", r.Header.Get("Content-Type"))
		}
		if r.ContentLength == -1 {
			t.Fatalf("Users.CreateUser request content-length is -1")
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`
    {
      "email": "user999@example.com",
      "id": 999,
      "name":"Firstname Lastname",
      "username":"user"
    }`))
	})

	user, _, err := client.Users.CreateUser(&CreateUserOptions{
		Email:    Ptr("user999@example.com"),
		Name:     Ptr("Firstname Lastname"),
		Username: Ptr("user"),
	})
	require.NoError(t, err)

	want := &User{
		Email:    "user999@example.com",
		ID:       999,
		Name:     "Firstname Lastname",
		Username: "user",
	}
	require.Equal(t, want, user)
}

func TestCreateUserAvatar(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/users"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
			t.Fatalf("Users.CreateUser request content-type %+v want multipart/form-data;", r.Header.Get("Content-Type"))
		}
		if r.ContentLength == -1 {
			t.Fatalf("Users.CreateUser request content-length is -1")
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`
    {
      "avatar_url":"http://localhost:3000/uploads/-/system/user/avatar/999/avatar.png",
      "email": "user999@example.com",
      "id": 999,
      "name":"Firstname Lastname",
      "username":"user"
    }`))
	})
	avatar := new(bytes.Buffer)
	userAvatar := &UserAvatar{
		Image:    avatar,
		Filename: "avatar.png",
	}
	user, _, err := client.Users.CreateUser(&CreateUserOptions{
		Avatar:   userAvatar,
		Email:    Ptr("user999@example.com"),
		Name:     Ptr("Firstname Lastname"),
		Username: Ptr("user"),
	})
	require.NoError(t, err)

	want := &User{
		AvatarURL: "http://localhost:3000/uploads/-/system/user/avatar/999/avatar.png",
		Email:     "user999@example.com",
		ID:        999,
		Name:      "Firstname Lastname",
		Username:  "user",
	}
	require.Equal(t, want, user)
}

func TestModifyUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/users/1"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			t.Fatalf("Users.ModifyUser request content-type %+v want application/json;", r.Header.Get("Content-Type"))
		}
		if r.ContentLength == -1 {
			t.Fatalf("Users.ModifyUser request content-length is -1")
		}
		fmt.Fprint(w, `{}`)
	})
	_, _, err := client.Users.ModifyUser(1, &ModifyUserOptions{})
	require.NoError(t, err)
}

func TestModifyUserAvatar(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/users/1"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data;") {
			t.Fatalf("Users.ModifyUser request content-type %+v want multipart/form-data;", r.Header.Get("Content-Type"))
		}
		if r.ContentLength == -1 {
			t.Fatalf("Users.ModifyUser request content-length is -1")
		}
		fmt.Fprint(w, `{}`)
	})
	avatar := new(bytes.Buffer)
	userAvatar := &UserAvatar{
		Image:    avatar,
		Filename: "avatar.png",
	}
	_, _, err := client.Users.ModifyUser(1, &ModifyUserOptions{Avatar: userAvatar})
	require.NoError(t, err)
}

func TestUploadAvatarUser(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/user/avatar", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data;") {
			t.Fatalf("Users.UploadAvatar request content-type %+v want multipart/form-data;", r.Header.Get("Content-Type"))
		}
		if r.ContentLength == -1 {
			t.Fatalf("Users.UploadAvatar request content-length is -1")
		}
		fmt.Fprint(w, `{}`)
	})

	avatar := new(bytes.Buffer)
	_, _, err := client.Users.UploadAvatar(avatar, "avatar.png")
	if err != nil {
		t.Fatalf("Users.UploadAvatar returns an error: %v", err)
	}
}

func TestListServiceAccounts(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/service_accounts"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_serviceaccounts.json")
	})

	serviceaccounts, _, err := client.Users.ListServiceAccounts(&ListServiceAccountsOptions{})
	require.NoError(t, err)
	want := []*ServiceAccount{
		{
			ID:       114,
			Username: "service_account_33",
			Name:     "Service account user",
		},
		{
			ID:       137,
			Username: "service_account_34",
			Name:     "john doe",
		},
	}
	require.Equal(t, want, serviceaccounts)
}

func TestDeleteUserIdentity(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/users/1/identities/google", func(_ http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Users.DeleteUserIdentity(1, "google")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestGetUserStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		uid     any
		path    string
		wantErr error
	}{
		{
			name: "numeric user ID",
			uid:  1,
			path: "/api/v4/users/1/status",
		},
		{
			name: "user name",
			uid:  "johndoe",
			path: "/api/v4/users/johndoe/status",
		},
		{
			name: "user name with @ prefix",
			uid:  "@johndoe",
			path: "/api/v4/users/johndoe/status",
		},
		{
			name:    "invalid uid type",
			uid:     User{ID: 1},
			path:    "/unused",
			wantErr: ErrInvalidIDType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mux, client := setup(t)

			mux.HandleFunc(tt.path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				mustWriteHTTPResponse(t, w, "testdata/get_user_status.json")
			})

			got, _, err := client.Users.GetUserStatus(tt.uid)
			require.ErrorIs(t, err, tt.wantErr)
			if tt.wantErr != nil {
				return
			}

			want := &UserStatus{
				Emoji:         "red_circle",
				Message:       "Duly swamped",
				Availability:  "busy",
				MessageHTML:   "Duly swamped",
				ClearStatusAt: Ptr(time.Date(2025, time.April, 24, 16, 56, 35, 0, time.UTC)),
			}
			require.Equal(t, want, got)
		})
	}
}

func TestSetUserStatus(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	opts := UserStatusOptions{
		Emoji:            Ptr("red_circle"),
		Message:          Ptr("Duly swamped"),
		Availability:     Ptr(Busy),
		ClearStatusAfter: Ptr(ClearStatusAfter30Minutes),
	}

	mux.HandleFunc("/api/v4/user/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testJSONBody(t, r, `
		{
			"emoji": "red_circle",
			"message": "Duly swamped",
			"availability": "busy",
			"clear_status_after": "30_minutes"
		}`)

		fmt.Fprint(w, `
		{
			"emoji": "red_circle",
			"message": "Duly swamped",
			"availability": "busy",
			"clear_status_at": "2025-04-24T15:02:02.000Z"
		}`)
	})

	got, _, err := client.Users.SetUserStatus(&opts)
	require.NoError(t, err)

	want := &UserStatus{
		Emoji:         "red_circle",
		Message:       "Duly swamped",
		Availability:  "busy",
		ClearStatusAt: Ptr(time.Date(2025, time.April, 24, 15, 2, 2, 0, time.UTC)),
	}
	require.Equal(t, 0, want.ClearStatusAt.Nanosecond())
	require.Equal(t, want, got)
}
