import { useEffect } from "react";
import { useParams } from "react-router-dom";

export default function Room() {
    const { id } = useParams();

    useEffect(() => {

        // Sends a message on join
        const ws = new WebSocket("ws://localhost:8080/ws");
        ws.onopen = () => ws.send(`Joined room ${id}`);
        ws.onmessage = (e) => console.log("Received:", e.data);

        return () => ws.close();
        
    }, [id]);

    return (
        <div>
            <h1>Room {id}</h1>
            <p>Connected to WebSocket server</p>
        </div>
    );
}