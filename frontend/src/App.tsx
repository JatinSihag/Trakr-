import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import './App.css'
import SignUp from './pages/Signup';
import Login from './pages/Login';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/signup" element={<SignUp />} />
        <Route path="/login" element={<Login />} />
        <Route path="/dashboard" element={<h1 style={{ textAlign: 'center', marginTop: '50px' }}>Dashboard (Coming Soon!)</h1>} />
        <Route path="/" element={<Navigate to="/login" replace />} />
      </Routes>
    </Router>
  )
}

export default App
