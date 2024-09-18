package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/vivekjha1213/EnvSafe/pkg/secretsafe"
)

var (
    storeFilePath string
    encryptionKey string
)

func init() {
    // Add persistent flags for customization
    rootCmd.PersistentFlags().StringVarP(&storeFilePath, "file", "f", "secrets.json", "Path to the secrets store file")
    rootCmd.PersistentFlags().StringVarP(&encryptionKey, "key", "k", "", "Encryption key for securing secrets")
}

var rootCmd = &cobra.Command{
    Use:   "envsafe",
    Short: "EnvSafe is a secure environment variable manager",
    Long:  `EnvSafe allows you to securely store, retrieve, and manage environment variables.`,
}

var setCmd = &cobra.Command{
    Use:   "set <key> <value>",
    Short: "Set a secret",
    Args:  cobra.ExactArgs(2), // Ensures exactly 2 arguments are provided
    Run: func(cmd *cobra.Command, args []string) {
        key := args[0]
        value := args[1]

        if encryptionKey == "" {
            fmt.Println("Error: Encryption key is required")
            os.Exit(1)
        }

        store, err := secretsafe.LoadSecretStore(storeFilePath)
        if err != nil && !os.IsNotExist(err) {
            fmt.Printf("Error loading secret store: %v\n", err)
            os.Exit(1)
        }

        if store == nil {
            store = secretsafe.NewSecretStore()
        }

        err = store.Set(key, value, []byte(encryptionKey))
        if err != nil {
            fmt.Printf("Error setting secret: %v\n", err)
            os.Exit(1)
        }

        if err := store.Save(storeFilePath); err != nil {
            fmt.Printf("Error saving secret store: %v\n", err)
            os.Exit(1)
        }

        fmt.Printf("Secret for key '%s' set successfully\n", key)
    },
}

var getCmd = &cobra.Command{
    Use:   "get <key>",
    Short: "Get a secret",
    Args:  cobra.ExactArgs(1), // Ensures exactly 1 argument is provided
    Run: func(cmd *cobra.Command, args []string) {
        key := args[0]

        if encryptionKey == "" {
            fmt.Println("Error: Encryption key is required")
            os.Exit(1)
        }

        store, err := secretsafe.LoadSecretStore(storeFilePath)
        if err != nil {
            fmt.Printf("Error loading secret store: %v\n", err)
            os.Exit(1)
        }

        value, err := store.Get(key, []byte(encryptionKey))
        if err != nil {
            fmt.Printf("Error retrieving secret for key '%s': %v\n", key, err)
            os.Exit(1)
        }

        fmt.Printf("Value for key '%s': %s\n", key, value)
    },
}

var loadFromEnvCmd = &cobra.Command{
    Use:   "load-env <prefix>",
    Short: "Load secrets from environment variables",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        prefix := args[0]

        store := secretsafe.NewSecretStore()
        err := store.LoadFromEnv(prefix)
        if err != nil {
            fmt.Printf("Error loading secrets from environment: %v\n", err)
            os.Exit(1)
        }

        err = store.Save(storeFilePath)
        if err != nil {
            fmt.Printf("Error saving secret store: %v\n", err)
            os.Exit(1)
        }

        fmt.Println("Secrets loaded from environment variables successfully")
    },
}

var exportToEnvCmd = &cobra.Command{
    Use:   "export-env <prefix>",
    Short: "Export secrets to environment variables",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        prefix := args[0]

        store, err := secretsafe.LoadSecretStore(storeFilePath)
        if err != nil {
            fmt.Printf("Error loading secret store: %v\n", err)
            os.Exit(1)
        }

        err = store.ExportToEnv(prefix)
        if err != nil {
            fmt.Printf("Error exporting secrets to environment: %v\n", err)
            os.Exit(1)
        }

        fmt.Println("Secrets exported to environment variables successfully")
    },
}

func main() {
    rootCmd.AddCommand(setCmd, getCmd, loadFromEnvCmd, exportToEnvCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Printf("Error executing command: %v\n", err)
        os.Exit(1)
    }
}
