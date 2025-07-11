<script lang="ts">
	interface Point {
		x: number;
		y: number;
	}

	interface Arrow {
		id: number;
		start: Point;
		startDirection: Direction;
		end: Point;
		endDirection: Direction;
	}
	6;
	enum Direction {
		Up,
		Down,
		Left,
		Right
	}

	let svgElement: SVGSVGElement;
	let containerElement: HTMLDivElement;
	let arrows: Arrow[] = $state([]);
	let firstPoint: Point | null = $state(null);
	let nextArrowId = 1;

	function handleContainerClick(event: MouseEvent) {
		if (event.target !== svgElement && event.target !== containerElement) {
			return;
		}

		const rect = containerElement.getBoundingClientRect();
		const x = event.clientX - rect.left;
		const y = event.clientY - rect.top;

		if (!firstPoint) {
			firstPoint = { x, y };
		} else {
			const newArrow: Arrow = {
				id: nextArrowId++,
				start: firstPoint,
				startDirection: Direction.Up,
				end: { x, y },
				endDirection: Direction.Down
			};
			arrows.push(newArrow);
			firstPoint = null;
		}
	}

	function handleArrowClick(arrow: Arrow, event: MouseEvent) {
		event.stopPropagation();
		console.log('Arrow clicked!', arrow);
	}

	function calculateArrowPath(
		start: Point,
		startDirection: Direction,
		end: Point,
		endDirection: Direction
	) {
		const dx = end.x - start.x;
		const dy = end.y - start.y;

		let segmentLength = 0;
		if (Math.abs(dx) > Math.abs(dy)) {
			segmentLength = dx / 4;
		} else {
			segmentLength = dy / 4;
		}

		const startExtended = movePoint(start, startDirection, segmentLength);
		const endExtended = movePoint(end, endDirection, segmentLength);

		// Calculate arrow head points
		const headSize = segmentLength / 3;
		let headPoints: Point[] = [];

		switch (endDirection) {
			case Direction.Down:
				headPoints = [
					{ x: end.x - headSize / 2, y: end.y + headSize },
					{ x: end.x + headSize / 2, y: end.y + headSize }
				];
				break;
			case Direction.Up:
				headPoints = [
					{ x: end.x - headSize / 2, y: end.y - headSize },
					{ x: end.x + headSize / 2, y: end.y - headSize }
				];
				break;
			case Direction.Left:
				headPoints = [
					{ x: end.x + headSize, y: end.y - headSize / 2 },
					{ x: end.x + headSize, y: end.y + headSize / 2 }
				];
				break;
			case Direction.Right:
				headPoints = [
					{ x: end.x - headSize, y: end.y - headSize / 2 },
					{ x: end.x - headSize, y: end.y + headSize / 2 }
				];
				break;
		}

		// Create integrated path with arrow head
		const linePath = `M ${start.x} ${start.y} L ${startExtended.x} ${startExtended.y} L ${endExtended.x} ${endExtended.y} L ${end.x} ${end.y}`;
		const arrowHeadPath = `M ${headPoints[0].x} ${headPoints[0].y} L ${end.x} ${end.y} L ${headPoints[1].x} ${headPoints[1].y}`;

		return {
			line: `${linePath} ${arrowHeadPath}`
		};
	}

	function movePoint(point: Point, direction: Direction, amount: number): Point {
		switch (direction) {
			case Direction.Up:
				return { x: point.x, y: point.y + amount };
			case Direction.Down:
				return { x: point.x, y: point.y - amount };
			case Direction.Left:
				return { x: point.x - amount, y: point.y };
			case Direction.Right:
				return { x: point.x + amount, y: point.y };
		}
	}
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<div class="w-screen h-screen" bind:this={containerElement} onclick={handleContainerClick}>
	<svg bind:this={svgElement} width="100%" height="100%">
		{#each arrows as arrow (arrow.id)}
			{@const path = calculateArrowPath(
				arrow.start,
				arrow.startDirection,
				arrow.end,
				arrow.endDirection
			)}
			<g onclick={(e) => handleArrowClick(arrow, e)}>
				<path d={path.line} class="arrow-body" />
				<path d={path.line} class="arrow-outline" />
			</g>
		{/each}
	</svg>
</div>

<style>
	.arrow-body {
		stroke: white;
		stroke-width: 20;
		fill: none;
	}

	.arrow-outline {
		stroke: red;
		stroke-width: 10;
		fill: none;
	}
</style>
