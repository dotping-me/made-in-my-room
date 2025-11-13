import { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";

interface Player {
    name: string;
}

export default function Room() {
    const { code } = useParams<{ code: string }>();
    const nav = useNavigate();

    const [username, setUsername] = useState<string>("");
    const [roomPlayers, setRoomPlayers] = useState<Player[]>([]);
    const [connected, setConnected] = useState<boolean>(false);

    // On Mount
    useEffect(() => {

        // Checks if username exists
        let u: string | null = sessionStorage.getItem("username");
        console.log(u);
        if (!u) { nav("/"); return; } // Returns to root if no username set

        setUsername(u);

        // Checks if room exists (Redirects player if not)
        fetch("/api/rooms/exists?room=" + code)
        .then((res) => { return res.json() })
        .then((json) => {
            if (!json.exists) {
                nav("/");
                return;
            }
        });
    }, []);

    // Establishes connection
    useEffect(() => {
        const ws = new WebSocket(`ws://localhost:8080/ws?room=${code}&name=${username}`);
        ws.onopen = () => {
            console.log(`Connected to room ${code}`);
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
                console.error("Invalcode WS message:", e.data);
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

    }, [username]);

    return (
        <div>
            <h1>Room {code}</h1>
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