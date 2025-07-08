<script lang="ts">
	import type { Snippet } from 'svelte';

	let {
		class: className = '',
		onClick,
		disabled = false,
		children
	}: {
		onClick: (
			e: MouseEvent & {
				currentTarget: EventTarget & HTMLButtonElement;
			}
		) => void;
		disabled?: boolean;
		children?: Snippet<[]>;
		class?: string;
	} = $props();
</script>

<button
	class={'text-bg-100 border-2 border-bg-300 bg-bg-600 outline-none font-pixel w-full p-4 mb-1 transition-all duration-150 ease-out shadow-[0px_4px_0px_0px_var(--color-bg-300)] cursor-pointer hover:-translate-y-1 hover:shadow-[0px_8px_0px_0px_var(--color-bg-300)] active:translate-y-1 active:shadow-[0px_0px_0px_0px_var(--color-bg-300)] disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-y-0 disabled:hover:shadow-[0px_4px_0px_0px_var(--color-bg-300)] ' +
		className}
	{disabled}
	onclick={(e) => onClick(e)}
>
	{#if children}
		{@render children()}
	{/if}
</button>
