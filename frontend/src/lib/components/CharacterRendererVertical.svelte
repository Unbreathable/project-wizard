<script lang="ts">
	import type { Action, Character } from '$lib/characters';
	import type { Arrow } from '$lib/components/arrows/arrows';
	import CharacterCard from '$lib/components/CharacterCard.svelte';
	import PixelArtImage from '$lib/components/PixelArtImage.svelte';
	import { LinkPreview } from 'bits-ui';
	import HealthIndicator from './HealthIndicator.svelte';

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
		<div onclick={handleClick} class="cursor-pointer flex border-2 border-bg-300">
			<div class="relative">
				<PixelArtImage url="character-frame-empty.png" class="w-16 h-16" />
				<PixelArtImage
					url={character.url}
					class="absolute top-0 left-0 w-16 h-16 pointer-events-none"
				/>
			</div>
			<div class="flex items-right justify-center flex-col p-2 pl-0 gap-2 text-left">
				<HealthIndicator health={100} />
				<div class="flex items-center gap-1">
					<PixelArtImage url="heart.png" class="w-4 h-4 aspect-auto" />
					<PixelArtImage url="heart.png" class="w-4 h-4 aspect-auto" />
					<PixelArtImage url="heart.png" class="w-4 h-4 aspect-auto" />
				</div>
			</div>
		</div>
	</LinkPreview.Trigger>
	<LinkPreview.Content class="bg-bg-800 border-2 border-bg-300 z-50" sideOffset={8} side="right">
		<CharacterCard {actionsClickable} {onAction} {character} />
	</LinkPreview.Content>
</LinkPreview.Root>
