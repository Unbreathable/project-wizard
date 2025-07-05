<script lang="ts">
	import { useEvent } from '$lib/connection.svelte';
	import { onMount } from 'svelte';

	let error = $state('');
	let isLoading = $state(false);

	onMount(() => {
		useEvent('lobby_info', (event) => {
			console.log(event);
		});
	});
</script>

<div class="flex items-center justify-center w-screen h-screen">
	<div
		class="flex flex-col p-6 gap-4 border-6 border-double border-bg-200 bg-bg-700 w-full max-w-md"
	>
		<h1 class="text-bg-100 text-center font-pixel text-xl mb-2">CHARACTER SELECTION</h1>

		{#if error}
			<p class="font-pixel text-error px-2 text-sm" class:opacity-0={!error}>
				{error || 'No error'}
			</p>
		{/if}
		<button
			disabled={isLoading}
			class="text-bg-100 border-2 border-bg-300 bg-bg-600 outline-none font-pixel w-full p-4 mb-1 transition-all duration-150 ease-out shadow-[0px_4px_0px_0px_var(--color-bg-300)] cursor-pointer hover:-translate-y-1 hover:shadow-[0px_8px_0px_0px_var(--color-bg-300)] active:translate-y-1 active:shadow-[0px_0px_0px_0px_var(--color-bg-300)] disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-y-0 disabled:hover:shadow-[0px_4px_0px_0px_var(--color-bg-300)]"
		>
			{isLoading ? '>> SELECTING <<' : '>> SELECT <<'}
		</button>
	</div>
</div>
