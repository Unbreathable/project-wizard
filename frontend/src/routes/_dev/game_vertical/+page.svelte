<script lang="ts">
	import { Element, getRandomElement, type Character } from '$lib/characters';
	import ArrowRenderer from '$lib/components/arrows/ArrowRenderer.svelte';
	import { Direction, getArrowColor, type Arrow } from '$lib/components/arrows/arrows';
	import CharacterRendererVertical from '$lib/components/CharacterRendererVertical.svelte';

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

	function handleClickLeft(e: MouseEvent, player: number, slot: number) {
		if (!(e.currentTarget instanceof HTMLElement)) {
			console.error('no html element wtf');
			return;
		}
		if (bottomElement != null) {
			arrows.push({
				start: bottomElement,
				startId: startId!,
				startDirection: Direction.Left,
				end: e.currentTarget as HTMLElement,
				endDirection: Direction.Left,
				colors: getArrowColor(Element.Earth, false, true)
			});
			console.log('arrow!!!');
			bottomElement = null;
			return;
		}
		startId = `${player}:${slot}`;
		bottomElement = e.currentTarget as HTMLElement;
	}

	function handleClickRight(e: MouseEvent, player: number, slot: number) {
		if (!(e.currentTarget instanceof HTMLElement)) {
			console.error('no html element wtf');
			return;
		}
		if (bottomElement != null) {
			arrows.push({
				start: bottomElement,
				startId: startId!,
				startDirection: Direction.Left,
				end: e.currentTarget as HTMLElement,
				endDirection: Direction.Right,
				colors: getArrowColor(getRandomElement(), false, true)
			});
			console.log('arrow!!!');
			bottomElement = null;
		}
	}
</script>

<div class="relative flex items-center justify-between w-screen min-h-screen h-full p-2 gap-16">
	<div class="flex flex-col gap-8">
		<div class="flex gap-4 items-center justify-center p-2">
			<p
				class="text-bg-100 font-pixel rotate-180"
				style="text-orientation: sideways; writing-mode: vertical-rl;"
			>
				[Your Name]
			</p>
			<div class="flex flex-col text-center gap-2">
				{#each { length: characterAmount } as _, index}
					<CharacterRendererVertical
						character={sampleCharacter}
						onClick={(e) => handleClickLeft(e, 3, index)}
						healthBar={true}
					/>
				{/each}
			</div>
		</div>
		<div class="flex gap-4 items-center justify-center p-2">
			<p
				class="text-bg-100 font-pixel rotate-180"
				style="text-orientation: sideways; writing-mode: vertical-rl;"
			>
				[Mate Name]
			</p>
			<div class="flex flex-col text-center gap-2">
				{#each { length: characterAmount } as _, index}
					<CharacterRendererVertical
						character={sampleCharacter}
						onClick={(e) => handleClickLeft(e, 4, index)}
						healthBar={true}
					/>
				{/each}
			</div>
		</div>
	</div>

	<div class="flex flex-col gap-16">
		<div class="flex gap-4 items-center justify-center p-2 relative">
			<div class="flex flex-col gap-2">
				{#each { length: characterAmount } as _, index}
					<CharacterRendererVertical
						character={sampleCharacter}
						onClick={(e) => handleClickRight(e, 1, index)}
						healthBar={true}
					/>
				{/each}
			</div>
			<p
				class="text-bg-100 font-pixel"
				style="text-orientation: sideways; writing-mode: vertical-rl;"
			>
				[Enemy Name]
			</p>
		</div>
		<div class="flex gap-4 items-center justify-center p-2">
			<div class="flex flex-col text-center gap-2">
				{#each { length: characterAmount } as _, index}
					<CharacterRendererVertical
						character={sampleCharacter}
						onClick={(e) => handleClickRight(e, 2, index)}
						healthBar={true}
					/>
				{/each}
			</div>
			<p
				class="text-bg-100 font-pixel"
				style="text-orientation: sideways; writing-mode: vertical-rl;"
			>
				[Enemy Name]
			</p>
		</div>
	</div>

	<ArrowRenderer {arrows} />
</div>
