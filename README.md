# Dokumentasi Implementasi Kriptografi AES-256-GCM

Repositori ini berisi implementasi lengkap algoritma kriptografi simetris **AES-256-GCM** (*Advanced Encryption Standard - Galois/Counter Mode*) untuk enkripsi dan dekripsi pesan teks.

Repositori ini menyediakan dua arsitektur implementasi mandiri yang saling kompatibel secara kriptografis (*cross-interoperable*):
1. **Arsitektur Decoupled (Go + Svelte)**: Backend REST API berbasis Golang (Gin) dan Frontend antarmuka pengguna berbasis Svelte 5 + TypeScript.
2. **Arsitektur Monolithic Fullstack (Python Flask)**: Aplikasi mandiri berbasis Flask yang menyajikan REST API sekaligus antarmuka web UI bawaan pada satu port.

---

## 1. Spesifikasi Teknis Algoritma

| Komponen | Spesifikasi | Keterangan |
| :--- | :--- | :--- |
| **Bidang** | Kriptografi | Keamanan Informasi & Data |
| **Kategori** | Kriptografi Simetris | Kunci yang sama digunakan untuk enkripsi dan dekripsi |
| **Algoritma** | AES-256-GCM | *Authenticated Encryption with Associated Data* (AEAD) |
| **Panjang Kunci** | 256-bit (32 bytes) | Direpresentasikan dalam string 64 karakter hexadecimal |
| **Nonce / IV** | 12 bytes (96-bit) | Dihasilkan secara acak via CSPRNG setiap kali enkripsi dilakukan |
| **Authentication Tag** | 16 bytes (128-bit) | Memverifikasi integritas data dan mendeteksi manipulasi teks |
| **Format Payload** | Base64 | Pengemasan: `Base64(Nonce [12B] + Ciphertext + Auth Tag [16B])` |

---

## 2. Struktur Direktori Proyek

```text
sourCode/
├── backend/                       # Implementasi 1: REST API Golang
│   ├── cmd/ atau main.go          # Entry point server Gin (Port 8080)
│   ├── internal/
│   │   ├── crypto/aes_gcm.go      # Core logic AES-256-GCM
│   │   ├── handler/               # Controller HTTP request & response
│   │   └── middleware/cors.go     # Middleware CORS
│   ├── docs/api-spec.json         # Dokumentasi kontrak OpenAPI 3.0.3
│   ├── bruno/                     # Koleksi pengujian endpoint API Bruno
│   └── go.mod
│
├── frontend/                      # Implementasi 1: Frontend Svelte
│   ├── src/
│   │   ├── components/            # Komponen UI (KeySection, EncryptSection, DecryptSection, Modal)
│   │   ├── services/cryptoApi.ts  # Client fetch ke backend Go (:8080)
│   │   ├── types/crypto.ts        # Kontrak tipe TypeScript
│   │   ├── App.svelte             # Root component (Svelte 5 Runes)
│   │   └── app.css                # Tailwind CSS v4 & konfigurasi font Roboto
│   ├── package.json
│   └── vite.config.ts
│
├── flask/                         # Implementasi 2: Fullstack Python Flask
│   ├── app.py                     # Entry point server Flask & routing API (Port 5000)
│   ├── crypto_service.py          # Core logic AES-256-GCM
│   ├── templates/index.html       # Antarmuka web UI mandiri (Jinja2)
│   └── requirements.txt           # Dependensi pustaka Python
│
└── README.md                      # Dokumentasi teknis sistem
```

---

## 3. Prasyarat Sistem

- **Golang**: Versi 1.22 atau yang lebih baru.
- **Node.js**: Versi 18 atau yang lebih baru (disertai `npm`).
- **Python**: Versi 3.10 atau yang lebih baru.

---

## 4. Panduan Menjalankan Sistem

Anda dapat menjalankan salah satu dari dua implementasi berikut sesuai kebutuhan:

### Solusi 1: Menjalankan Pasangan Golang Backend + Svelte Frontend

#### Langkah 1: Jalankan Backend Golang (Port 8080)
1. Buka terminal dan arahkan ke direktori `backend`:
   ```bash
   cd backend
   ```
2. Jalankan server:
   ```bash
   go run main.go
   ```
   Server REST API akan berjalan pada `http://localhost:8080`.

#### Langkah 2: Jalankan Frontend Svelte (Port 5173)
1. Buka terminal baru dan arahkan ke direktori `frontend`:
   ```bash
   cd frontend
   ```
2. Pasang dependensi jika belum dilakukan:
   ```bash
   npm install
   ```
3. Jalankan server pengembangan Vite:
   ```bash
   npm run dev
   ```
4. Buka peramban (*browser*) pada alamat:
   ```text
   http://localhost:5173
   ```

---

### Solusi 2: Menjalankan Fullstack Python Flask (Port 5000)

Implementasi ini beroperasi 100% mandiri, tidak memerlukan Svelte, dan menyajikan REST API sekaligus UI web di port yang sama.

1. Buka terminal dan arahkan ke direktori `flask`:
   ```bash
   cd flask
   ```
2. Pasang pustaka dependensi:
   ```bash
   pip install -r requirements.txt
   ```
3. Jalankan server Flask:
   ```bash
   python app.py
   ```
4. Buka peramban (*browser*) pada alamat:
   ```text
   http://localhost:5000
   ```
   Antarmuka web interaktif langsung tampil dan siap digunakan.

---

## 5. Dokumentasi Kontrak REST API

Kedua implementasi backend (Go pada port 8080 dan Flask pada port 5000) mengimplementasikan kontrak REST API baku yang identik.

### Format Baku Respon JSON
```json
{
  "code": 200,
  "status": "OK",
  "pesan": "Deskripsi status operasi",
  "data": [
    { ... }
  ]
}
```
Jika terjadi kegagalan validasi atau kesalahan kriptografi, `code` bernilai `400` atau `500`, `status` bernilai `"error"`, dan `data` berupa array kosong `[]`.

---

### Rincian Endpoint

#### 1. Generate Secret Key
- **Endpoint**: `POST /api/crypto/generate-key`
- **Request Body**: *Kosong*
- **Response Data (200 OK)**:
  ```json
  {
    "code": 200,
    "status": "OK",
    "pesan": "Secret Key AES-256 berhasil dibuat",
    "data": [
      {
        "key": "af0bd0f628a5ec17f8787d3fd33eba63ea5cabdede4f7884f2c1c4866591b3da",
        "length_bytes": 32,
        "bit_length": 256
      }
    ]
  }
  ```

#### 2. Enkripsi Pesan Teks
- **Endpoint**: `POST /api/crypto/encrypt`
- **Request Body**:
  ```json
  {
    "plaintext": "Pesan rahasia tugas kriptografi semester 5",
    "key": "af0bd0f628a5ec17f8787d3fd33eba63ea5cabdede4f7884f2c1c4866591b3da"
  }
  ```
- **Response Data (200 OK)**:
  ```json
  {
    "code": 200,
    "status": "OK",
    "pesan": "Pesan berhasil dienkripsi",
    "data": [
      {
        "ciphertext": "urLqZ14kk11DSWKJMtkWhKmPA4lgbOUi9FIFqyHg3bLoeh7uM2D1vN3qTvh//WkxgyoFzxAtYi0Epa0SI1ZoWNboBwzKNNkjMqHS1A4CzCMV6KDUlMEs7FlyaAk=",
        "algorithm": "AES-256-GCM"
      }
    ]
  }
  ```

#### 3. Dekripsi Pesan Teks
- **Endpoint**: `POST /api/crypto/decrypt`
- **Request Body**:
  ```json
  {
    "ciphertext": "urLqZ14kk11DSWKJMtkWhKmPA4lgbOUi9FIFqyHg3bLoeh7uM2D1vN3qTvh//WkxgyoFzxAtYi0Epa0SI1ZoWNboBwzKNNkjMqHS1A4CzCMV6KDUlMEs7FlyaAk=",
    "key": "af0bd0f628a5ec17f8787d3fd33eba63ea5cabdede4f7884f2c1c4866591b3da"
  }
  ```
- **Response Data (200 OK)**:
  ```json
  {
    "code": 200,
    "status": "OK",
    "pesan": "Pesan berhasil didekripsi",
    "data": [
      {
        "plaintext": "Pesan rahasia tugas kriptografi semester 5"
      }
    ]
  }
  ```

---

## 6. Pengujian API Melalui Bruno

Koleksi berkas pengujian API tersimpan pada direktori `backend/bruno/`.

### Menjalankan Koleksi via Bruno CLI
Pastikan backend target sedang menyala, lalu jalankan perintah berikut dari direktori utama:
```bash
npx @usebruno/cli run backend/bruno
```

### Menjalankan via Aplikasi Desktop Bruno
1. Buka aplikasi Bruno.
2. Pilih menu **Open Collection**.
3. Arahkan ke direktori `backend/bruno`.
4. Seluruh request (`Generate Key`, `Encrypt`, `Decrypt`) siap diuji secara langsung.

---

## 7. Interoperabilitas Kriptografi (Cross-Platform)

Seluruh implementasi (Go, Python, dan JavaScript) mematuhi standar NIST SP 800-38D:
- Ciphertext yang dihasilkan oleh **Golang** dapat didekripsi secara langsung oleh **Flask** (dan sebaliknya) asalkan kunci 256-bit yang digunakan identik.
- Mekanisme Galois Counter Mode menjamin integritas: perubahan satu karakter atau manipulasi bit pada ciphertext akan menyebabkan verifikasi *Authentication Tag* gagal dan sistem secara otomatis menolak proses dekripsi.
