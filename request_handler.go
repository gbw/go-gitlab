package gitlab

// Internal utility functions for handling common request/response patterns

// doRequest handles requests that return a single object.
// This is a generic utility function for GitLab API endpoints that return a single resource.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/
func doRequest[T any](
	client *Client,
	method, url string,
	body any,
	options ...RequestOptionFunc,
) (T, *Response, error) {
	var result T
	req, err := client.NewRequest(method, url, body, options)
	if err != nil {
		var zero T
		return zero, nil, err
	}
	resp, err := client.Do(req, &result)
	if err != nil {
		var zero T
		return zero, resp, err
	}
	return result, resp, nil
}

// doRequestSlice handles requests that return a slice of objects.
// This is a generic utility function for GitLab API endpoints that return a slice of objects.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/
func doRequestSlice[T any](
	client *Client,
	method, url string,
	body any,
	options ...RequestOptionFunc,
) ([]T, *Response, error) {
	var result []T
	req, err := client.NewRequest(method, url, body, options)
	if err != nil {
		return nil, nil, err
	}
	resp, err := client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// doRequestVoid handles requests that don't return data.
// This is a generic utility function for GitLab API endpoints that perform actions without returning response data.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/
func doRequestVoid(
	client *Client,
	method, url string,
	body any,
	options ...RequestOptionFunc,
) (*Response, error) {
	req, err := client.NewRequest(method, url, body, options)
	if err != nil {
		return nil, err
	}
	return client.Do(req, nil)
}
