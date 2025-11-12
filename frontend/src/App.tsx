import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Room from './pages/Room'
import Lobby from './pages/Lobby'

export default function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Lobby />} />        {/* Form to join room */}
                <Route path="/room/:id" element={<Room />} /> {/* Joins existing room */}
            </Routes>
        </BrowserRouter>
    )
}