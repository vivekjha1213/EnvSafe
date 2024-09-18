package secretsafe

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

type SecretStore struct {
    Secrets map[string]string `json:"secrets"`
    Version string            `json:"version"`
}

func NewSecretStore() *SecretStore {
    return &SecretStore{
        Secrets: make(map[string]string),
        Version: "1.0",
    }
}

func (s *SecretStore) Set(key, value string, encryptionKey []byte) error {
    encryptedValue, err := Encrypt(value, encryptionKey)
    if err != nil {
        return err
    }
    s.Secrets[key] = encryptedValue
    return nil
}

func (s *SecretStore) Get(key string, encryptionKey []byte) (string, error) {
    encryptedValue, ok := s.Secrets[key]
    if !ok {
        return "", nil
    }
    return Decrypt(encryptedValue, encryptionKey)
}

func (s *SecretStore) Save(filename string) error {
    data, err := json.MarshalIndent(s, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(filename, data, 0600)
}

func LoadSecretStore(filename string) (*SecretStore, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            return NewSecretStore(), nil
        }
        return nil, err
    }

    var store SecretStore
    err = json.Unmarshal(data, &store)
    if err != nil {
        return nil, err
    }

    return &store, nil
}