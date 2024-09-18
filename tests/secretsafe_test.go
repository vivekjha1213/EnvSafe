package secretsafe_test

import (
    "os"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/vivekjha1213/EnvSafe/pkg/secretsafe"
)

func TestSecretStore(t *testing.T) {
    store := secretsafe.NewSecretStore()
    key := []byte("testkey1234567890123456789012")

    t.Run("Set and Get Secret", func(t *testing.T) {
        err := store.Set("test_key", "test_value", key)
        assert.NoError(t, err)

        value, err := store.Get("test_key", key)
        assert.NoError(t, err)
        assert.Equal(t, "test_value", value)
    })

    t.Run("Get Non-existent Secret", func(t *testing.T) {
        value, err := store.Get("non_existent", key)
        assert.NoError(t, err)
        assert.Empty(t, value)
    })

    t.Run("Update Existing Secret", func(t *testing.T) {
        err := store.Set("test_key", "updated_value", key)
        assert.NoError(t, err)

        value, err := store.Get("test_key", key)
        assert.NoError(t, err)
        assert.Equal(t, "updated_value", value)
    })

    t.Run("Save and Load Store", func(t *testing.T) {
        tempFile := "temp_secrets.json"
        defer os.Remove(tempFile)

        err := store.Save(tempFile)
        assert.NoError(t, err)

        loadedStore, err := secretsafe.LoadSecretStore(tempFile)
        assert.NoError(t, err)

        assert.Equal(t, store.Version, loadedStore.Version)
        assert.Equal(t, len(store.Secrets), len(loadedStore.Secrets))

        for k, v := range store.Secrets {
            assert.Equal(t, v, loadedStore.Secrets[k])
        }
    })

    t.Run("Load From Environment", func(t *testing.T) {
        os.Setenv("TEST_ENV_SECRET1", "env_value1")
        os.Setenv("TEST_ENV_SECRET2", "env_value2")
        defer os.Unsetenv("TEST_ENV_SECRET1")
        defer os.Unsetenv("TEST_ENV_SECRET2")

        err := store.LoadFromEnv("TEST_ENV_")
        assert.NoError(t, err)

        value, err := store.Get("SECRET1", key)
        assert.NoError(t, err)
        assert.Equal(t, "env_value1", value)

        value, err = store.Get("SECRET2", key)
        assert.NoError(t, err)
        assert.Equal(t, "env_value2", value)
    })

    t.Run("Export To Environment", func(t *testing.T) {
        err := store.Set("EXPORT_SECRET", "export_value", key)
        assert.NoError(t, err)

        err = store.ExportToEnv("TEST_EXPORT_")
        assert.NoError(t, err)

        assert.Equal(t, "export_value", os.Getenv("TEST_EXPORT_EXPORT_SECRET"))
    })

    t.Run("Increment Version", func(t *testing.T) {
        initialVersion := store.Version
        err := store.IncrementVersion("minor")
        assert.NoError(t, err)
        assert.NotEqual(t, initialVersion, store.Version)
    })
}