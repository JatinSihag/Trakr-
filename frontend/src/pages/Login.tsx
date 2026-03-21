import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import api from "../services/api";



export default function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')

    try {
      const response = await api.post('/login', { email, password });
      const token = response.data.token;

      if (token) {
        localStorage.setItem('trakr_token', token);
        navigate('/dashboard');
      } else {
        setError('No token received from server.');
      }

    } catch (err: any) {
      setError(err.response?.data?.error || 'Invalid email or password.');
    }
  };
  return (
    <div style={{ maxWidth: '400px', margin: '50px auto', fontFamily: 'sans-serif' }}>
      <h2>Welcome Back to Trakr</h2>

      {error && (
        <div style={{ padding: '10px', marginBottom: '15px', backgroundColor: '#ffebee', color: '#c62828', borderRadius: '5px' }}>
          {error}
        </div>
      )}

      <form onSubmit={handleLogin} style={{ display: 'flex', flexDirection: 'column', gap: '15px' }}>
        <input
          type="email"
          placeholder="Email Address"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
          style={{ padding: '10px', fontSize: '16px' }}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          style={{ padding: '10px', fontSize: '16px' }}
        />
        <button
          type="submit"
          style={{ padding: '10px', fontSize: '16px', backgroundColor: '#28a745', color: 'white', border: 'none', cursor: 'pointer' }}
        >
          Log In
        </button>
      </form>

      <p style={{ marginTop: '20px', textAlign: 'center' }}>
        Don't have an account? <Link to="/signup">Sign up here</Link>
      </p>
    </div>
  );
}