<script lang="ts">
	import { onMount } from 'svelte';
	import {
		characters,
		initializeCharacters,
		type Character,
		Element,
		getElementColor,
		type Action
	} from '$lib/characters';
	import CharacterRenderComponent from './characterRender.svelte';

	let showCharacters = $state(false);

	onMount(() => {
		loadCharacters();
	});

	async function loadCharacters() {
		await initializeCharacters();
		showCharacters = true;
	}

	let characterForMenu = $state<Character | undefined>(undefined);
	let menuPosition = $state({ x: 0, y: 0 });
	let menuBelow = $state(false);

	let selectingCharacter = $state(false);
	let selectionCallback = (c: Character, o: boolean) => {};

	let selectedAction = $state<{ char: number; action: Action; targetOwn: string } | undefined>(
		undefined
	);
	let selectedOversight = $state<{ char: number; action: Action } | undefined>(undefined);

	function openActionMenu(character: Character, own: boolean, position: { x: number; y: number }) {
		if (selectingCharacter) {
			selectionCallback(character, own);
			return;
		}
		characterForMenu = character;
		menuPosition = position;
		menuBelow = !own;
	}

	function closeActionMenu(event?: MouseEvent) {
		// Only close if clicking on the background, not on character or menu
		if (event && event.target !== event.currentTarget) {
			return;
		}
		characterForMenu = undefined;
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	class="flex flex-col justify-between items-center w-screen h-screen p-4"
	onclick={closeActionMenu}
>
	<div class="flex flex-col gap-4 items-center">
		<h1 class="text-bg-100 font-pixel">Unbreathable</h1>
		{#if showCharacters}
			<div class="flex text-center p-4 gap-4 border-6 border-double border-bg-200 bg-bg-700">
				<CharacterRenderComponent
					character={characters[1]}
					openAbilityMenu={openActionMenu}
					own={false}
				/>
			</div>
		{/if}
	</div>

	<div class="flex flex-col gap-4 items-center">
		{#if showCharacters}
			<div class="flex text-center p-4 gap-4 border-6 border-double border-bg-200 bg-bg-700">
				<CharacterRenderComponent
					character={characters[1]}
					openAbilityMenu={openActionMenu}
					own={true}
				/>
			</div>
		{/if}
		<h1 class="text-bg-100 font-pixel">Your lineup</h1>
	</div>
</div>

<!-- Action menu -->
{#if characterForMenu}
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		class="fixed z-50 pointer-events-none transform -translate-x-1/2 {menuBelow
			? ''
			: '-translate-y-full'}"
		style="left: {menuPosition.x}px; top: {menuPosition.y}px;"
	>
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<div
			class="bg-bg-800 border-2 border-bg-300 rounded-lg p-4 shadow-lg max-w-sm pointer-events-auto"
			onclick={(e) => e.stopPropagation()}
		>
			<div class="mb-3">
				<h3 class="font-pixel text-bg-100 text-lg mb-1">
					{characterForMenu.name || `Character ${characterForMenu.id}`}
				</h3>
				{#if characterForMenu.origin}
					<p class="font-pixel text-bg-200 text-xs mb-2">from {characterForMenu.origin}</p>
				{/if}
				{#if characterForMenu.elements && characterForMenu.elements.length > 0}
					<div class="flex gap-1 mb-2">
						{#each characterForMenu.elements as element}
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

			{#if characterForMenu.actions}
				<div class="space-y-2">
					<h4 class="font-pixel text-bg-100 text-sm">Actions:</h4>
					{#each Object.values(characterForMenu.actions) as action}
						<div class="bg-bg-700 border border-bg-400 rounded p-2">
							<div class="flex justify-between items-start mb-1">
								<span class="font-pixel text-bg-100 text-sm">{action.name}</span>
								{#if action.element != Element.None}
									<span
										class="px-3 py-2 rounded text-xs font-pixel"
										style="background-color: {getElementColor(action.element)}; color: white;"
									>
										{action.element}
									</span>
								{/if}
							</div>
							<p class="font-pixel text-bg-200 text-xs">{action.description}</p>
							<div class="flex gap-3 text-xs font-pixel">
								{#if action.damage}
									<span class="text-red-400 mt-2">DMG: {action.damage}</span>
								{/if}
								{#if action.mana_cost}
									<span class="text-blue-400 mt-2">MANA: {action.mana_cost}</span>
								{/if}
								{#if action.oversight}
									<span class="text-yellow-400 mt-2">OVERSIGHT</span>
								{/if}
							</div>
						</div>
					{/each}
				</div>
			{/if}
		</div>
	</div>
{/if}
