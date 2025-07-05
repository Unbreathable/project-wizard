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
    Damage: number;
    ManaCost: number;
    Oversight: boolean;
}

export interface Character {
    id: number;
    url: string;
    name?: string;
    origin?: string;
    elements?: Element[];
    actions?: Action[];
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