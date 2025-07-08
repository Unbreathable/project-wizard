<script lang="ts">
	import type { Character } from '$lib/characters';
	import PixelArtImage from '$lib/components/PixelArtImage.svelte';
	import HealthIndicator from './healthIndicator.svelte';

	let {
		own,
		character,
		openAbilityMenu,
		healthBar = true
	}: {
		own: boolean;
		character: Character;
		healthBar?: boolean;
		openAbilityMenu: (char: Character, own: boolean, position: { x: number; y: number }) => void;
	} = $props();

	function handleClick(event: MouseEvent) {
		event.stopPropagation(); // Prevent the parent's closeActionMenu from being called
		const rect = (event.currentTarget as HTMLElement).getBoundingClientRect();
		const position = {
			x: rect.left + rect.width / 2,
			y: rect.top + rect.height / 2
		};
		openAbilityMenu(character, own, position);
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div onclick={handleClick} class="relative cursor-pointer">
	<PixelArtImage url="character-frame-empty.png" class="w-28 h-28" />
	<PixelArtImage url={character.url} class="absolute top-0 left-0 w-28 h-28 pointer-events-none" />
	{#if healthBar}
		<div class="absolute bottom-0 left-0 right-0 flex items-center justify-center">
			<HealthIndicator health={character.health!} />
		</div>
	{/if}
</div>
