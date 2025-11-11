import { useEffect, useState } from "react";

export default function Lobby() {
    const [lobbies, setLobbies] = useState<{ id: string; name: string }[]>([]);

    useEffect(() => {
        fetch("/api/lobby")
        .then((res) => res.json())
        .then(setLobbies);
    }, []);

    return (
        <div>
            <h1>Lobby List</h1>
            <ul>
                {lobbies.map((lobby) => (
                    <li key={lobby.id}>{lobby.name}</li>
                ))}
            </ul>
        </div>
    );
}