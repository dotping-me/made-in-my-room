// TODO: Make an endpoint (Backend && Frontend)
//       For example, "/rooms/new" that just creates a new room and redirects
//                    player to that directly

import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

// On "plain" Join -> A new room is created
// On "link"  Join -> Player is assigned to an existing room
//                 -> Also forwards username if already entered

export default function Lobby(this: any) {
    const [username, setUsername] = useState<string>("");
    const [roomCode, setRoomCode] = useState<string>("");
    const nav = useNavigate();

    // On Mount 
    // NOTE: Just temporary to make backend create rooms
    useEffect(() => {
        fetch("/api/rooms")
    }, []);

    useEffect(() => {

        // Checks whether room exists
        if (roomCode) {
            fetch("/api/rooms/exists?room=" + roomCode)
            .then((res) => { return res.json() })
            .then((json) => {

                // Disables "Join" button
                if (!json.exists) {
                    return;
                }

                // Enables button
            });
        }

    }, [roomCode]);

    // TODO: Handle errors, feedback and just reactive stuff...
    const joinRoom = () => {
        if (username.trim().length < 4) {
            console.log("Username must be at least 5 characters!");
            return;
        }

        nav(`/room/${roomCode}?username=${encodeURIComponent(username)}`);
    };

    return (
        <div>
            <h1>Made in My Room</h1>
            <form name="joinRoomForm">
                <div className="input-wrapper">
                    <label htmlFor="u">Username</label>
                    <input type="text" name="username" id="username" 
                        onInput={(e) => setUsername((e.target as HTMLInputElement).value.trim())} />
                </div>

                <div className="input-wrapper">
                    <label htmlFor="code">Room Code</label>
                    <input type="text" name="room" id="room"
                        onBlur={(e) => setRoomCode((e.target as HTMLInputElement).value.trim())} />
                </div>

                <input type="button" value="Join Room" 
                    onClick={joinRoom} />
            </form>
        </div>
    );
}