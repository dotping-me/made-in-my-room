import { useState, useEffect } from "react";
import { useParams, useSearchParams } from "react-router-dom";

interface Player {
    name: string;
}

export default function Room() {
    const { id } = useParams<{ id: string }>();
    const [searchParams] = useSearchParams();

    const [username, setUsername] = useState<string>("");
    const [roomPlayers, setRoomPlayers] = useState<Player[]>([]);
    const [connected, setConnected] = useState<boolean>(false);

    // Set username on load
    useEffect(() => {
        const nameFromQuery = searchParams.get("username");
        if (nameFromQuery) {
            setUsername(decodeURIComponent(nameFromQuery));
        }
    }, [searchParams]);

    useEffect(() => {
        if (!username) {
            return
        };

        const ws = new WebSocket(`ws://localhost:8080/ws?room=${id}&name=${username}`);
        ws.onopen = () => {
            console.log(`Connected to room ${id}`);
            setConnected(true);
        };

        // Handles messages received
        ws.onmessage = (e) => {
            try {
                const msg = JSON.parse(e.data);
                if (msg.type === "players") {
                    setRoomPlayers(msg.data);
                } 
                
                else if (msg.type === "info") {
                    console.log("Info:", msg.data);
                }
            
            } catch (err) {
                console.error("Invalid WS message:", e.data);
            }
        };

        ws.onclose = () => {
            console.log("Disconnected from server");
            setConnected(false);
        };

        ws.onerror = (err) => {
            console.error("WebSocket error:", err);
        };

        return () => ws.close();

    }, [id, username]);

    // Before joining, ask the user for a name
    if (!username) {
    return (
    <div className="p-4">
    <h1 className="text-xl mb-2">Enter your name to join Room {id}</h1>
    <input
    type="text"
    placeholder="Your name"
    className="border px-2 py-1 rounded"
    onKeyDown={(e) => {
    if (e.key === "Enter" && e.currentTarget.value.trim()) {
    setUsername(e.currentTarget.value.trim());
    }
    }}
    />
    </div>
    );
    }

    return (
        <div>
            <h1>Room {id}</h1>
            <p>
                {connected ? (
                    <span>Connected as {username}</span>
                ) : (
                    <span>Disconnected</span>
                )}
            </p>

            <h2>Players:</h2>
            <ul>
                {roomPlayers.length > 0 ? (
                    roomPlayers.map((p, i) => <li key={i}>{p.name}</li>)
                ) : (
                    <p>No one's here</p>
                )}
            </ul>
        </div>
    );
}