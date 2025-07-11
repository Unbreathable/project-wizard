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
				endDirection: Direction.Up
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
			segmentLength = Math.abs(dx / 4);
		} else {
			segmentLength = Math.abs(dy / 4);
		}

		const startExtended = movePoint(start, startDirection, segmentLength);
		const endExtended = movePoint(end, endDirection, segmentLength);

		const arrowheadSize = 20;
		const oppositeDirection = getOppositeDirection(endDirection);
		const arrowheadPosition = movePoint(end, oppositeDirection, arrowheadSize - 8);
		const arrowhead = calculateArrowhead(arrowheadPosition, endDirection, arrowheadSize);

		return {
			line: `M ${start.x} ${start.y} L ${startExtended.x} ${startExtended.y} L ${endExtended.x} ${endExtended.y} L ${end.x} ${end.y}`,
			arrowhead: `M ${arrowheadPosition.x} ${arrowheadPosition.y} L ${arrowhead.left.x} ${arrowhead.left.y} L ${arrowhead.right.x} ${arrowhead.right.y} Z`
		};
	}

	function calculateArrowhead(point: Point, direction: Direction, size: number) {
		// Since directions are inverse, we need to calculate the arrowhead
		// pointing in the opposite direction of the endDirection
		let angle: number;
		switch (direction) {
			case Direction.Up: // Arrow coming from below, pointing up
				angle = -Math.PI / 2; // Point upward
				break;
			case Direction.Down: // Arrow coming from above, pointing down
				angle = Math.PI / 2; // Point downward
				break;
			case Direction.Left: // Arrow coming from right, pointing left
				angle = Math.PI; // Point left
				break;
			case Direction.Right: // Arrow coming from left, pointing right
				angle = 0; // Point right
				break;
		}

		// Rotate by 180 degrees
		angle += Math.PI;

		const leftAngle = angle + Math.PI * 0.75;
		const rightAngle = angle - Math.PI * 0.75;

		return {
			left: {
				x: point.x + Math.cos(leftAngle) * size,
				y: point.y + Math.sin(leftAngle) * size
			},
			right: {
				x: point.x + Math.cos(rightAngle) * size,
				y: point.y + Math.sin(rightAngle) * size
			}
		};
	}

	function getOppositeDirection(direction: Direction): Direction {
		switch (direction) {
			case Direction.Up:
				return Direction.Down;
			case Direction.Down:
				return Direction.Up;
			case Direction.Left:
				return Direction.Right;
			case Direction.Right:
				return Direction.Left;
		}
	}

	function movePoint(point: Point, direction: Direction, amount: number): Point {
		switch (direction) {
			case Direction.Up:
				return { x: point.x, y: point.y - amount };
			case Direction.Down:
				return { x: point.x, y: point.y + amount };
			case Direction.Left:
				return { x: point.x + amount, y: point.y };
			case Direction.Right:
				return { x: point.x - amount, y: point.y };
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
				<path d={path.arrowhead} class="arrowhead-outline" />
				<path d={path.line} class="arrow-outline" />
				<path d={path.line} class="arrow-body" />
				<path d={path.arrowhead} class="arrowhead-body" />
			</g>
		{/each}
	</svg>
</div>

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
