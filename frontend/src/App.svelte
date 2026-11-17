<script lang="ts">
  import Header from './components/Header.svelte';
  import KeySection from './components/KeySection.svelte';
  import EncryptSection from './components/EncryptSection.svelte';
  import DecryptSection from './components/DecryptSection.svelte';
  import Modal from './components/Modal.svelte';
  import { cryptoApi } from './services/cryptoApi';

  // State
  let secretKey = $state('');
  let isGenerating = $state(false);
  let isEncrypting = $state(false);
  let isDecrypting = $state(false);

  let ciphertextResult = $state('');
  let ciphertextForDecrypt = $state('');
  let plaintextResult = $state('');

  // Modal Error Box
  let isModalOpen = $state(false);
  let modalTitle = $state('');
  let modalMessage = $state('');
  let modalDetails = $state('');

  function showError(title: string, message: string, details: string = '') {
    modalTitle = title;
    modalMessage = message;
    modalDetails = details;
    isModalOpen = true;
  }

  function validateKey(): boolean {
    const trimmed = secretKey.trim();
    if (!trimmed) {
      showError('Kunci Kosong', 'Masukkan atau generate secret key terlebih dahulu.');
      return false;
    }
    if (trimmed.length !== 64 || !/^[0-9a-fA-F]+$/.test(trimmed)) {
      showError('Kunci Tidak Valid', 'Kunci harus berupa 64 karakter hexadecimal.', `Panjang saat ini: ${trimmed.length}`);
      return false;
    }
    return true;
  }

  async function handleGenerateKey() {
    isGenerating = true;
    try {
      const data = await cryptoApi.generateKey();
      secretKey = data.key;
    } catch (err: any) {
      showError('Gagal Membuat Kunci', err.message || 'Terjadi kesalahan.');
    } finally {
      isGenerating = false;
    }
  }

  async function handleEncrypt(plaintext: string) {
    if (!validateKey()) return;

    isEncrypting = true;
    try {
      const data = await cryptoApi.encrypt({
        plaintext,
        key: secretKey.trim(),
      });
      ciphertextResult = data.ciphertext;
    } catch (err: any) {
      showError('Gagal Enkripsi', err.message || 'Terjadi kesalahan.');
    } finally {
      isEncrypting = false;
    }
  }

  async function handleDecrypt(ciphertext: string) {
    if (!validateKey()) return;

    isDecrypting = true;
    try {
      const data = await cryptoApi.decrypt({
        ciphertext: ciphertext.trim(),
        key: secretKey.trim(),
      });
      plaintextResult = data.plaintext;
    } catch (err: any) {
      showError('Gagal Dekripsi', 'Kunci salah atau ciphertext telah dimanipulasi.');
    } finally {
      isDecrypting = false;
    }
  }

  function handleSendToDecrypt(ciphertext: string) {
    ciphertextForDecrypt = ciphertext;
  }
</script>

<main class="min-h-screen bg-slate-950 text-slate-100 flex flex-col justify-between selection:bg-indigo-500/30 selection:text-indigo-200">
  <div class="max-w-4xl w-full mx-auto p-4 sm:p-6 space-y-5">
    <Header />

    <KeySection 
      bind:secretKey 
      {isGenerating} 
      onGenerate={handleGenerateKey} 
    />

    <div class="grid grid-cols-1 md:grid-cols-2 gap-5 items-stretch">
      <EncryptSection
        isLoading={isEncrypting}
        {ciphertextResult}
        onEncrypt={handleEncrypt}
        onSendToDecrypt={handleSendToDecrypt}
      />

      <DecryptSection
        isLoading={isDecrypting}
        bind:ciphertextInput={ciphertextForDecrypt}
        {plaintextResult}
        onDecrypt={handleDecrypt}
      />
    </div>

    <footer class="pt-4 border-t border-slate-900 text-center text-xs text-slate-600">
      <p>AES-256-GCM &bull; Symmetric Cryptography</p>
    </footer>
  </div>

  <Modal
    isOpen={isModalOpen}
    title={modalTitle}
    message={modalMessage}
    type="error"
    details={modalDetails}
    onClose={() => (isModalOpen = false)}
  />
</main>
