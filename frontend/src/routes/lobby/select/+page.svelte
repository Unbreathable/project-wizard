<script lang="ts">
	import { goto } from '$app/navigation';
	import {
		currentLobby,
		currentOpponent,
		currentPlayer,
		currentToken,
		setOpponent,
		useEvent
	} from '$lib/connection.svelte';
	import { characters, initializeCharacters, type Character } from '$lib/characters';
	import { onMount } from 'svelte';
	import { postRequestURL } from '$lib/requests';

	let error = $state('');
	let first = true;
	let isReady = $state(false);
	let showCharacters = $state(false);
	let isOpponentThere = $state(false);
	let isOpponentReady = $state(false);
	let isLoading = $state(true);
	let selectedCharacters = $state<number[]>([]);
	let maxCharacterAmount = $state(4);
	let hoveredCharacter = $state<Character | null>(null);
	let tooltipPosition = $state({ x: 0, y: 0 });

	onMount(() => {
		if (!currentLobby) {
			goto('/');
			return;
		}
		loadCharacters();

		useEvent('lobby_change', (event) => {
			maxCharacterAmount = event.data.character_amount;
			if (first) {
				isLoading = false;
				first = false;
			}

			if (event.data.player_1.player_id !== currentPlayer && event.data.player_1.name != '') {
				setOpponent(event.data.player_1);
				isOpponentThere = true;
				isOpponentReady = event.data.player_1.ready;
			} else if (
				event.data.player_2.player_id !== currentPlayer &&
				event.data.player_2.name != ''
			) {
				setOpponent(event.data.player_2);
				isOpponentThere = true;
				isOpponentReady = event.data.player_2.ready;
			}
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

	async function loadCharacters() {
		await initializeCharacters();
		showCharacters = true;
	}

	async function handleButton() {
		console.log(selectedCharacters);
		if (isLoading) {
			return;
		}
		if (isReady) {
			unready();
		} else {
			ready();
		}
	}

	async function ready() {
		isLoading = true;
		let res = await postRequestURL('/lobby/ready', {
			lobby_id: currentLobby,
			player_id: currentPlayer,
			token: currentToken,
			character_ids: selectedCharacters
		});
		isLoading = false;
		if (!res.success) {
			error = res.message;
			return;
		}
		isReady = true;
	}

	async function unready() {
		isLoading = true;
		let res = await postRequestURL('/lobby/unready', {
			lobby_id: currentLobby,
			player_id: currentPlayer,
			token: currentToken
		});
		isLoading = false;
		if (!res.success) {
			error = res.message;
			return;
		}
		isReady = false;
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

	function handleMouseEnter(character: Character, event: MouseEvent) {
		hoveredCharacter = character;
		const rect = (event.target as HTMLElement).getBoundingClientRect();
		tooltipPosition = {
			x: rect.left + rect.width / 2,
			y: rect.top - 10
		};
	}

	function handleMouseLeave() {
		hoveredCharacter = null;
	}

	function getElementColor(element: string): string {
		switch (element) {
			case 'fire':
				return '#ff6b35';
			case 'water':
				return '#4dabf7';
			case 'earth':
				return '#8bc34a';
			case 'air':
				return '#e0e0e0';
			case 'dark':
				return '#6c5ce7';
			case 'light':
				return '#fdcb6e';
			default:
				return '#95a5a6';
		}
	}
</script>

<div class="flex flex-col items-center justify-center w-screen h-screen gap-6">
	<div
		class="flex flex-col p-6 gap-6 border-6 border-double border-bg-200 bg-bg-700 w-full max-w-xl"
	>
		<h1 class="text-bg-100 text-center font-pixel text-xl">CHARACTER SELECTION</h1>

		<div class="grid grid-cols-4 gap-4 justify-items-center justify-center">
			{#if showCharacters}
				{#each Object.values(characters) as character}
					<div class="relative">
						<button
							class="w-24 h-24 relative cursor-pointer"
							onclick={() => select(character)}
							onmouseenter={(e) => handleMouseEnter(character, e)}
							onmouseleave={handleMouseLeave}
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
			{/if}
		</div>

		{#if error}
			<p class="font-pixel text-error px-2 text-sm" class:opacity-0={!error}>
				{error || 'No error'}
			</p>
		{/if}

		<button
			onclick={handleButton}
			disabled={isLoading || selectedCharacters.length != maxCharacterAmount}
			class="text-bg-100 border-2 border-bg-300 bg-bg-600 outline-none font-pixel w-full p-4 mb-1 transition-all duration-150 ease-out shadow-[0px_4px_0px_0px_var(--color-bg-300)] cursor-pointer hover:-translate-y-1 hover:shadow-[0px_8px_0px_0px_var(--color-bg-300)] active:translate-y-1 active:shadow-[0px_0px_0px_0px_var(--color-bg-300)] disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-y-0 disabled:hover:shadow-[0px_4px_0px_0px_var(--color-bg-300)]"
		>
			{#if isLoading}
				LOADING...
			{:else}
				{isReady
					? '>> UNREADY <<'
					: `>> READY ${selectedCharacters.length}/${maxCharacterAmount} <<`}
			{/if}
		</button>
	</div>

	<div
		class="flex flex-col p-6 gap-6 border-6 border-double border-bg-200 bg-bg-700 w-full max-w-xl"
	>
		<h1 class="text-bg-100 text-center font-pixel text-xl">LOBBY</h1>

		{#if !isOpponentThere}
			<p class="text-bg-100 font-pixel">
				No other player found. Copy the URL below to invite your friend.
			</p>

			<button
				onclick={() => copyUrl()}
				class="text-bg-100 border-2 border-bg-300 bg-bg-600 outline-none font-pixel w-full p-4 mb-1 transition-all duration-150 ease-out shadow-[0px_4px_0px_0px_var(--color-bg-300)] cursor-pointer hover:-translate-y-1 hover:shadow-[0px_8px_0px_0px_var(--color-bg-300)] active:translate-y-1 active:shadow-[0px_0px_0px_0px_var(--color-bg-300)] disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-y-0 disabled:hover:shadow-[0px_4px_0px_0px_var(--color-bg-300)]"
			>
				{'>> COPY URL <<'}
			</button>
		{:else}
			<p class="text-bg-100 font-pixel">
				Duel against {currentOpponent.name}.
			</p>

			<p class="text-bg-100 font-pixel">
				They are {#if isOpponentReady}ready{:else}not ready{/if}.
			</p>
		{/if}
	</div>
</div>

<!-- Character Tooltip -->
{#if hoveredCharacter && hoveredCharacter.actions}
	<div
		class="fixed z-50 pointer-events-none transform -translate-x-1/2 -translate-y-full"
		style="left: {tooltipPosition.x}px; top: {tooltipPosition.y}px;"
	>
		<div class="bg-bg-800 border-2 border-bg-300 rounded-lg p-4 shadow-lg max-w-sm">
			<div class="mb-3">
				<h3 class="font-pixel text-bg-100 text-lg mb-1">
					{hoveredCharacter.name || `Character ${hoveredCharacter.id}`}
				</h3>
				{#if hoveredCharacter.origin}
					<p class="font-pixel text-bg-200 text-xs mb-2">from {hoveredCharacter.origin}</p>
				{/if}
				{#if hoveredCharacter.elements && hoveredCharacter.elements.length > 0}
					<div class="flex gap-1 mb-2">
						{#each hoveredCharacter.elements as element}
							<span
								class="px-2 py-1 rounded text-xs font-pixel"
								style="background-color: {getElementColor(element)}; color: white;"
							>
								{element}
							</span>
						{/each}
					</div>
				{/if}
			</div>

			<div class="space-y-2">
				<h4 class="font-pixel text-bg-100 text-sm">Actions:</h4>
				{#each hoveredCharacter.actions as action}
					<div class="bg-bg-700 border border-bg-400 rounded p-2">
						<div class="flex justify-between items-start mb-1">
							<span class="font-pixel text-bg-100 text-sm">{action.name}</span>
							<span
								class="px-1 py-0.5 rounded text-xs font-pixel"
								style="background-color: {getElementColor(action.element)}; color: white;"
							>
								{action.element}
							</span>
						</div>
						<p class="font-pixel text-bg-200 text-xs mb-2">{action.description}</p>
						<div class="flex gap-3 text-xs font-pixel">
							<span class="text-red-400">DMG: {action.Damage}</span>
							<span class="text-blue-400">MANA: {action.ManaCost}</span>
							{#if action.Oversight}
								<span class="text-yellow-400">OVERSIGHT</span>
							{/if}
						</div>
					</div>
				{/each}
			</div>
		</div>
	</div>
{/if}
