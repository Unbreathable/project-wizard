<script lang="ts">
	import { Element, type Action } from '$lib/characters';
	import ElementChip from './ElementChip.svelte';

	let {
		action,
		clickable,
		onClick
	}: {
		action: Action;
		clickable?: boolean;
		onClick?: () => void;
	} = $props();
</script>

<div class="bg-bg-700 border border-bg-400 p-2 text-left">
	<div class="flex justify-between items-start mb-1">
		<span class="font-pixel text-bg-100 text-sm">{action.name}</span>
		{#if action.element != Element.None}
			<ElementChip element={action.element} />
		{/if}
	</div>
	<p class="font-pixel text-bg-200 text-xs">{action.description}</p>
	<div class="flex gap-6 text-xs font-pixel text-left">
		{#if action.damage}
			<span class="text-p-red-100 mt-2">DMG: {action.damage}</span>
		{/if}
		{#if action.mana_cost}
			<span class="text-p-blue-200 mt-2">MANA: {action.mana_cost}</span>
		{/if}
		{#if action.oversight}
			<span class="text-p-orange-100 mt-2">OVERSIGHT</span>
		{/if}
	</div>
</div>
