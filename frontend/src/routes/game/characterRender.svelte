<script lang="ts">
	import type { Action, Character } from '$lib/characters';
	import type { Arrow } from '$lib/components/arrows/arrows';
	import CharacterCard from '$lib/components/CharacterCard.svelte';
	import PixelArtImage from '$lib/components/PixelArtImage.svelte';
	import HealthIndicator from './healthIndicator.svelte';
	import { LinkPreview } from 'bits-ui';

	let {
		character,
		onClick,
		onAction,
		healthBar = true,
		actionsClickable = false
	}: {
		character: Character;
		onClick?: (e: MouseEvent) => void;
		onAction?: (action: Action) => void;
		healthBar?: boolean;
		actionsClickable?: boolean;
	} = $props();

	let previewOpen = $state(false);
	let changingBlocked = false;

	function changed(newOpen: boolean) {
		if (!newOpen) {
			changingBlocked = true;
			setTimeout(() => (changingBlocked = false), 500);
		}
	}

	function handleClick(e: MouseEvent) {
		if (onClick != null) {
			onClick(e);
			return;
		}
		if (changingBlocked) {
			return;
		}
		previewOpen = !previewOpen;
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<LinkPreview.Root
	bind:open={previewOpen}
	openDelay={999999999}
	onOpenChange={changed}
	closeDelay={0}
>
	<LinkPreview.Trigger class="cursor-pointer">
		<div onclick={handleClick} class="relative cursor-pointer">
			<PixelArtImage url="character-frame-empty.png" class="w-28 h-28" />
			<PixelArtImage
				url={character.url}
				class="absolute top-0 left-0 w-28 h-28 pointer-events-none"
			/>
			{#if healthBar}
				<div class="absolute bottom-0 left-0 right-0 flex items-center justify-center">
					<HealthIndicator health={character.health!} />
				</div>
			{/if}
		</div>
	</LinkPreview.Trigger>
	<LinkPreview.Content class="bg-bg-800 border-2 border-bg-300 z-50" sideOffset={0} side="bottom">
		<CharacterCard {actionsClickable} {onAction} {character} />
	</LinkPreview.Content>
</LinkPreview.Root>
