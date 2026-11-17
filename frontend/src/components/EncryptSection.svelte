<script lang="ts">
  import { Lock, ArrowRight, Copy, Check } from '@lucide/svelte';

  interface Props {
    isLoading: boolean;
    ciphertextResult: string;
    onEncrypt: (plaintext: string) => void;
    onSendToDecrypt?: (ciphertext: string) => void;
  }

  let { isLoading, ciphertextResult, onEncrypt, onSendToDecrypt }: Props = $props();

  let plaintext = $state('');
  let copied = $state(false);

  function handleSubmit(e: SubmitEvent) {
    e.preventDefault();
    if (!plaintext.trim()) return;
    onEncrypt(plaintext);
  }

  async function copyCiphertext() {
    if (!ciphertextResult) return;
    await navigator.clipboard.writeText(ciphertextResult);
    copied = true;
    setTimeout(() => {
      copied = false;
    }, 1500);
  }
</script>

<div class="bg-slate-900/80 border border-slate-800/80 rounded-xl p-4 flex flex-col justify-between space-y-4">
  <div class="space-y-3">
    <div class="flex items-center gap-2 pb-2 border-b border-slate-800">
      <Lock class="w-4 h-4 text-emerald-400" />
      <h2 class="text-xs font-semibold text-slate-200">Enkripsi</h2>
    </div>

    <form onsubmit={handleSubmit} class="space-y-2.5">
      <textarea
        bind:value={plaintext}
        rows="4"
        placeholder="Tulis pesan..."
        class="w-full bg-slate-950 border border-slate-800 rounded-lg p-3 text-xs text-slate-100 placeholder-slate-600 focus:outline-none focus:border-emerald-500 transition resize-none"
      ></textarea>

      <button
        type="submit"
        disabled={isLoading || !plaintext.trim()}
        class="w-full inline-flex items-center justify-center gap-1.5 px-3 py-2 text-xs font-medium rounded-lg bg-emerald-600 hover:bg-emerald-500 disabled:opacity-40 disabled:cursor-not-allowed text-white transition cursor-pointer"
      >
        <Lock class="w-3.5 h-3.5" />
        <span>{isLoading ? 'Memproses...' : 'Enkripsi'}</span>
      </button>
    </form>
  </div>

  <div class="space-y-1.5 pt-2 border-t border-slate-800">
    <div class="flex items-center justify-between text-xs">
      <span class="text-slate-400 text-[11px]">Ciphertext (Base64)</span>
      {#if ciphertextResult}
        <div class="flex items-center gap-2">
          {#if onSendToDecrypt}
            <button
              type="button"
              onclick={() => onSendToDecrypt(ciphertextResult)}
              class="inline-flex items-center gap-1 text-[11px] text-indigo-400 hover:text-indigo-300 transition cursor-pointer"
            >
              <span>Ke Dekripsi</span>
              <ArrowRight class="w-3 h-3" />
            </button>
          {/if}
          <button
            type="button"
            onclick={copyCiphertext}
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
        </div>
      {/if}
    </div>

    <div class="min-h-20 bg-slate-950 border border-slate-800 rounded-lg p-2.5 flex items-center justify-center">
      {#if ciphertextResult}
        <p class="text-xs font-mono text-emerald-300 break-all w-full select-all">
          {ciphertextResult}
        </p>
      {:else}
        <span class="text-[11px] text-slate-600 italic">Hasil enkripsi akan muncul di sini</span>
      {/if}
    </div>
  </div>
</div>
