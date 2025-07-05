import { goto } from "$app/navigation";
import { getServerUrl } from "$lib";

export interface Event {
    name: string;
    data: any;
}

export class Gateway {
    private socket: WebSocket | null = null;
    private eventHandlers = new Map<string, Set<(event: Event) => void>>();

    private constructor() {}

    static async connect(path: string, token: string, attachments: string): Promise<Gateway> {
        const gateway = new Gateway();
        
        const serverWs = getServerUrl() + path;
        return new Promise((resolve, reject) => {
            const ws = new WebSocket(serverWs);
            gateway.socket = ws;

            ws.onopen = async ()=> {
                await gateway.send({ "token": token, "attachments": attachments });
                resolve(gateway);
            };

            ws.onerror = () => reject(new Error('WebSocket connection failed'));

            ws.onmessage = async (event) => {
                try {
                    let blob = event.data as Blob;
                    const parsedEvent = JSON.parse(await blob.text());
                    gateway.notifyHandlers(parsedEvent);
                } catch(e) {
                    console.error("gate read failed", e)
                    gateway.close();
                }
            };

            ws.onclose = () => {
                gateway.socket = null;
                console.log("gateway connection lost")
                goto("/")
            };
        });
    }

    private notifyHandlers(event: Event): void {
        const handlers = this.eventHandlers.get(event.name);
        if (handlers) {
            handlers.forEach(handler => handler(event));
        }
    }

    addEventHandler(eventName: string, handler: (event: Event) => void): () => void {
        if (!this.eventHandlers.has(eventName)) {
            this.eventHandlers.set(eventName, new Set());
        }
        this.eventHandlers.get(eventName)!.add(handler);

        // Return cleanup function
        return () => {
            const handlers = this.eventHandlers.get(eventName);
            if (handlers) {
                handlers.delete(handler);
                if (handlers.size === 0) {
                    this.eventHandlers.delete(eventName);
                }
            }
        };
    }

    close(): void {
        if (this.socket) {
            this.socket.close();
            this.socket = null;
        }
    }

    private send(data: any): void {
        if (this.socket?.readyState === WebSocket.OPEN) {
            this.socket.send(JSON.stringify(data));
        }
    }
}