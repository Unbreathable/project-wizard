import { postRequestURL } from "./requests";

export enum Element {
    Fire = "fire",
    Water = "water",
    Earth = "earth",
    Air = "air",
    Dark = "dark",
    Light = "light",
    None = "none",
}

export interface Action {
    name: string;
    description: string;
    element: Element;
    damage: number;
    mana_cost: number;
    oversight: boolean;
}

export interface Character {
    id: number;
    url: string;
    name?: string;
    origin?: string;
    elements?: Element[];
    actions?: Record<number, Action>;
}

export let characters: Record<number, Character> = {
    1: {
        id: 1, 
        url: "chisato.png",
    },
    2: { 
        id: 2, 
        url: "takina.png" 
    },
    3: { 
        id: 3, 
        url: "rimuru.png" 
    },
    4: { 
        id: 4, 
        url: "shuna.png" 
    },
    5: { 
        id: 5, 
        url: "maomao.png" 
    },
    6: { 
        id: 6, 
        url: "wraith.png" 
    },
    7: { 
        id: 7, 
        url: "index.png" 
    },
};

export async function initializeCharacters() {
    try {
        const response = await postRequestURL("/info/characters", {});
        
        if (response.success && response.characters) {
            // Update the characters record with data from the backend
            for (const character of response.characters) {
                const existingCharacter = characters[character.id];
                if (existingCharacter) {
                    // Merge backend data with existing character data (keeping the URL)
                    characters[character.id] = {
                        ...existingCharacter,
                        name: character.name,
                        origin: character.origin,
                        elements: character.elements,
                        actions: character.actions
                    };
                }
            }
        } else {
            console.error("Failed to fetch characters:", response.message || "Unknown error");
        }
    } catch (error) {
        console.error("Error initializing characters:", error);
    }
}