package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListGroupWikis(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/wikis",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[
				{
					"content": "content here",
					"format": "markdown",
					"slug": "deploy",
					"title": "deploy title"
				}
			]`)
		})

	groupwikis, _, err := client.GroupWikis.ListGroupWikis(1, &ListGroupWikisOptions{})
	require.NoError(t, err)

	want := []*GroupWiki{
		{
			Content: "content here",
			Format:  WikiFormatMarkdown,
			Slug:    "deploy",
			Title:   "deploy title",
		},
	}
	assert.Equal(t, want, groupwikis)
}

func TestGetGroupWikiPage(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/wikis/deploy",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `{
				"content": "content here",
				"format": "asciidoc",
				"slug": "deploy",
				"title": "deploy title",
				"encoding": "UTF-8"
			}`)
		})

	groupwiki, _, err := client.GroupWikis.GetGroupWikiPage(1, "deploy", &GetGroupWikiPageOptions{})
	require.NoError(t, err)

	want := &GroupWiki{
		Content:  "content here",
		Encoding: "UTF-8",
		Format:   WikiFormatASCIIDoc,
		Slug:     "deploy",
		Title:    "deploy title",
	}
	assert.Equal(t, want, groupwiki)
}

func TestCreateGroupWikiPage(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/wikis",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{
				"content": "content here",
				"format": "rdoc",
				"slug": "deploy",
				"title": "deploy title"
			}`)
		})

	groupwiki, _, err := client.GroupWikis.CreateGroupWikiPage(1, &CreateGroupWikiPageOptions{
		Content: Ptr("content here"),
		Title:   Ptr("deploy title"),
		Format:  Ptr(WikiFormatRDoc),
	})
	require.NoError(t, err)

	want := &GroupWiki{
		Content: "content here",
		Format:  WikiFormatRDoc,
		Slug:    "deploy",
		Title:   "deploy title",
	}
	assert.Equal(t, want, groupwiki)
}

func TestEditGroupWikiPage(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/wikis/deploy",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			fmt.Fprint(w, `{
				"content": "content here",
				"format": "asciidoc",
				"slug": "deploy",
				"title": "deploy title"
			}`)
		})

	groupwiki, _, err := client.GroupWikis.EditGroupWikiPage(1, "deploy", &EditGroupWikiPageOptions{
		Content: Ptr("content here"),
		Title:   Ptr("deploy title"),
		Format:  Ptr(WikiFormatRDoc),
	})
	require.NoError(t, err)

	want := &GroupWiki{
		Content: "content here",
		Format:  WikiFormatASCIIDoc,
		Slug:    "deploy",
		Title:   "deploy title",
	}
	assert.Equal(t, want, groupwiki)
}

func TestDeleteGroupWikiPage(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/wikis/deploy",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			w.WriteHeader(204)
		})

	r, err := client.GroupWikis.DeleteGroupWikiPage(1, "deploy")
	require.NoError(t, err)
	assert.Equal(t, 204, r.StatusCode)
}
