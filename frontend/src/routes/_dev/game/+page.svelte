<script lang="ts">
	import { Element, type Character } from '$lib/characters';
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
			},
			3: {
				name: 'Earth Shield',
				description: 'Creates a protective barrier using earth magic.',
				element: Element.Earth,
				damage: 10,
				mana_cost: 10,
				oversight: false
			}
		}
	};
</script>

<div class="flex flex-col items-center justify-between w-screen h-screen p-2 py-3 gap-16">
	<div class="flex gap-16">
		<div class="flex flex-col gap-2 items-center justify-center p-2 relative">
			<p class="text-bg-100 font-pixel">[Enemy Name]</p>
			<div class="flex gap-2 w-min">
				{#each { length: characterAmount } as _, index}
					<CharacterRender character={sampleCharacter} healthBar={false} />
				{/each}
			</div>
		</div>
		<div class="flex flex-col gap-2 items-center justify-center p-2">
			<p class="text-bg-100 font-pixel">[Enemy Name]</p>
			<div class="flex text-center gap-2 w-min">
				{#each { length: characterAmount }}
					<CharacterRender character={sampleCharacter} healthBar={false} />
				{/each}
			</div>
		</div>
	</div>
	<div class="flex gap-16">
		<div class="flex flex-col gap-2 items-center justify-center p-2">
			<div class="flex text-center gap-2 w-min">
				{#each { length: characterAmount }}
					<CharacterRender character={sampleCharacter} healthBar={false} />
				{/each}
			</div>
			<p class="text-bg-100 font-pixel">[Your Name]</p>
		</div>
		<div class="flex flex-col gap-2 items-center justify-center p-2">
			<div class="flex text-center gap-2 w-min">
				{#each { length: characterAmount }}
					<CharacterRender character={sampleCharacter} healthBar={false} />
				{/each}
			</div>
			<p class="text-bg-100 font-pixel">[Mate Name]</p>
		</div>
	</div>
</div>
