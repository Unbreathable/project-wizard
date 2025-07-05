<script lang="ts">
	import { joinSession } from '$lib/connection.svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	let name = $state('');
	let error = $state('');
	let isLoading = $state(false);

	// Get lobby ID from URL params
	const lobbyId = $page.url.searchParams.get('lobby_id') || '';

	async function handleJoin() {
		if (!name.trim()) {
			error = 'Please enter your name!';
			return;
		}

		if (!lobbyId) {
			error = 'No lobby ID provided!';
			return;
		}

		isLoading = true;
		error = '';

		try {
			const result = await joinSession(lobbyId, name.trim());

			if (result) {
				error = result;
			} else {
				goto('/lobby/select');
			}
		} catch (err) {
			error = 'An unexpected error occurred';
		} finally {
			isLoading = false;
		}
	}

	function handleKeyPress(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			handleJoin();
		}
	}
</script>

<div class="flex items-center justify-center w-screen h-screen">
	<div
		class="flex flex-col p-6 gap-4 border-6 border-double border-bg-200 bg-bg-700 w-full max-w-md"
	>
		<h1 class="text-bg-100 text-center font-pixel text-xl mb-2">JOIN GAME</h1>
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
			onclick={handleJoin}
			disabled={isLoading}
			class="text-bg-100 border-2 border-bg-300 bg-bg-600 outline-none font-pixel w-full p-4 mb-1 transition-all duration-150 ease-out shadow-[0px_4px_0px_0px_var(--color-bg-300)] cursor-pointer hover:-translate-y-1 hover:shadow-[0px_8px_0px_0px_var(--color-bg-300)] active:translate-y-1 active:shadow-[0px_0px_0px_0px_var(--color-bg-300)] disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-y-0 disabled:hover:shadow-[0px_4px_0px_0px_var(--color-bg-300)]"
		>
			{isLoading ? '>> CONNECTING <<' : '>> CONNECT <<'}
		</button>
	</div>
</div>
