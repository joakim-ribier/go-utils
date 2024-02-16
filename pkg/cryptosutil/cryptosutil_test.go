package cryptosutil

import (
	"testing"
)

var secret = "Hello Word"

// TestMD5Hashing calls cryptosutil.MD5Hashing,
// checking for a valid return value.
func TestMD5Hashing(t *testing.T) {
	expected := "ed0a96e83ab7b0910fcbcc131b2e6b82"

	if hash := MD5Hashing("Hello Word"); hash != expected {
		t.Fatalf(`result: {%v} but expected: {%v}`, hash, expected)
	}
}

// TestEncrypt calls cryptosutil.Encrypt,
// checking for a valid return value.
func TestEncrypt(t *testing.T) {
	expected := "my-password-high-sensibility"

	if _, err := Encrypt([]byte(expected), secret); err != nil {
		t.Fatal(err)
	}
}

// TestDecryptNoEncryptedData calls cryptosutil.Decrypt,
// checking for a valid return value.
func TestDecryptNoEncryptedData(t *testing.T) {
	if r, err := Decrypt([]byte("at least 12 characters"), secret); err == nil {
		t.Fatalf(`result: {%v} but expected {%s}`, r, "cipher: message authentication failed")
	}
}

// TestEncrypt calls cryptosutil.Encrypt,
// checking for a valid return value.
func TestEncryptEqualsDecrypt(t *testing.T) {
	expected := "my-password-high-sensibility"

	encrypted, _ := Encrypt([]byte(expected), secret)

	if decrypted, _ := Decrypt(encrypted, secret); string(decrypted) != expected {
		t.Fatalf(`result: {%s} but expected {%s}`, string(decrypted), expected)
	}

	if _, err := Decrypt(encrypted, "bad secret"); err == nil {
		t.Fatalf(`result: {%s} but expected {%s}`, err.Error(), expected)
	}
}
