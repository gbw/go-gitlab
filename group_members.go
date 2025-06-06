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
	"time"
)

type (
	GroupMembersServiceInterface interface {
		GetGroupMember(gid any, user int, options ...RequestOptionFunc) (*GroupMember, *Response, error)
		GetInheritedGroupMember(gid any, user int, options ...RequestOptionFunc) (*GroupMember, *Response, error)
		AddGroupMember(gid any, opt *AddGroupMemberOptions, options ...RequestOptionFunc) (*GroupMember, *Response, error)
		ShareWithGroup(gid any, opt *ShareWithGroupOptions, options ...RequestOptionFunc) (*Group, *Response, error)
		DeleteShareWithGroup(gid any, groupID int, options ...RequestOptionFunc) (*Response, error)
		EditGroupMember(gid any, user int, opt *EditGroupMemberOptions, options ...RequestOptionFunc) (*GroupMember, *Response, error)
		RemoveGroupMember(gid any, user int, opt *RemoveGroupMemberOptions, options ...RequestOptionFunc) (*Response, error)
	}

	// GroupMembersService handles communication with the group members
	// related methods of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/members/
	GroupMembersService struct {
		client *Client
	}
)

var _ GroupMembersServiceInterface = (*GroupMembersService)(nil)

// GroupMember represents a GitLab group member.
//
// GitLab API docs: https://docs.gitlab.com/api/members/
type GroupMember struct {
	ID                int                      `json:"id"`
	Username          string                   `json:"username"`
	Name              string                   `json:"name"`
	State             string                   `json:"state"`
	AvatarURL         string                   `json:"avatar_url"`
	WebURL            string                   `json:"web_url"`
	CreatedAt         *time.Time               `json:"created_at"`
	ExpiresAt         *ISOTime                 `json:"expires_at"`
	AccessLevel       AccessLevelValue         `json:"access_level"`
	Email             string                   `json:"email,omitempty"`
	PublicEmail       string                   `json:"public_email,omitempty"`
	GroupSAMLIdentity *GroupMemberSAMLIdentity `json:"group_saml_identity"`
	MemberRole        *MemberRole              `json:"member_role"`
}

// GroupMemberSAMLIdentity represents the SAML Identity link for the group member.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#list-all-members-of-a-group-or-project
type GroupMemberSAMLIdentity struct {
	ExternUID      string `json:"extern_uid"`
	Provider       string `json:"provider"`
	SAMLProviderID int    `json:"saml_provider_id"`
}

// BillableGroupMember represents a GitLab billable group member.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#list-all-billable-members-of-a-group
type BillableGroupMember struct {
	ID             int        `json:"id"`
	Username       string     `json:"username"`
	Name           string     `json:"name"`
	State          string     `json:"state"`
	AvatarURL      string     `json:"avatar_url"`
	WebURL         string     `json:"web_url"`
	Email          string     `json:"email"`
	LastActivityOn *ISOTime   `json:"last_activity_on"`
	MembershipType string     `json:"membership_type"`
	Removable      bool       `json:"removable"`
	CreatedAt      *time.Time `json:"created_at"`
	IsLastOwner    bool       `json:"is_last_owner"`
	LastLoginAt    *time.Time `json:"last_login_at"`
}

// BillableUserMembership represents a Membership of a billable user of a group
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#list-memberships-for-a-billable-member-of-a-group
type BillableUserMembership struct {
	ID               int                 `json:"id"`
	SourceID         int                 `json:"source_id"`
	SourceFullName   string              `json:"source_full_name"`
	SourceMembersURL string              `json:"source_members_url"`
	CreatedAt        *time.Time          `json:"created_at"`
	ExpiresAt        *time.Time          `json:"expires_at"`
	AccessLevel      *AccessLevelDetails `json:"access_level"`
}

// ListGroupMembersOptions represents the available ListGroupMembers() and
// ListAllGroupMembers() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#list-all-members-of-a-group-or-project
type ListGroupMembersOptions struct {
	ListOptions
	Query   *string `url:"query,omitempty" json:"query,omitempty"`
	UserIDs *[]int  `url:"user_ids[],omitempty" json:"user_ids,omitempty"`
}

// ListGroupMembers get a list of group members viewable by the authenticated
// user. Inherited members through ancestor groups are not included.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#list-all-members-of-a-group-or-project
func (s *GroupsService) ListGroupMembers(gid any, opt *ListGroupMembersOptions, options ...RequestOptionFunc) ([]*GroupMember, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/members", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var gm []*GroupMember
	resp, err := s.client.Do(req, &gm)
	if err != nil {
		return nil, resp, err
	}

	return gm, resp, nil
}

// ListAllGroupMembers get a list of group members viewable by the authenticated
// user. Returns a list including inherited members through ancestor groups.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#list-all-members-of-a-group-or-project-including-inherited-and-invited-members
func (s *GroupsService) ListAllGroupMembers(gid any, opt *ListGroupMembersOptions, options ...RequestOptionFunc) ([]*GroupMember, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/members/all", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var gm []*GroupMember
	resp, err := s.client.Do(req, &gm)
	if err != nil {
		return nil, resp, err
	}

	return gm, resp, nil
}

// AddGroupMemberOptions represents the available AddGroupMember() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#add-a-member-to-a-group-or-project
type AddGroupMemberOptions struct {
	UserID       *int              `url:"user_id,omitempty" json:"user_id,omitempty"`
	Username     *string           `url:"username,omitempty" json:"username,omitempty"`
	AccessLevel  *AccessLevelValue `url:"access_level,omitempty" json:"access_level,omitempty"`
	ExpiresAt    *string           `url:"expires_at,omitempty" json:"expires_at"`
	MemberRoleID *int              `url:"member_role_id,omitempty" json:"member_role_id,omitempty"`
}

// GetGroupMember gets a member of a group.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#get-a-member-of-a-group-or-project
func (s *GroupMembersService) GetGroupMember(gid any, user int, options ...RequestOptionFunc) (*GroupMember, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/members/%d", PathEscape(group), user)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	gm := new(GroupMember)
	resp, err := s.client.Do(req, gm)
	if err != nil {
		return nil, resp, err
	}

	return gm, resp, nil
}

// GetInheritedGroupMember get a member of a group or project, including
// inherited and invited members
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#get-a-member-of-a-group-or-project-including-inherited-and-invited-members
func (s *GroupMembersService) GetInheritedGroupMember(gid any, user int, options ...RequestOptionFunc) (*GroupMember, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/members/all/%d", PathEscape(group), user)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	gm := new(GroupMember)
	resp, err := s.client.Do(req, gm)
	if err != nil {
		return nil, resp, err
	}

	return gm, resp, err
}

// ListBillableGroupMembersOptions represents the available
// ListBillableGroupMembers() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#list-all-billable-members-of-a-group
type ListBillableGroupMembersOptions struct {
	ListOptions
	Search *string `url:"search,omitempty" json:"search,omitempty"`
	Sort   *string `url:"sort,omitempty" json:"sort,omitempty"`
}

// ListBillableGroupMembers Gets a list of group members that count as billable.
// The list includes members in the subgroup or subproject.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#list-all-billable-members-of-a-group
func (s *GroupsService) ListBillableGroupMembers(gid any, opt *ListBillableGroupMembersOptions, options ...RequestOptionFunc) ([]*BillableGroupMember, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/billable_members", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var bgm []*BillableGroupMember
	resp, err := s.client.Do(req, &bgm)
	if err != nil {
		return nil, resp, err
	}

	return bgm, resp, nil
}

// ListMembershipsForBillableGroupMemberOptions represents the available
// ListMembershipsForBillableGroupMember() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#list-memberships-for-a-billable-member-of-a-group
type ListMembershipsForBillableGroupMemberOptions = ListOptions

// ListMembershipsForBillableGroupMember gets a list of memberships for a
// billable member of a group.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#list-memberships-for-a-billable-member-of-a-group
func (s *GroupsService) ListMembershipsForBillableGroupMember(gid any, user int, opt *ListMembershipsForBillableGroupMemberOptions, options ...RequestOptionFunc) ([]*BillableUserMembership, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/billable_members/%d/memberships", PathEscape(group), user)

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var bum []*BillableUserMembership
	resp, err := s.client.Do(req, &bum)
	if err != nil {
		return nil, resp, err
	}

	return bum, resp, nil
}

// RemoveBillableGroupMember removes a given group members that count as billable.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#remove-a-billable-member-from-a-group
func (s *GroupsService) RemoveBillableGroupMember(gid any, user int, options ...RequestOptionFunc) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s/billable_members/%d", PathEscape(group), user)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// AddGroupMember adds a user to the list of group members.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#add-a-member-to-a-group-or-project
func (s *GroupMembersService) AddGroupMember(gid any, opt *AddGroupMemberOptions, options ...RequestOptionFunc) (*GroupMember, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/members", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	gm := new(GroupMember)
	resp, err := s.client.Do(req, gm)
	if err != nil {
		return nil, resp, err
	}

	return gm, resp, nil
}

// ShareWithGroup shares a group with the group.
//
// GitLab API docs:
// https://docs.gitlab.com/api/groups/#create-a-link-to-share-a-group-with-another-group
func (s *GroupMembersService) ShareWithGroup(gid any, opt *ShareWithGroupOptions, options ...RequestOptionFunc) (*Group, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/share", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	g := new(Group)
	resp, err := s.client.Do(req, g)
	if err != nil {
		return nil, resp, err
	}

	return g, resp, nil
}

// DeleteShareWithGroup allows to unshare a group from a group.
//
// GitLab API docs:
// https://docs.gitlab.com/api/groups/#delete-the-link-that-shares-a-group-with-another-group
func (s *GroupMembersService) DeleteShareWithGroup(gid any, groupID int, options ...RequestOptionFunc) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s/share/%d", PathEscape(group), groupID)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// EditGroupMemberOptions represents the available EditGroupMember()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#edit-a-member-of-a-group-or-project
type EditGroupMemberOptions struct {
	AccessLevel  *AccessLevelValue `url:"access_level,omitempty" json:"access_level,omitempty"`
	ExpiresAt    *string           `url:"expires_at,omitempty" json:"expires_at,omitempty"`
	MemberRoleID *int              `url:"member_role_id,omitempty" json:"member_role_id,omitempty"`
}

// EditGroupMember updates a member of a group.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#edit-a-member-of-a-group-or-project
func (s *GroupMembersService) EditGroupMember(gid any, user int, opt *EditGroupMemberOptions, options ...RequestOptionFunc) (*GroupMember, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/members/%d", PathEscape(group), user)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	gm := new(GroupMember)
	resp, err := s.client.Do(req, gm)
	if err != nil {
		return nil, resp, err
	}

	return gm, resp, nil
}

// RemoveGroupMemberOptions represents the available options to remove a group member.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#remove-a-member-from-a-group-or-project
type RemoveGroupMemberOptions struct {
	SkipSubresources  *bool `url:"skip_subresources,omitempty" json:"skip_subresources,omitempty"`
	UnassignIssuables *bool `url:"unassign_issuables,omitempty" json:"unassign_issuables,omitempty"`
}

// RemoveGroupMember removes user from user team.
//
// GitLab API docs:
// https://docs.gitlab.com/api/members/#remove-a-member-from-a-group-or-project
func (s *GroupMembersService) RemoveGroupMember(gid any, user int, opt *RemoveGroupMemberOptions, options ...RequestOptionFunc) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s/members/%d", PathEscape(group), user)

	req, err := s.client.NewRequest(http.MethodDelete, u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
