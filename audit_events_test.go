package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuditEventsService_ListInstanceAuditEvents(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/audit_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 1,
				"author_id": 1,
				"entity_id": 6,
				"entity_type": "Project",
				"event_name": "project_archived",
				"details": {
				  "event_name": "project_archived",
				  "custom_message": "Project archived",
				  "author_name": "Venkatesh Thalluri",
				  "target_id": "flightjs/flight",
				  "target_type": "Project",
				  "target_details": "flightjs/flight",
				  "ip_address": "127.0.0.1",
				  "entity_path": "flightjs/flight"
				}
			  }
			]
		`)
	})

	want := []*AuditEvent{{
		ID:         1,
		AuthorID:   1,
		EntityID:   6,
		EntityType: "Project",
		EventName:  "project_archived",
		Details: AuditEventDetails{
			EventName:     "project_archived",
			CustomMessage: "Project archived",
			AuthorName:    "Venkatesh Thalluri",
			TargetID:      "flightjs/flight",
			TargetType:    "Project",
			TargetDetails: "flightjs/flight",
			IPAddress:     "127.0.0.1",
			EntityPath:    "flightjs/flight",
		},
	}}

	aes, resp, err := client.AuditEvents.ListInstanceAuditEvents(nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, aes)

	aes, resp, err = client.AuditEvents.ListInstanceAuditEvents(nil, errorOption)
	require.ErrorIs(t, err, errRequestOptionFunc)
	require.Nil(t, resp)
	require.Nil(t, aes)
}

func TestAuditEventsService_ListInstanceAuditEvents_StatusNotFound(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/audit_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
	})

	aes, resp, err := client.AuditEvents.ListInstanceAuditEvents(nil)
	require.Error(t, err)
	require.Nil(t, aes)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAuditEventsService_GetInstanceAuditEvent(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/audit_events/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 1,
			"author_id": 1,
			"entity_id": 6,
			"entity_type": "Project",
			"event_name": "project_archived",
			"details": {
			  "event_name": "project_archived",
			  "custom_message": "Project archived",
			  "author_name": "Venkatesh Thalluri",
			  "target_id": "flightjs/flight",
			  "target_type": "Project",
			  "target_details": "flightjs/flight",
			  "ip_address": "127.0.0.1",
			  "entity_path": "flightjs/flight"
			}
		  }
		`)
	})

	want := &AuditEvent{
		ID:         1,
		AuthorID:   1,
		EntityID:   6,
		EventName:  "project_archived",
		EntityType: "Project",
		Details: AuditEventDetails{
			EventName:     "project_archived",
			CustomMessage: "Project archived",
			AuthorName:    "Venkatesh Thalluri",
			TargetID:      "flightjs/flight",
			TargetType:    "Project",
			TargetDetails: "flightjs/flight",
			IPAddress:     "127.0.0.1",
			EntityPath:    "flightjs/flight",
		},
	}

	ae, resp, err := client.AuditEvents.GetInstanceAuditEvent(1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AuditEvents.GetInstanceAuditEvent(1, nil, errorOption)
	require.ErrorIs(t, err, errRequestOptionFunc)
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AuditEvents.GetInstanceAuditEvent(3, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAuditEventsService_ListGroupAuditEvents(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/6/audit_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 1,
				"author_id": 1,
				"entity_id": 6,
				"entity_type": "Group",
			    "event_name": "group_archived",
				"details": {
			      "event_name": "group_archived",
				  "custom_message": "Group archived",
				  "author_name": "Venkatesh Thalluri",
				  "target_id": "flightjs/flight",
				  "target_type": "Group",
				  "target_details": "flightjs/flight",
				  "ip_address": "127.0.0.1",
				  "entity_path": "flightjs/flight"
				}
			  }
			]
		`)
	})

	want := []*AuditEvent{{
		ID:         1,
		AuthorID:   1,
		EntityID:   6,
		EntityType: "Group",
		EventName:  "group_archived",
		Details: AuditEventDetails{
			EventName:     "group_archived",
			CustomMessage: "Group archived",
			AuthorName:    "Venkatesh Thalluri",
			TargetID:      "flightjs/flight",
			TargetType:    "Group",
			TargetDetails: "flightjs/flight",
			IPAddress:     "127.0.0.1",
			EntityPath:    "flightjs/flight",
		},
	}}

	aes, resp, err := client.AuditEvents.ListGroupAuditEvents(6, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, aes)

	aes, resp, err = client.AuditEvents.ListGroupAuditEvents(6.01, nil)
	require.ErrorIs(t, err, ErrInvalidIDType)
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AuditEvents.ListGroupAuditEvents(6, nil, errorOption)
	require.ErrorIs(t, err, errRequestOptionFunc)
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AuditEvents.ListGroupAuditEvents(3, nil)
	require.Error(t, err)
	require.Nil(t, aes)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAuditEventsService_GetGroupAuditEvent(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/6/audit_events/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 1,
			"author_id": 1,
			"entity_id": 6,
			"entity_type": "Group",
			"event_name": "group_archived",
			"details": {
			  "event_name": "group_archived",
			  "custom_message": "Group archived",
			  "author_name": "Venkatesh Thalluri",
			  "target_id": "flightjs/flight",
			  "target_type": "Group",
			  "target_details": "flightjs/flight",
			  "ip_address": "127.0.0.1",
			  "entity_path": "flightjs/flight"
			}
		  }
		`)
	})

	want := &AuditEvent{
		ID:         1,
		AuthorID:   1,
		EntityID:   6,
		EntityType: "Group",
		EventName:  "group_archived",
		Details: AuditEventDetails{
			EventName:     "group_archived",
			CustomMessage: "Group archived",
			AuthorName:    "Venkatesh Thalluri",
			TargetID:      "flightjs/flight",
			TargetType:    "Group",
			TargetDetails: "flightjs/flight",
			IPAddress:     "127.0.0.1",
			EntityPath:    "flightjs/flight",
		},
	}

	ae, resp, err := client.AuditEvents.GetGroupAuditEvent(6, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AuditEvents.GetGroupAuditEvent(6.01, 1, nil)
	require.ErrorIs(t, err, ErrInvalidIDType)
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AuditEvents.GetGroupAuditEvent(6, 1, nil, errorOption)
	require.ErrorIs(t, err, errRequestOptionFunc)
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AuditEvents.GetGroupAuditEvent(3, 1, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAuditEventsService_ListProjectAuditEvents(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/6/audit_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 1,
				"author_id": 1,
				"entity_id": 6,
				"entity_type": "Project",
				"event_name": "project_archived",
				"details": {
				  "event_name": "project_archived",
				  "custom_message": "Project archived",
				  "author_name": "Venkatesh Thalluri",
				  "target_id": "flightjs/flight",
				  "target_type": "Project",
				  "target_details": "flightjs/flight",
				  "ip_address": "127.0.0.1",
				  "entity_path": "flightjs/flight"
				}
			  }
			]
		`)
	})

	want := []*AuditEvent{{
		ID:         1,
		AuthorID:   1,
		EntityID:   6,
		EntityType: "Project",
		EventName:  "project_archived",
		Details: AuditEventDetails{
			EventName:     "project_archived",
			CustomMessage: "Project archived",
			AuthorName:    "Venkatesh Thalluri",
			TargetID:      "flightjs/flight",
			TargetType:    "Project",
			TargetDetails: "flightjs/flight",
			IPAddress:     "127.0.0.1",
			EntityPath:    "flightjs/flight",
		},
	}}

	aes, resp, err := client.AuditEvents.ListProjectAuditEvents(6, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, aes)

	aes, resp, err = client.AuditEvents.ListProjectAuditEvents(6.01, nil)
	require.ErrorIs(t, err, ErrInvalidIDType)
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AuditEvents.ListProjectAuditEvents(6, nil, errorOption)
	require.ErrorIs(t, err, errRequestOptionFunc)
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AuditEvents.ListProjectAuditEvents(3, nil)
	require.Error(t, err)
	require.Nil(t, aes)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAuditEventsService_GetProjectAuditEvent(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/6/audit_events/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 1,
			"author_id": 1,
			"entity_id": 6,
			"entity_type": "Project",
			"event_name": "project_archived",
			"details": {
			  "event_name": "project_archived",
			  "custom_message": "Project archived",
			  "author_name": "Venkatesh Thalluri",
			  "target_id": "flightjs/flight",
			  "target_type": "Project",
			  "target_details": "flightjs/flight",
			  "ip_address": "127.0.0.1",
			  "entity_path": "flightjs/flight"
			}
		  }
		`)
	})

	want := &AuditEvent{
		ID:         1,
		AuthorID:   1,
		EntityID:   6,
		EntityType: "Project",
		EventName:  "project_archived",
		Details: AuditEventDetails{
			EventName:     "project_archived",
			CustomMessage: "Project archived",
			AuthorName:    "Venkatesh Thalluri",
			TargetID:      "flightjs/flight",
			TargetType:    "Project",
			TargetDetails: "flightjs/flight",
			IPAddress:     "127.0.0.1",
			EntityPath:    "flightjs/flight",
		},
	}

	ae, resp, err := client.AuditEvents.GetProjectAuditEvent(6, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AuditEvents.GetProjectAuditEvent(6.01, 1, nil)
	require.ErrorIs(t, err, ErrInvalidIDType)
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AuditEvents.GetProjectAuditEvent(6, 1, nil, errorOption)
	require.ErrorIs(t, err, errRequestOptionFunc)
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AuditEvents.GetProjectAuditEvent(3, 1, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
