<script lang="ts">
  import { Shield } from '@lucide/svelte';
</script>

<header class="flex items-center justify-between pb-5 border-b border-slate-800/80">
  <div class="flex items-center gap-3">
    <div class="p-2 rounded-lg bg-slate-900 border border-slate-800 text-indigo-400">
      <Shield class="w-5 h-5" />
    </div>
    <div>
      <h1 class="text-lg font-bold text-slate-100 tracking-tight">AES-256-GCM</h1>
      <p class="text-xs text-slate-400">Enkripsi & Dekripsi Pesan Teks</p>
    </div>
  </div>
  <span class="text-xs font-mono px-2.5 py-1 rounded bg-slate-900 border border-slate-800 text-slate-400">
    256-bit
  </span>
</header>
