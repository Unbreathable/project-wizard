<script lang="ts">
	import RetroButton from '$lib/components/RetroButton.svelte';
</script>

<div class="flex flex-col items-center justify-between w-screen h-screen p-2 gap-2">
	<div class="flex items-center justify-center w-full h-full gap-8">
		<div
			class="flex flex-col p-6 gap-4 border-6 border-double border-bg-200 bg-bg-700 w-full max-w-md"
		>
			<h1 class="text-bg-100 text-center font-pixel text-xl mb-2">TEAM 1</h1>

			<RetroButton disabled={true} onClick={() => {}}>{'>> JOIN <<'}</RetroButton>
		</div>
		<div
			class="flex flex-col p-6 gap-4 border-6 border-double border-bg-200 bg-bg-700 w-full max-w-md"
		>
			<h1 class="text-bg-100 text-center font-pixel text-xl mb-2">TEAM 2</h1>

			<RetroButton disabled={true} onClick={() => {}}>{'>> JOIN <<'}</RetroButton>
		</div>
	</div>

	<div class="flex p-2 gap-4 justify-center items-center w-full">
		<RetroButton class="w-full max-w-sm" onClick={() => {}}>{'>> LEAVE TEAM <<'}</RetroButton>
		<RetroButton class="w-full max-w-sm" onClick={() => {}}>{'>> READY <<'}</RetroButton>
	</div>
</div>
