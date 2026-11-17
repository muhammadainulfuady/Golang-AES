// Format baku response dari backend Golang sesuai docs/api-spec.json
export interface ApiResponse<T> {
  code: number;
  status: 'OK' | 'error';
  pesan: string;
  data: T[];
}

// Data hasil pembuatan Secret Key
export interface KeyData {
  key: string;
  length_bytes: number;
  bit_length: number;
}

// Request & Response Enkripsi
export interface EncryptRequest {
  plaintext: string;
  key: string;
}

export interface EncryptData {
  ciphertext: string;
  algorithm: string;
}

// Request & Response Dekripsi
export interface DecryptRequest {
  ciphertext: string;
  key: string;
}

export interface DecryptData {
  plaintext: string;
}

// Tipe untuk status notifikasi UI (sukses / gagal)
export interface AlertNotification {
  type: 'success' | 'error' | 'info';
  message: string;
}
