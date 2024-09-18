package secretsafe

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "io"
)

// ValidateKeySize checks if the encryption key is valid for AES.
func ValidateKeySize(key []byte) error {
    switch len(key) {
    case 16, 24, 32:
        return nil
    default:
        return errors.New("invalid key size: must be 16, 24, or 32 bytes")
    }
}

// Encrypt encrypts plaintext using AES encryption with the provided key.
func Encrypt(plaintext string, key []byte) (string, error) {
    if err := ValidateKeySize(key); err != nil {
        return "", err
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

    return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts ciphertext using AES encryption with the provided key.
func Decrypt(ciphertext string, key []byte) (string, error) {
    if err := ValidateKeySize(key); err != nil {
        return "", err
    }

    ciphertextBytes, err := base64.URLEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    if len(ciphertextBytes) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }

    iv := ciphertextBytes[:aes.BlockSize]
    ciphertextBytes = ciphertextBytes[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertextBytes, ciphertextBytes)

    return string(ciphertextBytes), nil
}
