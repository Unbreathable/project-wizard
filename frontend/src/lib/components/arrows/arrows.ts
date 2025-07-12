import { Element } from "$lib/characters";

export enum Direction {
    Up,
    Down,
    Left,
    Right
}

export interface Arrow {
    start: HTMLElement;
    startId: string;
    startDirection: Direction;
    end: HTMLElement;
    endDirection: Direction;
    colors?: ArrowColors;
}

export interface ArrowPath {
    line: string;
    arrowHead: string;
}

export interface Point {
    x: number;
    y: number;
}

/**
 * Calculate the path for an arrow.
 */
export function calculateArrowPath(arrow: Arrow, index: number, arrows: number): ArrowPath {
    const gap = 8; // How much of a gap to leave from the html elements
    const arrowPadding = 16; // How big the arrow is to make sure we're pointing exactly to start and end
    const arrowWidth = 16;

    let totalLength = (arrowWidth + gap) * arrows - gap;
    let offset = (-totalLength / 2) + (arrowWidth + gap) * index + arrowWidth / 2;

    let start = getCenterAroundRect(arrow.start.getBoundingClientRect(), arrow.startDirection, offset);
    let end = getCenterAroundRect(arrow.end.getBoundingClientRect(), arrow.endDirection, 0);

    // Move to have some breathing room
    start = movePoint(start, arrow.startDirection, gap);
    end = movePoint(end, arrow.endDirection, gap);

    const dx = end.x - start.x;
    const dy = end.y - start.y;

    let segmentLength = 0;
    if (Math.abs(dx) > Math.abs(dy)) {
        segmentLength = Math.abs(dx / 4);
    } else {
        segmentLength = Math.abs(dy / 4);
    }

    const startExtended = movePoint(start, arrow.startDirection, segmentLength);
    const endExtended = movePoint(end, arrow.endDirection, segmentLength);

    const arrowDirection = getOppositeDirection(arrow.endDirection);
    end = movePoint(end, arrow.endDirection, arrowPadding / 2); // Move end a little bit to make sure the arrow points to end instead of above it

    const arrowHeadSize = 20;
    const arrowHeadPosition = movePoint(end, arrowDirection, arrowHeadSize - arrowPadding / 2);
    const arrowHead = calculateArrowHead(arrowHeadPosition, arrow.endDirection, arrowHeadSize);

    return {
        line: `M ${start.x} ${start.y} L ${startExtended.x} ${startExtended.y} L ${endExtended.x} ${endExtended.y} L ${end.x} ${end.y}`,
        arrowHead: `M ${arrowHeadPosition.x} ${arrowHeadPosition.y} L ${arrowHead.left.x} ${arrowHead.left.y} L ${arrowHead.right.x} ${arrowHead.right.y} Z`
    };
}

export function getCenterAroundRect(rect: DOMRect, direction: Direction, offset: number): Point {
    switch(direction) {
        case Direction.Up:
            return {
                x: rect.x + rect.width / 2 + offset,
                y: rect.y,
            };
        case Direction.Down:
            return {
                x: rect.x + rect.width / 2 + offset,
                y: rect.y + rect.height,
            };
        case Direction.Left:
            return  {
                x: rect.x,
                y: rect.y + rect.height / 2 + offset,
            }
        case Direction.Right:
            return {
                x: rect.x + rect.width,
                y: rect.y + rect.height / 2 + offset,
            };
    }
}

export function calculateArrowHead(point: Point, direction: Direction, size: number) {

    let angle: number;
    switch (direction) {
        case Direction.Up: // Arrow coming from below, pointing up
            angle = Math.PI / 2;
            break;
        case Direction.Down: // Arrow coming from above, pointing down
            angle = Math.PI + Math.PI / 2;
            break;
        case Direction.Left: // Arrow coming from right, pointing left
            angle = 2 * Math.PI;
            break;
        case Direction.Right: // Arrow coming from left, pointing right
            angle = Math.PI;
            break;
    }

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

export function getOppositeDirection(direction: Direction): Direction {
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

export function movePoint(point: Point, direction: Direction, amount: number): Point {
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

export interface ArrowColors {
    bodyColor: string;
    outlineColor: string;
}

export function getArrowColorOrDefault(arrow: Arrow): ArrowColors {
    return arrow.colors ?? defaultArrowColors();
}

export function getArrowColor(element: Element, swap: boolean, oversight: boolean): ArrowColors {
    if(element != Element.None) {
        switch(element) {
            case Element.Fire:
                return {
                    bodyColor: "var(--color-p-red-400)",
                    outlineColor: "var(--color-p-red-200)",
                };
            case Element.Water:
                return {
                    bodyColor: "var(--color-p-blue-400)",
                    outlineColor: "var(--color-p-blue-200)",
                };
            case Element.Earth:
                return {
                    bodyColor: "var(--color-p-brown-400)",
                    outlineColor: "var(--color-p-brown-200)",
                };
            case Element.Air:
                return defaultArrowColors();
            case Element.Dark:
                return {
                    bodyColor: "var(--color-p-purple-400)",
                    outlineColor: "var(--color-p-purple-200)",
                };
            case Element.Light:
                return {
                    bodyColor: "var(--color-p-orange-500)",
                    outlineColor: "var(--color-p-orange-300)",
                };
        }
    }

    if(oversight) {
        return {
            bodyColor: "var(--color-p-orange-500)",
            outlineColor: "var(--color-p-orange-300)",
        };
    }

    if(swap) {
        return defaultArrowColors();
    }
    
    return defaultArrowColors(); 
}

export function defaultArrowColors(): ArrowColors {
    return {
        bodyColor: "var(--color-bg-400)",
        outlineColor: "var(--color-bg-200)",
    };
}