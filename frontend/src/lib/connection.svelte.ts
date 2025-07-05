import type { Event } from "./gateway.svelte";
import { Gateway } from "./gateway.svelte";
import { getServerUrl } from "./index";
import { postRequestURL } from "./requests";

export let currentLobby: string | undefined = undefined;
export let currentPlayer: string | undefined = undefined;

let currentGate: Gateway | null = null;

/**
 * Connect to the gateway by joining a session.
 * 
 * Returns an error if there was one.
 */
export async function joinSession(lobbyId: string, name: string): Promise<string | undefined> {
    try {
        // Make API request to join lobby
        const response = await postRequestURL(`${getServerUrl()}/lobby/join`, {
            "lobby_id": lobbyId,
            "name": name
        });

        if (!response.success) {
            return response.message || "Failed to join lobby"
        }

        // Connect to gateway WebSocket
        currentGate = await Gateway.connect("/gateway/connect", response.token, JSON.stringify({
            "lobby_id": lobbyId,
            "player_id": response.player_id,
        }));
        currentLobby = lobbyId;
        currentPlayer = response.player_id;
        
        return undefined;
    } catch (error) {
        return `Failed to join session: ${error instanceof Error ? error.message : String(error)}`;
    }
}

/**
 * Connect to the gateway by creating a new session. 
 */
export async function createSession(name: string): Promise<string | undefined> {
    try {
        // Make API request to create lobby
        const response = await postRequestURL(`${getServerUrl()}/lobby/create`, {
            "name": name
        });

        if (!response.success) {
            return response.message || "Failed to create lobby";
        }

        // Connect to gateway WebSocket
        currentGate = await Gateway.connect("/gateway/connect", response.token, JSON.stringify({
            "lobby_id": response.lobby_id,
            "player_id": response.player_id,
        }));
        currentLobby = response.lobby_id;
        currentPlayer = response.player_id;
        
        return undefined;
    } catch (error) {
        return `Failed to create session: ${error instanceof Error ? error.message : String(error)}`;
    }
}

/**
 * Handle an event from the current gateway.
 * @param eventName The name of the event
 * @param handler The function that gets called when the event is received
 */
export function useEvent(eventName: string, handler: (event: Event) => void): void {
    if (!currentGate) return;
    
    const cleanup = currentGate.setEventHandler(eventName, handler);
    
    $effect(() => {
        return cleanup;
    });
}
