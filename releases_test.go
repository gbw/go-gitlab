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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReleasesService_ListReleases(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, exampleReleaseListResponse)
		})

	opt := &ListReleasesOptions{}
	releases, _, err := client.Releases.ListReleases(1, opt)
	require.NoError(t, err)
	assert.Len(t, releases, 2)
}

func TestReleasesService_GetRelease(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, exampleReleaseResponse)
		})

	release, _, err := client.Releases.GetRelease(1, exampleTagName)
	require.NoError(t, err)
	assert.Equal(t, exampleTagName, release.TagName)
}

func TestReleasesService_CreateRelease(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			testBodyJSON(t, r, map[string]string{
				"name":        "name",
				"tag_name":    exampleTagName,
				"description": "Description",
			})
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &CreateReleaseOptions{
		Name:        Ptr("name"),
		TagName:     Ptr(exampleTagName),
		Description: Ptr("Description"),
	}

	release, _, err := client.Releases.CreateRelease(1, opts)
	require.NoError(t, err)
	assert.Equal(t, exampleTagName, release.TagName)
}

func TestReleasesService_CreateReleaseWithAsset(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			testBodyJSON(t, r, map[string]any{
				"name":        "name",
				"tag_name":    exampleTagName,
				"description": "Description",
				"assets": map[string]any{
					"links": []any{
						map[string]any{
							"name":              "sldkf",
							"url":               "sldkfj",
							"filepath":          "sldkfh",
							"direct_asset_path": "direct-asset-path",
							"link_type":         "other",
						},
					},
				},
			})
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &CreateReleaseOptions{
		Name:        Ptr("name"),
		TagName:     Ptr(exampleTagName),
		Description: Ptr("Description"),
		Assets: &ReleaseAssetsOptions{
			Links: []*ReleaseAssetLinkOptions{
				{
					Name:            Ptr("sldkf"),
					URL:             Ptr("sldkfj"),
					FilePath:        Ptr("sldkfh"),
					DirectAssetPath: Ptr("direct-asset-path"),
					LinkType:        Ptr(OtherLinkType),
				},
			},
		},
	}

	release, _, err := client.Releases.CreateRelease(1, opts)
	require.NoError(t, err)
	assert.Equal(t, exampleTagName, release.TagName)
}

func TestReleasesService_CreateReleaseWithAssetAndNameMetadata(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			testBodyJSON(t, r, map[string]any{
				"name":        "name",
				"tag_name":    exampleTagNameWithMetadata,
				"description": "Description",
				"assets": map[string]any{
					"links": []any{
						map[string]any{
							"name":              "sldkf",
							"url":               "sldkfj",
							"filepath":          "sldkfh",
							"direct_asset_path": "direct-asset-path",
							"link_type":         "other",
						},
					},
				},
			})
			fmt.Fprint(w, exampleReleaseWithMetadataResponse)
		})

	opts := &CreateReleaseOptions{
		Name:        Ptr("name"),
		TagName:     Ptr(exampleTagNameWithMetadata),
		Description: Ptr("Description"),
		Assets: &ReleaseAssetsOptions{
			Links: []*ReleaseAssetLinkOptions{
				{
					Name:            Ptr("sldkf"),
					URL:             Ptr("sldkfj"),
					FilePath:        Ptr("sldkfh"),
					DirectAssetPath: Ptr("direct-asset-path"),
					LinkType:        Ptr(OtherLinkType),
				},
			},
		},
	}

	release, _, err := client.Releases.CreateRelease(1, opts)
	require.NoError(t, err)
	assert.Equal(t, exampleTagNameWithMetadata, release.TagName)
}

func TestReleasesService_CreateReleaseWithMilestones(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			testBodyJSON(t, r, map[string]any{
				"name":        "name",
				"tag_name":    exampleTagName,
				"description": "Description",
				"milestones":  []any{exampleTagName, "v0.1.0"},
			})
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &CreateReleaseOptions{
		Name:        Ptr("name"),
		TagName:     Ptr(exampleTagName),
		Description: Ptr("Description"),
		Milestones:  &[]string{exampleTagName, "v0.1.0"},
	}

	release, _, err := client.Releases.CreateRelease(1, opts)
	require.NoError(t, err)
	assert.Equal(t, exampleTagName, release.TagName)
}

func TestReleasesService_CreateReleaseWithReleasedAt(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			testBodyJSON(t, r, map[string]any{
				"name":        "name",
				"tag_name":    exampleTagName,
				"description": "Description",
				"released_at": "0001-01-01T00:00:00Z",
			})
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &CreateReleaseOptions{
		Name:        Ptr("name"),
		TagName:     Ptr(exampleTagName),
		Description: Ptr("Description"),
		ReleasedAt:  &time.Time{},
	}

	release, _, err := client.Releases.CreateRelease(1, opts)
	require.NoError(t, err)
	assert.Equal(t, exampleTagName, release.TagName)
}

func TestReleasesService_UpdateRelease(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			testBodyJSON(t, r, map[string]any{
				"name":        "name",
				"description": "Description",
			})
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &UpdateReleaseOptions{
		Name:        Ptr("name"),
		Description: Ptr("Description"),
	}

	release, _, err := client.Releases.UpdateRelease(1, exampleTagName, opts)
	require.NoError(t, err)
	assert.Equal(t, exampleTagName, release.TagName)
}

func TestReleasesService_UpdateReleaseWithMilestones(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			testBodyJSON(t, r, map[string]any{
				"name":        "name",
				"description": "Description",
				"milestones":  []any{exampleTagName, "v0.1.0"},
			})
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &UpdateReleaseOptions{
		Name:        Ptr("name"),
		Description: Ptr("Description"),
		Milestones:  &[]string{exampleTagName, "v0.1.0"},
	}

	release, _, err := client.Releases.UpdateRelease(1, exampleTagName, opts)
	require.NoError(t, err)
	assert.Equal(t, exampleTagName, release.TagName)
}

func TestReleasesService_UpdateReleaseWithReleasedAt(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			testBodyJSON(t, r, map[string]any{
				"name":        "name",
				"description": "Description",
				"released_at": "0001-01-01T00:00:00Z",
			})
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &UpdateReleaseOptions{
		Name:        Ptr("name"),
		Description: Ptr("Description"),
		ReleasedAt:  &time.Time{},
	}

	release, _, err := client.Releases.UpdateRelease(1, exampleTagName, opts)
	require.NoError(t, err)
	assert.Equal(t, exampleTagName, release.TagName)
}

func TestReleasesService_DeleteRelease(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			fmt.Fprint(w, exampleReleaseResponse)
		})

	release, _, err := client.Releases.DeleteRelease(1, exampleTagName)
	require.NoError(t, err)
	assert.Equal(t, exampleTagName, release.TagName)
}
