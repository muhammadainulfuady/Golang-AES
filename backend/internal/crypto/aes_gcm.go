package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

// GenerateKey menghasilkan kunci acak 256-bit (32 bytes) dalam format hexadecimal 64 karakter.
func GenerateKey() (string, error) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

// Encrypt mengenkripsi teks biasa (plaintext) menjadi ciphertext format Base64 menggunakan AES-256-GCM.
func Encrypt(plaintext, keyHex string) (string, error) {
	// 1. Decode kunci hex ke bytes
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return "", errors.New("format kunci tidak valid (harus hexadecimal)")
	}
	if len(key) != 32 {
		return "", errors.New("panjang kunci harus 32 bytes (64 karakter hex)")
	}

	// 2. Buat AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 3. Bungkus block dengan mode Galois/Counter Mode (GCM)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 4. Buat Nonce/IV unik berukuran 12 bytes
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 5. Enkripsi pesan (gcm.Seal otomatis menempelkan nonce di awal ciphertext)
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// 6. Return dalam format Base64 agar aman ditransmisikan
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt mendekripsi ciphertext Base64 kembali ke teks asli dan memverifikasi integritas payload.
func Decrypt(ciphertextB64, keyHex string) (string, error) {
	// 1. Decode kunci hex ke bytes
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return "", errors.New("format kunci tidak valid (harus hexadecimal)")
	}
	if len(key) != 32 {
		return "", errors.New("panjang kunci harus 32 bytes (64 karakter hex)")
	}

	// 2. Decode ciphertext dari Base64
	data, err := base64.StdEncoding.DecodeString(ciphertextB64)
	if err != nil {
		return "", errors.New("format ciphertext tidak valid (harus Base64)")
	}

	// 3. Buat AES cipher block dan mode GCM
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 4. Pastikan data minimal memiliki panjang Nonce (12 bytes)
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext terlalu pendek atau rusak")
	}

	// 5. Pisahkan Nonce (12 bytes pertama) dan Ciphertext
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// 6. Dekripsi dan otentikasi data
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.New("gagal mendekripsi: kunci salah atau ciphertext telah dimanipulasi")
	}

	return string(plaintext), nil
}
