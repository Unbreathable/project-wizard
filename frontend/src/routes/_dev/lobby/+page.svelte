<script lang="ts">
	import { Element, type Character } from '$lib/characters';
	import PixelArtImage from '$lib/components/PixelArtImage.svelte';
	import CharacterRender from '../../game/characterRender.svelte';

	const characterAmount = 4;

	let sampleCharacter: Character = {
		id: 1,
		url: 'shuna.png',
		name: 'Shuna',
		origin: 'Tensura',
		health: 100,
		elements: [Element.Light, Element.Earth],
		actions: {
			1: {
				name: 'Healing Light',
				description: 'Restores health to allies with divine light magic.',
				element: Element.Light,
				damage: 0,
				mana_cost: 15,
				oversight: false
			},
			2: {
				name: 'Earth Shield',
				description: 'Creates a protective barrier using earth magic.',
				element: Element.Earth,
				damage: 10,
				mana_cost: 12,
				oversight: false
			}
		}
	};
</script>

<div class="flex flex-col items-center justify-between w-screen h-screen p-2">
	<div class="flex flex-col gap-2 items-center justify-center p-2">
		<p class="text-bg-100 font-pixel">Your selection</p>
		<div class="flex text-center gap-2 w-min">
			{#each { length: characterAmount }}
				<CharacterRender character={sampleCharacter} openAbilityMenu={() => {}} own={false} />
			{/each}
		</div>
	</div>
	<div class="flex text-center p-4 gap-4 w-min"></div>
</div>
