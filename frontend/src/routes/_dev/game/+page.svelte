<script lang="ts">
	import { Element, type Character } from '$lib/characters';
	import ArrowRenderer from '$lib/components/arrows/ArrowRenderer.svelte';
	import { Direction, type Arrow } from '$lib/components/arrows/arrows';
	import CharacterRender from '../../game/characterRender.svelte';

	const characterAmount = 4;

	let arrows: Arrow[] = $state([]);

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

	let startId: string | null = null;
	let bottomElement: HTMLElement | null = null;

	function handleClickBottom(e: MouseEvent, player: number, slot: number) {
		if (!(e.currentTarget instanceof HTMLElement)) {
			console.error('no html element wtf');
			return;
		}
		if (bottomElement != null) {
			arrows.push({
				start: bottomElement,
				startId: startId!,
				startDirection: Direction.Up,
				end: e.currentTarget as HTMLElement,
				endDirection: Direction.Up
			});
			console.log('arrow!!!');
			bottomElement = null;
			return;
		}
		startId = `${player}:${slot}`;
		bottomElement = e.currentTarget as HTMLElement;
	}

	function handleClickTop(e: MouseEvent, player: number, slot: number) {
		if (!(e.currentTarget instanceof HTMLElement)) {
			console.error('no html element wtf');
			return;
		}
		if (bottomElement != null) {
			arrows.push({
				start: bottomElement,
				startId: startId!,
				startDirection: Direction.Up,
				end: e.currentTarget as HTMLElement,
				endDirection: Direction.Down
			});
			console.log('arrow!!!');
			bottomElement = null;
		}
	}
</script>

<div class="flex flex-col items-center justify-between w-screen h-screen p-2 py-3 gap-16">
	<div class="flex gap-16">
		<div class="flex flex-col gap-2 items-center justify-center p-2 relative">
			<p class="text-bg-100 font-pixel">[Enemy Name]</p>
			<div class="flex gap-2 w-min">
				{#each { length: characterAmount } as _, index}
					<CharacterRender
						character={sampleCharacter}
						onClick={(e) => handleClickTop(e, 1, index)}
						healthBar={false}
					/>
				{/each}
			</div>
		</div>
		<div class="flex flex-col gap-2 items-center justify-center p-2">
			<p class="text-bg-100 font-pixel">[Enemy Name]</p>
			<div class="flex text-center gap-2 w-min">
				{#each { length: characterAmount } as _, index}
					<CharacterRender
						character={sampleCharacter}
						onClick={(e) => handleClickTop(e, 2, index)}
						healthBar={false}
					/>
				{/each}
			</div>
		</div>
	</div>
	<div class="flex gap-16">
		<div class="flex flex-col gap-2 items-center justify-center p-2">
			<div class="flex text-center gap-2 w-min">
				{#each { length: characterAmount } as _, index}
					<CharacterRender
						character={sampleCharacter}
						onClick={(e) => handleClickBottom(e, 3, index)}
						healthBar={false}
					/>
				{/each}
			</div>
			<p class="text-bg-100 font-pixel">[Your Name]</p>
		</div>
		<div class="flex flex-col gap-2 items-center justify-center p-2">
			<div class="flex text-center gap-2 w-min">
				{#each { length: characterAmount } as _, index}
					<CharacterRender
						character={sampleCharacter}
						onClick={(e) => handleClickBottom(e, 4, index)}
						healthBar={false}
					/>
				{/each}
			</div>
			<p class="text-bg-100 font-pixel">[Mate Name]</p>
		</div>
	</div>
</div>

<ArrowRenderer {arrows} />
