import os
import base64
from cryptography.hazmat.primitives.ciphers.aead import AESGCM
from cryptography.exceptions import InvalidTag


class CryptoService:
    @staticmethod
    def generate_key() -> str:
        """
        Menghasilkan kunci acak AES-256 (32 bytes = 256 bits).
        Format pengembalian: string hexadecimal 64 karakter.
        """
        key_bytes = AESGCM.generate_key(bit_length=256)
        return key_bytes.hex()

    @staticmethod
    def encrypt(plaintext: str, key_hex: str) -> str:
        """
        Mengenkripsi pesan teks menjadi Base64 menggunakan AES-256-GCM.
        Format payload: Nonce (12B) + Ciphertext + AuthTag (16B).
        Format ini kompatibel 100% dengan backend Go dan Node.js.
        """
        key_bytes = bytes.fromhex(key_hex)
        aesgcm = AESGCM(key_bytes)

        # Nonce 12-byte standar aman AES-GCM
        nonce = os.urandom(12)

        plaintext_bytes = plaintext.encode("utf-8")
        # aesgcm.encrypt mengembalikan: ciphertext + 16-byte authentication tag
        ciphertext = aesgcm.encrypt(nonce, plaintext_bytes, None)

        # Gabungkan Nonce (12 bytes) di depan ciphertext
        encrypted_data = nonce + ciphertext
        return base64.b64encode(encrypted_data).decode("utf-8")

    @staticmethod
    def decrypt(ciphertext_b64: str, key_hex: str) -> str:
        """
        Mendekripsi ciphertext Base64 kembali ke teks asli.
        Memverifikasi otentikasi data (Auth Tag).
        """
        key_bytes = bytes.fromhex(key_hex)
        aesgcm = AESGCM(key_bytes)

        encrypted_data = base64.b64decode(ciphertext_b64)

        if len(encrypted_data) < 28:
            raise ValueError("Ciphertext terlalu pendek atau rusak.")

        # Pisahkan 12 byte pertama (nonce) dan sisanya (ciphertext + tag)
        nonce = encrypted_data[:12]
        ciphertext = encrypted_data[12:]

        decrypted_bytes = aesgcm.decrypt(nonce, ciphertext, None)
        return decrypted_bytes.decode("utf-8")
