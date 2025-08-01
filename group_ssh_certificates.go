package gitlab

import (
	"fmt"
	"net/http"
	"time"
)

type (
	// GroupSSHCertificatesServiceInterface defines methods for the GroupSSHCertificatesService.
	GroupSSHCertificatesServiceInterface interface {
		ListGroupSSHCertificates(gid any, options ...RequestOptionFunc) ([]*GroupSSHCertificate, *Response, error)
		CreateGroupSSHCertificate(gid any, opt *CreateGroupSSHCertificateOptions, options ...RequestOptionFunc) (*GroupSSHCertificate, *Response, error)
		DeleteGroupSSHCertificate(gid any, cert int, options ...RequestOptionFunc) (*Response, error)
	}

	// GroupSSHCertificatesService handles communication with the group
	// SSH certificate related methods of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/group_ssh_certificates/
	GroupSSHCertificatesService struct {
		client *Client
	}
)

var _ GroupSSHCertificatesServiceInterface = (*GroupSSHCertificatesService)(nil)

// GroupSSHCertificate represents a GitLab Group SSH certificate.
//
// GitLab API docs: https://docs.gitlab.com/api/group_ssh_certificates/
type GroupSSHCertificate struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Key       string     `json:"key"`
	CreatedAt *time.Time `json:"created_at"`
}

// ListGroupSSHCertificates gets a list of SSH certificates for a specified
// group.
//
// Gitlab API docs:
// https://docs.gitlab.com/api/group_ssh_certificates/#get-all-ssh-certificates-for-a-particular-group
func (s *GroupSSHCertificatesService) ListGroupSSHCertificates(gid any, options ...RequestOptionFunc) ([]*GroupSSHCertificate, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/ssh_certificates", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var certs []*GroupSSHCertificate
	resp, err := s.client.Do(req, &certs)
	if err != nil {
		return nil, resp, err
	}

	return certs, resp, nil
}

// CreateGroupSSHCertificateOptions represents the available
// CreateGroupSSHCertificate() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_ssh_certificates/#create-ssh-certificate
type CreateGroupSSHCertificateOptions struct {
	Key   *string `url:"key,omitempty" json:"key,omitempty"`
	Title *string `url:"title,omitempty" json:"title,omitempty"`
}

// CreateGroupSSHCertificate creates a new SSH certificate in the group.
//
// Gitlab API docs:
// https://docs.gitlab.com/api/group_ssh_certificates/#create-ssh-certificate
func (s *GroupSSHCertificatesService) CreateGroupSSHCertificate(gid any, opt *CreateGroupSSHCertificateOptions, options ...RequestOptionFunc) (*GroupSSHCertificate, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/ssh_certificates", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	cert := new(GroupSSHCertificate)
	resp, err := s.client.Do(req, cert)
	if err != nil {
		return nil, resp, err
	}

	return cert, resp, nil
}

// DeleteGroupSSHCertificate deletes a SSH certificate from a specified group.
//
// Gitlab API docs:
// https://docs.gitlab.com/api/group_ssh_certificates/#delete-group-ssh-certificate
func (s *GroupSSHCertificatesService) DeleteGroupSSHCertificate(gid any, cert int, options ...RequestOptionFunc) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s/ssh_certificates/%d", PathEscape(group), cert)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
