import { goto } from "$app/navigation";
import { getServerUrl } from "$lib";

export interface Event {
    name: string;
    data: any;
}

export class Gateway {
    private socket: WebSocket | null = null;
    private eventQueue: Event[] = [];
    private eventHandlers = new Map<string, (event: Event) => void>();

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
                    console.log("event", parsedEvent)
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
        const handler = this.eventHandlers.get(event.name);
        if (handler) {
            handler(event);
        } else {
            this.eventQueue.push(event);
        }
    }

    setEventHandler(eventName: string, handler: (event: Event) => void): () => void {
        this.eventHandlers.set(eventName, handler);
        
        // Process all queued events for this event name
        const remainingEvents: Event[] = [];
        while (this.eventQueue.length > 0) {
            const event = this.eventQueue.shift()!;
            if (event.name === eventName) {
                handler(event);
            } else {
                remainingEvents.push(event);
            }
        }
        this.eventQueue = remainingEvents;

        // Return cleanup function
        return () => {
            this.eventHandlers.delete(eventName);
        };
    }

    close(): void {
        if (this.socket) {
            this.socket.close();
            this.socket = null;
        }
        this.eventQueue = [];
        this.eventHandlers.clear();
    }

    private send(data: any): void {
        if (this.socket?.readyState === WebSocket.OPEN) {
            this.socket.send(JSON.stringify(data));
        }
    }
}