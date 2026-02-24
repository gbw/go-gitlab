package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectTemplatesService_ListTemplates(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/templates/issues.templates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, r.Pattern)
		fmt.Fprintf(w, `
			[
				{
					"key": "epl-1.0",
					"name": "Eclipse Public License 1.0"
				  },
				  {
					"key": "lgpl-3.0",
					"name": "GNU Lesser General Public License v3.0"
				  }
			]
		`)
	})

	want := []*ProjectTemplate{
		{
			Key:  "epl-1.0",
			Name: "Eclipse Public License 1.0",
		},
		{
			Key:  "lgpl-3.0",
			Name: "GNU Lesser General Public License v3.0",
		},
	}

	ss, resp, err := client.ProjectTemplates.ListTemplates(1, "issues.templates", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ss)
}

func TestProjectTemplatesService_GetProjectTemplate(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/templates/issues.templates/test_issue.template", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, r.Pattern)
		fmt.Fprintf(w, `
			{
			  "name": "test_issue.template",
			  "content": "## Test"
			}
		`)
	})

	want := &ProjectTemplate{
		Name:    "test_issue.template",
		Content: "## Test",
	}

	ss, resp, err := client.ProjectTemplates.GetProjectTemplate(1, "issues.templates", "test_issue.template", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ss)
}
