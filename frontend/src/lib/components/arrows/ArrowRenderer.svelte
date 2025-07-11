<script lang="ts">
	import { calculateArrowPath, type Arrow } from './arrows';
	import { onMount } from 'svelte';

	let {
		arrows
	}: {
		arrows: Arrow[];
	} = $props();

	let windowWidth = $state(0);
	let windowHeight = $state(0);

	onMount(() => {
		// Set initial values
		windowWidth = window.innerWidth;
		windowHeight = window.innerHeight;

		// Listen for resize events
		function handleResize() {
			windowWidth = window.innerWidth;
			windowHeight = window.innerHeight;
			console.log('resize');
		}

		window.addEventListener('resize', handleResize);

		// Cleanup
		return () => {
			window.removeEventListener('resize', handleResize);
		};
	});

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
						<g
							onclick={(e) => handleArrowClick(groupIndex * 100 + arrowIndex, e)}
							class="pointer-events-auto cursor-pointer"
						>
							<path d={path.arrowHead} class="arrowhead-outline" />
							<path d={path.line} class="arrow-outline" />
							<path d={path.line} class="arrow-body" />
							<path d={path.arrowHead} class="arrowhead-body" />
						</g>
					{/each}
				{/each}
			{/each}
		</svg>
	</div>
{/key}

<style>
	.arrow-body {
		stroke: var(--color-bg-300);
		stroke-width: 8;
		fill: none;
		stroke-linecap: square;
	}

	.arrow-outline {
		stroke: var(--color-bg-400);
		stroke-width: 16;
		fill: none;
		stroke-linecap: square;
	}

	.arrowhead-body {
		fill: var(--color-bg-300);
		stroke: none;
	}

	.arrowhead-outline {
		fill: var(--color-bg-400);
		stroke: var(--color-bg-400);
		stroke-width: 8;
		stroke-linecap: butt;
		stroke-linejoin: miter;
	}
</style>
