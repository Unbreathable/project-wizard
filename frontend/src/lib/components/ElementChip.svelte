<script lang="ts">
	import { Element, getElementColor } from '$lib/characters';

	let { element }: { element: Element } = $props();
</script>

<span
	class="px-2 py-1 rounded text-xs font-pixel w-min"
	style="background-color: {getElementColor(element)}; color: white;"
>
	{element}
</span>
