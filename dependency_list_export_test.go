package gitlab

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDependencyListExport(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/pipelines/1234/dependency_list_exports", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var content CreateDependencyListExportOptions
		err = json.Unmarshal(body, &content)
		assert.NoError(t, err)

		assert.Equal(t, "sbom", *content.ExportType)
		mustWriteHTTPResponse(t, w, "testdata/create_dependency_list_export.json")
	})

	d := &CreateDependencyListExportOptions{
		ExportType: Ptr("sbom"),
	}

	export, _, err := client.DependencyListExport.CreateDependencyListExport(1234, d)
	require.NoError(t, err)

	want := &DependencyListExport{
		ID:          5678,
		HasFinished: false,
		Self:        "http://gitlab.example.com/api/v4/dependency_list_exports/5678",
		Download:    "http://gitlab.example.com/api/v4/dependency_list_exports/5678/download",
	}
	require.Equal(t, want, export)
}

func TestGetDependencyListExport(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/dependency_list_exports/5678", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_dependency_list_export.json")
	})

	export, _, err := client.DependencyListExport.GetDependencyListExport(5678)
	require.NoError(t, err)

	want := &DependencyListExport{
		ID:          5678,
		HasFinished: true,
		Self:        "http://gitlab.example.com/api/v4/dependency_list_exports/5678",
		Download:    "http://gitlab.example.com/api/v4/dependency_list_exports/5678/download",
	}
	require.Equal(t, want, export)
}

func TestDownloadDependencyListExport(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/dependency_list_exports/5678/download", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/download_dependency_list_export.json")
	})

	sbomReader, _, err := client.DependencyListExport.DownloadDependencyListExport(5678)
	require.NoError(t, err)
	require.NotNil(t, sbomReader)

	expectedSbom, err := os.ReadFile("testdata/download_dependency_list_export.json")
	require.NoError(t, err)

	// GIVEN: A reader is returned
	// WHEN: We read the content
	gotContent, err := io.ReadAll(sbomReader)
	require.NoError(t, err)

	// THEN: The content should match the expected SBOM
	require.Equal(t, expectedSbom, gotContent)

	// Clean up: Close the reader if it's a ReadCloser
	if rc, ok := sbomReader.(io.Closer); ok {
		rc.Close()
	}
}
