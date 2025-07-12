<script lang="ts">
	import type { Arrow, ArrowPath } from './arrows';

	let {
		path,
		onClick
	}: {
		path: ArrowPath;
		onClick?: (e: MouseEvent) => void;
	} = $props();
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<g
	onclick={onClick}
	class="pointer-events-auto cursor-pointer"
	style="--arrow-color-body: var(--bg-color-300); --arrow-color-outline: var(--bg-color-100);"
>
	<path d={path.arrowHead} class="arrowhead-outline" />
	<path d={path.line} class="arrow-outline" />
	<path d={path.line} class="arrow-body" />
	<path d={path.arrowHead} class="arrowhead-body" />
</g>

<style>
	.arrow-body {
		stroke: var(--body-color, var(--color-bg-200));
		stroke-width: 8;
		fill: none;
		stroke-linecap: square;
	}

	.arrow-outline {
		stroke: var(--outline-color, var(--color-bg-300));
		stroke-width: 16;
		fill: none;
		stroke-linecap: square;
	}

	.arrowhead-body {
		fill: var(--body-color, var(--color-bg-200));
		stroke: none;
	}

	.arrowhead-outline {
		fill: var(--outline-color, var(--color-bg-300));
		stroke: var(--outline-color, var(--color-bg-300));
		stroke-width: 8;
		stroke-linecap: butt;
		stroke-linejoin: miter;
	}
</style>
