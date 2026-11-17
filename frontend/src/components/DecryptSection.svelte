<script lang="ts">
  import { Unlock, Copy, Check } from '@lucide/svelte';

  interface Props {
    isLoading: boolean;
    ciphertextInput: string;
    plaintextResult: string;
    onDecrypt: (ciphertext: string) => void;
  }

  let { isLoading, ciphertextInput = $bindable(''), plaintextResult, onDecrypt }: Props = $props();

  let copied = $state(false);

  function handleSubmit(e: SubmitEvent) {
    e.preventDefault();
    if (!ciphertextInput.trim()) return;
    onDecrypt(ciphertextInput);
  }

  async function copyPlaintext() {
    if (!plaintextResult) return;
    await navigator.clipboard.writeText(plaintextResult);
    copied = true;
    setTimeout(() => {
      copied = false;
    }, 1500);
  }
</script>

<div class="bg-slate-900/80 border border-slate-800/80 rounded-xl p-4 flex flex-col justify-between space-y-4">
  <div class="space-y-3">
    <div class="flex items-center gap-2 pb-2 border-b border-slate-800">
      <Unlock class="w-4 h-4 text-sky-400" />
      <h2 class="text-xs font-semibold text-slate-200">Dekripsi</h2>
    </div>

    <form onsubmit={handleSubmit} class="space-y-2.5">
      <textarea
        bind:value={ciphertextInput}
        rows="4"
        placeholder="Tempel ciphertext Base64..."
        class="w-full bg-slate-950 border border-slate-800 rounded-lg p-3 text-xs font-mono text-slate-100 placeholder-slate-600 focus:outline-none focus:border-sky-500 transition resize-none"
      ></textarea>

      <button
        type="submit"
        disabled={isLoading || !ciphertextInput.trim()}
        class="w-full inline-flex items-center justify-center gap-1.5 px-3 py-2 text-xs font-medium rounded-lg bg-sky-600 hover:bg-sky-500 disabled:opacity-40 disabled:cursor-not-allowed text-white transition cursor-pointer"
      >
        <Unlock class="w-3.5 h-3.5" />
        <span>{isLoading ? 'Memproses...' : 'Dekripsi'}</span>
      </button>
    </form>
  </div>

  <div class="space-y-1.5 pt-2 border-t border-slate-800">
    <div class="flex items-center justify-between text-xs">
      <span class="text-slate-400 text-[11px]">Hasil Plaintext</span>
      {#if plaintextResult}
        <button
          type="button"
          onclick={copyPlaintext}
          class="inline-flex items-center gap-1 text-[11px] text-slate-300 hover:text-white transition cursor-pointer"
        >
          {#if copied}
            <Check class="w-3 h-3 text-emerald-400" />
            <span>Tersalin</span>
          {:else}
            <Copy class="w-3 h-3" />
            <span>Salin</span>
          {/if}
        </button>
      {/if}
    </div>

    <div class="min-h-20 bg-slate-950 border border-slate-800 rounded-lg p-2.5 flex items-center justify-center">
      {#if plaintextResult}
        <p class="text-xs text-sky-200 break-words w-full select-all font-sans">
          {plaintextResult}
        </p>
      {:else}
        <span class="text-[11px] text-slate-600 italic">Hasil dekripsi akan muncul di sini</span>
      {/if}
    </div>
  </div>
</div>
