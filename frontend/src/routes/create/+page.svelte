<script lang="ts">
	import { createSession } from '$lib/connection.svelte';
	import { goto } from '$app/navigation';

	let name = $state('');
	let error = $state('');
	let isLoading = $state(false);

	async function handleCreate() {
		if (!name.trim() || name.length < 2) {
			error = 'Please enter your name!';
			return;
		}

		isLoading = true;
		error = '';

		try {
			const result = await createSession(name.trim());

			if (result) {
				// Error occurred
				error = result;
			} else {
				// Success - redirect to game
				goto('/game');
			}
		} catch (err) {
			error = 'An unexpected error occurred';
		} finally {
			isLoading = false;
		}
	}

	function handleKeyPress(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			handleCreate();
		}
	}
</script>

<div class="flex items-center justify-center w-screen h-screen">
	<div
		class="flex flex-col p-6 gap-4 border-6 border-double border-bg-200 bg-bg-700 w-full max-w-md"
	>
		<h1 class="text-bg-100 text-center font-pixel text-xl mb-2">CREATE GAME</h1>
		<input
			type="text"
			maxlength="20"
			placeholder="Enter your name..."
			bind:value={name}
			onkeypress={handleKeyPress}
			disabled={isLoading}
			class="text-bg-100 border-2 border-bg-300 bg-bg-600 outline-none font-pixel w-full p-4 transition-all duration-150 ease-out focus:border-bg-100 disabled:opacity-50"
		/>
		{#if error}
			<p class="font-pixel text-error px-2 text-sm" class:opacity-0={!error}>
				{error || 'No error'}
			</p>
		{/if}
		<button
			onclick={handleCreate}
			disabled={isLoading}
			class="text-bg-100 border-2 border-bg-300 bg-bg-600 outline-none font-pixel w-full p-4 mb-1 transition-all duration-150 ease-out shadow-[0px_4px_0px_0px_var(--color-bg-300)] cursor-pointer hover:-translate-y-1 hover:shadow-[0px_8px_0px_0px_var(--color-bg-300)] active:translate-y-1 active:shadow-[0px_0px_0px_0px_var(--color-bg-300)] disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-y-0 disabled:hover:shadow-[0px_4px_0px_0px_var(--color-bg-300)]"
		>
			{isLoading ? '>> CREATING <<' : '>> CREATE <<'}
		</button>
	</div>
</div>
