package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSecurityCategoriesService_CreateSecurityCategory(t *testing.T) {
	t.Parallel()

	// GIVEN a valid namespace ID and category options
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityCategoryCreate": {
						"securityCategory": {
							"id": "gid://gitlab/Security::Category/1",
							"name": "Exposure",
							"description": "Exposure level of the asset",
							"multipleSelection": false,
							"editableState": "EDITABLE",
							"templateType": null,
							"securityAttributes": []
						},
						"errors": []
					}
				}
			}
		`)
	})

	desc := "Exposure level of the asset"
	opt := &CreateSecurityCategoryOptions{
		Name:        "Exposure",
		Description: &desc,
	}

	// WHEN CreateSecurityCategory is called
	cat, _, err := client.SecurityCategories.CreateSecurityCategory(1, opt)

	// THEN it returns the created category without error
	require.NoError(t, err)
	want := &SecurityCategory{
		ID:                1,
		Name:              "Exposure",
		Description:       &desc,
		MultipleSelection: false,
		EditableState:     SecurityCategoryEditableStateEditable,
		TemplateType:      nil,
	}
	assert.Equal(t, want, cat)
}

func TestSecurityCategoriesService_CreateSecurityCategory_nilOpt(t *testing.T) {
	t.Parallel()

	// GIVEN a nil options argument
	_, client := setup(t)

	// WHEN CreateSecurityCategory is called with nil opt
	_, _, err := client.SecurityCategories.CreateSecurityCategory(1, nil)

	// THEN it returns an error without making an HTTP request
	assert.ErrorContains(t, err, "opt is required")
}

func TestSecurityCategoriesService_CreateSecurityCategory_errors(t *testing.T) {
	t.Parallel()

	// GIVEN the API returns mutation errors
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityCategoryCreate": {
						"securityCategory": null,
						"errors": ["Name has already been taken"]
					}
				}
			}
		`)
	})

	// WHEN CreateSecurityCategory is called
	opt := &CreateSecurityCategoryOptions{Name: "Exposure"}
	_, _, err := client.SecurityCategories.CreateSecurityCategory(1, opt)

	// THEN it returns the mutation error
	assert.ErrorContains(t, err, "Name has already been taken")
}

func TestSecurityCategoriesService_UpdateSecurityCategory(t *testing.T) {
	t.Parallel()

	// GIVEN a valid category ID, namespace ID, and update options
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityCategoryUpdate": {
						"securityCategory": {
							"id": "gid://gitlab/Security::Category/1",
							"name": "Updated Exposure",
							"description": null,
							"multipleSelection": true,
							"editableState": "EDITABLE",
							"templateType": null,
							"securityAttributes": []
						},
						"errors": []
					}
				}
			}
		`)
	})

	name := "Updated Exposure"
	opt := &UpdateSecurityCategoryOptions{Name: &name}

	// WHEN UpdateSecurityCategory is called
	cat, _, err := client.SecurityCategories.UpdateSecurityCategory(1, 2, opt)

	// THEN it returns the updated category without error
	require.NoError(t, err)
	want := &SecurityCategory{
		ID:                1,
		Name:              "Updated Exposure",
		Description:       nil,
		MultipleSelection: true,
		EditableState:     SecurityCategoryEditableStateEditable,
		TemplateType:      nil,
	}
	assert.Equal(t, want, cat)
}

func TestSecurityCategoriesService_UpdateSecurityCategory_notFound(t *testing.T) {
	t.Parallel()

	// GIVEN the API returns a null security category
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityCategoryUpdate": {
						"securityCategory": null,
						"errors": []
					}
				}
			}
		`)
	})

	// WHEN UpdateSecurityCategory is called with a non-existent ID
	_, _, err := client.SecurityCategories.UpdateSecurityCategory(999, 1, nil)

	// THEN it returns ErrNotFound
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestSecurityCategoriesService_UpdateSecurityCategory_errors(t *testing.T) {
	t.Parallel()

	// GIVEN the API returns mutation errors
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityCategoryUpdate": {
						"securityCategory": null,
						"errors": ["Name is too long"]
					}
				}
			}
		`)
	})

	// WHEN UpdateSecurityCategory is called
	name := "x"
	opt := &UpdateSecurityCategoryOptions{Name: &name}
	_, _, err := client.SecurityCategories.UpdateSecurityCategory(1, 2, opt)

	// THEN it returns the mutation error
	assert.ErrorContains(t, err, "Name is too long")
}

func TestSecurityCategoriesService_DestroySecurityCategory(t *testing.T) {
	t.Parallel()

	// GIVEN a valid category ID
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"data": {"securityCategoryDestroy": {"errors": []}}}`)
	})

	// WHEN DestroySecurityCategory is called
	_, err := client.SecurityCategories.DestroySecurityCategory(1)

	// THEN it returns no error
	assert.NoError(t, err)
}

func TestSecurityCategoriesService_DestroySecurityCategory_errors(t *testing.T) {
	t.Parallel()

	// GIVEN the API returns mutation errors
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"data": {"securityCategoryDestroy": {"errors": ["Record not found"]}}}`)
	})

	// WHEN DestroySecurityCategory is called with a non-existent ID
	_, err := client.SecurityCategories.DestroySecurityCategory(999)

	// THEN it returns the mutation error
	assert.ErrorContains(t, err, "Record not found")
}
