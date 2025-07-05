<script lang="ts">
	import type { Character } from '$lib/characters';
	import HealthIndicator from './healthIndicator.svelte';

	let {
		own,
		character,
		openAbilityMenu
	}: {
		own: boolean;
		character: Character;
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
	<div
		class="w-26 h-26 bg-center bg-no-repeat bg-contain"
		style="background-image: url('character-frame-empty.png'); image-rendering: pixelated; image-rendering: -moz-crisp-edges; image-rendering: crisp-edges;"
	></div>
	<div
		class="absolute top-0 left-0 w-26 h-26 bg-center bg-no-repeat bg-contain pointer-events-none"
		style="background-image: url('{character.url}'); image-rendering: pixelated; image-rendering: -moz-crisp-edges; image-rendering: crisp-edges;"
	></div>
	<div class="absolute bottom-0 left-0 right-0 flex items-center justify-center">
		<HealthIndicator health={character.health!} />
	</div>
</div>
