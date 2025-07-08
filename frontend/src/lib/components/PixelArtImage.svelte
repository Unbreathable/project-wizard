<script lang="ts">
	import type { Snippet } from 'svelte';

	let {
		url,
		class: className = '',
		children,
		...restProps
	}: {
		url: string;
		class?: string;
		children?: Snippet<[any]>;
		[key: string]: any;
	} = $props();
</script>

<div
	class={'bg-center bg-no-repeat bg-contain ' + className}
	style="background-image: url('/{url}'); image-rendering: pixelated; image-rendering: -moz-crisp-edges; image-rendering: crisp-edges;"
>
	{#if children}
		{@render children(restProps)}
	{/if}
</div>
