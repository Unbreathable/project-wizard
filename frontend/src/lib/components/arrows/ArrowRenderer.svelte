<script lang="ts">
	import SingleArrow from './SingleArrow.svelte';
	import { calculateArrowPath, getArrowColorOrDefault, type Arrow } from './arrows';
	import { onMount } from 'svelte';

	let {
		arrows
	}: {
		arrows: Arrow[];
	} = $props();

	let windowWidth = $state(0);
	let windowHeight = $state(0);

	onMount(() => {
		window.addEventListener('resize', handleResize);
		return () => window.removeEventListener('resize', handleResize);
	});

	function handleResize() {
		windowWidth = window.innerWidth;
		windowHeight = window.innerHeight;
	}

	function handleArrowClick(index: number, e: MouseEvent) {
		console.log('arrow clicked!', index);
	}

	function preProcessArrows(arrows: Arrow[]): Map<string, Arrow[]> {
		let map = new Map<string, Arrow[]>();
		for (let arrow of arrows) {
			if (map.has(arrow.startId)) {
				map.get(arrow.startId)?.push(arrow);
			} else {
				map.set(arrow.startId, [arrow]);
			}
		}

		return map;
	}
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
{#key windowHeight + windowWidth}
	<div class="absolute inset-0 z-20 w-screen h-screen pointer-events-none">
		<svg width="100%" height="100%">
			{#each { length: 1 }}
				{@const map = preProcessArrows(arrows)}
				{#each map.entries() as [_, arrowGroup], groupIndex}
					{#each arrowGroup as arrow, arrowIndex}
						{@const path = calculateArrowPath(arrow, arrowIndex, arrowGroup.length)}
						{@const colors = getArrowColorOrDefault(arrow)}
						<SingleArrow
							{path}
							onClick={(e) => handleArrowClick(groupIndex * 100 + arrowIndex, e)}
							--body-color={colors.bodyColor}
							--outline-color={colors.outlineColor}
						/>
					{/each}
				{/each}
			{/each}
		</svg>
	</div>
{/key}
