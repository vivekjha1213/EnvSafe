package secretsafe

import (
    "os"
    "strings"
)

// LoadFromEnv loads secrets from environment variables with a given prefix
func (s *SecretStore) LoadFromEnv(prefix string) error {
    for _, env := range os.Environ() {
        if key, value, found := strings.Cut(env, "="); found && strings.HasPrefix(key, prefix) {
            trimmedKey := strings.TrimPrefix(key, prefix)
            s.Secrets[trimmedKey] = value
        }
    }
    return nil
}

// ExportToEnv exports secrets to environment variables with a given prefix
func (s *SecretStore) ExportToEnv(prefix string) error {
    for key, value := range s.Secrets {
        if err := os.Setenv(prefix+key, value); err != nil {
            return err
        }
    }
    return nil
}

// LoadFromFile loads secrets from a file
func (s *SecretStore) LoadFromFile(filename string, encryptionKey []byte) error {
    loadedStore, err := LoadSecretStore(filename)
    if err != nil {
        return err
    }

    for key, encryptedValue := range loadedStore.Secrets {
        decryptedValue, err := Decrypt(encryptedValue, encryptionKey)
        if err != nil {
            return err
        }
        s.Secrets[key] = decryptedValue
    }

    s.Version = loadedStore.Version
    return nil
}

// ExportToFile exports secrets to a file
func (s *SecretStore) ExportToFile(filename string, encryptionKey []byte) error {
    exportStore := NewSecretStore()
    exportStore.Version = s.Version

    for key, value := range s.Secrets {
        encryptedValue, err := Encrypt(value, encryptionKey)
        if err != nil {
            return err
        }
        exportStore.Secrets[key] = encryptedValue
    }

    return exportStore.Save(filename)
}