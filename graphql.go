package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	// GraphQLAPIEndpoint defines the endpoint URI for the GraphQL backend
	GraphQLAPIEndpoint = "/api/graphql"
)

type (
	GraphQLInterface interface {
		Do(query GraphQLQuery, response any, options ...RequestOptionFunc) (*Response, error)
	}

	GraphQL struct {
		client *Client
	}

	GraphQLQuery struct {
		Query     string         `json:"query"`
		Variables map[string]any `json:"variables,omitempty"`
	}

	GenericGraphQLErrors struct {
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	GraphQLResponseError struct {
		Err    error
		Errors GenericGraphQLErrors
	}
)

var _ GraphQLInterface = (*GraphQL)(nil)

func (e *GraphQLResponseError) Error() string {
	if len(e.Errors.Errors) == 0 {
		return fmt.Sprintf("%s (no additional error messages)", e.Err)
	}

	var sb strings.Builder
	sb.WriteString(e.Err.Error())
	sb.WriteString(" (GraphQL errors: ")

	for i, err := range e.Errors.Errors {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(err.Message)
	}
	sb.WriteString(")")

	return sb.String()
}

// Do sends a GraphQL query and returns the response in the given response argument
// The response must be JSON serializable. The *Response return value is the HTTP response
// and must be used to retrieve additional HTTP information, like status codes and also
// error messages from failed queries.
//
// Example:
//
//	var response struct {
//		Data struct {
//			Project struct {
//				ID string `json:"id"`
//			} `json:"project"`
//		} `json:"data"`
//	}
//	_, err := client.GraphQL.Do(GraphQLQuery{Query: `query { project(fullPath: "gitlab-org/gitlab") { id } }`}, &response, gitlab.WithContext(ctx))
//
// Attention: This API is experimental and may be subject to breaking changes to improve the API in the future.
func (g *GraphQL) Do(query GraphQLQuery, response any, options ...RequestOptionFunc) (*Response, error) {
	request, err := g.client.NewRequest(http.MethodPost, "", query, options)
	if err != nil {
		return nil, fmt.Errorf("failed to create GraphQL request: %w", err)
	}
	// Overwrite the path of the existing request, as otherwise client-go appends /api/v4 instead.
	request.URL.Path = GraphQLAPIEndpoint
	resp, err := g.client.Do(request, response)
	if err != nil {
		// return error, details can be read from Response
		if errResp, ok := err.(*ErrorResponse); ok { //nolint:errorlint
			var v GenericGraphQLErrors
			if json.Unmarshal(errResp.Body, &v) == nil {
				return resp, &GraphQLResponseError{
					Err:    err,
					Errors: v,
				}
			}
		}
		return resp, fmt.Errorf("failed to execute GraphQL query: %w", err)
	}
	return resp, nil
}

type variableGQL struct {
	Name  string
	Type  string
	Value any
}

func (v variableGQL) definition() string {
	return fmt.Sprintf("$%s: %s", v.Name, v.Type)
}

func (v variableGQL) argument() string {
	return fmt.Sprintf("%s: $%s", v.Name, v.Name)
}

type variablesGQL []variableGQL

func (vs variablesGQL) asMap(base map[string]any) map[string]any {
	if base == nil {
		base = make(map[string]any)
	}

	for _, f := range vs {
		base[f.Name] = f.Value
	}

	return base
}

// Definitions generates the GraphQL query variable declarations for use in a query definition.
// It returns a comma-separated string of parameter declarations in the format "$name: Type".
// For example, if fieldsGQL contains fields with names "state" and "authorUsername" with types
// "IssuableState" and "String", it returns: "$state: IssuableState, $authorUsername: String".
// This is typically used in the query signature section of a GraphQL query.
func (vs variablesGQL) Definitions() string {
	defs := make([]string, len(vs))

	for i, v := range vs {
		defs[i] = v.definition()
	}

	return strings.Join(defs, ", ")
}

// Arguments generates the GraphQL argument assignments for use in a query body.
// It returns a comma-separated string of argument assignments in the format "name: $name".
// For example, if fieldsGQL contains fields with names "state" and "authorUsername", it returns:
// "state: $state, authorUsername: $authorUsername".
// This is typically used when passing variables to a GraphQL field or connection.
func (vs variablesGQL) Arguments() string {
	args := make([]string, len(vs))

	for i, v := range vs {
		args[i] = v.argument()
	}

	return strings.Join(args, ", ")
}

// gqlVariables extracts GraphQL variable definitions from a struct's fields.
// It accepts a pointer to a struct where each field is annotated with a `gql:"name type"` tag.
// The tag specifies the GraphQL variable name and type (e.g., `gql:"state IssuableState"`).
//
// Fields can be excluded using `gql:"-"`. Only non-zero fields are included in the result.
//
// Returns a variablesGQL slice containing the variable name, GraphQL type, and value for each field.
// This can be used to generate both variable definitions (for query signatures) and variable
// arguments (for field parameters) in GraphQL queries.
//
// Returns an error if:
//   - s is not a pointer to a struct
//   - any field is missing a `gql` tag
//   - a `gql` tag has invalid format (must be "name type", except those tagged with "-")
//
// Example:
//
//	type Options struct {
//	    State  *string `gql:"state IssuableState"`
//	    Author *string `gql:"authorUsername String"`
//	}
//	fields, err := gqlVariables(&Options{State: Ptr("opened")})
//	// Returns: [{Name: "state", Type: "IssuableState", Value: "opened"}]
func gqlVariables(s any) (variablesGQL, error) {
	if s == nil {
		return nil, nil
	}

	structValue := reflect.ValueOf(s)
	if structValue.Kind() != reflect.Ptr || structValue.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a pointer to a struct, got %T", s)
	}

	structValue = structValue.Elem() // Dereference the pointer to get the struct value
	structType := structValue.Type()

	var fields variablesGQL

	for i := range structType.NumField() {
		field := structType.Field(i)
		gqlTag := field.Tag.Get("gql")

		switch gqlTag {
		case "":
			return nil, fmt.Errorf("field %s.%s is missing a 'gql' tag", structType.Name(), field.Name)
		case "-":
			continue
		}

		name, typ, ok := strings.Cut(gqlTag, " ")
		if !ok {
			return nil, fmt.Errorf("invalid 'gql' tag format for field %s.%s: got %q, want \"name type\"", structType.Name(), field.Name, gqlTag)
		}

		fieldValue := structValue.Field(i)
		if fieldValue.IsZero() {
			continue
		}

		fields = append(fields, variableGQL{
			Name:  name,
			Type:  typ,
			Value: fieldValue.Interface(),
		})
	}

	return fields, nil
}

// gidGQL is a global ID. It is used by GraphQL to uniquely identify resources.
type gidGQL struct {
	Type  string
	Int64 int64
}

var gidGQLRegex = regexp.MustCompile(`^gid://gitlab/([^/]+)/(\d+)$`)

func (id *gidGQL) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	m := gidGQLRegex.FindStringSubmatch(s)
	if len(m) != 3 {
		return fmt.Errorf("invalid global ID format: %q", s)
	}

	i, err := strconv.ParseInt(m[2], 10, 64)
	if err != nil {
		return fmt.Errorf("failed parsing %q as numeric ID: %w", s, err)
	}

	id.Type = m[1]
	id.Int64 = i

	return nil
}

func (id gidGQL) String() string {
	return fmt.Sprintf("gid://gitlab/%s/%d", id.Type, id.Int64)
}

// iidGQL represents an int64 ID that is encoded by GraphQL as a string.
// This type is used unmarshal the string response into an int64 type.
type iidGQL int64

func (id *iidGQL) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("failed parsing %q as numeric ID: %w", s, err)
	}

	*id = iidGQL(i)
	return nil
}
