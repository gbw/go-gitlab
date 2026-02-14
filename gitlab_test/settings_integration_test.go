//go:build integration

package gitlab_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

func Test_Settings_AnonymousSearchesAllowed_Integration(t *testing.T) {
	client := SetupIntegrationClient(t)
	// get initial settings
	settings, _, err := client.Settings.GetSettings()
	require.NoError(t, err, "Failed to get initial settings")

	// get AnonymousSearchesAllowed value
	originalValue := settings.AnonymousSearchesAllowed

	t.Cleanup(func() {
		// restore original setting after test
		_, _, err := client.Settings.UpdateSettings(&gitlab.UpdateSettingsOptions{
			AnonymousSearchesAllowed: &originalValue,
		})
		require.NoError(t, err, "Failed to restore settings")
	})

	// Toggle the AnonymousSearchesAllowed setting
	newValue := !originalValue
	// update setting with the new value for AnonymousSearchesAllowed
	updatedSettings, _, err := client.Settings.UpdateSettings(&gitlab.UpdateSettingsOptions{
		AnonymousSearchesAllowed: &newValue,
	})
	require.NoError(t, err, "Failed to update settings")
	assert.Equal(t, newValue, updatedSettings.AnonymousSearchesAllowed)
}
