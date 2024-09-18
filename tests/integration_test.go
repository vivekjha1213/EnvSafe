package secretsafe_test

import (
    "os"
    "os/exec"
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
)

const binaryName = "envsafe"

func buildBinary(t *testing.T) {
    cmd := exec.Command("go", "build", "-o", binaryName, "../cmd/envsafe/main.go")
    output, err := cmd.CombinedOutput()
    if err != nil {
        t.Fatalf("could not build binary: %v\n%s", err, string(output))
    }
}

func cleanupBinary(t *testing.T) {
    err := os.Remove(binaryName)
    if err != nil {
        t.Fatalf("could not remove binary: %v", err)
    }
}

func runCommand(args ...string) (string, error) {
    cmd := exec.Command("./"+binaryName, args...)
    output, err := cmd.CombinedOutput()
    return strings.TrimSpace(string(output)), err
}

func TestCLI(t *testing.T) {
    buildBinary(t)
    defer cleanupBinary(t)

    t.Run("Set and Get Secret", func(t *testing.T) {
        _, err := runCommand("set", "test_key", "test_value")
        assert.NoError(t, err)

        output, err := runCommand("get", "test_key")
        assert.NoError(t, err)
        assert.Contains(t, output, "test_key=test_value")
    })

    t.Run("Get Non-existent Secret", func(t *testing.T) {
        output, err := runCommand("get", "non_existent")
        assert.NoError(t, err)
        assert.Contains(t, output, "No secret found for key 'non_existent'")
    })

    t.Run("Load From Environment", func(t *testing.T) {
        os.Setenv("TEST_ENV_SECRET", "env_value")
        defer os.Unsetenv("TEST_ENV_SECRET")

        _, err := runCommand("load-env", "TEST_ENV_")
        assert.NoError(t, err)

        output, err := runCommand("get", "SECRET")
        assert.NoError(t, err)
        assert.Contains(t, output, "SECRET=env_value")
    })

    t.Run("Export To Environment", func(t *testing.T) {
        _, err := runCommand("set", "EXPORT_SECRET", "export_value")
        assert.NoError(t, err)

        _, err = runCommand("export-env", "TEST_EXPORT_")
        assert.NoError(t, err)

        assert.Equal(t, "export_value", os.Getenv("TEST_EXPORT_EXPORT_SECRET"))
    })
}