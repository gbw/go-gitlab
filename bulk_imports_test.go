package gitlab

import (
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBulkImportsService_StartMigration(t *testing.T) {
	t.Parallel()
	startMigrationOptions := &BulkImportStartMigrationOptions{
		Configuration: &BulkImportStartMigrationConfiguration{
			URL:         Ptr("https://source-gitlab-instance.example.com"),
			AccessToken: Ptr("source-gitlab-instance-access-token"),
		},
		Entities: []BulkImportStartMigrationEntity{
			{
				SourceType:           Ptr("group_entity"),
				SourceFullPath:       Ptr("gitlab-org/gitlab"),
				DestinationSlug:      Ptr("destination_slug"),
				DestinationNamespace: Ptr("destination/namespace/path"),
				MigrateProjects:      Ptr(true),
				MigrateMemberships:   Ptr(true),
			},
		},
	}
	wantStartMigrationResponse := &BulkImportStartMigrationResponse{
		ID:          1337,
		Status:      "created",
		SourceType:  "group_entity",
		SourceURL:   "https://source-gitlab-instance.example.com",
		CreatedAt:   time.Date(2021, time.June, 18, 9, 45, 55, 358000000, time.UTC),
		UpdatedAt:   time.Date(2021, time.June, 18, 9, 46, 27, 3000000, time.UTC),
		HasFailures: false,
	}
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/bulk_imports", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		gotBody, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		wantBody, err := os.ReadFile("testdata/post_bulk_imports_request.json")
		assert.NoError(t, err)
		assert.JSONEq(t, string(wantBody), string(gotBody))
		mustWriteHTTPResponse(t, w, "testdata/post_bulk_imports_response.json")
	})

	gotStartMigrationResponse, _, err := client.BulkImports.StartMigration(startMigrationOptions, nil)

	require.NoError(t, err)
	assert.Equal(t, wantStartMigrationResponse, gotStartMigrationResponse)
}

func TestBulkImportsService_ListBulkImports(t *testing.T) {
	t.Parallel()

	wantResponse := []*BulkImport{
		{
			ID:          1,
			Status:      "finished",
			SourceType:  "gitlab",
			SourceURL:   "https://gitlab.example.com",
			CreatedAt:   time.Date(2021, time.June, 18, 9, 45, 55, 358000000, time.UTC),
			UpdatedAt:   time.Date(2021, time.June, 18, 9, 46, 27, 3000000, time.UTC),
			HasFailures: false,
		},
		{
			ID:          2,
			Status:      "started",
			SourceType:  "gitlab",
			SourceURL:   "https://gitlab.example.com",
			CreatedAt:   time.Date(2021, time.June, 18, 9, 47, 36, 581000000, time.UTC),
			UpdatedAt:   time.Date(2021, time.June, 18, 9, 47, 58, 286000000, time.UTC),
			HasFailures: false,
		},
	}

	mux, client := setup(t)
	mux.HandleFunc("/api/v4/bulk_imports", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_bulk_imports_response.json")
	})

	gotResponse, _, err := client.BulkImports.ListBulkImports(nil, nil)

	require.NoError(t, err)
	assert.Equal(t, wantResponse, gotResponse)
}

func TestBulkImportsService_ListBulkImportsEntities(t *testing.T) {
	t.Parallel()

	nsID1 := int64(1)
	wantResponse := []*BulkImportEntity{
		{
			ID:                   1,
			BulkImportID:         1,
			Status:               "finished",
			EntityType:           "group",
			SourceFullPath:       "source_group",
			DestinationFullPath:  "destination/full_path",
			DestinationName:      "destination_slug",
			DestinationSlug:      "destination_slug",
			DestinationNamespace: "destination_path",
			ParentID:             nil,
			NamespaceID:          &nsID1,
			ProjectID:            nil,
			CreatedAt:            time.Date(2021, time.June, 18, 9, 47, 37, 390000000, time.UTC),
			UpdatedAt:            time.Date(2021, time.June, 18, 9, 47, 51, 867000000, time.UTC),
			Failures:             []*BulkImportEntityFailure{},
			MigrateProjects:      true,
			MigrateMemberships:   true,
			HasFailures:          false,
			Stats: BulkImportEntityStats{
				Labels:     BulkImportEntityStatItem{Source: 10, Fetched: 10, Imported: 10},
				Milestones: BulkImportEntityStatItem{Source: 10, Fetched: 10, Imported: 10},
			},
		},
		{
			ID:                   2,
			BulkImportID:         2,
			Status:               "failed",
			EntityType:           "group",
			SourceFullPath:       "another_group",
			DestinationFullPath:  "destination/full_path",
			DestinationName:      "destination_slug",
			DestinationSlug:      "another_slug",
			DestinationNamespace: "another_namespace",
			ParentID:             nil,
			NamespaceID:          nil,
			ProjectID:            nil,
			CreatedAt:            time.Date(2021, time.June, 24, 10, 40, 20, 110000000, time.UTC),
			UpdatedAt:            time.Date(2021, time.June, 24, 10, 40, 46, 590000000, time.UTC),
			Failures: []*BulkImportEntityFailure{
				{
					Relation:           "group",
					Step:               "extractor",
					ExceptionMessage:   "Error!",
					ExceptionClass:     "Exception",
					CorrelationIDValue: "dfcf583058ed4508e4c7c617bd7f0edd",
					CreatedAt:          time.Date(2021, time.June, 24, 10, 40, 46, 495000000, time.UTC),
					PipelineClass:      "BulkImports::Groups::Pipelines::GroupPipeline",
					PipelineStep:       "extractor",
				},
			},
			MigrateProjects:    true,
			MigrateMemberships: true,
			HasFailures:        false,
			Stats:              BulkImportEntityStats{},
		},
	}

	mux, client := setup(t)
	mux.HandleFunc("/api/v4/bulk_imports/entities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_bulk_imports_entities_response.json")
	})

	gotResponse, _, err := client.BulkImports.ListBulkImportsEntities(nil, nil)

	require.NoError(t, err)
	assert.Equal(t, wantResponse, gotResponse)
}

func TestBulkImportsService_GetBulkImport(t *testing.T) {
	t.Parallel()

	wantResponse := &BulkImport{
		ID:          1,
		Status:      "finished",
		SourceType:  "gitlab",
		SourceURL:   "https://gitlab.example.com",
		CreatedAt:   time.Date(2021, time.June, 18, 9, 45, 55, 358000000, time.UTC),
		UpdatedAt:   time.Date(2021, time.June, 18, 9, 46, 27, 3000000, time.UTC),
		HasFailures: false,
	}

	mux, client := setup(t)
	mux.HandleFunc("/api/v4/bulk_imports/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_bulk_import_response.json")
	})

	gotResponse, _, err := client.BulkImports.GetBulkImport(1, nil)

	require.NoError(t, err)
	assert.Equal(t, wantResponse, gotResponse)
}

func TestBulkImportsService_ListBulkImportsEntitiesByID(t *testing.T) {
	t.Parallel()

	nsID1 := int64(1)
	wantResponse := []*BulkImportEntity{
		{
			ID:                   1,
			BulkImportID:         1,
			Status:               "finished",
			EntityType:           "group",
			SourceFullPath:       "source_group",
			DestinationFullPath:  "destination/full_path",
			DestinationName:      "destination_slug",
			DestinationSlug:      "destination_slug",
			DestinationNamespace: "destination_path",
			ParentID:             nil,
			NamespaceID:          &nsID1,
			ProjectID:            nil,
			CreatedAt:            time.Date(2021, time.June, 18, 9, 47, 37, 390000000, time.UTC),
			UpdatedAt:            time.Date(2021, time.June, 18, 9, 47, 51, 867000000, time.UTC),
			Failures: []*BulkImportEntityFailure{
				{
					Relation:           "group",
					Step:               "extractor",
					ExceptionMessage:   "Error!",
					ExceptionClass:     "Exception",
					CorrelationIDValue: "dfcf583058ed4508e4c7c617bd7f0edd",
					CreatedAt:          time.Date(2021, time.June, 24, 10, 40, 46, 495000000, time.UTC),
					PipelineClass:      "BulkImports::Groups::Pipelines::GroupPipeline",
					PipelineStep:       "extractor",
				},
			},
			MigrateProjects:    true,
			MigrateMemberships: true,
			HasFailures:        true,
			Stats: BulkImportEntityStats{
				Labels:     BulkImportEntityStatItem{Source: 10, Fetched: 10, Imported: 10},
				Milestones: BulkImportEntityStatItem{Source: 10, Fetched: 10, Imported: 10},
			},
		},
	}

	mux, client := setup(t)
	mux.HandleFunc("/api/v4/bulk_imports/1/entities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_bulk_import_entities_by_id_response.json")
	})

	gotResponse, _, err := client.BulkImports.ListBulkImportsEntitiesByID(1, nil, nil)

	require.NoError(t, err)
	assert.Equal(t, wantResponse, gotResponse)
}

func TestBulkImportsService_GetBulkImportEntity(t *testing.T) {
	t.Parallel()

	nsID1 := int64(1)
	wantResponse := &BulkImportEntity{
		ID:                   1,
		BulkImportID:         1,
		Status:               "finished",
		EntityType:           "group",
		SourceFullPath:       "source_group",
		DestinationFullPath:  "destination/full_path",
		DestinationName:      "destination_slug",
		DestinationSlug:      "destination_slug",
		DestinationNamespace: "destination_path",
		ParentID:             nil,
		NamespaceID:          &nsID1,
		ProjectID:            nil,
		CreatedAt:            time.Date(2021, time.June, 18, 9, 47, 37, 390000000, time.UTC),
		UpdatedAt:            time.Date(2021, time.June, 18, 9, 47, 51, 867000000, time.UTC),
		Failures: []*BulkImportEntityFailure{
			{
				Relation:           "group",
				Step:               "extractor",
				ExceptionMessage:   "Error!",
				ExceptionClass:     "Exception",
				CorrelationIDValue: "dfcf583058ed4508e4c7c617bd7f0edd",
				CreatedAt:          time.Date(2021, time.June, 24, 10, 40, 46, 495000000, time.UTC),
				PipelineClass:      "BulkImports::Groups::Pipelines::GroupPipeline",
				PipelineStep:       "extractor",
			},
		},
		MigrateProjects:    true,
		MigrateMemberships: true,
		HasFailures:        true,
		Stats: BulkImportEntityStats{
			Labels:     BulkImportEntityStatItem{Source: 10, Fetched: 10, Imported: 10},
			Milestones: BulkImportEntityStatItem{Source: 10, Fetched: 10, Imported: 10},
		},
	}

	mux, client := setup(t)
	mux.HandleFunc("/api/v4/bulk_imports/1/entities/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_bulk_import_entity_response.json")
	})

	gotResponse, _, err := client.BulkImports.GetBulkImportEntity(1, 1, nil)

	require.NoError(t, err)
	assert.Equal(t, wantResponse, gotResponse)
}

func TestBulkImportsService_GetBulkImportEntityFailures(t *testing.T) {
	t.Parallel()

	wantResponse := []*BulkImportEntityFailure{
		{
			Relation:           "issues",
			ExceptionMessage:   "Error!",
			ExceptionClass:     "StandardError",
			CorrelationIDValue: "06289e4b064329a69de7bb2d7a1b5a97",
			SourceURL:          "https://gitlab.example/project/full/path/-/issues/1",
			SourceTitle:        "Issue title",
		},
	}

	mux, client := setup(t)
	mux.HandleFunc("/api/v4/bulk_imports/1/entities/2/failures", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_bulk_import_entity_failures_response.json")
	})

	gotResponse, _, err := client.BulkImports.GetBulkImportEntityFailures(1, 2, nil)

	require.NoError(t, err)
	assert.Equal(t, wantResponse, gotResponse)
}

func TestBulkImportsService_CancelBulkImport(t *testing.T) {
	t.Parallel()

	wantResponse := &BulkImport{
		ID:          1,
		Status:      "canceled",
		SourceType:  "gitlab",
		SourceURL:   "https://gitlab.example.com",
		CreatedAt:   time.Date(2021, time.June, 18, 9, 45, 55, 358000000, time.UTC),
		UpdatedAt:   time.Date(2021, time.June, 18, 9, 46, 27, 3000000, time.UTC),
		HasFailures: false,
	}

	mux, client := setup(t)
	mux.HandleFunc("/api/v4/bulk_imports/1/cancel", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/post_bulk_import_cancel_response.json")
	})

	gotResponse, _, err := client.BulkImports.CancelBulkImport(1, nil)

	require.NoError(t, err)
	assert.Equal(t, wantResponse, gotResponse)
}
