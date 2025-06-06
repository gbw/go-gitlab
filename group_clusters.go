//
// Copyright 2021, Paul Shoemaker
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
	// Deprecated: in GitLab 14.5, to be removed in 19.0
	GroupClustersServiceInterface interface {
		// Deprecated: in GitLab 14.5, to be removed in 19.0
		ListClusters(pid any, options ...RequestOptionFunc) ([]*GroupCluster, *Response, error)
		// Deprecated: in GitLab 14.5, to be removed in 19.0
		GetCluster(pid any, cluster int, options ...RequestOptionFunc) (*GroupCluster, *Response, error)
		// Deprecated: in GitLab 14.5, to be removed in 19.0
		AddCluster(pid any, opt *AddGroupClusterOptions, options ...RequestOptionFunc) (*GroupCluster, *Response, error)
		// Deprecated: in GitLab 14.5, to be removed in 19.0
		EditCluster(pid any, cluster int, opt *EditGroupClusterOptions, options ...RequestOptionFunc) (*GroupCluster, *Response, error)
		// Deprecated: in GitLab 14.5, to be removed in 19.0
		DeleteCluster(pid any, cluster int, options ...RequestOptionFunc) (*Response, error)
	}

	// GroupClustersService handles communication with the
	// group clusters related methods of the GitLab API.
	// Deprecated: in GitLab 14.5, to be removed in 19.0
	//
	// GitLab API docs:
	// https://docs.gitlab.com/api/group_clusters/
	GroupClustersService struct {
		client *Client
	}
)

// Deprecated: in GitLab 14.5, to be removed in 19.0
var _ GroupClustersServiceInterface = (*GroupClustersService)(nil)

// GroupCluster represents a GitLab Group Cluster.
// Deprecated: in GitLab 14.5, to be removed in 19.0
//
// GitLab API docs: https://docs.gitlab.com/api/group_clusters/
type GroupCluster struct {
	ID                 int                 `json:"id"`
	Name               string              `json:"name"`
	Domain             string              `json:"domain"`
	CreatedAt          *time.Time          `json:"created_at"`
	Managed            bool                `json:"managed"`
	Enabled            bool                `json:"enabled"`
	ProviderType       string              `json:"provider_type"`
	PlatformType       string              `json:"platform_type"`
	EnvironmentScope   string              `json:"environment_scope"`
	ClusterType        string              `json:"cluster_type"`
	User               *User               `json:"user"`
	PlatformKubernetes *PlatformKubernetes `json:"platform_kubernetes"`
	ManagementProject  *ManagementProject  `json:"management_project"`
	Group              *Group              `json:"group"`
}

// Deprecated: in GitLab 14.5, to be removed in 19.0
func (v GroupCluster) String() string {
	return Stringify(v)
}

// ListClusters gets a list of all clusters in a group.
// Deprecated: in GitLab 14.5, to be removed in 19.0
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_clusters/#list-group-clusters
func (s *GroupClustersService) ListClusters(pid any, options ...RequestOptionFunc) ([]*GroupCluster, *Response, error) {
	group, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/clusters", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var pcs []*GroupCluster
	resp, err := s.client.Do(req, &pcs)
	if err != nil {
		return nil, resp, err
	}

	return pcs, resp, nil
}

// GetCluster gets a cluster.
// Deprecated: in GitLab 14.5, to be removed in 19.0
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_clusters/#get-a-single-group-cluster
func (s *GroupClustersService) GetCluster(pid any, cluster int, options ...RequestOptionFunc) (*GroupCluster, *Response, error) {
	group, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/clusters/%d", PathEscape(group), cluster)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	gc := new(GroupCluster)
	resp, err := s.client.Do(req, &gc)
	if err != nil {
		return nil, resp, err
	}

	return gc, resp, nil
}

// AddGroupClusterOptions represents the available AddCluster() options.
// Deprecated: in GitLab 14.5, to be removed in 19.0
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_clusters/#add-existing-cluster-to-group
type AddGroupClusterOptions struct {
	Name                *string                            `url:"name,omitempty" json:"name,omitempty"`
	Domain              *string                            `url:"domain,omitempty" json:"domain,omitempty"`
	ManagementProjectID *string                            `url:"management_project_id,omitempty" json:"management_project_id,omitempty"`
	Enabled             *bool                              `url:"enabled,omitempty" json:"enabled,omitempty"`
	Managed             *bool                              `url:"managed,omitempty" json:"managed,omitempty"`
	EnvironmentScope    *string                            `url:"environment_scope,omitempty" json:"environment_scope,omitempty"`
	PlatformKubernetes  *AddGroupPlatformKubernetesOptions `url:"platform_kubernetes_attributes,omitempty" json:"platform_kubernetes_attributes,omitempty"`
}

// AddGroupPlatformKubernetesOptions represents the available PlatformKubernetes options for adding.
// Deprecated: in GitLab 14.5, to be removed in 19.0
type AddGroupPlatformKubernetesOptions struct {
	APIURL            *string `url:"api_url,omitempty" json:"api_url,omitempty"`
	Token             *string `url:"token,omitempty" json:"token,omitempty"`
	CaCert            *string `url:"ca_cert,omitempty" json:"ca_cert,omitempty"`
	Namespace         *string `url:"namespace,omitempty" json:"namespace,omitempty"`
	AuthorizationType *string `url:"authorization_type,omitempty" json:"authorization_type,omitempty"`
}

// AddCluster adds an existing cluster to the group.
// Deprecated: in GitLab 14.5, to be removed in 19.0
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_clusters/#add-existing-cluster-to-group
func (s *GroupClustersService) AddCluster(pid any, opt *AddGroupClusterOptions, options ...RequestOptionFunc) (*GroupCluster, *Response, error) {
	group, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/clusters/user", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	gc := new(GroupCluster)
	resp, err := s.client.Do(req, gc)
	if err != nil {
		return nil, resp, err
	}

	return gc, resp, nil
}

// EditGroupClusterOptions represents the available EditCluster() options.
// Deprecated: in GitLab 14.5, to be removed in 19.0
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_clusters/#edit-group-cluster
type EditGroupClusterOptions struct {
	Name                *string                             `url:"name,omitempty" json:"name,omitempty"`
	Domain              *string                             `url:"domain,omitempty" json:"domain,omitempty"`
	EnvironmentScope    *string                             `url:"environment_scope,omitempty" json:"environment_scope,omitempty"`
	PlatformKubernetes  *EditGroupPlatformKubernetesOptions `url:"platform_kubernetes_attributes,omitempty" json:"platform_kubernetes_attributes,omitempty"`
	ManagementProjectID *string                             `url:"management_project_id,omitempty" json:"management_project_id,omitempty"`
}

// EditGroupPlatformKubernetesOptions represents the available PlatformKubernetes options for editing.
// Deprecated: in GitLab 14.5, to be removed in 19.0
type EditGroupPlatformKubernetesOptions struct {
	APIURL *string `url:"api_url,omitempty" json:"api_url,omitempty"`
	Token  *string `url:"token,omitempty" json:"token,omitempty"`
	CaCert *string `url:"ca_cert,omitempty" json:"ca_cert,omitempty"`
}

// EditCluster updates an existing group cluster.
// Deprecated: in GitLab 14.5, to be removed in 19.0
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_clusters/#edit-group-cluster
func (s *GroupClustersService) EditCluster(pid any, cluster int, opt *EditGroupClusterOptions, options ...RequestOptionFunc) (*GroupCluster, *Response, error) {
	group, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/clusters/%d", PathEscape(group), cluster)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	gc := new(GroupCluster)
	resp, err := s.client.Do(req, gc)
	if err != nil {
		return nil, resp, err
	}

	return gc, resp, nil
}

// DeleteCluster deletes an existing group cluster.
// Deprecated: in GitLab 14.5, to be removed in 19.0
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_clusters/#delete-group-cluster
func (s *GroupClustersService) DeleteCluster(pid any, cluster int, options ...RequestOptionFunc) (*Response, error) {
	group, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s/clusters/%d", PathEscape(group), cluster)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
