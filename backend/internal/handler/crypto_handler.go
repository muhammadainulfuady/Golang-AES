package handler

import (
	"crypto-backend/internal/crypto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// BaseResponse format baku respon JSON sesuai dokumen spesifikasi API
type BaseResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Pesan  string      `json:"pesan"`
	Data   interface{} `json:"data"`
}

// Request Models
type EncryptRequest struct {
	Plaintext string `json:"plaintext"`
	Key       string `json:"key"`
}

type DecryptRequest struct {
	Ciphertext string `json:"ciphertext"`
	Key        string `json:"key"`
}

// Response Data Models
type KeyData struct {
	Key         string `json:"key"`
	LengthBytes int    `json:"length_bytes"`
	BitLength   int    `json:"bit_length"`
}

type EncryptData struct {
	Ciphertext string `json:"ciphertext"`
	Algorithm  string `json:"algorithm"`
}

type DecryptData struct {
	Plaintext string `json:"plaintext"`
}

// Helper untuk mengirim respon error standar
func sendError(c *gin.Context, httpCode int, pesan string) {
	c.JSON(httpCode, BaseResponse{
		Code:   httpCode,
		Status: "error",
		Pesan:  pesan,
		Data:   []interface{}{},
	})
}

// GenerateKeyHandler menangani pembuatan kunci AES-256 acak baru
func GenerateKeyHandler(c *gin.Context) {
	keyHex, err := crypto.GenerateKey()
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Gagal menghasilkan kunci acak internal")
		return
	}

	c.JSON(http.StatusOK, BaseResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Pesan:  "Secret Key AES-256 berhasil dibuat",
		Data: []KeyData{
			{
				Key:         keyHex,
				LengthBytes: 32,
				BitLength:   256,
			},
		},
	})
}

// EncryptHandler menangani enkripsi pesan teks
func EncryptHandler(c *gin.Context) {
	var req EncryptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, http.StatusBadRequest, "Format JSON tidak valid")
		return
	}

	req.Key = strings.TrimSpace(req.Key)
	if req.Plaintext == "" {
		sendError(c, http.StatusBadRequest, "Pesan teks (plaintext) tidak boleh kosong")
		return
	}

	if len(req.Key) != 64 {
		sendError(c, http.StatusBadRequest, "Kunci harus berupa 64 karakter hexadecimal (256-bit)")
		return
	}

	ciphertext, err := crypto.Encrypt(req.Plaintext, req.Key)
	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, BaseResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Pesan:  "Pesan berhasil dienkripsi",
		Data: []EncryptData{
			{
				Ciphertext: ciphertext,
				Algorithm:  "AES-256-GCM",
			},
		},
	})
}

// DecryptHandler menangani dekripsi pesan teks
func DecryptHandler(c *gin.Context) {
	var req DecryptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, http.StatusBadRequest, "Format JSON tidak valid")
		return
	}

	req.Ciphertext = strings.TrimSpace(req.Ciphertext)
	req.Key = strings.TrimSpace(req.Key)

	if req.Ciphertext == "" {
		sendError(c, http.StatusBadRequest, "Ciphertext tidak boleh kosong")
		return
	}

	if len(req.Key) != 64 {
		sendError(c, http.StatusBadRequest, "Kunci harus berupa 64 karakter hexadecimal (256-bit)")
		return
	}

	plaintext, err := crypto.Decrypt(req.Ciphertext, req.Key)
	if err != nil {
		sendError(c, http.StatusBadRequest, "Gagal mendekripsi: Kunci salah atau ciphertext telah dimanipulasi")
		return
	}

	c.JSON(http.StatusOK, BaseResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Pesan:  "Pesan berhasil didekripsi",
		Data: []DecryptData{
			{
				Plaintext: plaintext,
			},
		},
	})
}
