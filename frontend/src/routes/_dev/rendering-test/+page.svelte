<script lang="ts">
	import { Element, type Action, type Character } from '$lib/characters';
	import ActionDescription from '$lib/components/ActionDescription.svelte';
	import CharacterCard from '$lib/components/CharacterCard.svelte';
	import ElementChip from '$lib/components/ElementChip.svelte';
	import PixelArtImage from '$lib/components/PixelArtImage.svelte';
	import CharacterRender from '../../game/characterRender.svelte';

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

<div class="flex flex-col gap-6">
	<PixelArtImage url="shuna.png" class="w-28 h-28" />

	<PixelArtImage url="character-frame-empty.png" class="w-28 h-28" />

	<!-- Character bar -->
	<div class="flex text-center p-4 gap-4 border-6 border-double w-min border-bg-200 bg-bg-700">
		<CharacterRender character={sampleCharacter} openAbilityMenu={() => {}} own={false} />
		<CharacterRender character={sampleCharacter} openAbilityMenu={() => {}} own={false} />
		<CharacterRender character={sampleCharacter} openAbilityMenu={() => {}} own={false} />
		<CharacterRender character={sampleCharacter} openAbilityMenu={() => {}} own={false} />
	</div>

	<div class="flex gap-2 p-6">
		<ElementChip element={Element.Fire} />
		<ElementChip element={Element.Water} />
		<ElementChip element={Element.Air} />
		<ElementChip element={Element.Earth} />
		<ElementChip element={Element.Dark} />
		<ElementChip element={Element.Light} />
	</div>

	<div class="flex flex-col gap-6 w-full max-w-sm px-6">
		<ActionDescription
			action={{
				damage: 20,
				mana_cost: 10,
				description: 'A great attack.',
				element: Element.Dark,
				name: 'Attack',
				oversight: false
			}}
		/>
		<ActionDescription
			action={{
				name: 'Oversight',
				description: 'Hello world',
				element: Element.None,
				damage: 0,
				mana_cost: 0,
				oversight: true
			}}
		/>
	</div>

	<div class="px-6">
		<CharacterCard
			character={{
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
			}}
		/>
	</div>
</div>
