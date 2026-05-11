package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSecurityAttributesService_CreateSecurityAttributes(t *testing.T) {
	t.Parallel()

	// GIVEN a valid namespace ID, category ID, and attribute options
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityAttributeCreate": {
						"securityAttributes": [
							{
								"id": "gid://gitlab/Security::Attribute/1",
								"name": "Internal",
								"color": "#ff0000",
								"description": "Internal exposure",
								"editableState": "EDITABLE",
								"securityCategory": {
									"id": "gid://gitlab/Security::Category/1",
									"name": "Exposure",
									"description": null,
									"multipleSelection": false,
									"editableState": "EDITABLE",
									"templateType": null
								}
							}
						],
						"errors": []
					}
				}
			}
		`)
	})

	opt := &CreateSecurityAttributesOptions{
		Attributes: Ptr([]SecurityAttributeInput{
			{Name: Ptr("Internal"), Description: Ptr("Internal exposure"), Color: Ptr("#ff0000")},
		}),
	}

	// WHEN CreateSecurityAttributes is called
	attrs, _, err := client.SecurityAttributes.CreateSecurityAttributes(1, 1, opt)

	// THEN it returns the created attributes without error
	require.NoError(t, err)
	require.Len(t, attrs, 1)

	want := &SecurityAttribute{
		ID:            1,
		Name:          "Internal",
		Color:         "#ff0000",
		Description:   "Internal exposure",
		EditableState: SecurityCategoryEditableStateEditable,
		SecurityCategory: &SecurityCategory{
			ID:                1,
			Name:              "Exposure",
			Description:       nil,
			MultipleSelection: false,
			EditableState:     SecurityCategoryEditableStateEditable,
			TemplateType:      nil,
		},
	}
	assert.Equal(t, want, attrs[0])
}

func TestSecurityAttributesService_CreateSecurityAttributes_nilOpt(t *testing.T) {
	t.Parallel()

	// GIVEN a nil options argument
	_, client := setup(t)

	// WHEN CreateSecurityAttributes is called with nil opt
	_, _, err := client.SecurityAttributes.CreateSecurityAttributes(1, 1, nil)

	// THEN it returns an error without making an HTTP request
	assert.ErrorContains(t, err, "opt is required")
}

func TestSecurityAttributesService_CreateSecurityAttributes_errors(t *testing.T) {
	t.Parallel()

	// GIVEN the API returns mutation errors
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityAttributeCreate": {
						"securityAttributes": [],
						"errors": ["Name has already been taken"]
					}
				}
			}
		`)
	})

	// WHEN CreateSecurityAttributes is called
	opt := &CreateSecurityAttributesOptions{
		Attributes: Ptr([]SecurityAttributeInput{
			{Name: Ptr("Internal"), Description: Ptr("desc"), Color: Ptr("#ff0000")},
		}),
	}
	_, _, err := client.SecurityAttributes.CreateSecurityAttributes(1, 1, opt)

	// THEN it returns the mutation error
	assert.ErrorContains(t, err, "Name has already been taken")
}

func TestSecurityAttributesService_UpdateSecurityAttribute(t *testing.T) {
	t.Parallel()

	// GIVEN a valid attribute ID and update options
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityAttributeUpdate": {
						"securityAttribute": {
							"id": "gid://gitlab/Security::Attribute/1",
							"name": "External",
							"color": "#00ff00",
							"description": null,
							"editableState": "EDITABLE",
							"securityCategory": null
						},
						"errors": []
					}
				}
			}
		`)
	})

	name := "External"
	color := "#00ff00"
	opt := &UpdateSecurityAttributeOptions{Name: &name, Color: &color}

	// WHEN UpdateSecurityAttribute is called
	attr, _, err := client.SecurityAttributes.UpdateSecurityAttribute(1, opt)

	// THEN it returns the updated attribute without error
	require.NoError(t, err)
	want := &SecurityAttribute{
		ID:               1,
		Name:             "External",
		Color:            "#00ff00",
		Description:      "",
		EditableState:    SecurityCategoryEditableStateEditable,
		SecurityCategory: nil,
	}
	assert.Equal(t, want, attr)
}

func TestSecurityAttributesService_UpdateSecurityAttribute_notFound(t *testing.T) {
	t.Parallel()

	// GIVEN the API returns a null security attribute
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityAttributeUpdate": {
						"securityAttribute": null,
						"errors": []
					}
				}
			}
		`)
	})

	// WHEN UpdateSecurityAttribute is called with a non-existent ID
	_, _, err := client.SecurityAttributes.UpdateSecurityAttribute(999, nil)

	// THEN it returns ErrNotFound
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestSecurityAttributesService_UpdateSecurityAttribute_errors(t *testing.T) {
	t.Parallel()

	// GIVEN the API returns mutation errors
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityAttributeUpdate": {
						"securityAttribute": null,
						"errors": ["Color is invalid"]
					}
				}
			}
		`)
	})

	// WHEN UpdateSecurityAttribute is called with an invalid color
	color := "notacolor"
	opt := &UpdateSecurityAttributeOptions{Color: &color}
	_, _, err := client.SecurityAttributes.UpdateSecurityAttribute(1, opt)

	// THEN it returns the mutation error
	assert.ErrorContains(t, err, "Color is invalid")
}

func TestSecurityAttributesService_DestroySecurityAttribute(t *testing.T) {
	t.Parallel()

	// GIVEN a valid attribute ID
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"data": {"securityAttributeDestroy": {"errors": []}}}`)
	})

	// WHEN DestroySecurityAttribute is called
	_, err := client.SecurityAttributes.DestroySecurityAttribute(1)

	// THEN it returns no error
	assert.NoError(t, err)
}

func TestSecurityAttributesService_DestroySecurityAttribute_errors(t *testing.T) {
	t.Parallel()

	// GIVEN the API returns mutation errors
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"data": {"securityAttributeDestroy": {"errors": ["Record not found"]}}}`)
	})

	// WHEN DestroySecurityAttribute is called with a non-existent ID
	_, err := client.SecurityAttributes.DestroySecurityAttribute(999)

	// THEN it returns the mutation error
	assert.ErrorContains(t, err, "Record not found")
}

func TestSecurityAttributesService_ProjectUpdateSecurityAttribute(t *testing.T) {
	t.Parallel()

	// GIVEN a valid project ID and update options with attribute IDs to add and remove
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityAttributeProjectUpdate": {
						"addedCount": 2,
						"removedCount": 1,
						"errors": []
					}
				}
			}
		`)
	})

	opt := &ProjectUpdateSecurityAttributeOptions{
		AddAttributeIDs:    Ptr([]int64{1, 2}),
		RemoveAttributeIDs: Ptr([]int64{3}),
	}

	// WHEN ProjectUpdateSecurityAttribute is called
	result, _, err := client.SecurityAttributes.ProjectUpdateSecurityAttribute(10, opt)

	// THEN it returns the added and removed counts without error
	require.NoError(t, err)
	want := &SecurityAttributeProjectUpdateResult{
		AddedCount:   2,
		RemovedCount: 1,
	}
	assert.Equal(t, want, result)
}

func TestSecurityAttributesService_ProjectUpdateSecurityAttribute_nilOpt(t *testing.T) {
	t.Parallel()

	// GIVEN a nil options argument
	_, client := setup(t)

	// WHEN ProjectUpdateSecurityAttribute is called with nil opt
	_, _, err := client.SecurityAttributes.ProjectUpdateSecurityAttribute(1, nil)

	// THEN it returns an error without making an HTTP request
	assert.ErrorContains(t, err, "opt is required")
}

func TestSecurityAttributesService_ProjectUpdateSecurityAttribute_errors(t *testing.T) {
	t.Parallel()

	// GIVEN the API returns mutation errors
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityAttributeProjectUpdate": {
						"addedCount": null,
						"removedCount": null,
						"errors": ["Project not found"]
					}
				}
			}
		`)
	})

	// WHEN ProjectUpdateSecurityAttribute is called with a non-existent project
	opt := &ProjectUpdateSecurityAttributeOptions{AddAttributeIDs: Ptr([]int64{1})}
	_, _, err := client.SecurityAttributes.ProjectUpdateSecurityAttribute(999, opt)

	// THEN it returns the mutation error
	assert.ErrorContains(t, err, "Project not found")
}

func TestSecurityAttributesService_BulkUpdateSecurityAttributes(t *testing.T) {
	t.Parallel()

	// GIVEN valid group IDs, project IDs, attribute IDs, and an ADD mode
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"data": {"bulkUpdateSecurityAttributes": {"errors": []}}}`)
	})

	opt := &BulkUpdateSecurityAttributesOptions{
		GroupIDs:     Ptr([]int64{1}),
		ProjectIDs:   Ptr([]int64{2, 3}),
		AttributeIDs: Ptr([]int64{10}),
		Mode:         Ptr(SecurityAttributeBulkUpdateModeAdd),
	}

	// WHEN BulkUpdateSecurityAttributes is called
	_, err := client.SecurityAttributes.BulkUpdateSecurityAttributes(opt)

	// THEN it returns no error
	assert.NoError(t, err)
}

func TestSecurityAttributesService_BulkUpdateSecurityAttributes_nilOpt(t *testing.T) {
	t.Parallel()

	// GIVEN a nil options argument
	_, client := setup(t)

	// WHEN BulkUpdateSecurityAttributes is called with nil opt
	_, err := client.SecurityAttributes.BulkUpdateSecurityAttributes(nil)

	// THEN it returns an error without making an HTTP request
	assert.ErrorContains(t, err, "opt is required")
}

func TestSecurityAttributesService_BulkUpdateSecurityAttributes_errors(t *testing.T) {
	t.Parallel()

	// GIVEN the API returns mutation errors
	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"data": {"bulkUpdateSecurityAttributes": {"errors": ["Insufficient permissions"]}}}`)
	})

	// WHEN BulkUpdateSecurityAttributes is called
	opt := &BulkUpdateSecurityAttributesOptions{
		AttributeIDs: Ptr([]int64{1}),
		Mode:         Ptr(SecurityAttributeBulkUpdateModeRemove),
	}
	_, err := client.SecurityAttributes.BulkUpdateSecurityAttributes(opt)

	// THEN it returns the mutation error
	assert.ErrorContains(t, err, "Insufficient permissions")
}
