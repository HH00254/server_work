package Util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// generateEncryptionKey generates a 256 bit (32 byte) AES encryption key and
// prints the base64 representation.
func GenEncryptionKey() error {
	// This is included for demonstration purposes. You should generate your own
	// key. Please remember that encryption keys should be handled with a
	// comprehensive security policy.
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return fmt.Errorf("rand.Read: %w", err)
	}
	encryptionKey := base64.StdEncoding.EncodeToString(key)
	fmt.Printf("Generated base64-encoded encryption key: %v\n", encryptionKey)
	return nil
}
