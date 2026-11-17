import os

from crypto_service import CryptoService
from cryptography.exceptions import InvalidTag

from flask import Flask, jsonify, make_response, request, render_template

app = Flask(__name__)


# ==============================================================================
# ROUTE HALAMAN UTAMA (FULLSTACK WEB UI)
# ==============================================================================
@app.route("/", methods=["GET"])
def index():
    return render_template("index.html")


# ==============================================================================
# MIDDLEWARE CORS
# ==============================================================================
@app.before_request
def handle_preflight():
    if request.method == "OPTIONS":
        res = make_response()
        res.headers["Access-Control-Allow-Origin"] = "*"
        res.headers["Access-Control-Allow-Methods"] = "GET, POST, OPTIONS, PUT, DELETE"
        res.headers["Access-Control-Allow-Headers"] = "Content-Type, Authorization"
        return res, 204


@app.after_request
def add_cors_headers(response):
    response.headers["Access-Control-Allow-Origin"] = "*"
    response.headers["Access-Control-Allow-Methods"] = "GET, POST, OPTIONS, PUT, DELETE"
    response.headers["Access-Control-Allow-Headers"] = "Content-Type, Authorization"
    return response


# ==============================================================================
# HELPER RESPONSE STANDAR
# ==============================================================================
def standard_response(code: int, status: str, pesan: str, data: list):
    return jsonify({"code": code, "status": status, "pesan": pesan, "data": data}), code


def error_response(code: int, pesan: str):
    return standard_response(code, "error", pesan, [])


# ==============================================================================
# ENDPOINTS REST API
# ==============================================================================
@app.route("/api/crypto/generate-key", methods=["POST"])
def generate_key_route():
    try:
        key_hex = CryptoService.generate_key()
        return standard_response(
            code=200,
            status="OK",
            pesan="Secret Key AES-256 berhasil dibuat",
            data=[{"key": key_hex, "length_bytes": 32, "bit_length": 256}],
        )
    except Exception as e:
        return error_response(500, f"Gagal menghasilkan kunci: {str(e)}")


@app.route("/api/crypto/encrypt", methods=["POST"])
def encrypt_route():
    body = request.get_json(silent=True)
    if not body:
        return error_response(400, "Format JSON tidak valid")

    plaintext = body.get("plaintext", "")
    key = str(body.get("key", "")).strip()

    if not plaintext:
        return error_response(400, "Pesan teks (plaintext) tidak boleh kosong")

    if len(key) != 64:
        return error_response(
            400, "Kunci harus berupa 64 karakter hexadecimal (256-bit)"
        )

    try:
        ciphertext = CryptoService.encrypt(plaintext, key)
        return standard_response(
            code=200,
            status="OK",
            pesan="Pesan berhasil dienkripsi",
            data=[{"ciphertext": ciphertext, "algorithm": "AES-256-GCM"}],
        )
    except ValueError:
        return error_response(400, "Format kunci hexadecimal tidak valid")
    except Exception as e:
        return error_response(500, f"Terjadi kesalahan saat enkripsi: {str(e)}")


@app.route("/api/crypto/decrypt", methods=["POST"])
def decrypt_route():
    body = request.get_json(silent=True)
    if not body:
        return error_response(400, "Format JSON tidak valid")

    ciphertext = str(body.get("ciphertext", "")).strip()
    key = str(body.get("key", "")).strip()

    if not ciphertext:
        return error_response(400, "Ciphertext tidak boleh kosong")

    if len(key) != 64:
        return error_response(
            400, "Kunci harus berupa 64 karakter hexadecimal (256-bit)"
        )

    try:
        plaintext = CryptoService.decrypt(ciphertext, key)
        return standard_response(
            code=200,
            status="OK",
            pesan="Pesan berhasil didekripsi",
            data=[{"plaintext": plaintext}],
        )
    except InvalidTag:
        return error_response(
            400, "Gagal mendekripsi: Kunci salah atau ciphertext telah dimanipulasi"
        )
    except ValueError:
        return error_response(400, "Format kunci atau ciphertext tidak valid")
    except Exception as e:
        return error_response(500, f"Terjadi kesalahan saat dekripsi: {str(e)}")


# ==============================================================================
# ENTRY POINT
# ==============================================================================
if __name__ == "__main__":
    port = int(os.environ.get("PORT", 5000))
    print(f"[*] Flask backend berjalan di http://localhost:{port}")
    app.run(host="0.0.0.0", port=port, debug=True)
