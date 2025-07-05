<script lang="ts">
	import { goto } from '$app/navigation';
	import { currentLobby, useEvent } from '$lib/connection.svelte';
	import { characters, type Character } from '$lib/characters';
	import { onMount } from 'svelte';

	let error = $state('');
	let isLoading = $state(false);
	let selectedCharacters = $state<number[]>([]);

	onMount(() => {
		if (!currentLobby) {
			goto('/');
			return;
		}

		useEvent('lobby_change', (event) => {
			console.log(event);
		});
	});

	async function select(char: Character) {
		if (selectedCharacters.includes(char.id)) {
			selectedCharacters.splice(selectedCharacters.indexOf(char.id), 1);
			return;
		}
		if (selectedCharacters.length == 4) {
			return;
		}

		selectedCharacters.push(char.id);
	}

	async function copyUrl() {
		if (!currentLobby) return;

		const url = `${window.location.origin}/lobby/join?lobby_id=${currentLobby}`;
		try {
			await navigator.clipboard.writeText(url);
			// You could add a toast notification here if needed
		} catch (err) {
			console.error('Failed to copy URL:', err);
		}
	}
</script>

<div class="flex flex-col items-center justify-center w-screen h-screen gap-6">
	<div
		class="flex flex-col p-6 gap-6 border-6 border-double border-bg-200 bg-bg-700 w-full max-w-xl"
	>
		<h1 class="text-bg-100 text-center font-pixel text-xl">CHARACTER SELECTION</h1>

		<div class="grid grid-cols-4 gap-4 justify-items-center justify-center">
			{#each Object.values(characters) as character}
				<div class="relative">
					<button
						class="w-24 h-24 relative cursor-pointer"
						onclick={() => select(character)}
						aria-label="Select {character.name || `Character ${character.id}`}"
					>
						<div
							class="absolute inset-0 bg-center bg-no-repeat bg-contain transition-transform duration-200 ease-out {selectedCharacters.includes(
								character.id
							)
								? ''
								: 'hover:scale-110'}"
							style="background-image: url('/{selectedCharacters.includes(character.id)
								? 'character-frame.png'
								: 'character-frame-empty.png'}'); image-rendering: pixelated; image-rendering: -moz-crisp-edges; image-rendering: crisp-edges;"
						></div>
						<div
							class="absolute inset-2 bg-center bg-no-repeat bg-contain pointer-events-none"
							style="background-image: url('/{character.url}'); image-rendering: pixelated; image-rendering: -moz-crisp-edges; image-rendering: crisp-edges;"
						></div>
					</button>
					{#if character.name}
						<p class="text-bg-100 font-pixel text-xs text-center mt-1">{character.name}</p>
					{/if}
				</div>
			{/each}
		</div>

		{#if error}
			<p class="font-pixel text-error px-2 text-sm" class:opacity-0={!error}>
				{error || 'No error'}
			</p>
		{/if}

		<button
			disabled={isLoading || selectedCharacters.length != 4}
			class="text-bg-100 border-2 border-bg-300 bg-bg-600 outline-none font-pixel w-full p-4 mb-1 transition-all duration-150 ease-out shadow-[0px_4px_0px_0px_var(--color-bg-300)] cursor-pointer hover:-translate-y-1 hover:shadow-[0px_8px_0px_0px_var(--color-bg-300)] active:translate-y-1 active:shadow-[0px_0px_0px_0px_var(--color-bg-300)] disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-y-0 disabled:hover:shadow-[0px_4px_0px_0px_var(--color-bg-300)]"
		>
			{isLoading ? '>> SELECTING <<' : `>> SELECT ${selectedCharacters.length}/4 <<`}
		</button>
	</div>

	<div
		class="flex flex-col p-6 gap-6 border-6 border-double border-bg-200 bg-bg-700 w-full max-w-xl"
	>
		<h1 class="text-bg-100 text-center font-pixel text-xl">LOBBY</h1>

		<p class="text-bg-100 font-pixel">
			No other player found. Copy the URL below to invite your friend.
		</p>

		<button
			onclick={() => copyUrl()}
			class="text-bg-100 border-2 border-bg-300 bg-bg-600 outline-none font-pixel w-full p-4 mb-1 transition-all duration-150 ease-out shadow-[0px_4px_0px_0px_var(--color-bg-300)] cursor-pointer hover:-translate-y-1 hover:shadow-[0px_8px_0px_0px_var(--color-bg-300)] active:translate-y-1 active:shadow-[0px_0px_0px_0px_var(--color-bg-300)] disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-y-0 disabled:hover:shadow-[0px_4px_0px_0px_var(--color-bg-300)]"
		>
			{'>> COPY URL <<'}
		</button>
	</div>
</div>
