<script lang="ts">
  import { X, CheckCircle, AlertCircle, Info } from '@lucide/svelte';

  interface Props {
    isOpen: boolean;
    title: string;
    message: string;
    type?: 'success' | 'error' | 'info';
    details?: string;
    onClose: () => void;
  }

  let { isOpen, title, message, type = 'info', details = '', onClose }: Props = $props();

  function handleKeydown(e: KeyboardEvent) {
    if (e.key === 'Escape' && isOpen) {
      onClose();
    }
  }
</script>

<svelte:window onkeydown={handleKeydown} />

{#if isOpen}
  <!-- Backdrop -->
  <div 
    class="fixed inset-0 z-50 bg-slate-950/75 backdrop-blur-xs flex items-center justify-center p-4"
    role="presentation"
  >
    <!-- Modal Dialog Container -->
    <div 
      class="bg-slate-900 border border-slate-800 rounded-xl shadow-2xl max-w-md w-full overflow-hidden text-slate-100 animate-in fade-in zoom-in-95 duration-150"
      role="dialog"
      aria-modal="true"
      aria-labelledby="modal-title"
      tabindex="-1"
    >
      <!-- Header -->
      <div class="flex items-center justify-between px-5 py-4 border-b border-slate-800">
        <div class="flex items-center gap-2.5">
          {#if type === 'success'}
            <CheckCircle class="w-5 h-5 text-emerald-400" />
          {:else if type === 'error'}
            <AlertCircle class="w-5 h-5 text-rose-400" />
          {:else}
            <Info class="w-5 h-5 text-sky-400" />
          {/if}
          <h3 id="modal-title" class="font-semibold text-slate-100 text-base">{title}</h3>
        </div>
        <button 
          type="button"
          onclick={onClose}
          class="p-1 rounded-md text-slate-400 hover:text-slate-100 hover:bg-slate-800 transition cursor-pointer"
          aria-label="Tutup modal"
        >
          <X class="w-4 h-4" />
        </button>
      </div>

      <!-- Content -->
      <div class="p-5 space-y-3">
        <p class="text-sm text-slate-300 leading-relaxed">{message}</p>
        
        {#if details}
          <div class="bg-slate-950 border border-slate-800/80 rounded-lg p-3 max-h-48 overflow-y-auto">
            <pre class="text-xs font-mono text-slate-300 break-all whitespace-pre-wrap">{details}</pre>
          </div>
        {/if}
      </div>

      <!-- Footer -->
      <div class="px-5 py-3.5 bg-slate-950/50 border-t border-slate-800 flex justify-end">
        <button
          type="button"
          onclick={onClose}
          class="px-4 py-2 text-xs font-medium rounded-lg bg-slate-800 hover:bg-slate-700 text-slate-200 transition cursor-pointer border border-slate-700"
        >
          Tutup
        </button>
      </div>
    </div>
  </div>
{/if}
