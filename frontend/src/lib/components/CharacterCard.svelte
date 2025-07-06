<script lang="ts">
	import { type Character } from '$lib/characters';
	import ActionDescription from './ActionDescription.svelte';
	import ElementChip from './ElementChip.svelte';

	let { character }: { character: Character } = $props();
</script>

<div class="bg-bg-800 border-2 border-bg-300 p-4 max-w-sm">
	<div class="mb-3">
		<div class="flex justify-between">
			<h3 class="font-pixel text-bg-100 text-lg mb-1">
				{character.name || `Character ${character.id}`}
			</h3>
			{#if character.health}
				<p class="text-p-red-100 font-pixel">{character.health} HP</p>
			{/if}
		</div>
		{#if character.origin}
			<p class="font-pixel text-bg-200 text-xs mb-2">from {character.origin}</p>
		{/if}
		{#if character.elements && character.elements.length > 0}
			<div class="flex gap-3 mb-2">
				{#each character.elements as element}
					<ElementChip {element} />
				{/each}
			</div>
		{/if}
	</div>

	{#if character.actions}
		<div class="flex flex-col mt-2 gap-2 space-y-2">
			{#each Object.values(character.actions) as action}
				<ActionDescription {action} />
			{/each}
		</div>
	{/if}
</div>
