<script lang="ts">
  import { Key, RefreshCw, Copy, Check } from '@lucide/svelte';

  interface Props {
    secretKey: string;
    isGenerating: boolean;
    onGenerate: () => void;
  }

  let { secretKey = $bindable(''), isGenerating, onGenerate }: Props = $props();

  let copied = $state(false);

  async function copyToClipboard() {
    if (!secretKey) return;
    await navigator.clipboard.writeText(secretKey);
    copied = true;
    setTimeout(() => {
      copied = false;
    }, 1500);
  }

  let keyStatus = $derived(() => {
    if (!secretKey) return { label: 'Kosong', color: 'text-slate-400 bg-slate-900 border-slate-800' };
    const isValid = secretKey.length === 64 && /^[0-9a-fA-F]+$/.test(secretKey);
    if (isValid) {
      return { label: '64 Hex Valid', color: 'text-emerald-400 bg-emerald-950/40 border-emerald-800/60' };
    }
    return { label: `${secretKey.length}/64 Hex`, color: 'text-amber-400 bg-amber-950/40 border-amber-800/60' };
  });
</script>

<div class="bg-slate-900/80 border border-slate-800/80 rounded-xl p-4 space-y-3">
  <div class="flex items-center justify-between gap-3">
    <div class="flex items-center gap-2">
      <Key class="w-4 h-4 text-indigo-400" />
      <span class="text-xs font-semibold text-slate-200">Secret Key</span>
    </div>
    <span class={`px-2 py-0.5 text-[11px] font-mono rounded border ${keyStatus().color}`}>
      {keyStatus().label}
    </span>
  </div>

  <div class="flex flex-col sm:flex-row gap-2">
    <input
      type="text"
      bind:value={secretKey}
      placeholder="64 karakter hex..."
      class="flex-1 bg-slate-950 border border-slate-800 rounded-lg px-3 py-2 text-xs font-mono text-slate-100 placeholder-slate-600 focus:outline-none focus:border-indigo-500 transition"
    />

    <div class="flex gap-2">
      <button
        type="button"
        onclick={copyToClipboard}
        disabled={!secretKey}
        class="inline-flex items-center gap-1.5 px-3 py-2 text-xs font-medium rounded-lg bg-slate-800 hover:bg-slate-700 disabled:opacity-40 disabled:cursor-not-allowed text-slate-200 border border-slate-700/80 transition cursor-pointer"
      >
        {#if copied}
          <Check class="w-3.5 h-3.5 text-emerald-400" />
          <span>Tersalin</span>
        {:else}
          <Copy class="w-3.5 h-3.5" />
          <span>Salin</span>
        {/if}
      </button>

      <button
        type="button"
        onclick={onGenerate}
        disabled={isGenerating}
        class="inline-flex items-center gap-1.5 px-3.5 py-2 text-xs font-medium rounded-lg bg-indigo-600 hover:bg-indigo-500 disabled:opacity-40 text-white transition cursor-pointer"
      >
        <RefreshCw class={`w-3.5 h-3.5 ${isGenerating ? 'animate-spin' : ''}`} />
        <span>Generate</span>
      </button>
    </div>
  </div>
</div>
