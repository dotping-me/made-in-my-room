import { BrowserRouter, Routes, Route, Link } from 'react-router-dom'
import Lobby from './pages/Lobby'
import Room from './pages/Room'

function App() {
    return (
        <BrowserRouter>
            <nav style={{ padding: "1rem", background: "#eee" }}>
                <Link to="/">Lobby</Link> | <Link to="/room/1">Room</Link>
            </nav>

            <Routes>
                <Route path="/" element={<Lobby />} />
                <Route path="/room/:id" element={<Room />} />
            </Routes>
        </BrowserRouter>
    )
}

export default App
