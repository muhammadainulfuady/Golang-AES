import type { ApiResponse, KeyData, EncryptRequest, EncryptData, DecryptRequest, DecryptData } from '../types/crypto';

const BASE_URL = 'http://localhost:8080/api/crypto';

class CryptoApiService {
  /**
   * Helper private untuk menangani fetch dan error handling standar
   */
  private async request<T>(endpoint: string, options?: RequestInit): Promise<ApiResponse<T>> {
    try {
      const response = await fetch(`${BASE_URL}${endpoint}`, {
        headers: {
          'Content-Type': 'application/json',
          ...(options?.headers || {}),
        },
        ...options,
      });

      const json: ApiResponse<T> = await response.json();

      if (!response.ok || json.status === 'error') {
        throw new Error(json.pesan || `Terjadi kesalahan (HTTP ${response.status})`);
      }

      return json;
    } catch (err: any) {
      if (err.name === 'TypeError' && err.message.includes('fetch')) {
        throw new Error('Gagal terhubung ke backend Go. Pastikan server di port 8080 aktif.');
      }
      throw err;
    }
  }

  /**
   * POST /api/crypto/generate-key
   */
  async generateKey(): Promise<KeyData> {
    const res = await this.request<KeyData>('/generate-key', {
      method: 'POST',
    });
    if (!res.data || res.data.length === 0) {
      throw new Error('Data kunci tidak ditemukan dalam respon server.');
    }
    return res.data[0];
  }

  /**
   * POST /api/crypto/encrypt
   */
  async encrypt(payload: EncryptRequest): Promise<EncryptData> {
    const res = await this.request<EncryptData>('/encrypt', {
      method: 'POST',
      body: JSON.stringify(payload),
    });
    if (!res.data || res.data.length === 0) {
      throw new Error('Data ciphertext tidak ditemukan dalam respon server.');
    }
    return res.data[0];
  }

  /**
   * POST /api/crypto/decrypt
   */
  async decrypt(payload: DecryptRequest): Promise<DecryptData> {
    const res = await this.request<DecryptData>('/decrypt', {
      method: 'POST',
      body: JSON.stringify(payload),
    });
    if (!res.data || res.data.length === 0) {
      throw new Error('Data plaintext tidak ditemukan dalam respon server.');
    }
    return res.data[0];
  }
}

export const cryptoApi = new CryptoApiService();
