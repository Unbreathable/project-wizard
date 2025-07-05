<script lang="ts">
	let { health }: { health: number } = $props();

	let previousHealth = health;
	let animatingDifference = $state<number | null>(null);
	let animationKey = $state(0);

	$effect(() => {
		if (health !== previousHealth) {
			const difference = health - previousHealth;
			animatingDifference = difference;
			animationKey++;
			previousHealth = health;

			// Clear the animation after it completes
			setTimeout(() => {
				animatingDifference = null;
			}, 1000);
		}
	});
</script>

<div class="flex items-center gap-1.5 bg-bg-700 border-2 border-bg-300 p-1 relative">
	<img
		src="heart.png"
		alt="health"
		class="w-4 aspect-auto"
		style="image-rendering: pixelated; image-rendering: -moz-crisp-edges; image-rendering: crisp-edges;"
	/>
	<span class="text-bg-100 font-pixel text-xs">{health < 0 ? 0 : health}</span>

	{#if animatingDifference !== null}
		{#key animationKey}
			<span
				class="absolute font-pixel text-sm font-bold pointer-events-none animate-health-change z-10"
				class:text-green-300={animatingDifference > 0}
				class:text-red-300={animatingDifference < 0}
				class:drop-shadow-glow-green={animatingDifference > 0}
				class:drop-shadow-glow-red={animatingDifference < 0}
				style="left: 2rem; top: -0.25rem; text-shadow: 1px 1px 2px rgba(0,0,0,0.8);"
			>
				{animatingDifference > 0 ? '+' : ''}{animatingDifference}
			</span>
		{/key}
	{/if}
</div>

<style>
	@keyframes health-change {
		0% {
			opacity: 1;
			transform: translateY(0) scale(1.2);
		}
		100% {
			opacity: 0;
			transform: translateY(-1.5rem);
		}
	}

	.animate-health-change {
		animation: health-change 1s ease-out forwards;
	}

	.drop-shadow-glow-green {
		filter: drop-shadow(0 0 4px rgba(34, 197, 94, 0.6));
	}

	.drop-shadow-glow-red {
		filter: drop-shadow(0 0 4px rgba(239, 68, 68, 0.6));
	}
</style>
