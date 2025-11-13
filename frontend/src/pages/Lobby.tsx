// TODO: Make an endpoint (Backend && Frontend)
//       For example, "/rooms/new" that just creates a new room and redirects
//                    player to that directly

// TODO: Handle errors, feedback and just reactive stuff...

import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

// On "plain" Join -> A new room is created
// On "link"  Join -> Player is assigned to an existing room
//                 -> Also forwards username if already entered

export default function Lobby(this: any) {
    const [username, setUsername] = useState<string>("");
    const [roomCode, setRoomCode] = useState<string>("");
    const nav = useNavigate();

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

    // Makes a request to backend to create room, receives code and redirects
    const createRoom = () => {
        if (username.trim().length < 2) {
            console.log("Username must be at least 3 characters!");
            return;
        }

        // Request to create room
        fetch("/api/rooms/new")
        .then((res) => { return res.json() })
        .then((json) => {
            console.log(json);

            // Handles JSON
            if (json.error) {
                console.log(json.error);
                return;
            }

            // Redirects
            sessionStorage.setItem("username", username.trim());
            nav(`/room/${json.code}`);

        })
        .catch((err) => {
            console.log(err);
        }); 
    }

    // Validation, then redirects to room
    const joinRoom = () => {
        if (username.trim().length < 2) {
            console.log("Username must be at least 3 characters!");
            return;
        }

        sessionStorage.setItem("username", username.trim());
        nav(`/room/${roomCode}`);
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

                <input type="button" value="Create Room" 
                    onClick={createRoom} />

                <input type="button" value="Join Room" 
                    onClick={joinRoom} />
            </form>
        </div>
    );
}