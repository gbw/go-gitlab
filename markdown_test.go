package gitlab

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		options    *RenderOptions
		wantBody   map[string]any
		wantHTML   string
		statusCode int
	}{
		{
			name: "Basic Markdown",
			options: &RenderOptions{
				Text: Ptr("# Testing"),
			},
			wantBody: map[string]any{
				"text": "# Testing",
			},
			wantHTML:   "<h1>Testing</h1>",
			statusCode: http.StatusOK,
		},
		{
			name: "With GFM and project",
			options: &RenderOptions{
				Text:                    Ptr("**bold**"),
				GitlabFlavouredMarkdown: Ptr(true),
				Project:                 Ptr("group/project"),
			},
			wantBody: map[string]any{
				"text":    "**bold**",
				"gfm":     true,
				"project": "group/project",
			},
			wantHTML:   "<p><strong>bold</strong></p>",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mux, client := setup(t)

			mux.HandleFunc("/api/v4/markdown", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				testBodyJSON(t, r, tc.wantBody)
				w.WriteHeader(tc.statusCode)
				_ = json.NewEncoder(w).Encode(Markdown{HTML: tc.wantHTML})
			})

			md, resp, err := client.Markdown.Render(tc.options)
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.Equal(t, tc.statusCode, resp.StatusCode)
			require.NotNil(t, md)
			assert.Equal(t, tc.wantHTML, md.HTML)
		})
	}
}
